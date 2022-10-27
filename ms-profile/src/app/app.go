package app

import (
	"time"

	"github.com/my-storage/ms-profile/src/presentation"
)

type App struct {
	StartAt time.Time

	Presentation presentation.Presentation
}

func New() App {
	return App{
		StartAt:      time.Now(),
		Presentation: presentation.New(),
	}
}

func (srv *App) Start() {
	srv.Presentation.HttpServer.Start()
}

func (srv *App) Stop() {
	srv.Presentation.Destroy()
}
