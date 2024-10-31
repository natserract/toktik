package params

import (
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/shared/store"
)

type FeedsRouteParams struct {
	FeedsGroup *echo.Group `name:"feed-echo-group"`
	Store      *store.Store
}
