/**
 * Created by Goland
 * @file   common.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/13 13:47
 * @desc   common.go
 */

package widgets

import (
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"image"
	"image/color"
)

type CommonWidget struct {
	click       gesture.Click
	state       state
	borderColor color.NRGBA
	bgColor     color.NRGBA
	hint        string
	radius      unit.Dp
}

type state uint8

const (
	inactive state = iota
	hovered
	activated
	focused
)

func (t *CommonWidget) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	gtx.Constraints.Min.X = gtx.Constraints.Max.X
	gtx.Constraints.Min.Y = 0
	macro := op.Record(gtx.Ops)
	dims := widget(gtx)
	call := macro.Stop()
	defer pointer.PassOp{}.Push(gtx.Ops).Pop()
	defer clip.Rect(image.Rectangle{Max: dims.Size}).Push(gtx.Ops).Pop()
	t.click.Add(gtx.Ops)
	// event.Op(gtx.Ops, &t.textEditor)
	call.Add(gtx.Ops)
	return dims
}
