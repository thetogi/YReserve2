package repository

import (
	"github.com/thetogi/YReserve2/metrics"
	"github.com/thetogi/YReserve2/repository/sqlRepository"

	"github.com/thetogi/YReserve2/logger"
	"github.com/thetogi/YReserve2/model"
)

type PersistentRepository struct {
	SqlRepository *sqlRepository.SqlRepository
	Log           logger.Logger
	Config        *model.Config
	Metrics       metrics.Metrics

	UserRepository UserRepository
}

func NewPersistentRepository(log logger.Logger, config *model.Config, metrics metrics.Metrics) *PersistentRepository {
	repository := &PersistentRepository{
		Log:     log,
		Config:  config,
		Metrics: metrics,
	}

	repository.SqlRepository = sqlRepository.NewSqlRepository(log, config.DatabaseConfig)

	return repository
}

func (s *PersistentRepository) User() UserRepository {
	return s.SqlRepository.UserRepository
}

func (s *PersistentRepository) Close() error {
	return s.SqlRepository.Close()
}
