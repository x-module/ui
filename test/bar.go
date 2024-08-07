package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
)

func main() {
	var clickable widget.Clickable
	var th = theme.New(material.NewTheme(), true)
	//var bar = component.AppBar{
	//	NavigationButton: widget.Clickable{},
	//	NavigationIcon:   widgets.DeleteIcon,
	//	Title:            "hello",
	//}
	var tip = component.Tooltip{
		Bg:           theme.LightBlue,
		CornerRadius: unit.Dp(5),
		Inset:        layout.Inset{},
		Text:         material.Label(th.Material(), unit.Sp(12), "hello"),
	}
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
					fmt.Println("clicked")
				}
				//==============================================
				layout.Stack{Alignment: layout.Center}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return widgets.Button(th, &clickable, "click me", unit.Dp(100)).Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								//return bar.Layout(gtx, th.Material(), "hello", "asdfasdf")
								return tip.Layout(gtx)
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
