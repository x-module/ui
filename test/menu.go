package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"github.com/x-module/ui/widgets"
	"image"
)

var clickable widget.Clickable
var th = theme.New(material.NewTheme(), true)

// w := new(app.Window)
var menu = widgets.NewMenu(th)
var ops op.Ops
var menuItemOptions = []widgets.MenuItemOption{
	{Icon: widgets.SwapHoriz, Text: "ACCOUNT  ", MarginRight: 0,
		SubMenu: []widgets.MenuItemOption{
			{Text: "LOGIN   ", MarginRight: 0},
			{Text: "LOGOUT", MarginRight: 0},
		},
	},
	{Icon: widgets.MenuIcon, Text: "RPC          ", MarginRight: 0},
	{Icon: widgets.LogsIcon, Text: "GENERATE", MarginRight: 0,
		SubMenu: []widgets.MenuItemOption{
			{Text: "CONFIG", MarginRight: 0},
			{Text: "CODE   ", MarginRight: 0},
		},
	},
	{Icon: widgets.WorkspacesIcon, Text: "MATCH     ", MarginRight: 0},
}
var clicks [][]*widget.Clickable
var menuItems []widgets.MenuItem

func generateMenu(gtx layout.Context, key int, th *theme.Theme) layout.Dimensions {
	return layout.Inset{Left: unit.Dp(0), Right: unit.Dp(0), Top: unit.Dp(0), Bottom: unit.Dp(0)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						icon := widgets.ButtonWithIcon(th, clicks[key][0], widgets.MenuIcon, 0, menuItemOptions[key].Text, unit.Dp(160))
						return icon.SetCornerRadius(unit.Dp(0)).Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.DrawLine(gtx, th.Palette.Fg, unit.Dp(1), unit.Dp(160))
					}),
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if clicked != key {
					return layout.Dimensions{}
				} else {
					if menuItemOptions[key].SubMenu != nil {
						var subMenu []layout.FlexChild
						for i := range menuItemOptions[key].SubMenu {
							subMenu = append(subMenu, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											icon := widgets.ButtonWithIcon(th, clicks[key][i+1], menuItemOptions[key].SubMenu[i].Icon, 0, "    |-- "+menuItemOptions[key].SubMenu[i].Text, unit.Dp(160))
											return icon.SetCornerRadius(unit.Dp(0)).Layout(gtx)
										})
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return widgets.DrawLine(gtx, th.Palette.Fg, unit.Dp(1), unit.Dp(160))
									}),
								)
							}))
						}
						return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							utils.ColorBackground(gtx, image.Point{
								X: gtx.Constraints.Max.X,
								Y: gtx.Dp(unit.Dp(35)),
							}, th.Palette.Fg)
							return layout.Inset{Left: unit.Dp(0)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Vertical}.Layout(gtx, subMenu...)
							})
						})

					}
					return layout.Dimensions{}
				}
			}),
			//layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			//	utils.ColorBackground(gtx, gtx.Constraints.Min, th.Palette.Fg)
			//	return layout.Dimensions{}
			//}),
		)
	})
}

var clicked = 100

func main() {
	for key := range menuItemOptions {
		scli := []*widget.Clickable{
			new(widget.Clickable),
		}
		if menuItemOptions[key].SubMenu != nil {
			for range menuItemOptions[key].SubMenu {
				scli = append(scli, new(widget.Clickable))
			}
		}
		clicks = append(clicks, scli)
	}

	for key := range menuItemOptions {
		menuItems = append(menuItems, widgets.MenuItem{
			Theme: th,
			Content: func(gtx layout.Context) layout.Dimensions {
				return generateMenu(gtx, key, th)
			},
		})
	}

	go func() {
		w := new(app.Window)
		menu.SetItems(menuItems)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				for i, cs := range clicks {
					for s, c := range cs {
						if c.Hovered() {
							fmt.Println("==========================hoasdf======================= i:", i, " s:", s)
						}
						for c.Clicked(gtx) {
							if s == 0 {
								if clicked == i {
									clicked = 100
								} else {
									clicked = i
								}
							}
							fmt.Println("===============clicked================ i:", i, " s:", s)
						}
					}
				}
				//==============================================
				layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						utils.ColorBackground(gtx, gtx.Constraints.Max, th.Palette.Bg)
						return menu.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.DrawLine(gtx, th.Palette.Fg, unit.Dp(gtx.Constraints.Max.Y), unit.Dp(1))
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						//utils.ColorBackground(gtx, gtx.Constraints.Max, th.Palette.Bg)
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
