import axios from 'axios'

const client = axios.create({
    headers: {Authorization: `Bearer ${localStorage.getItem('_token')}`}
})

export type Paginated<T> = {
    list: T[]
    total: number
}

export type Project = {
    id: number
    name: string
    repo_address: string
    default_branch: string
}

export function getProjects(page: number = 1) {
    return client.get<any, Paginated<Project>>("/api/projects?page=" + page)
}