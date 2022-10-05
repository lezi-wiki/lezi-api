//go:build wireinject
// +build wireinject

package model

import (
	"github.com/google/wire"
)

func initializeClient() *DataClient {
	wire.Build(initDB, NewDataClient, NewSettingService, NewTextService)
	return &DataClient{}
}
