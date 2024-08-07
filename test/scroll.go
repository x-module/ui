package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
)

func main() {

	//var clickable widget.Clickable

	var th = theme.New(material.NewTheme(), true)
	var elements []layout.Widget
	//var elements []widgets.Notice
	var scroll = widgets.NewScroll(th)
	//var scroll = widgets.NewScrollNoBar(th)
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
				if len(elements) < 10 {
					for i := 0; i < 100; i++ {
						elements = append(elements, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									gtx.Constraints.Max.Y = 100
									return widgets.Label(th, "n.Name").Layout(gtx)
								}),
							)
						})
					}
				}

				scroll.SetElementList(elements)
				//==============================================
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return scroll.Layout(gtx)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
