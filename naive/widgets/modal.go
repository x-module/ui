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
	"github.com/x-module/ui/naive/resource"
	theme2 "github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"github.com/x-module/ui/widgets"
	"image"
)

type Modal struct {
	visible   bool
	theme     *theme2.Theme
	content   layout.Widget
	closeIcon *IconButton
	title     string
	height    int
	width     int
}

func NewModal(th *theme2.Theme) *Modal {
	// bkColor := color.NRGBA{}
	bkColor := resource.ModalBgGrayColor
	hoveredColor := utils.Hovered(bkColor)
	iconSize := unit.Dp(16)
	modal := &Modal{
		theme:  th,
		height: 300,
		width:  500,
		title:  "modal",
	}
	modal.closeIcon = &IconButton{
		Icon:                 widgets.CloseIcon,
		Color:                resource.IconGrayColor,
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
func (m *Modal) SetWidth(width int) {
	m.width = width
}
func (m *Modal) Visible() bool {
	return m.visible
}
func (m *Modal) SetTitle(title string) {
	m.title = title
}

func (m *Modal) SetHeight(height int) {
	m.height = height
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
		paint.Fill(gtx.Ops, resource.DefaultMaskBgColor)
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
			paint.FillShape(gtx.Ops, resource.DefaultContentBgGrayColor, rc.Op(gtx.Ops))
			// Center the text inside the Modal
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: 0, Right: 10, Bottom: 10, Top: 10}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Stack{Alignment: layout.N}.Layout(gtx,
								layout.Stacked(func(gtx layout.Context) layout.Dimensions {
									label := material.Label(m.theme.Material(), unit.Sp(16), m.title)
									label.Color = resource.DefaultTextWhiteColor
									return label.Layout(gtx)
								}),
								layout.Expanded(func(gtx layout.Context) layout.Dimensions {
									return layout.Inset{Left: 470}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										return m.closeIcon.Layout(gtx)
									})
								}),
							)
						})

					})
				}),
				utils.DrawLineFlex(resource.DefaultLineColor, unit.Dp(1), unit.Dp(m.width)),
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
