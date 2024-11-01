package queries

import (
	"context"
	"fmt"

	"github.com/natserract/toktik/internal/feeds/features/get_feed_by_id/v1/dtos"
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
	currValue, err := iterator.Value()
	if err == nil {
		fmt.Print(currValue)
	}

	return &dtos.GetFeedByIdResponseDto{}, nil
}
