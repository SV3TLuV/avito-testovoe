package http

import "github.com/labstack/echo/v4"

type AppController interface {
	Ping(echo.Context) error
}

type TenderController interface {
	GetList(echo.Context) error
	GetMy(echo.Context) error
	GetStatus(echo.Context) error
	Create(echo.Context) error
	UpdateStatus(echo.Context) error
	Rollback(echo.Context) error
	Edit(echo.Context) error
}

type BidController interface {
	GetMy(echo.Context) error
	GetStatus(echo.Context) error
	GetOffers(echo.Context) error
	GetReviews(echo.Context) error
	Create(echo.Context) error
	UpdateStatus(echo.Context) error
	SubmitDecision(echo.Context) error
	Feedback(echo.Context) error
	Rollback(echo.Context) error
	Edit(echo.Context) error
}
