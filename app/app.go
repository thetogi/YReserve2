package app

import (
	"github.com/thetogi/YReserve2/bi"
	"github.com/thetogi/YReserve2/logger"
	"github.com/thetogi/YReserve2/metrics"
	"github.com/thetogi/YReserve2/model"
	"github.com/thetogi/YReserve2/repository"
	"github.com/thetogi/YReserve2/setting"
)

type App struct {
	Repository     repository.Repository
	Config         *model.Config
	Setting        *setting.Setting
	Metrics        metrics.Metrics
	Log            logger.Logger
	BiEventHandler bi.EventHandler
	RequestID      string
	UserSession    *model.UserSession
}

func NewApp(appOption *AppOption) *App {
	app := &App{
		Repository:     appOption.Repository,
		Config:         appOption.Config,
		Setting:        appOption.Setting,
		Metrics:        appOption.Metrics,
		Log:            appOption.Log,
		BiEventHandler: appOption.BiEventHandler,
	}
	return app
}
