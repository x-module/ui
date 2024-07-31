/**
 * Created by Goland
 * @file   list_menu.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/27 23:25
 * @desc   list_menu.go
 */

package widgets

import (
	"fmt"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/x/component"
	"github.com/x-module/ui/theme"
	"image"
	"image/color"
)

type ListMenu struct {
	Label           string
	optionsItems    []string
	theme           *theme.Theme
	menuContextArea component.ContextArea
	menuState       component.MenuState
	labelWidth      unit.Dp
	menuWidth       unit.Dp
	clickFun        func(key int, menu string)
	options         []*ListMenuOption
}

type ListMenuOption struct {
	Text      string
	Value     string
	clickable widget.Clickable

	Icon      *widget.Icon
	IconColor color.NRGBA
}

func NewListMenu(th *theme.Theme, label string, options []*ListMenuOption) *ListMenu {
	listMenu := &ListMenu{
		theme:      th,
		Label:      label,
		labelWidth: unit.Dp(100),
		options:    options,
		menuContextArea: component.ContextArea{
			Activation:       pointer.ButtonPrimary,
			AbsolutePosition: true,
		},
		// menuState: component.MenuState{
		//	OptionList: layout.List{Axis: layout.Vertical},
		//	Options:    []func(gtx layout.Context) layout.Dimensions{},
		// },
	}
	return listMenu
}

func (l *ListMenu) SetMenuWidth(width unit.Dp) {
	l.menuWidth = width
}
func (l *ListMenu) SetLabelWidth(width unit.Dp) {
	l.labelWidth = width
}
func (l *ListMenu) Clicked(fun func(key int, menu string)) {
	l.clickFun = fun
}

func (l *ListMenu) Layout(gtx layout.Context) layout.Dimensions {
	for i, opt := range l.options {
		for opt.clickable.Clicked(gtx) {
			fmt.Println("clicked,i:", i)
			l.clickFun(i, opt.Value)
		}
	}
	l.updateMenuItems()
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{}.Layout(gtx,
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						label := Button(l.theme, new(widget.Clickable), nil, 1, l.Label, l.labelWidth)
						label.SetBackground(l.theme.Dark.Bg)
						return label.Layout(gtx, l.theme)
					})
				}),
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					return l.menuContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							offset := layout.Inset{
								Top:  unit.Dp(30),
								Left: unit.Dp(2),
							}
							return offset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								gtx.Constraints.Min = image.Point{}
								menu := component.Menu(l.theme.Material(), &l.menuState)
								menu.SurfaceStyle.Fill = l.theme.MenuBgColor
								return menu.Layout(gtx)
							})
						})
					})
				}),
			)
		}),
	)
}

// updateMenuItems creates or updates menu items based on options and calculates minWidth.
func (c *ListMenu) updateMenuItems() {
	c.menuState.Options = c.menuState.Options[:0]
	for _, opt := range c.options {
		opt := opt
		c.menuState.Options = append(c.menuState.Options, func(gtx layout.Context) layout.Dimensions {
			itm := component.MenuItem(c.theme.Material(), &opt.clickable, opt.Text)
			if opt.Icon != nil {
				itm.Icon = opt.Icon
				itm.IconColor = opt.IconColor
			}
			itm.Label.Color = theme.White
			return itm.Layout(gtx)
		})
	}
}
