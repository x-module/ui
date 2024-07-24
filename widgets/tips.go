/**
 * Created by Goland
 * @file   tips.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/20 23:29
 * @desc   tips.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	theme2 "github.com/x-module/ui/theme"
	"image"
	"image/color"
	"sync"
)

type Tips struct {
	sync.Mutex
	visible bool
	message string
	theme   *theme2.Theme
}

func NewTips(th *theme2.Theme) *Tips {
	return &Tips{
		theme: th,
	}
}

func (t *Tips) Notify(message string) {
	t.message = message
	t.visible = true
}
func (t *Tips) Close() {
	t.visible = false
}

func (t *Tips) Layout(gtx layout.Context) layout.Dimensions {
	if !t.visible {
		return layout.Dimensions{}
	}
	if t.visible {
		// 绘制全屏半透明遮罩层
		// fill := color.NRGBA{R: 0, G: 0, B: 0, A: 240} // 半透明黑色
		// paint.FillShape(gtx.Ops, fill, clip.Rect{Max: gtx.Constraints.Max}.Op())
		paint.Fill(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 0, A: 240})
	}
	// Define the size of the Tips
	width := gtx.Dp(250)
	height := gtx.Dp(50)
	return layout.Inset{
		Top: unit.Dp(0),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Center the Tips in the parent container
		return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// Set the size of the Tips
			gtx.Constraints = layout.Exact(image.Point{X: width, Y: height})
			// 绘制全屏半透明遮罩层
			// fill := color.NRGBA{R: 255, G: 255, B: 255, A: 255} // 半透明黑色
			// paint.FillShape(gtx.Ops, fill, clip.Rect{Max: gtx.Constraints.Max}.Op())
			rc := clip.RRect{
				Rect: image.Rectangle{Max: image.Point{
					X: gtx.Constraints.Min.X,
					Y: gtx.Constraints.Min.Y,
				}},
				NW: 20, NE: 20, SE: 20, SW: 20,
			}
			paint.FillShape(gtx.Ops, t.theme.Palette.Bg, rc.Op(gtx.Ops))
			// Center the text inside the Tips

			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				l := material.Label(t.theme.Material(), unit.Sp(16), "")
				l.Font.Typeface = "MaterialIcons"
				l.Color = t.theme.Material().Palette.Fg
				l.Text = t.message
				return l.Layout(gtx)
			})
		})
	})
}
