package queries

import (
	"context"
)

type SearchFeedsHandler struct{}

func NewSearchFeedsHandler() *SearchFeedsHandler {
	return &SearchFeedsHandler{}
}

func (c *SearchFeedsHandler) Handle(
	ctx context.Context,
	query *SearchFeeds,
) (interface {}, error) {
	return nil, nil
}
