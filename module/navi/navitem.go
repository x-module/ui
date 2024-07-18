package navi

import (
	"fmt"
	"github.com/x-module/ui/module/misc"
	widget2 "github.com/x-module/ui/widget"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"

	"github.com/x-module/ui/module/list"
	"github.com/x-module/ui/module/menu"
	"github.com/x-module/ui/module/view"
)

var (
	moreIcon, _ = widget.NewIcon(icons.NavigationMoreHoriz)
)

var NavItemPadding = layout.Inset{
	Left:   unit.Dp(8),
	Right:  unit.Dp(0),
	Top:    unit.Dp(14),
	Bottom: unit.Dp(14),
}

type NavSection interface {
	Title() string
	Layout(gtx C, th *widget2.Theme) D
	Attach(d *NavDrawer)
}

type NavItem interface {
	OnSelect(gtx layout.Context) view.Intent
	Icon() *widget.Icon
	Layout(gtx layout.Context, th *widget2.Theme, textColor color.NRGBA) D
	// when there's menu options, a context menu should be attached to this navItem.
	// The returned boolean value suggest the position of the popup menu should be at
	// fixed position or not. NavItemStyle should place a clickable icon to guide user interactions.
	ContextMenuOptions(gtx layout.Context) ([][]menu.MenuOption, bool)
	Children() []NavItem
}

type NavItemStyle struct {
	drawer      *NavDrawer
	item        NavItem
	label       *list.InteractiveLabel
	menu        *menu.ContextMenu
	fixMenuPos  bool
	showMenuBtn widget.Clickable

	childList layout.List
	children  []*NavItemStyle
}

func (n *NavItemStyle) IsSelected() bool {
	return n.label.IsSelected()
}

func (n *NavItemStyle) Unselect() {
	n.label.Unselect()
}

func (n *NavItemStyle) Update(gtx C) bool {
	if n.menu == nil {
		menuOpts, fixPos := n.item.ContextMenuOptions(gtx)
		if len(menuOpts) > 0 {
			n.menu = menu.NewContextMenu(menuOpts, fixPos)
			n.menu.PositionHint = layout.N
			n.fixMenuPos = fixPos
		}
	}

	// handle naviitem events
	if n.label.Update(gtx) {
		n.drawer.OnItemSelected(gtx, n)
		return true
	}

	return false
}

func (n *NavItemStyle) layoutRoot(gtx layout.Context, th *widget2.Theme) layout.Dimensions {
	macro := op.Record(gtx.Ops)
	dims := layout.Inset{Bottom: unit.Dp(0)}.Layout(gtx, func(gtx C) D {
		return n.label.Layout(gtx, th, func(gtx C, color color.NRGBA) D {
			return NavItemPadding.Layout(gtx, func(gtx C) D {
				return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						if n.item.Icon() == nil {
							return layout.Dimensions{}
						}
						return layout.Inset{Right: unit.Dp(6)}.Layout(gtx, func(gtx C) D {
							iconColor := th.ContrastBg
							if n.label.IsSelected() {
								iconColor = th.ContrastFg
							}
							return misc.Icon{Icon: n.item.Icon(), Color: iconColor}.Layout(gtx, th)
						})
					}),
					layout.Flexed(1, func(gtx C) D {
						return layout.W.Layout(gtx, func(gtx C) D {
							return n.item.Layout(gtx, th, color)
						})
					}),

					layout.Rigid(func(gtx C) D {
						if n.menu == nil || !n.fixMenuPos {
							return D{}
						}
						return material.Clickable(gtx, &n.showMenuBtn, func(gtx C) D {
							alpha := 0xb6
							if n.showMenuBtn.Hovered() {
								alpha = 0xff
							}
							dims := misc.Icon{Icon: moreIcon, Color: misc.WithAlpha(color, uint8(alpha))}.Layout(gtx, th)
							// a tricky way to let contextual menu show up just near the button.
							n.menu.Layout(gtx, th)
							return dims
						})

					}),
				)
			})
		})
	})
	c := macro.Stop()
	defer clip.Rect(image.Rectangle{Max: dims.Size}).Push(gtx.Ops).Pop()
	c.Add(gtx.Ops)

	// if menu is not fixed position, let it follow the pointer.
	if n.menu != nil && !n.fixMenuPos {
		n.menu.Layout(gtx, th)
	}

	return dims
}

func (n *NavItemStyle) Layout(gtx C, th *widget2.Theme) D {
	if n.label == nil {
		n.label = &list.InteractiveLabel{}
	}

	n.Update(gtx)

	itemChildren := n.item.Children()
	if len(n.item.Children()) <= 0 {
		return n.layoutRoot(gtx, th)
	}

	if len(n.children) != len(itemChildren) {
		n.children = n.children[:0]
		for _, child := range itemChildren {
			n.children = append(n.children, NewNavItem(child, n.drawer))
		}
	}

	n.childList.Axis = layout.Vertical

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Inset{
				Top:    unit.Dp(8),
				Bottom: unit.Dp(8),
				Left:   unit.Dp(10),
				Right:  unit.Dp(0),
			}.Layout(gtx, func(gtx C) D {

				fmt.Println("======444==============")
				fmt.Println("======444==============")
				fmt.Println("======444==============")
				return n.childList.Layout(gtx, len(n.children), func(gtx C, index int) D {
					return n.children[index].Layout(gtx, th)
				})
			})
		}),
	)

}

func NewNavItem(item NavItem, drawer *NavDrawer) *NavItemStyle {
	style := &NavItemStyle{
		item:       item,
		label:      &list.InteractiveLabel{},
		drawer:     drawer,
		fixMenuPos: false,
	}

	return style
}

type simpleItemSection struct {
	item *NavItemStyle
}

type simpleNavItem struct {
	icon        *widget.Icon
	name        string
	targetView  view.ViewID
	openAsModal bool
}

func (item simpleNavItem) OnSelect(gtx C) view.Intent {
	return view.Intent{
		Target:      item.targetView,
		ShowAsModal: item.openAsModal,
	}
}

func (item simpleNavItem) Icon() *widget.Icon {
	return item.icon
}

func (item simpleNavItem) Layout(gtx C, th *widget2.Theme, textColor color.NRGBA) D {
	label := material.Label(th.Theme, th.TextSize, item.name)
	label.Color = textColor
	label.TextSize = unit.Sp(16) // todo
	return label.Layout(gtx)
}

func (item simpleNavItem) ContextMenuOptions(gtx C) ([][]menu.MenuOption, bool) {
	return nil, false
}

func (item simpleNavItem) Children() []NavItem {
	return nil
}

func (ss simpleItemSection) Title() string {
	return ""
}

func (ss simpleItemSection) Layout(gtx C, th *widget2.Theme) D {
	return ss.item.Layout(gtx, th)
}

func (ss simpleItemSection) Attach(d *NavDrawer) {
	ss.item.drawer = d
}

func SimpleItemSection(icon *widget.Icon, name string, targetView view.ViewID, openAsModal bool) NavSection {
	item := NewNavItem(simpleNavItem{icon: icon, name: name, targetView: targetView, openAsModal: openAsModal}, nil)
	return simpleItemSection{item: item}
}
func ChildSimpleItemSection(icon *widget.Icon, name string, targetView view.ViewID, openAsModal bool) *simpleNavItem {
	return &simpleNavItem{icon: icon, name: name, targetView: targetView, openAsModal: openAsModal}
}
