package home_handler

import (
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/tempest"
	
	"lib/feature/counter_feature"
)

func HomePage(c hiro.Ctx) Node {
	return Div(
		tempest.Class().Grid().PlaceItemsCenter().W("screen").H("screen"),
		c.Create().Component(&counter_feature.Counter{}),
	)
}
