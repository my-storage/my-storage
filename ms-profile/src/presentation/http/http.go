package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/my-storage/ms-profile/src/app/config"
	v1 "github.com/my-storage/ms-profile/src/presentation/http/controllers/v1"
	"github.com/my-storage/ms-profile/src/shared/infra/http/gin/helpers"
)

func New() *HttpServer {
	setupConfig()

	httpServer := createServer()
	httpServer.SetupBaseMiddlewares()

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
	router := gin.New()

	config := config.GetInstance()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", config.ServerAddress, config.HttpPort),
		Handler: router,
	}

	return &HttpServer{
		Server: srv,
		Router: router,
	}
}

func setupConfig() {
	config := config.GetInstance()

	if config.ApiMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func (httpServer *HttpServer) Start() {
	go func() {
		log.Printf("Server listen at %v\n", httpServer.Server.Addr)

		if err := httpServer.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server could not start: %v", err)
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

func (httpServer *HttpServer) SetupBaseMiddlewares() {
	httpServer.Router.Use(helpers.LogFormatter())
	httpServer.Router.Use(helpers.ErrorHandler())
}
