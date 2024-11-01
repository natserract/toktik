package endpoints

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
	"github.com/natserract/toktik/internal/feeds/features/search_feeds/v1/dtos"
	"github.com/natserract/toktik/internal/feeds/features/search_feeds/v1/queries"
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

		query := queries.SearchFeeds{
			Keywords: request.Keywords,
		}
		if err := query.Validate(); err != nil {
			return c.String(http.StatusBadRequest, "query validation failed")
		}

		var result *dtos.SearchFeedsResponseDTO

		// Cached
		err := ep.Store.Feeds.GetFeeds(request.Keywords, &result)
		if err != nil {
			q := queries.NewSearchFeedsHandler()
			queryResult, err := q.Handle(ctx, &query)
			if err != nil {
				return c.String(http.StatusBadRequest, "error in sending SearchFeeds")
			}

			result = queryResult

			// Store the result
			if err := ep.Store.Feeds.SetFeeds(request.Keywords, queryResult); err != nil {
				return c.String(http.StatusBadRequest, "can't set feeds cache")
			}
		}

		return c.JSON(http.StatusOK, result)
	}
}
