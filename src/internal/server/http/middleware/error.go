package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"tender_api/src/internal/model"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		status := http.StatusInternalServerError
		if err != nil {
			switch {
			case errors.Is(err, model.ErrBadRequest), errors.Is(err, echo.ErrBadRequest):
				status = http.StatusBadRequest
			case errors.Is(err, model.ErrNotFound):
				status = http.StatusNotFound
			case errors.Is(err, model.ErrForbidden), errors.Is(err, echo.ErrForbidden):
				status = http.StatusForbidden
			case errors.Is(err, model.ErrUserNotExists), errors.Is(err, echo.ErrNotFound):
				status = http.StatusUnauthorized
			}

			return c.JSON(status, echo.Map{"reason": err.Error()})
		}

		return nil
	}
}
