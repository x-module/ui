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
	"github.com/x-module/ui/naive/utils"
	"github.com/x-module/ui/naive/widgets"
	"github.com/x-module/ui/theme"
)

func main() {
	var clickable widget.Clickable
	var clickable1 widget.Clickable
	var th = theme.New(material.NewTheme(), true)
	// w := new(app.Window)
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
					utils.SystemNotice("登录成功")
				}
				if clickable1.Clicked(gtx) {
					utils.AppNotice("登录成功")
				}
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, resource.DefaultWindowBgGrayColor, rect.Op())
				// ==============================================
				layout.Stack{Alignment: layout.Center}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return widgets.DefaultButton(th, &clickable, "system notice", unit.Dp(100)).Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return widgets.DefaultButton(th, &clickable1, "application notice", unit.Dp(100)).Layout(gtx)
							}),
						)
					}),
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min = gtx.Constraints.Max
						return utils.NotificationController.Layout(gtx, th)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
