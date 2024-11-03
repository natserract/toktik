package v1

import (
	"context"
	"log"
	"strings"

	"github.com/natserract/toktik/embedding"
	"github.com/natserract/toktik/internal/recommendations/features/get_recommendation/v1/dtos"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
	"github.com/natserract/toktik/internal/user_interests_embedding/shared/util"
	"github.com/natserract/toktik/pkg/config"
	"github.com/natserract/toktik/pkg/scraper"
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
	// If current store hasn't data
	// Set based on current trending topic (rapid.api)
	userInterestsCache := c.inMemoryRepository.Store.UserInterests.Cache
	userInterestsEmbeddingCache := c.inMemoryRepository.Store.UserInterestsEmbedding.Cache
	if userInterestsCache.Len() == 0 || userInterestsEmbeddingCache.Len() == 0 {
		cfg := config.GetConfig()
		s := scraper.NewScraper(cfg.RapidApiKey, cfg.RapiApiHost)

		trendings, err := s.Trendings(scraper.TrendingsParams{
			Count:  "10",
			Region: "us",
		})
		if err != nil {
			return nil, err
		}

		var tags []string
		var titles []string
		for _, trending := range trendings.Data {
			if trending.Title != "" {
				textSplitted := util.TextSplitter(trending.Title)
				if len(textSplitted.Tags) != 0 && len(textSplitted.Titles) != 0 {
					tags = append(tags, strings.Join(textSplitted.Tags, " "))
					titles = append(titles, strings.Join(textSplitted.Titles, " "))
				}
			}
		}
		// Sometimes trendings value is empty
		if len(tags) == 0 || len(titles) == 0 {
			tags = append(tags, []string{
				"#Health #Wellness #Lifestyle",
				"#Economy #Finance #GlobalTrends",
				" #Education #Learning #Reform",
			}...)
			titles = append(titles, []string{
				"Health and Wellness",
				"Economic Trends",
				"Education Reform",
			}...)
		}

		return &dtos.GetRecommendationsResponseDTO{
			Data: dtos.GetRecommendationsData{
				Tags:   helper.SafeSubslice(tags, 3),
				Titles: helper.SafeSubslice(titles, 3),
			},
		}, nil
	}

	// Give user recommendations based on provided user profile
	// I wondering it would have several ways:
	// - Based on provided user profile
	// - Based on user keywords
	// etc...
	//
	// I just find the easest solution...
	//
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
