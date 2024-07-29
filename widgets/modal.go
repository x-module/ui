/**
 * Created by Goland
 * @file   Modal.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/22 18:56
 * @desc   Modal.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	theme2 "github.com/x-module/ui/theme"
	"image"
	"image/color"
	"sync"
)

type Modal struct {
	sync.Mutex
	visible   bool
	theme     *theme2.Theme
	content   layout.Widget
	closeIcon *IconButton
	title     string
	height    int
	width     int
}

func NewModal(th *theme2.Theme) *Modal {
	bkColor := color.NRGBA{}
	hoveredColor := Hovered(bkColor)
	iconSize := unit.Dp(16)
	modal := &Modal{
		theme:  th,
		height: 300,
		width:  500,
		title:  "modal",
	}
	modal.closeIcon = &IconButton{
		Icon:                 CloseIcon,
		Color:                th.ContrastFg,
		BackgroundColor:      bkColor,
		BackgroundColorHover: hoveredColor,
		Size:                 iconSize,
		Clickable:            &widget.Clickable{},
	}
	modal.closeIcon.Hovered = false
	modal.closeIcon.OnClick = func() {
		modal.visible = false
	}
	return modal
}
func (t *Modal) SetWidth(width int) {
	t.width = width
}
func (t *Modal) Visible() bool {
	return t.visible
}
func (t *Modal) SetTitle(title string) {
	t.title = title
}

func (t *Modal) SetHeight(height int) {
	t.height = height
}

func (m *Modal) Display(content layout.Widget) {
	m.content = content
	m.visible = true
}
func (m *Modal) Close() {
	m.visible = false
}

func (m *Modal) Layout(gtx layout.Context) layout.Dimensions {
	if !m.visible {
		return layout.Dimensions{}
	}
	if m.visible {
		// 绘制全屏半透明遮罩层
		paint.Fill(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 0, A: 240})
	}
	width := gtx.Dp(unit.Dp(m.width))
	height := gtx.Dp(unit.Dp(m.height))
	return layout.Inset{
		Top: unit.Dp(0),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// Set the size of the Modal
			gtx.Constraints = layout.Exact(image.Point{X: width, Y: height})
			rc := clip.RRect{
				Rect: image.Rectangle{Max: image.Point{
					X: gtx.Constraints.Min.X,
					Y: gtx.Constraints.Min.Y,
				}},
				NW: 20, NE: 20, SE: 20, SW: 20,
			}
			paint.FillShape(gtx.Ops, m.theme.Palette.Bg, rc.Op(gtx.Ops))
			// Center the text inside the Modal
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: 0, Right: 10, Bottom: 10, Top: 10}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Stack{Alignment: layout.N}.Layout(gtx,
								layout.Stacked(func(gtx layout.Context) layout.Dimensions {
									label := material.Label(m.theme.Material(), unit.Sp(16), m.title)
									label.Color = m.theme.Material().Palette.Fg
									return label.Layout(gtx)
								}),
								layout.Expanded(func(gtx layout.Context) layout.Dimensions {
									return layout.Inset{Left: 470}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										return m.closeIcon.Layout(gtx, m.theme)
									})
								}),
							)
						})

					})
				}),
				DrawLineFlex(m.theme.SeparatorColor, unit.Dp(1), unit.Dp(m.width)),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: 30, Right: 30, Bottom: 30, Top: 30}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return m.content(gtx)
						})
					})
				}),
			)

		})
	})
}
