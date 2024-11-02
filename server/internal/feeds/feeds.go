package feeds

import (
	"encoding/gob"

	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
	getFeedByIdV1 "github.com/natserract/toktik/internal/feeds/features/get_feed_by_id/v1"
	searchFeedsV1 "github.com/natserract/toktik/internal/feeds/features/search_feeds/v1"
	streamFeedsV1 "github.com/natserract/toktik/internal/feeds/features/streams/v1"
	createUserInterest "github.com/natserract/toktik/internal/user_interests/features/create_user_interest/v1"
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
	// https://stackoverflow.com/questions/21934730/gob-type-not-registered-for-interface-mapstringinterface
	gob.Register(createUserInterest.CreateUserMetadata{})

	e.RouteBuilder().RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
		group := v1.Group("/feeds")

		params := params.FeedsRouteParams{
			FeedsGroup: group,
			Store:      s.Store,
		}

		// Register feed endpoints
		searchFeedsV1.NewSearchFeedsEndpoint(params).MapEndpoint()
		getFeedByIdV1.NewGetFeedByIdEndpoint(params).MapEndpoint()
		streamFeedsV1.NewStreamsEndpoint(params).MapEndpoint()
	})
}
