import {get, post} from "../utils.ts";

export type Key = {
    id: number
    name: string
    public_key: string
    private_key: string
    created_at: string
    updated_at: string
}

export function getKeys() {
    return get<Key[]>('/api/manage/keys')
}

export function createKey(key: Key) {
    return post<Key>('/api/manage/key', key)
}
