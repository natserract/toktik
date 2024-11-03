package v1

import (
	"context"
	"log"

	"github.com/natserract/toktik/embedding"
	"github.com/natserract/toktik/internal/recommendations/features/get_recommendation/v1/dtos"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
	"github.com/natserract/toktik/internal/user_interests_embedding/shared/util"
	helper "github.com/natserract/toktik/shared/util"
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
) (*dtos.GetRecommendationsResponseDTO, error) {
	var result *dtos.GetRecommendationsResponseDTO

	results, err := c.inMemoryRepository.GetAllUserInterestsEmbedding()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	embed := embedding.NewVectorEmbedding()
	user_profile_vec, err := embed.LoadQueryVector("sample/user_profile.json")
	if err == nil {
		tags, tagsSimilarityScores, err := util.FindMostSimilar(
			user_profile_vec.Values,
			results,
			0.7,
			"tags",
		)
		if err != nil {
			return nil, err
		}
		log.Println("Found tags similarities: ", tags, tagsSimilarityScores)

		titles, titleSimilarityScores, err := util.FindMostSimilar(
			user_profile_vec.Values,
			results,
			0.7,
			"title",
		)
		if err != nil {
			return nil, err
		}

		log.Println("Found title similarities: ", titles, titleSimilarityScores)
		result = &dtos.GetRecommendationsResponseDTO{
			Data: dtos.GetRecommendationsData{
				Tags:   helper.SafeSubslice(tags, 3),
				Titles: helper.SafeSubslice(titles, 3),
			},
		}
	}

	return result, nil
}
