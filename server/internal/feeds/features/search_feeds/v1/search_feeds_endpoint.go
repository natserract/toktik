package v1

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
	"github.com/natserract/toktik/internal/feeds/data/repositories"
	"github.com/natserract/toktik/internal/feeds/features/search_feeds/v1/dtos"
	userInterestsRepo "github.com/natserract/toktik/internal/user_interests/data/repositories"
	createUserInterest "github.com/natserract/toktik/internal/user_interests/features/create_user_interest/v1"
)

type searchFeedsEndpoint struct {
	params.FeedsRouteParams
}

func NewSearchFeedsEndpoint(
	params params.FeedsRouteParams,
) *searchFeedsEndpoint {
	return &searchFeedsEndpoint{
		FeedsRouteParams: params,
	}
}

func (ep *searchFeedsEndpoint) MapEndpoint() {
	ep.FeedsGroup.GET("/search", ep.handler())
}

func (ep *searchFeedsEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		request := &dtos.SearchFeedsRequestDTO{}
		if err := c.Bind(request); err != nil {
			return c.String(http.StatusBadRequest, "error in the binding request")
		}

		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			return c.String(http.StatusBadRequest, "error in the binding request")
		}

		query := SearchFeeds{
			Keywords: request.Keywords,
			Count:    request.Count,
		}
		if err := query.Validate(); err != nil {
			return c.String(http.StatusBadRequest, "query validation failed")
		}

		feedsRepo := repositories.NewFeedsRepository(ep.Store)
		searchFeedsHandler := NewSearchFeedsHandler(feedsRepo)
		queryResult, err := searchFeedsHandler.Handle(ctx, &query)
		if err != nil {
			return c.String(http.StatusBadRequest, "error in sending SearchFeeds")
		}

		// Collect user keywords to user interests
		userInterestsRepo := userInterestsRepo.NewUserInterestsRepository(ep.Store)
		userInterestsHandler := createUserInterest.NewCreateUserInterestHandler(userInterestsRepo)
		err = userInterestsHandler.Handle(ctx, &createUserInterest.CreateUserInterest{
			PageContent: query.Keywords,
			Metadata:    createUserInterest.CreateUserMetadata{},
		})
		if err != nil {
			fmt.Println("error in collecting user interests", err)
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
