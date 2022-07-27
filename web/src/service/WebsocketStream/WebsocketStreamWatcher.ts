import { WebsocketMessage, WebsocketMessageData } from "./types";
import { ICE_SERVERS_URLS } from "./consts";

class BroadcastWatcher {
    readonly videoTag: HTMLMediaElement;
    readonly pc: RTCPeerConnection;

    constructor(videoTag: HTMLMediaElement) {
        this.videoTag = videoTag;
        this.pc = this.newPeerConnection();
    }

    public SetRemoteOffer(remoteSDP: RTCSessionDescription): Promise<void> {
        return new Promise((resolve, _) => {
            this.pc.setRemoteDescription(remoteSDP).then(resolve);
        });
    }

    public NewAnswer(): Promise<RTCSessionDescription> {
        return new Promise((resolve, _) => {
            this.pc.onicecandidate = (event) => {
                if (event.candidate === null && this.pc.localDescription) {
                    console.info("@pc [ice null candidate]:", event);
                    resolve(this.pc.localDescription);
                }
            };

            this.pc.createAnswer().then((answer) => {
                this.pc.setLocalDescription(answer);
            });
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

        pc.ontrack = (event) => {
            console.info("@pc [track]:", event);
            this.videoTag.srcObject = event.streams[0];
        };

        return pc;
    }
}

class WebsocketStreamWatcher {
    private readonly bcWatcher: BroadcastWatcher;
    private readonly ws: WebSocket;

    constructor(videoTag: HTMLMediaElement) {
        this.bcWatcher = new BroadcastWatcher(videoTag);
        this.ws = this.newWSConn();
    }

    private newWSConn(): WebSocket {
        const ws = new WebSocket(
            `ws://${location.hostname}:${
                import.meta.env.VITE_API_PORT
            }/api/v1/ws/stream/watch`
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

            if (json.error) return console.log(json.error);
            else if (json.data) this.sendAnswer(json.data);
        };

        return ws;
    }

    private async sendAnswer(data: WebsocketMessageData): Promise<void> {
        let offer = new RTCSessionDescription({
            type: data.type,
            sdp: data.sdp,
        });

        await this.bcWatcher.SetRemoteOffer(offer);

        this.bcWatcher.NewAnswer().then((answer) => {
            let msg: WebsocketMessageData = {
                broadcastID: data.broadcastID,
                type: answer.type,
                sdp: answer.sdp,
            };

            this.ws.send(JSON.stringify(msg));
        });
    }
}

export { WebsocketStreamWatcher };
