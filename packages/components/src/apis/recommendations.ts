import { BaseClient } from './client'
import { ENDPOINTS } from '../constants/endpoints'

export class RecommendationsAPI extends BaseClient {
  public async listTags(): Promise<any> {
    return this.getListResource(`${ENDPOINTS.RECOMMENDATIONS}/tags`, -1, -1)
  }

  public async listKeywords(): Promise<any> {
    return this.getListResource(`${ENDPOINTS.RECOMMENDATIONS}/keywords`, -1, -1)
  }
}
const recommendationsAPI = new RecommendationsAPI()

export default recommendationsAPI
