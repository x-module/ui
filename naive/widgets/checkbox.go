package widgets

import (
	"gioui.org/font"
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"image"
	"image/color"
)

type Checkbox struct {
	CheckBox           *widget.Bool
	Label              string
	Color              color.NRGBA
	Font               font.Font
	TextSize           unit.Sp
	IconColor          color.NRGBA
	Size               unit.Dp
	shaper             *text.Shaper
	checkedStateIcon   *widget.Icon
	uncheckedStateIcon *widget.Icon
}

func CheckBox(th *theme.Theme, checkBox *widget.Bool, label string) Checkbox {
	c := Checkbox{
		CheckBox:           checkBox,
		Label:              label,
		Color:              resource.DefaultTextWhiteColor,
		IconColor:          resource.BorderLightGrayColor,
		TextSize:           resource.DefaultTextSize,
		Size:               resource.Medium.IconSize,
		shaper:             th.Shaper,
		checkedStateIcon:   th.Icon.CheckBoxChecked,
		uncheckedStateIcon: th.Icon.CheckBoxUnchecked,
	}
	c.Font.Typeface = th.Face
	return c
}

// SetSize 设置Size
func (c *Checkbox) SetSize(size resource.Size) {
	c.Size = size.IconSize
	c.TextSize = size.TextSize
}

// Layout updates the checkBox and displays it.
func (c Checkbox) Layout(gtx layout.Context) layout.Dimensions {

	return c.CheckBox.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.CheckBox.Add(gtx.Ops)
		var icon *widget.Icon
		if c.CheckBox.Value {
			icon = c.checkedStateIcon
			c.IconColor = resource.HoveredBorderBlueColor
		} else {
			icon = c.uncheckedStateIcon
		}
		if c.CheckBox.Hovered() {
			c.IconColor = resource.HoveredBorderBlueColor
		}
		return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					size := gtx.Dp(c.Size)
					col := c.IconColor
					if !gtx.Enabled() {
						col = utils.Disabled(col)
					}
					gtx.Constraints.Min = image.Point{X: size}
					icon.Layout(gtx, col)
					return layout.Dimensions{
						Size: image.Point{X: size, Y: size},
					}
				})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					colMacro := op.Record(gtx.Ops)
					paint.ColorOp{Color: c.Color}.Add(gtx.Ops)
					return widget.Label{}.Layout(gtx, c.shaper, c.Font, c.TextSize, c.Label, colMacro.Stop())
				})
			}),
		)
	})
}
