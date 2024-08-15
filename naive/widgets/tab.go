package widgets

import (
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/internal/safemap"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"image"
)

type Tabs struct {
	list     layout.List
	tabs     []*Tab
	selected int

	theme *theme.Theme

	// number of characters to show in the tab title
	onSelectedChange func(int)

	width unit.Dp
}

type Tab struct {
	btn        widget.Clickable
	Title      string
	Identifier string

	CloseClickable *widget.Clickable

	isDataChanged bool

	Meta *safemap.Map[string]
}

func NewTabs(theme *theme.Theme, items []*Tab, onSelectedChange func(int)) *Tabs {
	t := &Tabs{
		theme:            theme,
		tabs:             items,
		selected:         0,
		onSelectedChange: onSelectedChange,
	}
	return t
}

func (tabs *Tabs) SetWidth(width unit.Dp) {
	tabs.width = width
}

func (tabs *Tabs) Selected() int {
	return tabs.selected
}

func (tabs *Tabs) SelectedTab() *Tab {
	if len(tabs.tabs) == 0 {
		return nil
	}
	return tabs.tabs[tabs.selected]
}

func (tab *Tab) GetIdentifier() string {
	if tab == nil {
		return ""
	}
	return tab.Identifier
}

func (tabs *Tabs) AddTab(tab *Tab) int {
	tabs.tabs = append(tabs.tabs, tab)
	return len(tabs.tabs) - 1
}

func (tabs *Tabs) findTabByID(id string) *Tab {
	for _, t := range tabs.tabs {
		if t.Identifier == id {
			return t
		}
	}
	return nil
}

func (tabs *Tabs) SetSelected(index int) {
	tabs.selected = index
}

func (tabs *Tabs) SetSelectedByID(id string) {
	for i, t := range tabs.tabs {
		if t.Identifier == id {
			tabs.selected = i
			return
		}
	}
}

func (tabs *Tabs) SetTabs(items []*Tab) {
	tabs.tabs = items
}

func (tab *Tab) SetDataChanged(changed bool) {
	tab.isDataChanged = changed
}

func (tab *Tab) SetIdentifier(id string) {
	tab.Identifier = id
}

func (tab *Tab) IsDataChanged() bool {
	return tab.isDataChanged
}

// decoration.
func (tabs *Tabs) clickable(gtx layout.Context, button *widget.Clickable, w layout.Widget) layout.Dimensions {
	return button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				tr := 10
				tl := 10
				br := 0
				bl := 0
				defer clip.RRect{
					Rect: image.Rectangle{Max: image.Point{
						X: gtx.Constraints.Min.X,
						Y: gtx.Constraints.Min.Y,
					}},
					NW: tl, NE: tr, SE: br, SW: bl,
				}.Push(gtx.Ops).Pop()
				// defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
				// if button.Hovered() {
				// 	paint.Fill(gtx.Ops, utils.Hovered(color.NRGBA{}))
				// }
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			w,
		)
	})
}

func (tabs *Tabs) Layout(gtx layout.Context) layout.Dimensions {
	// update tabs with new items
	tabItems := make([]*Tab, 0)
	for _, ot := range tabs.tabs {
		tabItems = append(tabItems, ot)
	}

	tabs.tabs = tabItems
	if tabs.selected > len(tabs.tabs)-1 {
		if len(tabs.tabs) > 0 {
			tabs.selected = len(tabs.tabs) - 1
		} else {
			tabs.selected = 0
		}
	}

	if len(tabs.tabs) == 1 {
		tabs.selected = 0
	}
	return layout.Stack{Alignment: layout.W}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(36)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return utils.DrawLine(gtx, resource.DefaultLineColor, unit.Dp(1), tabs.width)
			})
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return tabs.list.Layout(gtx, len(tabs.tabs), func(gtx layout.Context, tabIdx int) layout.Dimensions {
				if tabs.width == 0 {
					tabs.width = unit.Dp(gtx.Constraints.Max.X)
				}
				if tabIdx > len(tabs.tabs)-1 {
					tabIdx = len(tabs.tabs) - 1
				}

				t := tabs.tabs[tabIdx]

				if t.btn.Clicked(gtx) {
					tabs.selected = tabIdx
					if tabs.onSelectedChange != nil {
						go tabs.onSelectedChange(tabIdx)
						gtx.Execute(op.InvalidateCmd{})
					}
				}

				var tabWidth int
				return layout.Stack{Alignment: layout.S}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						var dims layout.Dimensions
						dims = tabs.clickable(gtx, &t.btn, func(gtx layout.Context) layout.Dimensions {
							// return layout.UniformInset(unit.Dp(12)).Layout(gtx,
							return layout.Inset{Top: unit.Dp(8), Bottom: unit.Dp(12), Left: unit.Dp(12), Right: unit.Dp(12)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								textColor := resource.DefaultTextWhiteColor
								if t.btn.Hovered() {
									textColor = resource.GreenColor
								}
								if tabs.selected == tabIdx {
									textColor = resource.GreenColor
								}

								label := material.Label(tabs.theme.Material(), resource.DefaultTextSize, t.Title)
								label.Color = textColor
								return label.Layout(gtx)
							})
						})
						tabWidth = dims.Size.X
						return dims
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						if tabs.selected != tabIdx {
							return layout.Dimensions{}
						}
						tabHeight := gtx.Dp(unit.Dp(2))
						tabRect := image.Rect(0, 0, tabWidth, tabHeight)
						paint.FillShape(gtx.Ops, resource.DefaultBorderBlueColor, clip.Rect(tabRect).Op())
						return layout.Dimensions{
							Size: image.Point{X: tabWidth, Y: tabHeight},
						}
					}),
				)

			})
		}),
	)
}
