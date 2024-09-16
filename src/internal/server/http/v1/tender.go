package v1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"tender_api/src/internal/converter"
	"tender_api/src/internal/model"
	"tender_api/src/internal/model/enum"
	def "tender_api/src/internal/server/http"
	"tender_api/src/internal/server/http/v1/requests/tender"
	"tender_api/src/internal/service"
)

var _ def.TenderController = (*tenderController)(nil)

type tenderController struct {
	tenderService service.TenderService
}

func NewTenderController(tenderService service.TenderService) *tenderController {
	return &tenderController{
		tenderService: tenderService,
	}
}

func (controller *tenderController) GetList(ctx echo.Context) error {
	var request tender.GetListRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	request.SetDefaults()

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	tenders, err := controller.tenderService.GetList(
		context,
		uint(request.Limit),
		uint(request.Offset),
		request.ServiceType)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, converter.ToTenderViewsFromTender(tenders))
}

func (controller *tenderController) GetMy(ctx echo.Context) error {
	var request tender.GetMyRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	request.SetDefaults()

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	tenders, err := controller.tenderService.GetMy(
		context,
		uint(request.Limit),
		uint(request.Offset),
		request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, converter.ToTenderViewsFromTender(tenders))
}

func (controller *tenderController) GetStatus(ctx echo.Context) error {
	var request tender.GetStatusRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	tenderID, err := uuid.Parse(ctx.Param("tenderId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.TenderID = tenderID

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	status, err := controller.tenderService.GetStatus(context, request.TenderID, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, status)
}

func (controller *tenderController) Create(ctx echo.Context) error {
	var request tender.CreateRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	created, err := controller.tenderService.Create(context, model.Tender{
		Name:           request.Name,
		Description:    request.Description,
		ServiceType:    request.ServiceType,
		OrganizationID: request.OrganizationID,
	}, request.CreatorUsername)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, converter.ToTenderViewFromTender(*created))
}

func (controller *tenderController) UpdateStatus(ctx echo.Context) error {
	var request tender.UpdateStatusRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	request.Username = ctx.QueryParam("username")
	request.Status = enum.TenderStatus(ctx.QueryParam("status"))

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	updated, err := controller.tenderService.UpdateStatus(context, request.TenderID, request.Status, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, converter.ToTenderViewFromTender(*updated))
}

func (controller *tenderController) Rollback(ctx echo.Context) error {
	var request tender.RollbackRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	tenderID, err := uuid.Parse(ctx.Param("tenderId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.TenderID = tenderID

	version, err := strconv.ParseUint(ctx.Param("version"), 10, 64)
	if err != nil {
		return model.ErrBadRequest
	}
	request.Version = version
	request.Username = ctx.QueryParam("username")

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	updated, err := controller.tenderService.Rollback(context, request.TenderID, request.Version, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, converter.ToTenderViewFromTender(*updated))
}

func (controller *tenderController) Edit(ctx echo.Context) error {
	var request tender.EditRequest
	if err := ctx.Bind(&request); err != nil {
		return model.ErrBadRequest
	}

	tenderID, err := uuid.Parse(ctx.Param("tenderId"))
	if err != nil {
		return model.ErrBadRequest
	}
	request.TenderID = tenderID
	request.Username = ctx.QueryParam("username")

	if err := ctx.Validate(&request); err != nil {
		return model.ErrBadRequest
	}

	context := ctx.Request().Context()
	updated, err := controller.tenderService.Edit(context, model.Tender{
		ID:          request.TenderID,
		Name:        request.Name,
		Description: request.Description,
		ServiceType: request.ServiceType,
	}, request.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, converter.ToTenderViewFromTender(*updated))
}
