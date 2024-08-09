package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"github.com/x-module/ui/widgets"
)

var th = theme.New(material.NewTheme(), true)

var menu = widgets.NewMenu(th)
var ops op.Ops
var menuItemOptions = []widgets.MenuItemOption{
	{Icon: widgets.ActionPermIdentityIcon, Text: "ACCOUNT  ", MarginRight: 0,
		SubMenu: []widgets.MenuItemOption{
			{Icon: widgets.NavigationSubdirectoryArrowRightIcon, Text: "LOGIN   "},
			{Icon: widgets.NavigationSubdirectoryArrowRightIcon, Text: "LOGOUT"},
		},
	},
	{Icon: widgets.EditorFunctionsIcon, Text: "RPC          ", MarginRight: 0},
	{Icon: widgets.EditorBorderAllIcon, Text: "GENERATE", MarginRight: 0,
		SubMenu: []widgets.MenuItemOption{
			{Icon: widgets.NavigationSubdirectoryArrowRightIcon, Text: "CONFIG"},
			{Icon: widgets.NavigationSubdirectoryArrowRightIcon, Text: "CODE   "},
		},
	},
	{Icon: widgets.MapsDirectionsRunIcon, Text: "MATCH     ", MarginRight: 0},
}

func main() {
	menu.SetClickCallback(func(main int, sub int) {
		fmt.Println("===============clicked================ main:", main, " sub:", sub)
	})
	menu.SetMenuItemOptions(menuItemOptions)

	go func() {
		w := new(app.Window)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						utils.ColorBackground(gtx, gtx.Constraints.Max, th.Palette.Bg)
						return menu.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.DrawLine(gtx, th.Palette.Fg, unit.Dp(gtx.Constraints.Max.Y), unit.Dp(1))
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						// utils.ColorBackground(gtx, gtx.Constraints.Max, th.Palette.Bg)
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.BlueLabel(th, "Hello World").Layout(gtx)
						})
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
