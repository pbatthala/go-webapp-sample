package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/logger"
)

// APIError represents
type APIError struct {
	Code    int
	Message string
}

// JSONErrorHandler is cumstomize error handler
func JSONErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}
	if c.Echo().Debug {
		msg = err.Error()
	}

	var apierr APIError
	apierr.Code = code
	apierr.Message = msg

	if !c.Response().Committed {
		if err := c.JSON(code, apierr); err != nil {
			logger.GetZapLogger().Errorf(err.Error())
		}
	}
	logger.GetZapLogger().Debugf(err.Error())
}
