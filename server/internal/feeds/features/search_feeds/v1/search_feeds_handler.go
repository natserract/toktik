package v1

import (
	"context"

	"github.com/natserract/toktik/internal/feeds/data/repositories"
	"github.com/natserract/toktik/internal/feeds/features/search_feeds/v1/dtos"
	"github.com/natserract/toktik/pkg/config"
	"github.com/natserract/toktik/pkg/scraper"
	"github.com/natserract/toktik/shared/store"
)

type SearchFeedsHandler struct {
	inMemoryRepository repositories.FeedsRepository
}

func NewSearchFeedsHandler(r repositories.FeedsRepository) *SearchFeedsHandler {
	return &SearchFeedsHandler{
		inMemoryRepository: r,
	}
}

func (c *SearchFeedsHandler) Handle(
	ctx context.Context,
	query *SearchFeeds,
) (*dtos.SearchFeedsResponseDTO, error) {
	key := c.inMemoryRepository.Instance().Key(store.SearchFeedsActor, query.Keywords, query.Count)
	results, err := c.inMemoryRepository.GetAllFeeds(key)
	if err != nil {
		cfg := config.GetConfig()
		s := scraper.NewScraper(cfg.RapidApiKey, cfg.RapiApiHost)

		feeds, err := s.SearchVideos(scraper.SearchVideosParams{
			Keywords: query.Keywords,
			Count:    query.Count,
			Region:   "us",
		})
		if err != nil {
			return nil, err
		}

		var videos *[]scraper.VideoInfo
		videos = &feeds.Data.Videos

		// Store to the cache
		if err := c.inMemoryRepository.SaveFeeds(key, videos); err != nil {
			return nil, err
		}

		return &dtos.SearchFeedsResponseDTO{Data: *videos}, nil
	}

	return &dtos.SearchFeedsResponseDTO{Data: *results}, nil
}
