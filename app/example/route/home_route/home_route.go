package home_route

import (
	"net/http"

	"github.com/daarlabs/arcanum/mirage"

	"example/handler/home_handler"
)

func HomeRoutes(app mirage.Mirage) {
	app.Route(
		"/",
		home_handler.Home(),
		mirage.Method(http.MethodGet),
		mirage.Name("home"),
	)
}
