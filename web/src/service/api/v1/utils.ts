import { request, StandartResponse } from "./_request";

interface GetLocalIPBody {
    local_ip: string;
}

function GetLocalIP(): Promise<GetLocalIPBody> {
    return new Promise((resolve, reject) => {
        request({
            method: "GET",
            url: "/utils/local_ip",
        })
            .then((response) =>
                resolve((response.data as StandartResponse).data as GetLocalIPBody)
            )
            .catch(reject);
    });
}

export { GetLocalIP };
