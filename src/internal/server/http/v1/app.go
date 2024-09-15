package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	def "tender_api/src/internal/server/http"
)

var _ def.AppController = (*appController)(nil)

type appController struct{}

func NewAppController() *appController {
	return &appController{}
}

func (controller *appController) Ping(ctx echo.Context) error {
	if err := ctx.JSON(http.StatusOK, "ok"); err != nil {
		return echo.ErrInternalServerError
	}
	return nil
}
