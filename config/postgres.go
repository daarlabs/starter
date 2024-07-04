package config

import (
	"github.com/daarlabs/hirokit/env"
	"github.com/daarlabs/hirokit/esquel"
)

func Postgres() *esquel.DB {
	host := "localhost"
	if env.Development() {
		host = "starter-postgres"
	}
	return esquel.MustConnect(
		esquel.WithPostgres(),
		esquel.WithHost(host),
		esquel.WithPort(5432),
		esquel.WithDbname("starter"),
		esquel.WithUser("starter"),
		esquel.WithPassword("starter"),
		esquel.WithSslDisable(),
		esquel.WithLog(true),
	)
}
