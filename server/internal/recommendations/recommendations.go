package recommendations

import (
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/recommendations/contracts/params"
	getRecommendationV1 "github.com/natserract/toktik/internal/recommendations/features/get_recommendation/v1"
	"github.com/natserract/toktik/pkg/http/contracts"
	"github.com/natserract/toktik/shared/store"
)

type Recommendations struct {
	Store *store.Store
}

func NewRecommendations(s *store.Store) *Recommendations {
	return &Recommendations{Store: s}
}

func (s *Recommendations) Mount(e contracts.EchoHttpServer) {
	e.RouteBuilder().RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
		group := v1.Group("/recommendations")

		params := params.RecommendationsRouteParams{
			RecommendationsGroup: group,
			Store:                s.Store,
		}

		// Register endpoints
		getRecommendationV1.NewGetRecommendationEndpoint(params).MapEndpoint()
	})
}
