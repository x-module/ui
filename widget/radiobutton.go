// SPDX-License-Identifier: Unlicense OR MIT

package widget

import (
	"image/color"

	"gioui.org/widget"
	"gioui.org/widget/material"
)

type RadioButton struct {
	material.RadioButtonStyle
}

// RadioButton returns a RadioButton with a label. The key specifies
// the value for the Enum.
func (t *Theme) RadioButton(group *widget.Enum, key, label string, color, icon color.NRGBA) RadioButton {
	rb := RadioButton{material.RadioButton(t.Theme, group, key, label)}
	rb.Color = color
	rb.IconColor = icon
	return rb
}
