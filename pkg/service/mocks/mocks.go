package mocks

import (
	"context"

	"go.keploy.io/server/pkg/models"
	"go.uber.org/zap"
)

func NewMockService(c models.MockDB, log *zap.Logger) *Mocks {
	return &Mocks{
		sdb: c,
		log: log,
	}
}

type Mocks struct {
	sdb models.MockDB
	log *zap.Logger
}

func (s *Mocks) Insert(ctx context.Context, doc models.Mock) error {
	if count, err := s.sdb.CountDocs(ctx, doc.AppID, doc.TestName); err == nil && count > 0 {
		return s.sdb.UpdateArr(ctx, doc.AppID, doc.TestName, doc)
	}
	return s.sdb.Insert(ctx, doc)
}

func (s *Mocks) Get(ctx context.Context, app string, testName string) ([]models.Mock, error) {
	return s.sdb.Get(ctx, app, testName)
}
