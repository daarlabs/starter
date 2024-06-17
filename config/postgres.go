package config

import (
	"github.com/daarlabs/arcanum/env"
	"github.com/daarlabs/arcanum/quirk"
)

func Postgres() *quirk.DB {
	host := "localhost"
	if env.Development() {
		host = "starter-postgres"
	}
	return quirk.MustConnect(
		quirk.WithPostgres(),
		quirk.WithHost(host),
		quirk.WithPort(5432),
		quirk.WithDbname("starter"),
		quirk.WithUser("starter"),
		quirk.WithPassword("starter"),
		quirk.WithSslDisable(),
		quirk.WithLog(true),
	)
}
