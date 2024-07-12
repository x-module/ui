// SPDX-License-Identifier: Unlicense OR MIT

package test

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	widget2 "github.com/x-module/ui/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image/color"
	"os"
	"testing"
)

func TestButton(T *testing.T) {
	go func() {
		window := new(app.Window)
		window.Invalidate()
		err := runButton(window)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runButton(window *app.Window) error {
	var ops op.Ops
	theme := widget2.NewTheme(window)
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			bgColor := color.NRGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xFF}
			paint.ColorOp{Color: bgColor}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)

			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return theme.Button("Button").Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return theme.OutlineButton("OutlineButton").Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					button := theme.DangerButton("click me")
					button.SetClickable(&click)
					button.SetMargin(layout.UniformInset(unit.Dp(20)))
					// button.SetEnabled(false)
					return button.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return theme.IconButton(widget2.MustIcon(widget.NewIcon(icons.ActionHome))).Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					butt := theme.TextAndIconButton("TextAndIconButton", widget2.MustIcon(widget.NewIcon(icons.ActionHome)))
					return butt.Layout(gtx)
				}),
			)
			e.Frame(gtx.Ops)
		}

	}
}
