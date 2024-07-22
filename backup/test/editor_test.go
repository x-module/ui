// SPDX-License-Identifier: Unlicense OR MIT

package test

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	widget2 "github.com/x-module/ui/widget"
	"os"
	"testing"
)

func TestEditor(T *testing.T) {
	go func() {
		window := new(app.Window)
		window.Invalidate()
		err := runEditor(window)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runEditor(window *app.Window) error {
	var ops op.Ops
	theme := widget2.NewTheme(window)
	var (
		editor1 = widget.Editor{SingleLine: true}
		editor2 = widget.Editor{SingleLine: true}
	)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(layout.Spacer{Height: 10}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return theme.Editor(&editor1, "editor is focused").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return theme.Button("click me").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// return theme.Editor(&editor2, "editor is focused").Layout(gtx, theme.Base)
					return theme.EditorPassword(&editor2, "ssssssssss").Layout(gtx)
				}),
			)
			// theme.Editor(new(widget.Editor), "editor is focused").Layout(gtx, theme.Base)
			e.Frame(gtx.Ops)
		}
	}
}
