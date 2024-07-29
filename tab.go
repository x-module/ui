package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	theme2 "github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
)

//type MenuState struct {
//	OptionList layout.List
//	Options    []func(gtx C) D
//}

var tabs = widgets.NewTabs([]*widgets.Tab{
	{Title: "Params"},
	{Title: "Body"},
	{Title: "Auth"},
	{Title: "Headers"},
	//	{Title: "Pre Request"},
	{Title: "Post Request"},
}, nil)

func main() {
	//w := new(app.Window)

	var ops op.Ops
	th := theme2.New(material.NewTheme(), true)
	go func() {
		w := new(app.Window)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				Layout(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
func Layout(gtx layout.Context, theme *theme2.Theme) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(10)}
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Start,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return tabs.Layout(gtx, theme)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				switch tabs.SelectedTab().Title {
				// case "Pre Request":
				//	return r.PreRequest.Layout(gtx, theme)
				case "Post Request":
					return widgets.H6(*theme, "Post Request").Layout(gtx)
				case "Params":
					return widgets.H6(*theme, "Params").Layout(gtx)
				case "Headers":
					return widgets.H6(*theme, "Headers").Layout(gtx)
				case "Auth":
					return widgets.H6(*theme, "Auth").Layout(gtx)
				case "Body":
					return widgets.H6(*theme, "Body").Layout(gtx)
				default:
					return layout.Dimensions{}
				}
			}),
		)
	})
}
