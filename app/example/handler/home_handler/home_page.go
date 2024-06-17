package home_handler

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/mirage"
	"github.com/daarlabs/arcanum/tempest"

	"lib/feature/counter_feature"
)

func HomePage(c mirage.Ctx) Node {
	return Div(
		tempest.Class().Grid().PlaceItemsCenter().W("screen").H("screen"),
		c.Create().Component(&counter_feature.Counter{}),
	)
}
