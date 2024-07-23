package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/swisslockTech/go-http/http-server/commons"
	"github.com/swisslockTech/go-http/http-server/config"
	"github.com/swisslockTech/go-http/http-server/logging"
	"github.com/swisslockTech/go-http/http-server/monitoring"
	"github.com/swisslockTech/go-http/http-server/rest"
)

var (
	monitoringServer *monitoring.Server
	restServer       *rest.Server
)

func main() {
	initLogging()

	logging.SugaredLog.Infof("Start %s", commons.ServiceName)

	cfg := loadConfig()

	if cfg.EnableMonitoring {
		monitoringServer = startMonitoringServer()
	}

	restServer = startRestServer()

	logging.SugaredLog.Infof("%s up and running", commons.ServiceName)

	startSysCallChannel()

	shutdownAndWait(1)
}

func initLogging() {
	err := logging.InitGlobalLogger()
	if err != nil {
		logging.SugaredLog.Errorf("Logging setup failed: %s", err.Error())
		os.Exit(501)
	}
}

func loadConfig() *config.Config {
	logging.Log.Debug("Load configurations")
	return config.LoadConfig()
}

func startMonitoringServer() *monitoring.Server {
	logging.Log.Debug("Start monitoring")
	server := monitoring.New()
	logging.Log.Debug("Monitoring server successfully created")

	server.Start()
	logging.Log.Debug("Monitoring successfully started")

	return server
}

func startRestServer() *rest.Server {
	logging.Log.Debug("Start REST server")

	server, newErr := rest.New()
	if newErr != nil {
		logging.SugaredLog.Errorf("REST server creation failed: %s", newErr.Error())
		os.Exit(501)
	}
	logging.Log.Debug("REST server successfully created")

	startErr := server.Start()
	if startErr != nil {
		logging.SugaredLog.Errorf("REST server start failed: %s", startErr.Error())
		os.Exit(502)
	}
	logging.Log.Debug("REST server successfully started")

	rest.RegisterCustomMetrics()

	return server
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if restServer != nil {
		restServer.Shutdown(timeout)
	}

	if monitoringServer != nil {
		monitoringServer.Shutdown(timeout)
	}

	time.Sleep(time.Duration(timeout+1) * time.Second)
}
