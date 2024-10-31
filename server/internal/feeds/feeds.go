package feeds

import (
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
	getFeedById "github.com/natserract/toktik/internal/feeds/features/get_feed_by_id/v1/endpoints"
	searchFeeds "github.com/natserract/toktik/internal/feeds/features/search_feeds/v1/endpoints"
	"github.com/natserract/toktik/pkg/http/contracts"
	"github.com/natserract/toktik/shared/store"
)

type Feeds struct {
	Store *store.Store
}

func NewFeeds(store *store.Store) *Feeds {
	return &Feeds{Store: store}
}

func (s *Feeds) Mount(e contracts.EchoHttpServer) {
	e.RouteBuilder().RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
		group := v1.Group("/feeds")

		// Register feed endpoints
		params := params.FeedsRouteParams{
			FeedsGroup: group,
			Store:      s.Store,
		}
		searchFeeds.NewSearchFeedsEndpoint(params).MapEndpoint()
		getFeedById.NewGetFeedByIdEndpoint(params).MapEndpoint()
	})
}
