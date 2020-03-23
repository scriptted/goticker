package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gitlab.com/wpetit/goweb/service"
)

// ServiceName name of this service
const ServiceName service.Name = "repository"

// Repository service structure
type Repository struct {
	db *gorm.DB
}

// Balance repository exposed
func (s *Repository) Balance() *BalanceRepository {
	return NewBalanceRepository(s.db)
}

// ServiceProvider -
func ServiceProvider(db *gorm.DB) service.Provider {
	repositoryService := &Repository{db}
	return func(ctn *service.Container) (interface{}, error) {
		return repositoryService, nil
	}
}

// From retrieves the repository service in the given container
func From(container *service.Container) (*Repository, error) {
	service, err := container.Service(ServiceName)
	if err != nil {
		return nil, errors.Wrapf(err, "error while retrieving '%s' service", ServiceName)
	}
	repositoryService, ok := service.(*Repository)
	if !ok {
		return nil, errors.Errorf("retrieved service is not a valid '%s' service", ServiceName)
	}
	return repositoryService, nil
}

// Must retrieves the repository service in the given container or panic otherwise
func Must(container *service.Container) *Repository {
	repositoryService, err := From(container)
	if err != nil {
		panic(err)
	}
	return repositoryService
}
