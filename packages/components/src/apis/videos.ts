import { BaseClient } from './client'
import { ENDPOINTS } from '../constants/endpoints'

export class VideosAPI extends BaseClient {
  public async searchVideos(keywords: string, count: number): Promise<any[]> {
    return this.getResources(`${ENDPOINTS.VIDEOS}/search`, {
      keywords,
      count,
    })
  }

  public async getVideoById(id: string): Promise<any> {
    return this.getResource(ENDPOINTS.VIDEOS, id)
  }

  public async getStreamVideoById(id: string): Promise<any> {
    return this.getResource(`${ENDPOINTS.VIDEOS}/stream`, id)
  }
}
const videosAPI = new VideosAPI()

export default videosAPI
