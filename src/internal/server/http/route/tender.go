package route

import (
	"github.com/labstack/echo/v4"
	def "tender_api/src/internal/server/http"
)

func InitTenderRoutes(group *echo.Group, controller def.TenderController) {
	g := group.Group("/tenders")

	g.GET("", controller.GetList)
	g.GET("/my", controller.GetMy)
	g.GET("/:tenderId/status", controller.GetStatus)
	g.POST("/new", controller.Create)
	g.PUT("/:tenderId/status", controller.UpdateStatus)
	g.PUT("/:tenderId/rollback/:version", controller.Rollback)
	g.PATCH("/:tenderId/edit", controller.Edit)
}
