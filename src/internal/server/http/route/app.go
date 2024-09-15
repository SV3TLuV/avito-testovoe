package route

import (
	"github.com/labstack/echo/v4"
	def "tender_api/src/internal/server/http"
)

func InitAppRoutes(group *echo.Group, controller def.AppController) {
	group.GET("/ping", controller.Ping)
}
