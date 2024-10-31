package params

import "github.com/labstack/echo/v4"

type FeedsRouteParams struct {
	FeedsGroup *echo.Group `name:"feed-echo-group"`
}
