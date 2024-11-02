package v1

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
	"github.com/natserract/toktik/internal/feeds/data/repositories"
	"github.com/natserract/toktik/internal/feeds/features/get_feed_by_id/v1/dtos"
	userInterestsRepo "github.com/natserract/toktik/internal/user_interests/data/repositories"
	createUserInterestV1 "github.com/natserract/toktik/internal/user_interests/features/create_user_interest/v1"
	"github.com/natserract/toktik/shared/store"
)

type getFeedByIdEndpoint struct {
	params.FeedsRouteParams
}

func NewGetFeedByIdEndpoint(
	params params.FeedsRouteParams,
) *getFeedByIdEndpoint {
	return &getFeedByIdEndpoint{
		FeedsRouteParams: params,
	}
}

func (ep *getFeedByIdEndpoint) MapEndpoint() {
	ep.FeedsGroup.GET("/:id", ep.handler())
}

func (ep *getFeedByIdEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		request := &dtos.GetFeedByIdRequestDto{}

		if err := c.Bind(request); err != nil {
			return c.String(http.StatusBadRequest, "error in the binding request")
		}

		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			return c.String(http.StatusBadRequest, "error in the binding request")
		}

		query := GetFeedById{
			Id: request.Id,
		}
		if err := query.Validate(); err != nil {
			return c.String(http.StatusBadRequest, "query validation failed")
		}

		feedsRepo := repositories.NewFeedsRepository(ep.Store)
		getFeedByIdHandler := NewGetFeedByIdHandler(feedsRepo)
		queryResult, err := getFeedByIdHandler.Handle(ctx, &query)
		if err != nil {
			return c.String(http.StatusBadRequest, "error in sending SearchFeeds")
		}

		// Collect user watched to user interests
		actor := ep.Store.UserInterests.Key(store.WatchUserInterestsActor, request.Id)
		pageContent := getFeedByIdHandler.ToPageContent(queryResult)
		userInterestQuery := createUserInterestV1.CreateUserInterest{
			Actor:        actor,
			PageContents: []string{pageContent},
		}
		if err = userInterestQuery.Validate(); err != nil {
			return c.String(http.StatusBadRequest, "query validation failed")
		}

		userInterestsRepo := userInterestsRepo.NewUserInterestsRepository(ep.Store)
		userInterestsHandler := createUserInterestV1.NewCreateUserInterestHandler(userInterestsRepo)
		if userInterestsHandler.Handle(ctx, userInterestQuery); err != nil {
			fmt.Println("error in collecting user interests", err)
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
