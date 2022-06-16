//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	api "github.com/thnkrn/go-gin-template/pkg/api"
	handler "github.com/thnkrn/go-gin-template/pkg/api/handler"
	repoAdapter "github.com/thnkrn/go-gin-template/pkg/repository/adapter"
	usecaseAdapter "github.com/thnkrn/go-gin-template/pkg/usecase/adapter"
)

func InitializeAPI() (*api.ServerHTTP, error) {
	wire.Build(repoAdapter.NewRepo, usecaseAdapter.NewUseCase, handler.NewHandler, HTTPSet)

	return &api.ServerHTTP{}, nil
}
