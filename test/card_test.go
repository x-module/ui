/**
 * Created by Goland
 * @file   card_test.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/11 16:48
 * @desc   card_test.go
 */

package test

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/x-module/ui/widget"
	"image/color"
	"os"
	"testing"
)

func TestCard(T *testing.T) {
	go func() {
		window := new(app.Window)
		err := runCard(window)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runCard(window *app.Window) error {
	// theme := material.NewTheme()
	theme := widget.NewTheme(window)
	var ops op.Ops
	// margin := layout.UniformInset(unit.Dp(10))
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// 设置背景颜色为浅灰色
			bgColor := color.NRGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xFF}
			paint.ColorOp{Color: bgColor}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			theme.Card().Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				margin := layout.UniformInset(unit.Dp(10))
				return margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H5(theme.Theme, "Hello, Gio!").Layout(gtx)
				})
			})
			e.Frame(gtx.Ops)
		}
	}
}
