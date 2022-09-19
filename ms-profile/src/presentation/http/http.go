package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	v1 "github.com/my-storage/ms-profile/src/presentation/http/controllers/v1"
)

func New() *HttpServer {
	httpServer := createServer()
	router := httpServer.Router.Group("/api")

	v1.Register(router)

	return httpServer
}

type HttpServer struct {
	Server     *http.Server
	Router     *gin.Engine
	StartAt    time.Time
	SignalChan chan os.Signal
}

func createServer() *HttpServer {
	router := gin.Default()

	// gin.SetMode(gin.ReleaseMode)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	return &HttpServer{
		Server: srv,
		Router: router,
	}
}

func (httpServer *HttpServer) Start() {
	go func() {
		if err := httpServer.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	httpServer.SignalChan = signalChan

	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Turning off Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	} else {
		log.Println("Server Shutdown gracefully")
	}
}

func (httpServer *HttpServer) Stop() {
	httpServer.SignalChan <- syscall.SIGTERM
}

// TODO Metodo de Setup de Middlewares
