package app

import (
	"context"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"tender_api/src/internal/config"
	"tender_api/src/internal/server/http/middleware"
	"tender_api/src/internal/server/http/route"
	v1 "tender_api/src/internal/server/http/v1"
	"tender_api/src/internal/server/http/validator"
)

type App struct {
	provider   *ServiceProvider
	httpServer *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	ctx := context.Background()
	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	return a.runHttpServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHttpServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	return config.Load()
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.provider = NewServiceProvider()
	return nil
}

func (a *App) initHttpServer(_ context.Context) error {
	a.httpServer = echo.New()
	group := a.httpServer.Group("/api")

	a.httpServer.Use(middleware2.Recover())
	a.httpServer.Use(middleware2.Logger())
	a.httpServer.Use(middleware.ErrorHandlerMiddleware)

	a.httpServer.Validator = validator.NewRequestValidator()

	route.InitAppRoutes(group, v1.NewAppController())
	route.InitTenderRoutes(group, v1.NewTenderController(a.provider.TenderService()))
	route.InitBidRoutes(group, v1.NewBidController(a.provider.BidService()))

	return nil
}

func (a *App) runHttpServer() error {
	err := a.httpServer.Start(a.provider.Config().ServerAddress)
	if err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}
