package widget

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ScrollbarStyle configures the presentation of a scrollbar.
type ScrollbarStyle struct {
	material.ScrollbarStyle
}

// ListStyle configures the presentation of a layout.List with a scrollbar.
type ListStyle struct {
	material.ListStyle
}

func (t *Theme) Scrollbar(state *widget.Scrollbar) ScrollbarStyle {
	return ScrollbarStyle{material.Scrollbar(t.Theme, state)}
}

//c.Gray1 = argb(0x99FFFFFF)
//c.Gray2 = rgb(0x3D3D3D)
//c.Gray3 = rgb(0x8997a5)
//c.Gray4 = rgb(0x121212)
//c.Gray5 = rgb(0x363636)
//c.Surface = rgb(0x252525)

func (t *Theme) List(state *widget.List) ListStyle {
	list := ListStyle{material.List(t.Theme, state)}
	//list.Indicator.Color = t.Color.Gray3
	//list.Indicator.HoverColor = t.Color.Gray2
	return list
}

// layout the scroll track and indicator.
func (s ScrollbarStyle) layout(gtx layout.Context, axis layout.Axis, viewportStart, viewportEnd float32) layout.Dimensions {
	gtx.Constraints.Min.X = gtx.Constraints.Max.X
	return s.ScrollbarStyle.Layout(gtx, axis, viewportStart, viewportEnd)
}

// Layout the list and its scrollbar.
func (l ListStyle) Layout(gtx layout.Context, length int, w layout.ListElement) layout.Dimensions {
	gtx.Constraints.Min.X = gtx.Constraints.Max.X
	return l.ListStyle.Layout(gtx, length, w)
}
