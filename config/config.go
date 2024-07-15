package config

import (
	"os"
	
	"github.com/daarlabs/hirokit/cache/memory"
	"github.com/daarlabs/hirokit/config"
	"github.com/daarlabs/hirokit/env"
	"github.com/daarlabs/hirokit/esquel"
	"github.com/daarlabs/hirokit/form"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
)

func init() {
	tempest.GlobalConfig = &tempest.Config{
		FontFamily: "Sora, sans-serif",
		Font: map[string]tempest.Font{
			"sora": {
				Value: "Sora, sans-serif",
				Url:   "https://fonts.googleapis.com/css2?family=Sora:wght@100..800&display=swap",
			},
		},
		Styles:  []string{},
		Scripts: []string{},
	}
	tempest.Start()
}

func New() config.Config {
	name := os.Getenv("APP_NAME")
	return config.Config{
		App: config.App{
			Name:           "starter-" + name,
			PublicUrlPath:  "/public/",
			PublicLocalDir: "/app/" + name + "/public/",
		},
		Cache: config.Cache{
			Redis:  Redis(),
			Memory: memory.New(".cache"),
		},
		Database: map[string]*esquel.DB{
			hiro.Main: Postgres(),
		},
		Dev: config.Dev{
			LiveReload: env.Development(),
			Tool:       env.Development(),
		},
		Form: form.Config{Limit: 256},
		Router: config.Router{
			Recover: true,
		},
	}
}
