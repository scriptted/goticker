package main

import (
	"github.com/jinzhu/gorm"
	"github.com/scriptted/goticker/internal/config"
	"github.com/scriptted/goticker/internal/repository"
	"gitlab.com/wpetit/goweb/service"
)

func getServiceContainer(conf *config.Config, db *gorm.DB) (*service.Container, error) {
	// Initialize and configure service container
	ctn := service.NewContainer()

	// Create and expose repository service provider
	ctn.Provide(repository.ServiceName, repository.ServiceProvider(db))

	return ctn, nil
}
