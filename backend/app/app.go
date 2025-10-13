package app

import (
	"context"
	"github.com/kishansakhiya/wails-demo/backend/app/handlers"
)

type App struct {
	Ctx           context.Context
	UserHandler   *handlers.UserHandler
	SystemHandler *handlers.SystemHandler
}

func NewApp() *App {
	return &App{
		UserHandler:   &handlers.UserHandler{},
		SystemHandler: &handlers.SystemHandler{},
	}
}

func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}
