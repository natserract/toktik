package endpoints

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/internal/feeds/contracts/params"
	"github.com/natserract/toktik/internal/feeds/features/streams/v1/dtos"
	"github.com/natserract/toktik/internal/feeds/features/streams/v1/queries"
)

type streamsEndpoint struct {
	params.FeedsRouteParams
}

func NewStreamsEndpoint(
	params params.FeedsRouteParams,
) *streamsEndpoint {
	return &streamsEndpoint{
		FeedsRouteParams: params,
	}
}

func (ep *streamsEndpoint) MapEndpoint() {
	ep.FeedsGroup.GET("/:id", ep.handler())
}

func (ep *streamsEndpoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		request := &dtos.StreamsRequestDto{}

		if err := c.Bind(request); err != nil {
			return c.String(http.StatusBadRequest, "error in the binding request")
		}

		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			return c.String(http.StatusBadRequest, "error in the binding request")
		}

		query := queries.Streams{
			Id: request.Id,
		}
		if err := query.Validate(); err != nil {
			return c.String(http.StatusBadRequest, "query validation failed")
		}

		stream := queries.NewStreamsHandler()
		streamReq, err := stream.Handle(ctx, &query)
		if err != nil {
			return c.String(http.StatusBadRequest, "error in sending SearchFeeds")
		}

		rangeHeader := c.Request().Header.Get("Range")
		if rangeHeader != "" {
			streamReq.Header.Set("Range", rangeHeader)
		}

		resp, err := http.DefaultClient.Do(streamReq)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to retrieve remote video")
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				c.Response().Header().Add(key, value)
			}
		}

		// Set status code and stream the response body to the client
		c.Response().WriteHeader(resp.StatusCode)

		// Serve the video data
		return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
	}
}
