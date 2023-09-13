import {get, post} from "./utils.ts";

export type Paginated<T> = {
    list: T[]
    total: number
}

export type Project = {
    id: number
    name: string
    group_id: number
    key_id: number
    repo_address: string
    project_path: string
    default_branch: string
}

export function getProjects(page: number = 1) {
    return get<Paginated<Project>>("/api/projects?page=" + page)
}

export function createProject(project: Project) {
    return post<Project>("/api/project", project)
}