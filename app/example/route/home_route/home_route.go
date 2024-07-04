package home_route

import (
	"net/http"
	
	. "github.com/daarlabs/hirokit/hiro"
	
	"example/handler/home_handler"
)

func HomeRoutes(app Hiro) {
	app.Route(
		"/",
		home_handler.Home(),
		Method(http.MethodGet),
		Name("home"),
	)
}
