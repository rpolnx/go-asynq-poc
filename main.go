package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hibiken/asynq"
	"github.com/rpolnx/go-asynq-poc/internal/configs"
	"github.com/rpolnx/go-asynq-poc/internal/daemon"
	"github.com/rpolnx/go-asynq-poc/internal/repository"
	handler "github.com/rpolnx/go-asynq-poc/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	appConfig, err := configs.InitEnvVariables()
	if err != nil {
		log.Fatalln(err)
	}

	configs.InitLogger()

	cacheClient := repository.NewCacheClient(appConfig)
	defer cacheClient.Close()

	server, err := handler.InitializeServer(appConfig)

	if err != nil {
		logrus.Fatal("Error initializing server", err)
	}

	enqueuer := daemon.NewEnqueuer(appConfig, cacheClient)

	//continue create jobs
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for idx := 0; ; idx++ {
			time.Sleep(time.Second * time.Duration(5))
			t1, err := enqueuer.NewEmailDeliveryTask(idx)
			if err != nil {
				logrus.Errorln(err)
			}
			t2, err := enqueuer.NewImageResizeTask(idx, "https://cdn2.thecatapi.com/images/TBA3JzB9P.jpg")
			if err != nil {
				logrus.Errorln(err)
			}

			randomInt := r.Intn(101)
			t3, err := enqueuer.NewScheduledJob(idx, randomInt)
			if err != nil {
				logrus.Errorln(err)
			}

			err = enqueuer.EnqueueAll(t1, t2, t3)
			if err != nil {
				logrus.Errorf("could not enqueue task: %v", err)
			}
		}
	}()

	processorServer := repository.NewCacheServer(appConfig)

	processor := daemon.NewProcessor(appConfig, processorServer)

	mux := asynq.NewServeMux()
	mux.HandleFunc(daemon.TypeEmailDelivery, processor.HandleEmailDeliveryTask)
	mux.HandleFunc(daemon.TypeImageResize, processor.HandleImageResizeTask)
	mux.HandleFunc(daemon.TypeScheduledJob, processor.HandleScheduledJob)

	go func() {
		if err := processorServer.Run(mux); err != nil {
			logrus.Fatalf("could not run server: %v", err)
		}
	}()

	serverHost := fmt.Sprintf("%s:%d", configs.GlobalAppConfig.Host, configs.GlobalAppConfig.Port)
	if err = server.Engine.Run(serverHost); err != nil {
		logrus.Fatalln(err)
	}

	go func() {
		time.Sleep(time.Minute * time.Duration(5))
		processorServer.Shutdown()
	}()

}
