import axios, { type AxiosError } from 'axios'
import { type AxiosCacheInstance, type CacheAxiosResponse, type InternalCacheRequestConfig, setupCache } from 'axios-cache-interceptor'

type ObjectValue<T> = T[keyof T]
type Endpoint = ObjectValue<any>

export class BaseClient {
  private api: AxiosCacheInstance

  constructor() {
    this.api = setupCache(
      axios.create({
        baseURL: 'https://jsonplaceholder.typicode.com',
        headers: {
          'Content-Type': 'application/json',
        },
      })
    )

    this.api.interceptors.request.use(
      (config: InternalCacheRequestConfig) => config,
      (error: AxiosError<string>) => Promise.reject(error)
    )

    this.api.interceptors.response.use(
      (response: CacheAxiosResponse) => response,
      (error: AxiosError<string>) => Promise.reject(error)
    )
  }

  protected async getResource<T>(endpoint: string, identifier?: string | number): Promise<T> {
    return (await this.api.get<T>(`${endpoint}/${identifier || identifier === 0 ? identifier : ''}`)).data
  }

  protected async getListResource(endpoint: Endpoint, offset = 0, limit = 20): Promise<any> {
    return (await this.api.get<any>(`${endpoint}?offset=${offset}&limit=${limit}`)).data
  }
}
