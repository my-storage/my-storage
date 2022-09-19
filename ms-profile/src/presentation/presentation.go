package presentation

import (
	"time"

	"github.com/my-storage/ms-profile/src/presentation/http"
)

type Presentation struct {
	StartAt    time.Time
	HttpServer *http.HttpServer
}

func New() Presentation {
	httpServer := http.New()

	return Presentation{
		StartAt:    time.Now(),
		HttpServer: httpServer,
	}
}

func (p *Presentation) Destroy() {
	p.HttpServer.Stop()
}
