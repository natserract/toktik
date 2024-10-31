package endpoints

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
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
		request := &dtos.GetFeedByIdRequestDto{}

		if err := c.Bind(request); err != nil {
			return c.String(http.StatusBadRequest, "error in the binding request")
		}

		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			return c.String(http.StatusBadRequest, "error in the binding request")
		}

		return c.JSON(http.StatusOK, request.Id)
	}
}
