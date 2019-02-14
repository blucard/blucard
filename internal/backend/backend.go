package backend

import (
	"github.com/blucard/blucard/internal/pkg/dynamo/dao"
	"github.com/gin-gonic/gin"
)

// Backend represents all the actions in API layer
type Backend interface {
	GetRecord(ctx *gin.Context)
	SetRecord(ctx *gin.Context)

	GetQRCode(ctx *gin.Context)
}

type backendImpl struct {
	recordDB dao.RecordDao
}

var _ Backend = &backendImpl{}

// NewBackendInput pulls in the dependencies needed to setup the backend
type NewBackendInput struct {
	RecordDB dao.RecordDao
}

// NewBackend creates a new backend instance
func NewBackend(input *NewBackendInput) Backend {
	return &backendImpl{
		recordDB: input.RecordDB,
	}
}
