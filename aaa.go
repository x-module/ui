package main

import (
	"gioui.org/app"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"image"
)

//type MenuState struct {
//	OptionList layout.List
//	Options    []func(gtx C) D
//}

var th = material.NewTheme()

var menuState = &component.MenuState{
	OptionList: layout.List{Axis: layout.Vertical},
	Options:    []func(gtx layout.Context) layout.Dimensions{},
}

type MenuItem struct {
	Label     string
	Clickable widget.Clickable
}

type Menu struct {
	Items []MenuItem
}

func NewMenu(items []string) *Menu {
	menuItems := make([]MenuItem, len(items))
	for i, item := range items {
		menuItems[i] = MenuItem{Label: item}
	}
	return &Menu{Items: menuItems}
}

func (m *Menu) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	menu := component.Menu(th, menuState)
	return menu.Layout(gtx)
}

var menuContextArea = component.ContextArea{
	Activation:       pointer.ButtonPrimary,
	AbsolutePosition: true,
}

func main() {
	//w := new(app.Window)
	menu := NewMenu([]string{"Item 1", "Item 2", "Item 3"})

	menuContextArea.Active()
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
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						//gtx.Constraints.Max.Y = 30
						return layout.Stack{}.Layout(gtx,
							layout.Stacked(func(gtx layout.Context) layout.Dimensions {
								return material.Label(th, unit.Sp(20), "open").Layout(gtx)
							}),
							layout.Expanded(func(gtx layout.Context) layout.Dimensions {
								return menuContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									menuState.Options = []func(gtx layout.Context) layout.Dimensions{
										func(gtx layout.Context) layout.Dimensions {
											itm := component.MenuItem(th, new(widget.Clickable), "aaa")
											//itm.Label.Color = theme2.White
											return itm.Layout(gtx)
										},
										func(gtx layout.Context) layout.Dimensions {
											itm := component.MenuItem(th, new(widget.Clickable), "bbb")
											//itm.Label.Color = theme2.White
											return itm.Layout(gtx)
										}, func(gtx layout.Context) layout.Dimensions {
											itm := component.MenuItem(th, new(widget.Clickable), "ccc")
											//itm.Label.Color = theme2.White
											return itm.Layout(gtx)
										},
										func(gtx layout.Context) layout.Dimensions {
											itm := component.MenuItem(th, new(widget.Clickable), "dddd")
											//itm.Label.Color = theme2.White
											return itm.Layout(gtx)
										},
									}
									offset := layout.Inset{
										Top:  unit.Dp(20),
										Left: unit.Dp(4),
									}
									return offset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										gtx.Constraints.Min = image.Point{}
										//m := component.Menu(theme.Material(), &node.menu)
										//m.SurfaceStyle.Fill = theme.MenuBgColor
										//return m.Layout(gtx)
										//
										return menu.Layout(gtx, th)
									})
								})
							}),
						)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Sp(20), "adfasdfasd").Layout(gtx)
					}), layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Sp(20), "adfasdfasd").Layout(gtx)
					}), layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Sp(20), "adfasdfasd").Layout(gtx)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
