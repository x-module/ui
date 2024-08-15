package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/naive/widgets"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	widgets2 "github.com/x-module/ui/widgets"
)

var th = theme.New(material.NewTheme(), true)

var menu = widgets.NewMenu(th)
var ops op.Ops
var menuItemOptions = []widgets.MenuItemOption{
	{Icon: widgets2.ActionPermIdentityIcon, Text: "ACCOUNT  ", MarginRight: 0,
		SubMenu: []widgets.MenuItemOption{
			{Icon: widgets2.NavigationSubdirectoryArrowRightIcon, Text: "LOGIN   "},
			{Icon: widgets2.NavigationSubdirectoryArrowRightIcon, Text: "LOGOUT"},
		},
	},
	{Icon: widgets2.EditorFunctionsIcon, Text: "RPC          ", MarginRight: 0},
	{Icon: widgets2.EditorBorderAllIcon, Text: "GENERATE", MarginRight: 0,
		SubMenu: []widgets.MenuItemOption{
			{Icon: widgets2.NavigationSubdirectoryArrowRightIcon, Text: "CONFIG"},
			{Icon: widgets2.NavigationSubdirectoryArrowRightIcon, Text: "CODE   "},
		},
	},
	{Icon: widgets2.MapsDirectionsRunIcon, Text: "MATCH     ", MarginRight: 0},
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
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, resource.DefaultWindowBgGrayColor, rect.Op())
				layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return menu.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return utils.DrawLine(gtx, th.Palette.Fg, unit.Dp(gtx.Constraints.Max.Y), unit.Dp(1))
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						// utils.ColorBackground(gtx, gtx.Constraints.Max, th.Palette.Bg)
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.Label(th, "Hello World").Layout(gtx)
						})
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
