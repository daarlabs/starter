package counter_feature

import (
	"github.com/daarlabs/farah/ui/button_ui"
	"github.com/daarlabs/farah/ui/icon_ui"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/hx"
	"github.com/daarlabs/hirokit/tempest"
)

type Counter struct {
	hiro.Component
	Count int `json:"count"`
}

func (c *Counter) Name() string {
	return "counter"
}

func (c *Counter) Mount() {}

func (c *Counter) Node() Node {
	return Div(
		tempest.Class().Grid().Gap(4).TextCenter(),
		c.createCounter(),
		Div(
			tempest.Class().Grid().GridCols(2).Gap(4),
			c.createAddButton(),
			c.createResetButton(),
		),
	)
}

func (c *Counter) HandleAdd() error {
	c.Count++
	return c.Response().Render(c.createCounter())
}

func (c *Counter) HandleReset() error {
	c.Count = 0
	return c.Response().Render(c.createCounter())
}

func (c *Counter) createCounter() Node {
	return Div(
		Id(hx.Id("counter")),
		Span(
			tempest.Class().FontBold(),
			Text("Counter: "),
		),
		Span(Text(c.Count)),
	)
}

func (c *Counter) createAddButton() Node {
	return button_ui.EmeraldButton(
		button_ui.Props{
			Type: "button",
			Icon: icon_ui.Add,
		},
		hx.Get(c.Generate().Action("HandleAdd")),
		hx.Trigger("click"),
		hx.Target(hx.HashId("counter")),
		hx.Swap(hx.SwapOuterHtml),
		Text("Add"),
	)
}

func (c *Counter) createResetButton() Node {
	return button_ui.RedButton(
		button_ui.Props{
			Type: "button",
			Icon: icon_ui.Refresh,
		},
		hx.Get(c.Generate().Action("HandleReset")),
		hx.Trigger("click"),
		hx.Target(hx.HashId("counter")),
		hx.Swap(hx.SwapOuterHtml),
		Text("Reset"),
	)
}
