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

		if err != nil {
			switch {
			case errors.Is(err, model.ErrBadRequest), errors.Is(err, echo.ErrBadRequest):
				return c.JSON(http.StatusBadRequest, echo.Map{"reason": err.Error()})
			case errors.Is(err, model.ErrUserNotExists):
				return c.JSON(http.StatusNotFound, echo.Map{"reason": err.Error()})
			case errors.Is(err, model.ErrForbidden), errors.Is(err, echo.ErrForbidden):
				return c.JSON(http.StatusForbidden, echo.Map{"reason": err.Error()})
			case errors.Is(err, model.ErrNotFound), errors.Is(err, echo.ErrNotFound):
				return c.JSON(http.StatusNotFound, echo.Map{"reason": err.Error()})
			default:
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"reason": "internal server error: " + err.Error(),
				})
			}
		}

		return nil
	}
}
