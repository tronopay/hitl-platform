package service

import (
	"tronopay/config"
	"tronopay/internal/repo"

	"github.com/sirupsen/logrus"
)

type (
	ServicesDependencies struct {
		App   *config.App
		Log   *logrus.Logger
		Repos *repo.Repositories
		// other deps & components
	}

	Services struct {
		App           *config.App
		Log           *logrus.Logger
		TaskService   *TaskService
		WalletService *WalletService
	}
)

func NewServices(deps *ServicesDependencies) *Services {
	return &Services{
		App:           deps.App,
		Log:           deps.Log,
		WalletService: NewWalletService(),
		TaskService:   NewTaskService(deps),
	}
}
