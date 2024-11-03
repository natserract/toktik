package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/recommendations/contracts/params"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
)

type getRecommendationTagsEndpoint struct {
	params.RecommendationsRouteParams
}

func NewGetRecommendationTagsEndpoint(
	params params.RecommendationsRouteParams,
) *getRecommendationTagsEndpoint {
	return &getRecommendationTagsEndpoint{
		RecommendationsRouteParams: params,
	}
}

func (ep *getRecommendationTagsEndpoint) MapEndpoint() {
	ep.RecommendationsGroup.GET("/tags", ep.handler())
}

func (ep *getRecommendationTagsEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		repo := repositories.NewUserInterestsEmbeddingRepository(ep.Store)
		handler := NewGetRecommendationTagsHandler(repo)
		_, err := handler.Handle(ctx, nil)
		if err != nil {
			return c.String(http.StatusBadRequest, "error in sending GetRecommendations")
		}

		return c.JSON(http.StatusOK, "Okay")
	}
}
