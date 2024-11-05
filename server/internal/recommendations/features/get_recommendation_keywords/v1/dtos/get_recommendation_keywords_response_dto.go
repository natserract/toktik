package dtos

type GetRecommendationsData struct {
	Titles []string `json:"titles"`
}

type GetRecommendationsResponseDTO struct {
	Data GetRecommendationsData `json:"data"`
}
