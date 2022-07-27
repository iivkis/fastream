import { WebsocketMessage, WebsocketMessageData } from "./types";

class WebsocketChat {
    private readonly ws: WebSocket;

    constructor() {
        this.ws = this.newConnection();
    }

    public OnMessage(message: WebsocketMessageData) {}

    public Send(message: WebsocketMessageData) {
        this.ws.send(JSON.stringify(message));
    }

    private newConnection(): WebSocket {
        const conf = {
            host: location.hostname,
            port: import.meta.env.VITE_API_PORT,
            url: "api/v1/ws/chat",
        };

        let ws = new WebSocket(`ws://${conf.host}:${conf.port}/${conf.url}`);

        ws.onerror = (err) => {
            console.error(`@ws error:`, err);
        };

        ws.onopen = () => {
            console.info(`@ws: connected to chat`);
        };

        ws.onmessage = ({ data }) => {
            let json = JSON.parse(data) as WebsocketMessage;

            if (json.error) return console.log("@ws chat error:", json);
            else if (json.data) this.OnMessage(json.data);
        };

        return ws;
    }
}

export { WebsocketChat };
