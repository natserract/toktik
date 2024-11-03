package v1

import (
	"context"
	"fmt"
	"log"

	"github.com/natserract/toktik/embedding"
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
) (interface{}, error) {
	results, err := c.inMemoryRepository.GetAllUserInterestsEmbedding()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if results != nil {
		embed := embedding.NewVectorEmbedding()

		queryVector, err := embed.CreateVector("OBama", ctx)
		if err != nil {
			return nil, err
		}

		bestText, maxSimilarity, err := util.FindMostSimilar(queryVector, results, "tags")
		if err != nil {
			return nil, err
		}
		fmt.Println("Results", bestText, maxSimilarity)
	}

	return nil, nil
}
