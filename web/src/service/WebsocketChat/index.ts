export interface WebsocketMessageData {
    username: string;
    content: string;
}

interface WebsocketMessage {
    data?: WebsocketMessageData;
    error?: string;
}

class WebsocketChat {}

export { WebsocketChat };
