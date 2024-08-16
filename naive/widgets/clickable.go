package widgets

import (
	"gioui.org/io/input"
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/utils"
	"image"
	"image/color"
)

type Clickable struct {
	bgColor      color.NRGBA
	bgColorHover color.NRGBA
	Clickable    widget.Clickable
	OnClick      func()
	widget       layout.Widget
}

func NewClickable() *Clickable {
	return &Clickable{}
}

func (c *Clickable) SetBgColor(bgColor color.NRGBA) {
	c.bgColor = bgColor
}
func (c *Clickable) SetBgColorHover(bgColorHover color.NRGBA) {
	c.bgColorHover = bgColorHover
}
func (c *Clickable) SetWidget(widget layout.Widget) {
	c.widget = widget
}

func (c *Clickable) SetOnClick(onClick func()) {
	c.OnClick = onClick
}

func (c *Clickable) Layout(gtx layout.Context) layout.Dimensions {
	if c.bgColor == (color.NRGBA{}) {
		c.bgColor = resource.DefaultBgGrayColor
	}

	if c.bgColorHover == (color.NRGBA{}) {
		c.bgColorHover = resource.HoveredBorderBlueColor
	}

	for c.Clickable.Clicked(gtx) {
		if c.OnClick != nil {
			c.OnClick()
		}
	}

	return c.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rect := clip.UniformRRect(image.Rectangle{Max: image.Point{
					X: gtx.Constraints.Min.X,
					Y: gtx.Constraints.Min.Y,
				}}, gtx.Dp(resource.DefaultElementRadiusSize))
				defer rect.Push(gtx.Ops).Pop()
				if gtx.Source == (input.Source{}) {
					paint.Fill(gtx.Ops, utils.Disabled(c.bgColorHover))
				} else if c.Clickable.Hovered() {
					paint.Fill(gtx.Ops, c.bgColorHover)
				}
				if gtx.Focused(c.Clickable) {
					paint.Fill(gtx.Ops, c.bgColorHover)
				}
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			c.widget,
		)
	})
}
