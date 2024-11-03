package dtos

type GetRecommendationsData struct {
	Tags   []string `json:"tags"`
	Titles []string `json:"titles"`
}

type GetRecommendationsResponseDTO struct {
	Data GetRecommendationsData `json:"data"`
}
