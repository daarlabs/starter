package main

import (
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/farah/farah"
	"github.com/daarlabs/farah/ui/layout_ui"

	"config"
	"example/route/home_route"
)

func main() {
	cfg := config.New()
	app := mirage.New(cfg)
	app.Plugin(farah.Plugin())
	app.Static(cfg.App.Public, cfg.App.Assets)
	app.Layout().Add(mirage.Main, layout_ui.Layout)
	home_route.HomeRoutes(app)
	app.Run("0.0.0.0:80")
}
