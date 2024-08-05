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
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
)

func main() {
	var clickable widget.Clickable
	var th = theme.New(material.NewTheme(), true)
	var girl = widgets.CheckBox(th, &widget.Bool{Value: true}, "handler")
	var boy = widgets.CheckBox(th, &widget.Bool{Value: true}, "handler")
	//w := new(app.Window)
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
				if clickable.Clicked(gtx) {
				}
				//==============================================
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, th.Palette.Fg, rect.Op())

				//==============================================
				layout.Stack{Alignment: layout.Center}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {

						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return widgets.Button(th, &clickable, "click me", unit.Dp(100)).Layout(gtx, th)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return girl.Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return boy.Layout(gtx)
							}),
						)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
