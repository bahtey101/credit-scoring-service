package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/bahtey101/credit-scoring-service/internal/app/scoringservice"
	"github.com/bahtey101/credit-scoring-service/internal/config"
	"github.com/bahtey101/credit-scoring-service/internal/http/handlers"
)

func Run(cfg *config.Config) error {
	_, cancel := context.WithCancel(context.Background())
	scoringService := scoringservice.NewScoringService(cfg)

	handler := handlers.NewHandler(
		cfg.GoogleAPIKey,
		scoringService,
	)

	server := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler.InitRoutes(),
		MaxHeaderBytes: 1 << 28, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		logrus.Infof("Starting listening http server at %s", cfg.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error service http server %v", err)
		}
	}()

	gracefulShotdown(server, cancel)

	return nil
}

func gracefulShotdown(s *http.Server, cancel context.CancelFunc) {
	const waitTime = 60 * time.Second

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	logrus.Infof("Received shutdown signal: %v. Initiating graceful shutdown...", sig)

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), waitTime)
	defer cancelShutdown()

	if err := s.Shutdown(ctxShutdown); err != nil {
		logrus.Errorf("error shutting down server: %v", err)
	}

	cancel()
	logrus.Info("Graceful shutdown completed.")

}
