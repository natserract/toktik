package dtos

type GetRecommendationsData struct {
	Keywords []string `json:"keywords"`
}

type GetRecommendationsResponseDTO struct {
	Data GetRecommendationsData `json:"data"`
}
