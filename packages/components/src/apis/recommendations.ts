import { BaseClient } from './client'
import { ENDPOINTS } from '../constants/endpoints'
import { RecommendationTagsAPI as RecommendationTagsAPIContract, RecommendationKeywordsAPI as RecommendationKeywordsAPIContract } from '@toktik/contracts'

export class RecommendationsAPI extends BaseClient {
  public async listTags(): Promise<RecommendationTagsAPIContract.GetAllResponse> {
    return this.getListResource(`${ENDPOINTS.RECOMMENDATIONS}/tags`, -1, -1)
  }

  public async listKeywords(): Promise<RecommendationKeywordsAPIContract.GetAllResponse> {
    return this.getListResource(`${ENDPOINTS.RECOMMENDATIONS}/keywords`, -1, -1)
  }
}
