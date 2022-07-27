export interface WebsocketMessageData {
    username: string;
    content: string;
}

export interface WebsocketMessage {
    data?: WebsocketMessageData;
    error?: string;
}

export {};
