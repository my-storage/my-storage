package service

import (
	"time"

	"github.com/my-storage/ms-profile/src/presentation"
)

type IService interface {
	Start()
	Stop()
}

type Service struct {
	IService
	StartAt time.Time

	Presentation presentation.Presentation
}

func New() Service {
	return Service{
		StartAt:      time.Now(),
		Presentation: presentation.New(),
	}
}

func (srv *Service) Start() {
	srv.Presentation.HttpServer.Start()
}

func (srv *Service) Stop() {
	srv.Presentation.Destroy()
}
