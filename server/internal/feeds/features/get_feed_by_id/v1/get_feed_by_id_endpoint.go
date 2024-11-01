package v1

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
	"github.com/natserract/toktik/internal/feeds/data/repositories"
	"github.com/natserract/toktik/internal/feeds/features/get_feed_by_id/v1/dtos"
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

		repo := repositories.NewFeedsRepository(ep.Store)
		q := NewGetFeedByIdHandler(repo)
		queryResult, err := q.Handle(ctx, &query)
		if err != nil {
			return c.String(http.StatusBadRequest, "error in sending SearchFeeds")
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
