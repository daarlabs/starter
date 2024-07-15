package main

import (
	"os"
	
	"github.com/daarlabs/farah/farah"
	"github.com/daarlabs/farah/ui/layout_ui"
	"github.com/daarlabs/hirokit/dd"
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
	dd.Print(os.Getenv("APP_NAME"))
	dd.Print(os.Getenv("APP_PORT"))
	dd.Print(os.Getenv("APP_ENV"))
	app.Run("0.0.0.0:" + os.Getenv("APP_PORT"))
}
