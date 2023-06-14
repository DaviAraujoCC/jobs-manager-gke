package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hurbcom/jobs-manager-gke/config"
	"github.com/hurbcom/jobs-manager-gke/server"
	"github.com/hurbcom/jobs-manager-gke/utils"

	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		logrus.Info(err)
	}

	utils.SetLoggerLevel(viper.GetString("LOG_LEVEL"))

	s := server.NewServer(cfg)
	s.Listen()
	logrus.Println("Listening on port", viper.GetString("HTTP_PORT"))

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGHUP)

	<-shutdown
	s.Shutdown()
}
