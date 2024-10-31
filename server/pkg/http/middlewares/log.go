package middlewares

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func LogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			start := time.Now()

			if err := next(c); err != nil {
				c.Error(err)
			}

			res := c.Response()
			stop := time.Now()

			fields := map[string]interface{}{
				"request_id":      res.Header().Get(echo.HeaderXRequestID),
				"ip":              c.RealIP(),
				"host":            req.Host,
				"uri":             req.RequestURI,
				"method":          req.Method,
				"user_agent":      req.UserAgent(),
				"status":          res.Status,
				"roundtrip":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
				"roundtrip_human": stop.Sub(start).String(),
			}

			// We will add log Form Values and Query String if...
			if res.Status == http.StatusInternalServerError {
				if !strings.HasPrefix(req.Header.Get(echo.HeaderContentType), echo.MIMEMultipartForm) {
					qs := c.QueryString()

					forms, err := c.FormParams()
					if err != nil {
						c.Error(err)
					}

					fields["query_string"] = qs
					fields["form_values"] = forms
				}
			}

			return nil
		}
	}
}
