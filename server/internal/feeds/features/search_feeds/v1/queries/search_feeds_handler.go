package queries

import (
	"context"

	"github.com/natserract/toktik/internal/feeds/features/search_feeds/v1/dtos"
	"github.com/natserract/toktik/pkg/config"
	"github.com/natserract/toktik/pkg/scraper"
)

type SearchFeedsHandler struct{}

func NewSearchFeedsHandler() *SearchFeedsHandler {
	return &SearchFeedsHandler{}
}

func (c *SearchFeedsHandler) Handle(
	ctx context.Context,
	query *SearchFeeds,
) (*dtos.SearchFeedsResponseDTO, error) {
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

	return &dtos.SearchFeedsResponseDTO{Data: feeds.Data.Videos}, nil
}
