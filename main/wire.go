//go:build wireinject
// +build wireinject

package main

import (
	"todotom/go-gin-restful-api/forum"
	"todotom/go-gin-restful-api/postgres"
	"todotom/go-gin-restful-api/rest"

	"github.com/google/wire"
)

func InitializeContainer() Container {
	wire.Build(
		postgres.ProvidePostgresClient,
		postgres.ProvideThreadPostgresTable,
		postgres.ProvideThreadRepository,
		forum.ProvideGetThreadsService,
		rest.ProvideGetThreadsHandler,
		rest.ProvideController,
		ProvideContainer,
	)
	return Container{}
}
