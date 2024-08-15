package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/naive/widgets"
	"github.com/x-module/ui/theme"
)

func main() {
	var th = theme.New(material.NewTheme(), true)

	// w := new(app.Window)
	var ops op.Ops
	dropDown := widgets.NewDropDown(th, []string{"a", "b", "c", "d"}...)
	dropDown2 := widgets.NewDropDown(th, []string{"aaa", "bbb", "cccc", "dddd"}...)
	card := widgets.NewCard(th)
	// dropDown.SetOptions(options)
	dropDown.SetWidth(unit.Dp(300))
	go func() {
		w := new(app.Window)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				// ==============================================
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, resource.DefaultWindowBgGrayColor, rect.Op())

				layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return dropDown.Layout(gtx, th)
							})
						}),
						layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return dropDown2.Layout(gtx, th)
							})
						}),
						layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
					)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
