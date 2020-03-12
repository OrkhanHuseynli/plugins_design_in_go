package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	serviceName string
	plugins []Plugin
}

func (a *App) Register(plugin Plugin)  {
	a.plugins = append(a.plugins, plugin)
}

func New(serviceName string)  *App{
	return &App{serviceName:serviceName}
}

func (a *App) Start(ctx context.Context, ctxCancel context.CancelFunc)  {
	defer ctxCancel()
	chErrors := make(chan error, 1)
	for _, plugin := range a.plugins {
		err := plugin.Initialize(ctx)
		if err != nil {
			log.Println("Failed to initialize %s plugin", plugin.Name())
			chErrors <- err
			break
		}
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("One of the plugins has failed")
	case e := <-chErrors:
		log.Printf("The application is aborting : %v", e)
	case <-signals:
		log.Println("A signal was triggered")
	}

	log.Println("Terminating...")

}