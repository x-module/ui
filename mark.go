package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/x-module/ui/theme"
)

// MaskWidget 是遮罩层的自定义 Widget。
type MaskWidget struct {
	widget.Clickable
}

func (w *MaskWidget) Layout(gtx layout.Context) layout.Dimensions {
	// 绘制遮罩层
	paint.ColorOp{Color: gtx.Palette.Surface.MulAlpha(0.5)}.Add(gtx.Ops)
	clip.Rect{
		Max: gtx.Constraints.Max.Size,
	}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	// 检查触摸事件
	if w.Clickable.Clicked() {
		// 处理点击事件
		// 这里可以添加你需要的逻辑
	}

	return layout.Dimensions{
		Size: gtx.Constraints.Max.Size,
	}
}

func main() {
	go func() {
		w := app.NewWindow(app.Title("Mask Demo"))
		th := theme.New(w.Context())

		var maskWidget MaskWidget

		gtx := layout.NewContext(&th.Theme, w)
		for {
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					maskWidget.Layout(gtx)
					return layout.Dimensions{}
				}),
			)

			if err := w.Update(); err != nil {
				return
			}
		}
	}()
	app.Main()
}
