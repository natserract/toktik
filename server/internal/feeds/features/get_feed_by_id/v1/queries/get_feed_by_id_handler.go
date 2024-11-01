package queries

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/natserract/toktik/internal/feeds/features/get_feed_by_id/v1/dtos"
	"github.com/natserract/toktik/pkg/config"
	"github.com/natserract/toktik/pkg/scraper"

	"github.com/natserract/toktik/shared/store"
)

type GetFeedByIdHandler struct {
	Store *store.Store
}

func NewGetFeedByIdHandler(s *store.Store) *GetFeedByIdHandler {
	return &GetFeedByIdHandler{
		Store: s,
	}
}

func (c *GetFeedByIdHandler) Handle(
	ctx context.Context,
	query *GetFeedById,
) (*dtos.GetFeedByIdResponseDto, error) {
	var result *scraper.VideoInfo

	// Cached
	iterator := c.Store.Feeds.Cache.Iterator()
	for iterator.SetNext() {
		current, err := iterator.Value()
		if err != nil {
			return nil, err
		}

		var videos *[]scraper.VideoInfo
		c.Store.Feeds.GetFeeds(current.Key(), &videos)

		// Find feed in cache
		for _, video := range *videos {
			if video.ID == query.Id || video.VideoID == query.Id {
				result = &video
				break
			}
		}
	}

	// Otherwise fetch by Id (video_id)
	// Then append to the existing feed caches
	if result == nil {
		cfg := config.GetConfig()
		s := scraper.NewScraper(cfg.RapidApiKey, cfg.RapiApiHost)

		feed, err := s.GetVideo(query.Id)
		if err != nil {
			return nil, err
		}

		result = feed.Data.VideoInfo

		// Store to the cache
		var videos []scraper.VideoInfo
		videos = append(videos, *result)

		uuid, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}

		key := c.Store.Feeds.Key(store.GetFeedByIdActor, uuid.String())
		if err := c.Store.Feeds.SetFeeds(key, videos); err != nil {
			return nil, err
		}
	}

	return &dtos.GetFeedByIdResponseDto{Data: result}, nil
}
