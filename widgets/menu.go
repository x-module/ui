/**
 * Created by Goland
 * @file   menu.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/7 14:20
 * @desc   menu.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/x-module/ui/theme"
)

type Menu struct {
	theme *theme.Theme
	list  *widget.List
	Items []MenuItem
}

func NewMenu(th *theme.Theme) *Menu {
	return &Menu{
		theme: th,
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
				// ScrollToEnd: true,
			},
		},
	}
}

type MenuItemOption struct {
	Icon        *widget.Icon
	Text        string
	MarginRight int
	SubMenu     []MenuItemOption
}

type MenuItem struct {
	Theme     *theme.Theme
	Content   layout.Widget
	Clickable *widget.Clickable
}

func (m *MenuItem) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return m.Content(gtx)
		}),
	)
}

func (m *Menu) copyMenuItem(item MenuItem) MenuItem {
	return MenuItem{
		Theme:     item.Theme,
		Clickable: item.Clickable,
		Content:   item.Content,
	}
}

// SetItems 设置items
func (m *Menu) SetItems(items []MenuItem) {
	m.Items = items
}

// Layout
func (m *Menu) Layout(gtx layout.Context) layout.Dimensions {
	return m.list.Layout(gtx, len(m.Items), func(gtx layout.Context, index int) layout.Dimensions {
		return m.Items[index].Layout(gtx)
	})
}
