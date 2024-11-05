package dtos

type GetRecommendationTagsData struct {
	Tags []string `json:"tags"`
}

type GetRecommendationTagsResponseDTO struct {
	Data GetRecommendationTagsData `json:"data"`
}
