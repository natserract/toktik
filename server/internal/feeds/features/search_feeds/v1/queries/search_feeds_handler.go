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
	var result *dtos.SearchFeedsResponseDTO

	// Cached
	err := c.Store.Feeds.GetFeeds(query.Keywords, &result)
	if err != nil {
		cfg := config.GetConfig()

		s := scraper.NewScraper(cfg.RapidApiKey, cfg.RapiApiHost)
		feeds, err := s.SearchVideos(scraper.SearchVideoParams{
			Keywords: query.Keywords,
			Count:    query.Count,
			Region:   "us",
		})
		if err != nil {
			return nil, err
		}

		result = &dtos.SearchFeedsResponseDTO{Data: feeds.Data.Videos}

		// Store the result
		if err := c.Store.Feeds.SetFeeds(query.Keywords, result); err != nil {
			return nil, err
		}
	}

	return result, nil
}
