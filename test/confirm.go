package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
)

func main() {
	var confirm *widgets.Confirm
	var clickable widget.Clickable
	var th = theme.New(material.NewTheme(), true)

	//w := new(app.Window)
	var ops op.Ops
	confirm = widgets.NewConfirm(th)
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
					confirm.Message("确定退出吗?")
				}
				//==============================================
				layout.Stack{Alignment: layout.Center}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(400))
								return widgets.BlueLabel(th, "&clickable, nil, 0,  unit.Dp(100)").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return widgets.Button(th, &clickable, "click me", unit.Dp(100)).Layout(gtx, th)
							}),
						)
					}),
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						if confirm.Visible() {
							return confirm.Layout(gtx)
						}
						return layout.Dimensions{}
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
