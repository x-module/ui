package widgets

import (
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"image"
	"image/color"

	"gioui.org/io/input"
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
)

type IconButton struct {
	Icon                 *widget.Icon
	Size                 unit.Dp
	Color                color.NRGBA
	BackgroundColor      color.NRGBA
	BackgroundColorHover color.NRGBA

	SkipFocus bool
	Hovered   bool
	Clickable *widget.Clickable

	OnClick func()
}

func (ib *IconButton) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	if ib.BackgroundColor == (color.NRGBA{}) {
		ib.BackgroundColor = resource.DefaultContentBgGrayColor
	}

	if ib.BackgroundColorHover == (color.NRGBA{}) {
		ib.BackgroundColorHover = utils.Hovered(ib.BackgroundColor)
	}

	for ib.Clickable.Clicked(gtx) {
		if ib.OnClick != nil {
			ib.OnClick()
		}
	}

	return ib.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 4).Push(gtx.Ops).Pop()
				background := ib.BackgroundColor
				if gtx.Source == (input.Source{}) {
					background = utils.Disabled(ib.BackgroundColor)
				} else if (ib.Clickable.Hovered() || (gtx.Focused(ib.Clickable) && !ib.SkipFocus)) && ib.Hovered {
					background = ib.BackgroundColorHover
				}
				paint.Fill(gtx.Ops, background)
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = gtx.Dp(ib.Size)
				return ib.Icon.Layout(gtx, ib.Color)
			},
		)
	})
}
