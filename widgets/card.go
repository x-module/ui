/**
 * Created by Goland
 * @file   card.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/22 11:54
 * @desc   card.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	theme2 "github.com/x-module/ui/theme"
	"image"
	"image/color"
)

type Card struct {
	theme *theme2.Theme
}

func NewCard(theme *theme2.Theme) *Card {
	return &Card{
		theme: theme,
	}
}
func fill(gtx layout.Context, col color.NRGBA) layout.Dimensions {
	cs := gtx.Constraints
	d := image.Point{X: cs.Max.X, Y: cs.Min.Y}
	track := image.Rectangle{
		Max: d,
	}
	defer clip.Rect(track).Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, col)
	return layout.Dimensions{Size: d}
}

func (c *Card) Layout(gtx layout.Context, children layout.Widget) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			tr := 20
			tl := 20
			br := 20
			bl := 20
			defer clip.RRect{
				Rect: image.Rectangle{Max: image.Point{
					X: gtx.Constraints.Max.X,
					Y: gtx.Constraints.Min.Y,
				}},
				NW: tl, NE: tr, SE: br, SW: bl,
			}.Push(gtx.Ops).Pop()
			return fill(gtx, c.theme.Palette.Fg)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(20).Layout(gtx, children)
		}),
	)
}
