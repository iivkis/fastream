const ICE_SERVERS_URLS = [
    "stun:stun.l.google.com:19302",
    "stun:stun1.l.google.com:19302",
    "stun:stun2.l.google.com:19302",
    "stun:stun3.l.google.com:19302",
    "stun:stun4.l.google.com:19302",
];

export interface WebsocketMessageData {
    broadcastID: string;
    type: RTCSdpType;
    sdp: string;
}

export interface WebsocketMessage {
    data?: WebsocketMessageData;
    error?: string;
}

export {ICE_SERVERS_URLS};
