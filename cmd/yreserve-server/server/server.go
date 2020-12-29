package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/thetogi/YReserve2/api/middleware"

	"github.com/thetogi/YReserve2/bi"
	"github.com/thetogi/YReserve2/eventdispatcher"

	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
	"github.com/thetogi/YReserve2/api"
	"github.com/thetogi/YReserve2/app"
	"github.com/thetogi/YReserve2/bus"
	"github.com/thetogi/YReserve2/logger"
	"github.com/thetogi/YReserve2/metrics"
	"github.com/thetogi/YReserve2/model"
	"github.com/thetogi/YReserve2/repository"
	"github.com/thetogi/YReserve2/setting"
)

type Server struct {
	API        *api.API
	App        *app.App
	Repository repository.Repository
	Router     *mux.Router
	Config     *model.Config
	Metrics    metrics.Metrics
	Log        logger.Logger
	httpServer *http.Server
}

func NewServer(settingData *setting.Setting) *Server {
	config := setting.GetConfig()
	setting.WatcherConfig()
	logger := logger.NewLogger(config)
	setting.AddConfigChangeListener(logger)
	bus := bus.NewBus(logger)
	metrics := metrics.NewMetrics()
	router := mux.NewRouter()
	repository := repository.NewPersistentCacheRepository(logger, config, metrics)
	eventDispatcher := eventdispatcher.NewEventDispatcher(logger, bus, 10, 2)
	biEventHandler := bi.NewBiEventHandler(eventDispatcher)

	appOption := &app.AppOption{
		Config:         config,
		Setting:        settingData,
		Log:            logger,
		Metrics:        metrics,
		Repository:     repository,
		BiEventHandler: biEventHandler,
		Bus:            bus,
	}

	api := api.NewAPI(router, appOption, config, metrics, logger)
	server := &Server{
		API:        api,
		Log:        logger,
		Metrics:    metrics,
		Config:     config,
		Router:     router,
		Repository: repository,
	}

	return server
}

func (s *Server) StartServer() {
	n := negroni.New()
	n.UseFunc(middleware.NewLoggerMiddleware(s.Log).GetMiddlewareHandler())
	if s.Config.ZipkinConfig.IsEnable {
		n.UseFunc(middleware.NewZipkinMiddleware(s.Log, "project", s.Config.ZipkinConfig).GetMiddlewareHandler())
	}

	n.UseHandler(s.Router)

	listenAddr := (":" + s.Config.ServerConfig.Port)
	s.Log.Debug("Staring server", logger.String("address", listenAddr))
	s.httpServer = &http.Server{
		Handler:      n,
		Addr:         listenAddr,
		ReadTimeout:  s.Config.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: s.Config.ServerConfig.WriteTimeout * time.Second,
	}

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			s.Log.Error("Error starting server ", logger.Error(err))
			return
		}
	}()
}

func (s *Server) StopServer() {
	s.Repository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)

	setting.DeleteConfigChangeListener(s.Log)
	os.Exit(0)
}
