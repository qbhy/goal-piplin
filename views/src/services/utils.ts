import axios, {AxiosResponse} from "axios";

export function get<T>(uri: string) {
    return axios.get<any, AxiosResponse<T>>(uri, {
        headers: {Authorization: `Bearer ${localStorage.getItem('_token')}`}
    }).then(res => {
        return res.data
    }).catch(function (res) {
        if (res.response.status == 401) {
            location.href = '/login'
        }
        throw res
    })
}

export function post<T>(uri: string, data: any) {
    return axios.post<any, AxiosResponse<T>>(uri, data, {
        headers: {Authorization: `Bearer ${localStorage.getItem('_token')}`}
    }).then(res => {
        return res.data
    }).catch(function (res) {
        if (res.response.status == 401) {
            location.href = '/login'
        }
        throw res
    })
}

