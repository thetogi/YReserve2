package app

import (
	"github.com/thetogi/YReserve2/bi"
	"github.com/thetogi/YReserve2/bus"
	"github.com/thetogi/YReserve2/logger"
	"github.com/thetogi/YReserve2/metrics"
	"github.com/thetogi/YReserve2/model"
	"github.com/thetogi/YReserve2/repository"
	"github.com/thetogi/YReserve2/setting"
)

type AppOption struct {
	Repository     repository.Repository
	Config         *model.Config
	Setting        *setting.Setting
	Metrics        metrics.Metrics
	Log            logger.Logger
	BiEventHandler bi.EventHandler
	Bus            bus.Bus
}
