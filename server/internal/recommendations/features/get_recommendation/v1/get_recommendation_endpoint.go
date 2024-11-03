package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/recommendations/contracts/params"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
)

type getRecommendationEndpoint struct {
	params.RecommendationsRouteParams
}

func NewGetRecommendationEndpoint(
	params params.RecommendationsRouteParams,
) *getRecommendationEndpoint {
	return &getRecommendationEndpoint{
		RecommendationsRouteParams: params,
	}
}

func (ep *getRecommendationEndpoint) MapEndpoint() {
	ep.RecommendationsGroup.GET("/tags", ep.handler())
}

func (ep *getRecommendationEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		repo := repositories.NewUserInterestsEmbeddingRepository(ep.Store)
		handler := NewGetRecommendationTagsHandler(repo)
		queryResult, err := handler.Handle(ctx, nil)
		if err != nil {
			return c.String(http.StatusBadRequest, "error in sending GetRecommendations")
		}
		return c.JSON(http.StatusOK, queryResult)
	}
}
