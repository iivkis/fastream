export interface WebsocketMessageData {
    broadcastID: string;
    type: RTCSdpType;
    sdp: string;
}

export interface WebsocketMessage {
    data?: WebsocketMessageData;
    error?: string;
}

export {};
