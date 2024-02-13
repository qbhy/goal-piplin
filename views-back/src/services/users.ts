import {get, post} from "./utils.ts";

export type Paginated<T> = {
    list: T[]
    total: number
}

export type User = {
    id: number
    name: string
}

export function getUsers(page: number = 1) {
    return get<Paginated<User>>("/api/projects?page=" + page)
}

export function createUser(user: User) {
    return post<User>("/api/project", user)
}