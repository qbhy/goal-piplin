import {get, post} from "./utils.ts";

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
    const res = await post<LoginData>('/api/login', data)
    if (res.token) {
        localStorage.setItem("_token", res.token)
        localStorage.setItem("_user", JSON.stringify(res.user))
    }
    return res
}

export async function getMyself(): Promise<User> {
    let localUser = localStorage.getItem("_user")
    if (localUser && localUser != "") {
        return JSON.parse(localUser)
    }
    const res = await get<User>("/api/myself")
    localStorage.setItem('_user', JSON.stringify(res))
    return res
}