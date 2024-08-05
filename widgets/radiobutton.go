package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
)

type RadioButtonStyle struct {
	key   string
	label string
	Group *widget.Enum
	theme *theme.Theme
}

func CustomRadioButton(th *theme.Theme, enum *widget.Enum, text, value string) material.RadioButtonStyle {
	rb := material.RadioButton(th.Material(), enum, value, text)
	// 设置白色的选中颜色
	rb.IconColor = th.TextColor
	return rb
}

// RadioButton returns a RadioButton with a label. The key specifies
// the value for the Enum.
func RadioButton(th *theme.Theme, group *widget.Enum, key, label string) *RadioButtonStyle {
	r := &RadioButtonStyle{
		theme: th,
		Group: group,
		key:   key,
		label: label,
	}
	return r
}

// Layout updates enum and displays the radio button.
func (r *RadioButtonStyle) Layout(gtx layout.Context) layout.Dimensions {
	rb := material.RadioButton(r.theme.Material(), r.Group, r.key, r.label)
	// 设置白色的选中颜色
	// rb.IconColor = r.theme.Dark.Fg
	rb.IconColor = r.theme.TextColor
	rb.Color = r.theme.TextColor
	rb.Size = unit.Dp(20)
	return rb.Layout(gtx)
}
