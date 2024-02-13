import {get, post} from "../utils.ts";

export type Group = {
    id: number
    name: string
    creator_id: number
    created_at: string
    updated_at: string
}

export function getGroups() {
    return get<Group[]>('/api/manage/groups')
}

export function createGroup(group: Group) {
    return post<Group[]>('/api/manage/group', group)
}