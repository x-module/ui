// SPDX-License-Identifier: Unlicense OR MIT

package widgets

import (
	"gioui.org/font"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
)

func H1(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.H1(th.Material(), txt))
}

func H2(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.H2(th.Material(), txt))
}

func H3(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.H3(th.Material(), txt))
}

func H4(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.H4(th.Material(), txt))
}

func H5(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.H5(th.Material(), txt))
}

func H6(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.H6(th.Material(), txt))
}

func Body1(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.Body1(th.Material(), txt))
}

func Body2(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.Body2(th.Material(), txt))
}

func Caption(th theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.Caption(th.Material(), txt))
}

func ErrorLabel(th theme.Theme, txt string) material.LabelStyle {
	label := Caption(th, txt)
	label.Color = theme.LightRed
	return label
}
func Label(th *theme.Theme, txt string) material.LabelStyle {
	label := material.Label(th.Material(), th.TextSize, txt)
	label.Color = th.TextColor
	return label
}
func BoldLabel(th *theme.Theme, txt string) material.LabelStyle {
	label := material.Label(th.Material(), th.TextSize, txt)
	label.Color = th.TextColor
	label.Font.Weight = font.Bold
	return label
}
func BlueLabel(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(material.Label(th.Material(), unit.Sp(14), txt))
}
func SizeLabel(th theme.Theme, txt string, size unit.Sp) material.LabelStyle {
	return labelWithDefaultColor(material.Label(th.Material(), size, txt))
}
func labelWithDefaultColor(entry material.LabelStyle) material.LabelStyle {
	entry.Color = theme.LightBlue
	return entry
}
