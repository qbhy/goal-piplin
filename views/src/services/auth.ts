import axios, {AxiosResponse} from 'axios'

const client = axios.create({
    headers: {Authorization: `Bearer ${localStorage.getItem('_token')}`}
})

export type LoginData = {
    token?: string
    user?: User
    msg?: string
}

export type User = {
    avatar: string
    created_at: string
    id: string
    nickname: string
    password: string
    role: string
    updated_at: string
    username: string
}

export async function login(data: any): Promise<LoginData> {
    const res = await client.post<any, AxiosResponse<LoginData>>('/api/login', data)
    if (res.data.token) {
        localStorage.setItem("_token", res.data.token)
        localStorage.setItem("_user", JSON.stringify(res.data.user))
        client.defaults.headers.Authorization = `Bearer ${localStorage.getItem('_token')}`
    }
    return res.data
}

export async function getMyself(): Promise<User> {
    let localUser = localStorage.getItem("_user")
    if (localUser && localUser != "") {
        return JSON.parse(localUser)
    }
    const res = await client.get<any, AxiosResponse<User>>("/api/myself")
    localStorage.setItem('_user', JSON.stringify(res.data))
    return res.data
}