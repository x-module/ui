package widgets

import (
	"gioui.org/io/semantic"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/x-module/ui/internal/safemap"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"image"
	"image/color"
	"unicode"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
)

type Tabs struct {
	list     layout.List
	tabs     []*Tab
	selected int

	theme *theme.Theme

	// number of characters to show in the tab title
	maxTitleWidth    int
	onSelectedChange func(int)
}

type Tab struct {
	btn        widget.Clickable
	Title      string
	Identifier string

	Closable       bool
	CloseClickable *widget.Clickable

	isDataChanged bool
	onClose       func(t *Tab)
	isClosed      bool

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

func (tabs *Tabs) SetMaxTitleWidth(maxWidth int) {
	tabs.maxTitleWidth = maxWidth
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

func (tab *Tab) SetOnClose(f func(t *Tab)) {
	tab.onClose = f
}

func (tabs *Tabs) AddTab(tab *Tab) int {
	tabs.tabs = append(tabs.tabs, tab)
	return len(tabs.tabs) - 1
}

func (tabs *Tabs) RemoveTabByID(id string) {
	tab := tabs.findTabByID(id)
	if tab == nil {
		return
	}

	tab.isClosed = true
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
func (tabs *Tabs) Clickable(gtx layout.Context, button *widget.Clickable, w layout.Widget) layout.Dimensions {
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
				if button.Hovered() {
					paint.Fill(gtx.Ops, utils.Hovered(color.NRGBA{}))
				}
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
		if !ot.isClosed {
			tabItems = append(tabItems, ot)
		}
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

	return tabs.list.Layout(gtx, len(tabs.tabs), func(gtx layout.Context, tabIdx int) layout.Dimensions {
		if tabIdx > len(tabs.tabs)-1 {
			tabIdx = len(tabs.tabs) - 1
		}

		t := tabs.tabs[tabIdx]
		if t.Closable && t.onClose != nil && t.CloseClickable.Clicked(gtx) {
			t.onClose(t)
			gtx.Execute(op.InvalidateCmd{})
		}

		if t.btn.Clicked(gtx) {
			tabs.selected = tabIdx
			if tabs.onSelectedChange != nil {
				go tabs.onSelectedChange(tabIdx)
				gtx.Execute(op.InvalidateCmd{})
			}
		}

		if t.btn.Hovered() {
			paint.FillShape(gtx.Ops, tabs.theme.Palette.ContrastBg, clip.Rect{Max: gtx.Constraints.Min}.Op())
		}

		var tabWidth int
		return layout.Stack{Alignment: layout.S}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				var dims layout.Dimensions
				if t.Closable {
					dims = tabs.Clickable(gtx, &t.btn, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(12)).Layout(gtx,
									Label(tabs.theme, ellipticalTruncate(t.Title, tabs.maxTitleWidth)).Layout,
								)
							}),
							layout.Rigid(layout.Spacer{Width: unit.Dp(2)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								bkColor := color.NRGBA{}
								hoveredColor := utils.Hovered(bkColor)
								if t.btn.Hovered() {
									bkColor = hoveredColor
								}
								iconColor := tabs.theme.ContrastFg
								closeIcon := CloseIcon
								iconSize := unit.Dp(16)
								padding := unit.Dp(4)
								if t.isDataChanged {
									// yellow
									iconColor = color.NRGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
									closeIcon = CircleIcon
									iconSize = unit.Dp(10)
									padding = unit.Dp(8)
								}

								ib := &IconButton{
									Icon:                 closeIcon,
									Color:                iconColor,
									BackgroundColor:      bkColor,
									BackgroundColorHover: hoveredColor,
									Size:                 iconSize,
									Clickable:            t.CloseClickable,
								}
								return layout.UniformInset(padding).Layout(gtx,
									func(gtx layout.Context) layout.Dimensions {
										return ib.Layout(gtx, tabs.theme)
									},
								)
							}),
						)
					})
				} else {
					dims = tabs.Clickable(gtx, &t.btn, func(gtx layout.Context) layout.Dimensions {
						// return layout.UniformInset(unit.Dp(12)).Layout(gtx,
						return layout.Inset{Top: unit.Dp(8), Bottom: unit.Dp(8), Left: unit.Dp(12), Right: unit.Dp(12)}.Layout(gtx,
							Label(tabs.theme, t.Title).Layout,
						)
					})
				}

				tabWidth = dims.Size.X
				return dims
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				if tabs.selected != tabIdx {
					return layout.Dimensions{}
				}
				tabHeight := gtx.Dp(unit.Dp(2))
				tabRect := image.Rect(0, 0, tabWidth, tabHeight)
				paint.FillShape(gtx.Ops, tabs.theme.TabInactiveColor, clip.Rect(tabRect).Op())
				return layout.Dimensions{
					Size: image.Point{X: tabWidth, Y: tabHeight},
				}
			}),
		)
	})
}

func ellipticalTruncate(text string, maxLen int) string {
	if maxLen == 0 {
		return text
	}

	lastSpaceIx := maxLen
	l := 0
	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIx = i
		}
		l++
		if l > maxLen {
			return text[:lastSpaceIx] + "..."
		}
	}
	return text
}
