package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/svbnbyrk/wallet/pkg/http"

	"github.com/gin-gonic/gin"
	"github.com/svbnbyrk/wallet/config"
	v1 "github.com/svbnbyrk/wallet/internal/controller/http/v1"
	"github.com/svbnbyrk/wallet/internal/usecase"
	"github.com/svbnbyrk/wallet/internal/usecase/repository"
	"github.com/svbnbyrk/wallet/pkg/db"
	"github.com/svbnbyrk/wallet/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := db.New(cfg.Postgre.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	transactionUseCase := usecase.NewTransactionUseCase(
		repository.NewTransactionRepository(pg),
		repository.NewWalletRepository(pg),
		repository.NewExchangeRepository(pg),
	)

	userUseCase := usecase.NewUserUseCase(
		repository.NewUserRepository(pg),
	)

	walletUseCase := usecase.NewWalletUseCase(
		repository.NewWalletRepository(pg),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, transactionUseCase, userUseCase, walletUseCase)
	httpServer := http.New(handler)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		}
	}
}
