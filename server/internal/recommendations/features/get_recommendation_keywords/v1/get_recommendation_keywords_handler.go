package v1

import (
	"context"
	"log"

	"github.com/natserract/toktik/config"
	"github.com/natserract/toktik/embedding"
	"github.com/natserract/toktik/internal/recommendations/features/get_recommendation_keywords/v1/dtos"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
	"github.com/natserract/toktik/internal/user_interests_embedding/shared/util"
	"github.com/natserract/toktik/pkg/scraper"
	helper "github.com/natserract/toktik/shared/util"
)

type GetRecommendationKeywordsHandler struct {
	inMemoryRepository repositories.UserInterestsEmbeddingRepository
}

func NewGetRecommendationKeywordsHandler(r repositories.UserInterestsEmbeddingRepository) *GetRecommendationKeywordsHandler {
	return &GetRecommendationKeywordsHandler{
		inMemoryRepository: r,
	}
}

func (c *GetRecommendationKeywordsHandler) Handle(
	ctx context.Context,
	query interface{},
) (*dtos.GetRecommendationsResponseDTO, error) {
	// If current store hasn't data
	// Set based on current trending topic (rapid.api)
	var trendingKeywords []string

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

		for _, trending := range trendings.Data {
			if trending.Title != "" {
				textSplitted := util.TextSplitter(trending.Title)
				if len(textSplitted.Tags) != 0 && len(textSplitted.Titles) != 0 {
					trendingKeywords = append(trendingKeywords, textSplitted.Titles...)
				}
			}
		}
		// Sometimes trendings value is empty
		if len(trendings.Data) == 0 {
			trendingKeywords = append(trendingKeywords, []string{
				"AI Art Challenge: Watch Me Turn My Doodles into Masterpieces!",
				"Zero-Waste Kitchen Hacks: Save Money and the Planet!",
				"10-Minute Morning Yoga Routine for a Stress-Free Day",
				"Thrift Flip: Transforming a $5 Dress into a Fashion Statement",
				"Virtual Tour: Exploring the Hidden Gems of Kyoto",
				"Meet My Rescue Dog: From Shelter to Superstar!",
				"AI vs. Human: Who Can Create the Best Song?",
				"Cooking with Grandma: Secret Family Recipes Revealed",
				"Quick HIIT Workout: Get Fit in Just 15 Minutes!",
			}...)
		}

		return &dtos.GetRecommendationsResponseDTO{
			Data: dtos.GetRecommendationsData{
				Keywords: helper.SafeSubslice(trendingKeywords, 10),
			},
		}, nil
	}

	log.Println("Trending keywords: ", trendingKeywords)
	if len(trendingKeywords) < 4 {
		trendingKeywords = append(trendingKeywords, []string{
			"AI Art Challenge: Watch Me Turn My Doodles into Masterpieces!",
			"Zero-Waste Kitchen Hacks: Save Money and the Planet!",
			"10-Minute Morning Yoga Routine for a Stress-Free Day",
			"Thrift Flip: Transforming a $5 Dress into a Fashion Statement",
			"Virtual Tour: Exploring the Hidden Gems of Kyoto",
			"Meet My Rescue Dog: From Shelter to Superstar!",
			"AI vs. Human: Who Can Create the Best Song?",
			"Cooking with Grandma: Secret Family Recipes Revealed",
			"Quick HIIT Workout: Get Fit in Just 15 Minutes!",
		}...)
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
		titles, titleSimilarityScores, err := util.FindMostSimilar(
			user_profile_vec.Values,
			results,
			0.7,
			"title",
		)
		if err != nil {
			return nil, err
		}

		log.Println("Title similarities: ", titles, titleSimilarityScores)
		if len(titles) != 0 {
			result = &dtos.GetRecommendationsResponseDTO{
				Data: dtos.GetRecommendationsData{
					Keywords: helper.SafeSubslice(titles, 10),
				},
			}
		} else {
			result = &dtos.GetRecommendationsResponseDTO{
				Data: dtos.GetRecommendationsData{
					Keywords: helper.SafeSubslice(trendingKeywords, 10),
				},
			}
		}
	}

	return result, nil
}
