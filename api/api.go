package api

import (
	"github.com/gorilla/mux"
	"github.com/thetogi/YReserve2/logger"

	"github.com/thetogi/YReserve2/app"
	"github.com/thetogi/YReserve2/metrics"
	"github.com/thetogi/YReserve2/model"
)

type API struct {
	MainRouter *mux.Router
	AppOption  *app.AppOption
	Config     *model.Config
	Metrics    metrics.Metrics
	Log        logger.Logger
	Router     *Router
}

func NewAPI(router *mux.Router, appOption *app.AppOption, config *model.Config, metrics metrics.Metrics, logger logger.Logger) *API {
	api := &API{
		MainRouter: router,
		AppOption:  appOption,
		Config:     config,
		Metrics:    metrics,
		Log:        logger,
		Router:     &Router{},
	}
	api.setupRoutes()
	return api
}
