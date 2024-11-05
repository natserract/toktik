import { BaseClient } from './client'
import { ENDPOINTS } from '../constants/endpoints'
import { RecommendationTagsAPI, RecommendationKeywordsAPI } from '@toktik/contracts'

export class RecommendationsAPI extends BaseClient {
  public async listTags(): Promise<RecommendationTagsAPI.GetAllResponse> {
    return this.getListResource(`${ENDPOINTS.RECOMMENDATIONS}/tags`, -1, -1)
  }

  public async listKeywords(): Promise<RecommendationKeywordsAPI.GetAllResponse> {
    return this.getListResource(`${ENDPOINTS.RECOMMENDATIONS}/keywords`, -1, -1)
  }
}
const recommendationsAPI = new RecommendationsAPI()

export default recommendationsAPI
