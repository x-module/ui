package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/naive/widgets"
	"github.com/x-module/ui/theme"
	widgets2 "github.com/x-module/ui/widgets"
)

func main() {
	var clickable1 widget.Clickable
	var th = theme.New(material.NewTheme(), true)
	card := widgets.NewCard(th)
	// w := new(app.Window)
	image := widgets.NewImage(th, "welcome.jpg")
	var ops op.Ops
	go func() {
		w := new(app.Window)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, resource.DefaultWindowBgGrayColor, rect.Op())
				// ==============================================
				layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return image.Layout(gtx)
							})
						}),
						layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return widgets.DefaultButton(th, &clickable1, "default", unit.Dp(100)).SetIcon(widgets2.DeleteIcon, 0).Layout(gtx)
							})
						}),
					)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
