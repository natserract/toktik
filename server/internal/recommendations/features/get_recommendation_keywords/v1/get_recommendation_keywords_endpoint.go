package v1

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/recommendations/contracts/params"
	"github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
)

type getRecommendationKeywordsEndpoint struct {
	params.RecommendationsRouteParams
}

func NewGetRecommendationKeywordsEndpoint(
	params params.RecommendationsRouteParams,
) *getRecommendationKeywordsEndpoint {
	return &getRecommendationKeywordsEndpoint{
		RecommendationsRouteParams: params,
	}
}

func (ep *getRecommendationKeywordsEndpoint) MapEndpoint() {
	ep.RecommendationsGroup.Add("GET", "/keywords", ep.handler())
}

func (ep *getRecommendationKeywordsEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		repo := repositories.NewUserInterestsEmbeddingRepository(ep.Store)
		handler := NewGetRecommendationKeywordsHandler(repo)
		queryResult, err := handler.Handle(ctx, nil)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "error in sending GetRecommendationKeywords")
		}
		return c.JSON(http.StatusOK, queryResult)
	}
}
