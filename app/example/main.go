package main

import (
	"github.com/daarlabs/farah/farah"
	"github.com/daarlabs/farah/ui/layout_ui"
	"github.com/daarlabs/hirokit/hiro"
	
	"config"
	"example/route/home_route"
)

func main() {
	cfg := config.New()
	app := hiro.New(cfg)
	app.Plugin(farah.Plugin())
	app.Static(cfg.App.PublicUrlPath, cfg.App.PublicLocalDir)
	app.Layout().Add(hiro.Main, layout_ui.Layout)
	home_route.HomeRoutes(app)
	app.Run("0.0.0.0:80")
}
