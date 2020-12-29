package app

import (
	"fmt"
	"time"

	"github.com/thetogi/YReserve2/repository"
	"github.com/thetogi/YReserve2/repository/docker"

	"github.com/thetogi/YReserve2/bus"

	"github.com/thetogi/YReserve2/bi"

	"github.com/thetogi/YReserve2/logger"
	"github.com/thetogi/YReserve2/metrics"
	"github.com/thetogi/YReserve2/setting"
)

type AppTest struct {
	App         *App
	MySQLDocker *docker.MysqlDocker
}

type AppTestOption struct {
	AppOption   *AppOption
	MySQLDocker *docker.MysqlDocker
}

var appTestOption *AppTestOption

func SetupAppTestOption() *AppTestOption {
	fmt.Println("Starting setup app test option")
	now := time.Now()
	startTime := fmt.Sprintf("%d", now.Unix())
	settingData := setting.NewSetting("1", "1", "NA", "master", startTime)
	config := setting.GetConfigFromFile("test")
	logger := logger.NewTestLogger(config)
	metrics := metrics.NewTestMetrics()
	mysqlDocker := &docker.MysqlDocker{
		ContainerName: "mysql-app-container",
	}
	mysqlDocker.StartMysqlDocker()
	// Wait for docker mysql server to start
	time.Sleep(10 * time.Second)
	repository := repository.NewPersistentRepository(logger, config, metrics)
	appOption := &AppOption{
		Config:         config,
		Setting:        settingData,
		Log:            logger,
		Metrics:        metrics,
		Repository:     repository,
		BiEventHandler: bi.NewBiTestEventHandler(),
		Bus:            bus.NewTestBus(),
	}

	appTestOption = &AppTestOption{
		AppOption:   appOption,
		MySQLDocker: mysqlDocker,
	}
	fmt.Println("App test option created")
	return appTestOption
}

func (ato *AppTestOption) Cleanup() {
	fmt.Println("App test option cleanup called")
	ato.MySQLDocker.Stop()
	time.Sleep(10 * time.Second)
}

func SetupAppTest() *AppTest {
	app := NewApp(appTestOption.AppOption)
	return &AppTest{
		App: app,
	}
}

func (at *AppTest) CleanupAppTest() {
}
