/**
 * Created by Goland
 * @file   common.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/11 16:17
 * @desc   common.go
 */

package widget

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func fill(gtx layout.Context, col color.NRGBA) layout.Dimensions {
	cs := gtx.Constraints
	d := image.Point{X: cs.Min.X, Y: cs.Min.Y}
	track := image.Rectangle{
		Max: d,
	}
	defer clip.Rect(track).Push(gtx.Ops).Pop()
	// 设置成白色背景
	// col = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	paint.Fill(gtx.Ops, col)

	return layout.Dimensions{Size: d}
}
