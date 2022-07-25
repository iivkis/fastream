import axios from 'axios'

export interface StandartResponse {
    data?: any;
    error?: string;
} 

const request = axios;
request.defaults.baseURL = `http://${location.hostname}:${import.meta.env.VITE_API_PORT}/api/v1/`

export {request};