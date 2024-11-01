package queries

import (
	"context"

	"github.com/natserract/toktik/internal/feeds/features/search_feeds/v1/dtos"
	"github.com/natserract/toktik/pkg/config"
	"github.com/natserract/toktik/pkg/scraper"
	"github.com/natserract/toktik/shared/store"
)

type SearchFeedsHandler struct {
	Store *store.Store
}

func NewSearchFeedsHandler(s *store.Store) *SearchFeedsHandler {
	return &SearchFeedsHandler{
		Store: s,
	}
}

func (c *SearchFeedsHandler) Handle(
	ctx context.Context,
	query *SearchFeeds,
) (*dtos.SearchFeedsResponseDTO, error) {
	var videos *[]scraper.VideoInfo

	// Cached
	err := c.Store.Feeds.GetFeeds(query.Keywords, &videos)
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

		videos = &feeds.Data.Videos

		// Store to the cache
		if err := c.Store.Feeds.SetFeeds(query.Keywords, videos); err != nil {
			return nil, err
		}
	}

	return &dtos.SearchFeedsResponseDTO{Data: *videos}, nil
}
