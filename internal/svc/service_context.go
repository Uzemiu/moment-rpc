package svc

import (
	"github.com/xh-polaris/moment-rpc/internal/config"
	"github.com/xh-polaris/moment-rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	model.MomentModel
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		MomentModel: model.NewMomentModel(c.Mongo.URL, c.Mongo.DB, model.MomentCollectionName, c.CacheConf),
		Redis:       c.Redis.NewRedis(),
	}
}
