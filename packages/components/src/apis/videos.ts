import { BaseClient } from './client'

export class VideosAPI extends BaseClient {
  public async listPosts(offset?: number, limit?: number): Promise<any> {
    return this.getListResource('/posts', offset, limit)
  }
}
const videosAPI = new VideosAPI()

export default videosAPI
