import { WebsocketMessage, WebsocketMessageData } from "./types";
import { ICE_SERVERS_URLS } from "./consts";
import { v4 as uuidv4 } from "uuid";

class BroadcastStorage {
    private broadcasts: Array<BroadcastCreator>;

    constructor() {
        this.broadcasts = [];
    }

    public Add(broadcast: BroadcastCreator): void {
        this.broadcasts.push(broadcast);
    }

    public GetByID(id: string): BroadcastCreator | null {
        for (let i = 0; i < Object.keys(this.broadcasts).length; i++)
            if (this.broadcasts[i].ID === id) return this.broadcasts[i];
        return null;
    }
}

// BroadcastCreator исользуется для каждого подключения по WebRTC,
// чтобы установить peer2peer соединение со стримером и зрителем.
// Передача SDP из RTCPeerConnection (WebRTC) между стримером и зрителем переходит
// по веб-сокетам (на уровне выше) через сервер.
class BroadcastCreator {
    // ID (uuid) - уникальное значение каждой трансляции.
    // Нужно, чтобы одновременно можно было подключить несколько зрителей,
    // иначе offer sdp и answer sdp будут путаться между собой.
    public readonly ID: string;
    private readonly pc: RTCPeerConnection;

    constructor(stream: MediaStream) {
        this.ID = uuidv4();

        this.pc = this.newPeerConnection();
        this.addStreamToPeerConn(stream);
    }

    public Conn(): RTCPeerConnection {
        return this.pc;
    }

    //создает новый webRTC offer, оторый будет отправлен зрителю
    public NewOffer(): Promise<RTCSessionDescription> {
        return new Promise((resolve, _) => {
            this.pc.onicecandidate = (event) => {
                if (event.candidate === null && this.pc.localDescription) {
                    console.info("@pc [ice null candidate]:", event);
                    resolve(this.pc.localDescription);
                }
            };

            this.pc.createOffer().then((sdp) => {
                this.pc.setLocalDescription(sdp);
            });
        });
    }

    //устанвливает WebRTC answer зрителя
    public SetRemoteAnswer(remoteSDP: RTCSessionDescription): Promise<void> {
        return new Promise((resolve, _) => {
            this.pc.setRemoteDescription(remoteSDP).then(resolve);
        });
    }

    private newPeerConnection(): RTCPeerConnection {
        let pc = new RTCPeerConnection({
            iceServers: [
                {
                    urls: ICE_SERVERS_URLS,
                },
            ],
        });

        pc.oniceconnectionstatechange = () => {
            console.info(
                "@pc [ice connection state change]:",
                pc.iceConnectionState
            );
        };

        pc.onicecandidateerror = (err) => {
            console.error("@pc [ice candidate error]:", err);
        };

        return pc;
    }

    private addStreamToPeerConn(stream: MediaStream): void {
        stream.getTracks().forEach((track) => {
            this.pc.addTrack(track, stream);
        });
    }
}

// Веб-сокет обрабочтик, который обменивается SDP.
class WebsocketStreamCreator {
    private readonly stream: MediaStream;
    private readonly broadcastStorage: BroadcastStorage;
    private readonly ws: WebSocket;

    constructor(stream: MediaStream) {
        this.stream = stream;
        this.broadcastStorage = new BroadcastStorage();
        this.ws = this.newWSConn();
    }

    public Conn(): WebSocket {
        return this.ws;
    }

    private newWSConn(): WebSocket {
        const ws = new WebSocket(
            `ws://${location.hostname}:${
                import.meta.env.VITE_API_PORT
            }/api/v1/ws/stream/create`
        );
        console.info("@ws: connectiong...");

        ws.onerror = (err) => {
            console.error("@ws:", err);
        };

        ws.onopen = () => {
            console.info("@ws: success connected to server");
        };

        ws.onmessage = ({ data }) => {
            let json = JSON.parse(data) as WebsocketMessage;
            console.info("@ws [message]:", json);

            if (json.error) return console.error(json.error);
            else if (json.data?.broadcastID) this.setRemoteAnswer(json.data);
            else this.sendOffer();
        };

        return ws;
    }

    private sendOffer(): void {
        let broadcast = new BroadcastCreator(this.stream);

        broadcast.NewOffer().then((sdp) => {
            let msg: WebsocketMessageData = {
                broadcastID: broadcast.ID,
                type: sdp.type,
                sdp: sdp.sdp,
            };

            this.ws.send(JSON.stringify(msg));
        });

        this.broadcastStorage.Add(broadcast);
    }

    private setRemoteAnswer(data: WebsocketMessageData): void {
        let broadcast = this.broadcastStorage.GetByID(data.broadcastID);

        if (broadcast) {
            let sdp = new RTCSessionDescription({
                type: data.type,
                sdp: data.sdp,
            });

            broadcast.SetRemoteAnswer(sdp);
        } else {
            console.error("undefined broadcastID", data.broadcastID);
        }
    }
}

export { WebsocketStreamCreator };
