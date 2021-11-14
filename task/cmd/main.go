package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"github.com/p12s/uber-popug/task/pkg/broker"
	"github.com/p12s/uber-popug/task/pkg/handler"
	"github.com/p12s/uber-popug/task/pkg/repository"
	"github.com/p12s/uber-popug/task/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s\n", err.Error())
	}

	db, err := repository.NewSqlite3DB(repository.Config{
		Driver: os.Getenv("DB_DRIVER"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s\n", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	kafka, err := broker.NewKafka()
	if err != nil {
		logrus.Fatalf("kafka error: %s\n", err.Error())
	}
	go kafka.Subscribe(repos)
	handlers := handler.NewHandler(services, kafka)

	srv := new(Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error while running http server: %s\n", err.Error())
		}
	}()
	logrus.Print("😀 task app started with port: ", os.Getenv("PORT"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("task app shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close: %s", err.Error())
	}
}

// Server - http server
type Server struct {
	httpServer *http.Server
}

// Run - start
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown - grace-full
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
