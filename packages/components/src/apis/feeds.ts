import { BaseClient } from './client'
import { ENDPOINTS } from '../constants/endpoints'
import { FeedsAPI as FeedsAPIContract } from '@toktik/contracts'

export class FeedsAPI extends BaseClient {
  public async searchVideos(keywords: string, count: number): Promise<FeedsAPIContract.SearchResponse> {
    return this.getResources<FeedsAPIContract.SearchResponse>(`${ENDPOINTS.FEEDS}/search`, {
      keywords,
      count,
    })
  }

  public async getVideoById(id: string): Promise<any> {
    return this.getResource(ENDPOINTS.FEEDS, id)
  }

  public getStreamVideoUrl(id: string): string {
    return `${this.BASE_URL}${ENDPOINTS.FEEDS}/${id}/stream`
  }

  public async getStreamVideoById(id: string): Promise<any> {
    return this.getResource(`${ENDPOINTS.FEEDS}/stream`, id)
  }
}
