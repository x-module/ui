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
	var clickable widget.Clickable
	var th = theme.New(material.NewTheme(), true)
	var dirSelect = widgets.NewDirSelector("请选择目录...")
	dirSelect.SetWidth(unit.Dp(400))
	// w := new(app.Window)
	var ops op.Ops
	confirm := widgets.NewConfirm(th)
	confirm.Confirm(func() {
		fmt.Println("确定...")
	})
	confirm.Cancel(func() {
		fmt.Println("取消...")
	})
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
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, resource.DefaultWindowBgGrayColor, rect.Op())
				// ==============================================
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(400))
						return widgets.Label(th, "&clickable, nil, 0,  unit.Dp(100)").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return dirSelect.Layout(gtx, th)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
