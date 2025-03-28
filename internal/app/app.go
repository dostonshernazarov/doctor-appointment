package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dostonshernazarov/doctor-appointment/config"
	"github.com/dostonshernazarov/doctor-appointment/internal/repo/persistent"
	"github.com/dostonshernazarov/doctor-appointment/internal/usecase/common"
	"github.com/dostonshernazarov/doctor-appointment/pkg/httpserver"
	"github.com/dostonshernazarov/doctor-appointment/pkg/logger"
	"github.com/dostonshernazarov/doctor-appointment/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	usecaseCommon := common.NewUseCase(
		persistent.NewUser(pg),
		persistent.NewDoctor(pg),
		persistent.NewAppointment(pg),
	)

	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	v1.NewRouter(httpServer.App, cfg, l, usecaseCommon)

	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
