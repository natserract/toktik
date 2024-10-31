package feeds

import (
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
	searchFeeds "github.com/natserract/toktik/internal/feeds/features/search_feeds/v1/endpoints"
	"github.com/natserract/toktik/pkg/http/contracts"
)

type Feeds struct {
}

func NewFeeds() *Feeds {
	return &Feeds{}
}

func (s *Feeds) Mount(e contracts.EchoHttpServer) {
	e.RouteBuilder().RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
		group := v1.Group("/feed")

		// Register feed endpoints
		group.GET("/search",
			searchFeeds.NewSearchFeedsEndpoint(params.FeedsRouteParams{}).Handler(),
		)
	})
}
