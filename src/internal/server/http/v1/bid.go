package v1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"tender_api/src/internal/model"
	def "tender_api/src/internal/server/http"
	"tender_api/src/internal/server/http/v1/requests/bid"
	"tender_api/src/internal/service"
)

var _ def.BidController = (*bidController)(nil)

type bidController struct {
	bidService service.BidService
}

func NewBidController(bidService service.BidService) *bidController {
	return &bidController{
		bidService: bidService,
	}
}

func (controller *bidController) GetMy(ctx echo.Context) error {
	var request bid.GetMyRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	request.SetDefaults()

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	bids, err := controller.bidService.GetMy(
		context,
		uint(request.Limit),
		uint(request.Offset),
		request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, bids)
}

func (controller *bidController) GetStatus(ctx echo.Context) error {
	var request bid.GetStatusRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	bidID, err := uuid.Parse(ctx.Param("bidId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.BidID = bidID

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	status, err := controller.bidService.GetStatus(context, request.BidID, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, status)
}

func (controller *bidController) GetOffers(ctx echo.Context) error {
	var request bid.GetListRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	tenderID, err := uuid.Parse(ctx.Param("tenderId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.TenderID = tenderID

	request.SetDefaults()

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	bids, err := controller.bidService.GetTenderList(
		context,
		request.TenderID,
		uint(request.Limit),
		uint(request.Offset),
		request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, bids)
}

func (controller *bidController) GetReviews(ctx echo.Context) error {
	var request bid.GetReviewsRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	tenderID, err := uuid.Parse(ctx.Param("tenderId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.TenderID = tenderID

	request.SetDefaults()

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	reviews, err := controller.bidService.GetTenderReviews(
		context,
		uint(request.Limit),
		uint(request.Offset),
		request.TenderID,
		request.AuthorUsername,
		request.RequesterUsername,
	)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, reviews)
}

func (controller *bidController) Create(ctx echo.Context) error {
	var request bid.CreateRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	createdBid, err := controller.bidService.Create(context, model.Bid{
		Name:        request.Name,
		Description: request.Description,
		TenderID:    request.TenderID,
		AuthorType:  request.AuthorType,
		AuthorID:    request.AuthorID,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, createdBid)
}

func (controller *bidController) UpdateStatus(ctx echo.Context) error {
	var request bid.UpdateStatusRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	bidID, err := uuid.Parse(ctx.Param("bidId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.BidID = bidID

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	updatedBid, err := controller.bidService.UpdateStatus(context, request.BidID, request.Status, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, updatedBid)
}

func (controller *bidController) SubmitDecision(ctx echo.Context) error {
	var request bid.SubmitDecisionRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	bidID, err := uuid.Parse(ctx.Param("bidId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.BidID = bidID

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	updatedBid, err := controller.bidService.SubmitDecision(context, request.BidID, request.Decision, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, updatedBid)
}

func (controller *bidController) Feedback(ctx echo.Context) error {
	var request bid.FeedbackRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	bidID, err := uuid.Parse(ctx.Param("bidId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.BidID = bidID

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	updatedBid, err := controller.bidService.Feedback(context, request.BidID, request.BidFeedback, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, updatedBid)
}

func (controller *bidController) Rollback(ctx echo.Context) error {
	var request bid.RollbackRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	bidID, err := uuid.Parse(ctx.Param("bidId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.BidID = bidID

	version, err := strconv.ParseUint(ctx.Param("version"), 10, 64)
	if err != nil {
		return model.ErrBadRequest
	}
	request.Version = version

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	rolledBackBid, err := controller.bidService.Rollback(context, request.BidID, request.Version, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, rolledBackBid)
}

func (controller *bidController) Edit(ctx echo.Context) error {
	var request bid.EditRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	bidID, err := uuid.Parse(ctx.Param("bidId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.BidID = bidID

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	editedBid, err := controller.bidService.Edit(context, model.Bid{
		ID:          request.BidID,
		Name:        request.Name,
		Description: request.Description,
	}, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, editedBid)
}
