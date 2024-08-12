package main

import (
	"gioui.org/app"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
	"image"
	"image/color"
)

func main() {
	go func() {
		w := new(app.Window)
		th := material.NewTheme()
		editor := widget.Editor{}

		var hovered bool

		for e := range w.Events() {
			switch e := e.(type) {
			case app.DestroyEvent:
				return
			case system.FrameEvent:
				gtx := layout.NewContext(&e.Queue, e)

				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					// 监听鼠标事件
					pointer.Rect(image.Rectangle{Max: gtx.Constraints.Max}).Add(gtx.Ops)
					pointer.InputOp{
						Tag:   &editor,
						Types: pointer.Enter | pointer.Leave,
					}.Add(gtx.Ops)

					for _, ev := range gtx.Events(&editor) {
						if ev, ok := ev.(pointer.Event); ok {
							switch ev.Type {
							case pointer.Enter:
								hovered = true
							case pointer.Leave:
								hovered = false
							}
						}
					}

					color := color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
					if hovered {
						color = color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
					}

					return material.Editor(th, &editor, "Hover over me").Layout(gtx)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
func main() {
	//var clickable widget.Clickable
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
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					// 监听鼠标事件
					pointer.Rect(image.Rectangle{Max: gtx.Constraints.Max}).Add(gtx.Ops)
					pointer.InputOp{
						Tag:   &editor,
						Types: pointer.Enter | pointer.Leave,
					}.Add(gtx.Ops)

					for _, ev := range gtx.Event(&editor) {
						if ev, ok := ev.(pointer.Event); ok {
							switch ev.Type {
							case pointer.Enter:
								hovered = true
							case pointer.Leave:
								hovered = false
							}
						}
					}

					color := color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
					if hovered {
						color = color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
					}

					return material.Editor(th, &editor, "Hover over me").Layout(gtx)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
