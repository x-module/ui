/**
 * Created by Goland
 * @file   utils.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/7 14:05
 * @desc   utils.go
 */

package utils

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image"
	"image/color"
)

func ColorBackground(gtx layout.Context, point image.Point, color color.NRGBA) {
	rect := clip.Rect{
		Max: point,
	}
	paint.FillShape(gtx.Ops, color, rect.Op())
}

func WithAlpha(c color.NRGBA, a uint8) color.NRGBA {
	return color.NRGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: a,
	}
}
