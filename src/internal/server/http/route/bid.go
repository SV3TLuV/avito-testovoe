package route

import (
	"github.com/labstack/echo/v4"
	def "tender_api/src/internal/server/http"
)

func InitBidRoutes(group *echo.Group, controller def.BidController) {
	g := group.Group("/bids")

	g.GET("/my", controller.GetMy)
	g.GET("/:bidId/status", controller.GetStatus)
	g.GET("/:tenderId/list", controller.GetOffers)
	g.GET("/:tenderId/reviews", controller.GetReviews)
	g.POST("/new", controller.Create)
	g.PUT("/:bidId/status", controller.UpdateStatus)
	g.PUT("/:bidId/submit_decision", controller.SubmitDecision)
	g.PUT("/:bidId/feedback", controller.Feedback)
	g.PUT("/:bidId/rollback/:version", controller.Rollback)
	g.PATCH("/:bidId/edit", controller.Edit)
}
