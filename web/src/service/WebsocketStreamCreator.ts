class BroadcastCreator {
    private readonly stream: MediaStream;
    private readonly pc: RTCPeerConnection;

    constructor(stream: MediaStream) {
        this.stream = stream;

        this.pc = this.NewPeerConnection();
    }

    public Conn(): RTCPeerConnection {
        return this.pc;
    }

    private NewPeerConnection(): RTCPeerConnection {
        let pc = new RTCPeerConnection({
            iceServers: [
                {
                    urls: [
                        "stun:stun.l.google.com:19302",
                        "stun:stun1.l.google.com:19302",
                        "stun:stun2.l.google.com:19302",
                        "stun:stun3.l.google.com:19302",
                        "stun:stun4.l.google.com:19302",
                    ],
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
}

interface WebSocketMessage {
    error: string;
    type: string;
    sdp: string;
}

class WebsocketStreamCreator {
    private readonly stream: MediaStream;
    private readonly ws: WebSocket;

    constructor(stream: MediaStream) {
        this.stream = stream;
        this.ws = this.NewWebsocketConnection();
    }

    public Conn(): WebSocket {
        return this.ws;
    }

    private NewWebsocketConnection(): WebSocket {
        const ws = new WebSocket("/api/v1/ws/stream/create");

        ws.onerror = (err) => {
            console.error("@ws:", err);
        };

        ws.onopen = () => {
            console.info("@ws: success connected to server");
        };

        ws.onmessage = ({ data }) =>
            this.OnMessage(JSON.parse(data) as WebSocketMessage);

        return ws;
    }

    private OnMessage(json: WebSocketMessage): void {
        if (json.error) return console.error(json.error);

        // if (json.type)
    }
}

export { WebsocketStreamCreator };
