package main

import (
	"fmt"
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
)

func main() {
	var labelInput *widgets.LabeledInput
	var clickable widget.Clickable
	var th = theme.New(material.NewTheme(), true)

	// w := new(app.Window)
	var ops op.Ops
	labelInput = widgets.NewLabeledInput(th, "用户名", "请输入用户名")
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
				// =============================================

				if clickable.Clicked(gtx) {
					fmt.Println(labelInput.GetText())
				}
				// ==============================================
				layout.Stack{Alignment: layout.Center}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return labelInput.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Spacer{Height: unit.Dp(100)}.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return widgets.DefaultButton(th, &clickable, "default", unit.Dp(100)).Layout(gtx)
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
