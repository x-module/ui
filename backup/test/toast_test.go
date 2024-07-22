package test

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	widget2 "gioui.org/widget"
	"github.com/x-module/ui/widget"
	"os"
	"testing"
)

func TestToast(t *testing.T) {
	go func() {
		window := new(app.Window)
		err := runToast(window)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var click widget2.Clickable

func prepareToDisplayUI(evt app.FrameEvent, theme *widget.Theme) (*op.Ops, layout.Context) {
	// Use a StackLayout to write the above UI components into an operations
	// list via a graphical context that is linked to the ops.
	ops := &op.Ops{}
	gtx := app.NewContext(ops, evt)
	layout.Stack{Alignment: layout.N}.Layout(
		gtx,
		layout.Stacked(theme.Toast.Layout),
	)
	return ops, gtx
}

func runToast(window *app.Window) error {
	// theme := material.NewTheme()
	theme := widget.NewTheme(window)

	// toast.NotifyError("Hello, Gio!", notification.Long)
	// toast.Layout(gtx)

	// var ops op.Ops

	// margin := layout.UniformInset(unit.Dp(10))
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops, gtx := prepareToDisplayUI(e, theme)
			// gtx := app.NewContext(ops, e)
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// layout.Rigid(toast.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if click.Clicked(gtx) {
						theme.Toast.NotifyError("Hello, Gio!")
						// theme.Card().Layout(gtx, func(gtx C) D {
						// 	msg := theme.Body1("message")
						// 	msg.Color = theme.Color.Surface
						// 	return msg.Layout(gtx)
						// })
						theme.Reload()
					}

					return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						button := theme.Button("click me")
						button.SetClickable(&click)
						return button.Layout(gtx)
					})
				}),
			)
			e.Frame(ops)
		}
	}
}
