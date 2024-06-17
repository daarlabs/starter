package home_handler

import "github.com/daarlabs/arcanum/mirage"

func Home() mirage.Handler {
	return func(c mirage.Ctx) error {
		c.Page().Set().Title("Starter - Home")
		return c.Response().Render(HomePage(c))
	}
}
