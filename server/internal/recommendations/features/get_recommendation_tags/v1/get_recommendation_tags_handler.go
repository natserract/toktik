package v1

import (
	"context"
	"fmt"
	"log"

	"github.com/natserract/toktik/embedding"
	"github.com/natserract/toktik/internal/recommendations/features/get_recommendation_tags/v1/dtos"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
	"github.com/natserract/toktik/internal/user_interests_embedding/shared/util"
)

type GetRecommendationTagsHandler struct {
	inMemoryRepository repositories.UserInterestsEmbeddingRepository
}

func NewGetRecommendationTagsHandler(r repositories.UserInterestsEmbeddingRepository) *GetRecommendationTagsHandler {
	return &GetRecommendationTagsHandler{
		inMemoryRepository: r,
	}
}

func (c *GetRecommendationTagsHandler) Handle(
	ctx context.Context,
	query interface{},
) (*dtos.GetRecommendationsTagsResponseDTO, error) {
	results, err := c.inMemoryRepository.GetAllUserInterestsEmbedding()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	embed := embedding.NewVectorEmbedding()
	user_profile_vec, err := embed.LoadQueryVector("sample/user_profile.json")
	if err == nil {
		bestText, similarityScore, err := util.FindMostSimilar(user_profile_vec.Values, results, "tags")
		if err != nil {
			return nil, err
		}
		fmt.Println("Found similarity: ", bestText, similarityScore)
	}

	return nil, nil
}
