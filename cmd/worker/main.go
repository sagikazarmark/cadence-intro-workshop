package main

import (
	"math/rand"
	"time"

	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/worker"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	address        = "127.0.0.1:7933"
	domain         = "workshop"
	taskListName   = "workshop"
	clientName     = "workshop"
	cadenceService = "cadence-frontend"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.InfoLevel)

	logger, err := config.Build()
	if err != nil {
		panic("failed to setup logger")
	}

	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(clientName))
	if err != nil {
		panic("failed to setup tchannel")
	}

	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: clientName,
		Outbounds: yarpc.Outbounds{
			cadenceService: {Unary: ch.NewSingleOutbound(address)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("failed to start dispatcher")
	}

	// Create RPC service
	service := workflowserviceclient.New(dispatcher.ClientConfig(cadenceService))

	workerOptions := worker.Options{
		Logger: logger,
	}

	// Create worker
	worker := worker.New(service, domain, taskListName, workerOptions)

	// Register workflows and activities
	register(worker)

	// Run worker
	err = worker.Run()
	if err != nil {
		panic("failed to start worker")
	}
}
