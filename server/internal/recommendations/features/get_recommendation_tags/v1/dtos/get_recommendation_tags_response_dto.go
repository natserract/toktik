package dtos

type GetRecommendationsTagsData struct {
	Tags []string `json:"tags"`
}

type GetRecommendationsTagsResponseDTO struct {
	Data GetRecommendationsTagsData `json:"data"`
}
