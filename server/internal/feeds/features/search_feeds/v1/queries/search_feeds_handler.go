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

	key := c.Store.Feeds.Key(store.SearchFeedsActor, query.Keywords, query.Count)
	err := c.Store.Feeds.GetFeeds(key, &videos)
	if err != nil {
		cfg := config.GetConfig()
		s := scraper.NewScraper(cfg.RapidApiKey, cfg.RapiApiHost)

		feeds, err := s.SearchVideos(scraper.SearchVideosParams{
			Keywords: key,
			Count:    query.Count,
			Region:   "us",
		})
		if err != nil {
			return nil, err
		}

		videos = &feeds.Data.Videos

		// Store to the cache
		if err := c.Store.Feeds.SetFeeds(key, videos); err != nil {
			return nil, err
		}
	}

	return &dtos.SearchFeedsResponseDTO{Data: *videos}, nil
}
