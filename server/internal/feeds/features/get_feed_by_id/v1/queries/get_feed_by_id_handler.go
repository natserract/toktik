package queries

import (
	"context"
	"fmt"

	"github.com/natserract/toktik/internal/feeds/features/get_feed_by_id/v1/dtos"
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
	// var result *dtos.GetFeedByIdResponseDto

	// Cached
	iterator := c.Store.Feeds.Cache.Iterator()
	for iterator.SetNext() {
		current, err := iterator.Value()

		if err != nil {
			fmt.Println(err)
		}

		var retrievedData *[]scraper.VideoInfo
		c.Store.Feeds.GetFeeds(current.Key(), &retrievedData)

		fmt.Println(current.Key())
		fmt.Println(retrievedData)
	}

	return &dtos.GetFeedByIdResponseDto{}, nil
}
