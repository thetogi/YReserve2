package repository

import (
	"github.com/thetogi/YReserve2/metrics"
	"github.com/thetogi/YReserve2/repository/redisRepository"

	"github.com/thetogi/YReserve2/logger"
	"github.com/thetogi/YReserve2/model"
)

type PersistentCacheRepository struct {
	*PersistentRepository
	RedisRepository *redisRepository.RedisRepository
}

func NewPersistentCacheRepository(log logger.Logger, config *model.Config, metrics metrics.Metrics) *PersistentCacheRepository {
	repository := &PersistentCacheRepository{
		PersistentRepository: NewPersistentRepository(log, config, metrics),
	}

	repository.RedisRepository = redisRepository.NewRedisRepository(log, config.CacheConfig, repository.SqlRepository)
	return repository
}

func (s *PersistentCacheRepository) User() UserRepository {
	return s.RedisRepository.UserRepository
}
