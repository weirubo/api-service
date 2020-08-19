package service

import (
	"context"

	"github.com/weirubo/api-service/dao"
	"github.com/weirubo/api-service/global"
)

type Service struct {
	c   context.Context
	dao *dao.Dao
}

func New(c context.Context) Service {
	svc := Service{c: c}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
