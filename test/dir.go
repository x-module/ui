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
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
)

func main() {
	var dir *widgets.DirSelector
	var clickable widget.Clickable
	var th = theme.New(material.NewTheme(), true)

	//w := new(app.Window)
	var ops op.Ops
	dir = widgets.NewDirSelector("/usr/local/go/src/", "请选择文件...")
	go func() {
		w := new(app.Window)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				//==============================================
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, th.Palette.Fg, rect.Op())
				//=============================================
				dir.SetOnSelectDir(func(dir string) {
					fmt.Println(dir)
				})
				if clickable.Clicked(gtx) {
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
								return dir.Layout(gtx, th)
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
