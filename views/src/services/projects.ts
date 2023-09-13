import {get} from "./utils.ts";

export type Paginated<T> = {
    list: T[]
    total: number
}

export type Project = {
    id: number
    name: string
    group_id: number
    repo_address: string
    default_branch: string
}

export function getProjects(page: number = 1) {
    return get<Paginated<Project>>("/api/projects?page=" + page)
}