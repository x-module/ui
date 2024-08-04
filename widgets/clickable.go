package widgets

import (
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/x-module/ui/utils"
	"image"
	"image/color"
)

// Clickable lays out a rectangular clickable widget without further
// decoration.
func Clickable(gtx layout.Context, button *widget.Clickable, w layout.Widget) layout.Dimensions {
	return button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				tr := 10
				tl := 10
				br := 10
				bl := 10
				defer clip.RRect{
					Rect: image.Rectangle{Max: image.Point{
						X: gtx.Constraints.Min.X,
						Y: gtx.Constraints.Min.Y,
					}},
					NW: tl, NE: tr, SE: br, SW: bl,
				}.Push(gtx.Ops).Pop()
				//defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
				if button.Hovered() {
					paint.Fill(gtx.Ops, utils.Hovered(color.NRGBA{}))
				}
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			w,
		)
	})
}
