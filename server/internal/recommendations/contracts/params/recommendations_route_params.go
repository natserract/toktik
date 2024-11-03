package params

import (
	"github.com/labstack/echo/v4"
	"github.com/natserract/toktik/shared/store"
)

type RecommendationsRouteParams struct {
	RecommendationsGroup *echo.Group `name:"recommendations-echo-group"`
	Store                *store.Store
}
