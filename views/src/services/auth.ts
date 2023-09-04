import axios from 'axios'

export function login(data: any) {
    return axios.post('/api/login', data)
}