import { BaseClient } from './client'
import { ENDPOINTS } from '../constants/endpoints'
import { FeedsAPI } from '@toktik/contracts'

export class VideosAPI extends BaseClient {
  public async searchVideos(keywords: string, count: number): Promise<FeedsAPI.SearchResponse> {
    return this.getResources<FeedsAPI.SearchResponse>(`${ENDPOINTS.FEEDS}/search`, {
      keywords,
      count,
    })
  }

  public async getVideoById(id: string): Promise<any> {
    return this.getResource(ENDPOINTS.FEEDS, id)
  }

  public async getStreamVideoById(id: string): Promise<any> {
    return this.getResource(`${ENDPOINTS.FEEDS}/stream`, id)
  }
}
const videosAPI = new VideosAPI()

export default videosAPI
