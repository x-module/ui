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
	widgets2 "github.com/x-module/ui/widgets"
)

func main() {
	var clickable widget.Clickable
	var clickable1 widget.Clickable
	var clickable2 widget.Clickable
	var clickable3 widget.Clickable
	var clickable4 widget.Clickable
	var clickable5 widget.Clickable
	var clickable6 widget.Clickable
	var clickable7 widget.Clickable
	var clickable8 widget.Clickable
	var clickable9 widget.Clickable
	var clickable10 widget.Clickable
	var clickable11 widget.Clickable
	var clickable12 widget.Clickable
	var clickable13 widget.Clickable
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
					fmt.Println("clicked")
				}
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, resource.WindowBgColor, rect.Op())
				// ==============================================
				layout.Stack{Alignment: layout.Center}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.DefaultButton(th, &clickable, "default", unit.Dp(100)).Layout(gtx)
									}),
									layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.DefaultButton(th, &clickable1, "default", unit.Dp(100)).SetIcon(widgets2.DeleteIcon, 0).Layout(gtx)
									}),
								)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.TertiaryButton(th, &clickable2, "Tertiary", unit.Dp(100)).Layout(gtx)
									}),
									layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.TertiaryButton(th, &clickable3, "Tertiary", unit.Dp(100)).SetIcon(widgets2.DeleteIcon, 0).Layout(gtx)
									}),
								)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.PrimaryButton(th, &clickable4, "Primary", unit.Dp(100)).Layout(gtx)
									}),
									layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.PrimaryButton(th, &clickable5, "Primary", unit.Dp(100)).SetIcon(widgets2.DeleteIcon, 0).Layout(gtx)
									}),
								)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.InfoButton(th, &clickable6, "Info", unit.Dp(100)).Layout(gtx)
									}),
									layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.InfoButton(th, &clickable7, "Info", unit.Dp(100)).SetIcon(widgets2.DeleteIcon, 0).Layout(gtx)
									}),
								)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.SuccessButton(th, &clickable8, "Success", unit.Dp(100)).Layout(gtx)
									}),
									layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.SuccessButton(th, &clickable9, "Success", unit.Dp(100)).SetIcon(widgets2.DeleteIcon, 0).Layout(gtx)
									}),
								)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.WarningButton(th, &clickable10, "Warning", unit.Dp(100)).Layout(gtx)
									}),
									layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.WarningButton(th, &clickable11, "Warning", unit.Dp(100)).SetIcon(widgets2.DeleteIcon, 0).Layout(gtx)
									}),
								)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.ErrorButton(th, &clickable12, "Error", unit.Dp(100)).Layout(gtx)
									}),
									layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.ErrorButton(th, &clickable13, "Error", unit.Dp(100)).SetIcon(widgets2.DeleteIcon, 0).Layout(gtx)
									}),
								)
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
