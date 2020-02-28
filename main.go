package main

import (
	//"github.com/micro/go-micro/v2/util/log"
	"os"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/logger/zerolog/v2"
	"gitlab.com/krayohu/logtest/handler"
	"gitlab.com/krayohu/logtest/subscriber"

	logtest "gitlab.com/krayohu/logtest/proto/logtest"
)

func main() {
	// Setup log
	log.DefaultLogger = zerolog.NewLogger(
		log.WithOutput(os.Stdout),
		log.WithLevel(log.DebugLevel),
		zerolog.ReportCaller(),
		//zerolog.ReportCallerWithCount(4),
		zerolog.WithDevelopmentMode(),
		//zerolog.WithProductionMode(),
	)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.logtest"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	log.Info("Info")
	log.Debug("Debug")
	log.Error("Error")
	log.Log(log.InfoLevel, "Log Info")

	// Register Handler
	logtest.RegisterLogtestHandler(service.Server(), new(handler.Logtest))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("ps.clientele.svc.logtest", service.Server(), new(subscriber.Logtest))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
