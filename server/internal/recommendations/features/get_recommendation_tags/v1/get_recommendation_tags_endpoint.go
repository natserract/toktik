package v1

import (
	"log"
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
		queryResult, err := handler.Handle(ctx, nil)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "error in sending GetRecommendationTags")
		}
		return c.JSON(http.StatusOK, queryResult)
	}
}
