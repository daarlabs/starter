package home_handler

import (
	"github.com/daarlabs/hirokit/hiro"
)

func Home() hiro.Handler {
	return func(c hiro.Ctx) error {
		c.Page().Set().Title("Starter - Home")
		return c.Response().Render(HomePage(c))
	}
}
