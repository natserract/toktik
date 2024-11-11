import axios, { type AxiosError, AxiosInstance, AxiosResponse, InternalAxiosRequestConfig } from 'axios'
import { ENDPOINTS } from '../constants/endpoints'

type ObjectValue<T> = T[keyof T]
type Endpoint = ObjectValue<typeof ENDPOINTS>

export class BaseClient {
  BASE_URL: string
  private api: AxiosInstance

  constructor(baseUrl: string) {
    this.BASE_URL = baseUrl
    if (!this.BASE_URL) {
      throw new Error('Please provide base url!')
    }

    this.api = axios.create({
      baseURL: this.BASE_URL,
    })
    this.api.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => config,
      (error: AxiosError<string>) => Promise.reject(error)
    )

    this.api.interceptors.response.use(
      (response: AxiosResponse) => response,
      (error: AxiosError<string>) => Promise.reject(error)
    )
  }

  protected async getResource<T>(endpoint: string, identifier?: string | number): Promise<T> {
    return (await this.api.get<T>(`${endpoint}/${identifier || identifier === 0 ? identifier : ''}`)).data
  }

  protected async getResources<T>(endpoint: string, params: object = {}): Promise<T> {
    return (
      await this.api.get<T>(endpoint, {
        params: { ...params },
      })
    ).data
  }

  protected async getListResource<T>(endpoint: Endpoint, offset = 0, limit = 20): Promise<T> {
    if (offset < 0 || limit < 0) {
      return (await this.api.get<T>(`${endpoint}`)).data
    }

    return (await this.api.get<T>(`${endpoint}?offset=${offset}&limit=${limit}`)).data
  }
}
