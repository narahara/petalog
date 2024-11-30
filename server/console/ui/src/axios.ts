import axios from 'axios'
import { AxiosInstance } from 'axios'

export class ApiClient{
    private apiClient: AxiosInstance;

    constructor(baseUrl: string) {
        this.apiClient = axios.create(
            {baseURL: baseUrl}
        );
        this.apiClient.interceptors.request.use((config: any) => {
            try {
                config.headers['Authorization'] = 'Bearer ' + "1234565";
            } catch (e) {
                return Promise.resolve(config);
            }
            return config;
        })
    }
    
    async get(url: string, data: any = {}) {
        return await this.apiClient.get(url, data);
    }

    async post(url: string, data: any, config: any = {}) {
        return await this.apiClient.post(url, data, config);
    }

}
