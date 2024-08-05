/**
 * Created by Goland
 * @file   Confirm.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/22 18:56
 * @desc   Confirm.go
 */

package widgets

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	theme2 "github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"image"
	"image/color"
)

type Confirm struct {
	visible          bool
	theme            *theme2.Theme
	closeIcon        *IconButton
	title            string
	message          string
	height           int
	width            int
	cancelClickable  widget.Clickable
	confirmClickable widget.Clickable
	cancelFunc       func()
	confirmFunc      func()
}

func NewConfirm(th *theme2.Theme) *Confirm {
	bkColor := color.NRGBA{}
	hoveredColor := utils.Hovered(bkColor)
	iconSize := unit.Dp(6)
	modal := &Confirm{
		theme:  th,
		height: 130,
		width:  300,
		title:  "操作确认",
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

func (c *Confirm) Confirm(fun func()) {
	c.confirmFunc = fun
}
func (c *Confirm) Cancel(fun func()) {
	c.cancelFunc = fun
}

func (c *Confirm) SetWidth(width int) {
	c.width = width
}
func (c *Confirm) Visible() bool {
	return c.visible
}
func (c *Confirm) SetTitle(title string) {
	c.title = title
}

func (c *Confirm) SetHeight(height int) {
	c.height = height
}

func (c *Confirm) Message(message string) {
	c.message = message
	c.visible = true
}
func (c *Confirm) Close() {
	c.visible = false
}

func (c *Confirm) Layout(gtx layout.Context) layout.Dimensions {
	if !c.visible {
		return layout.Dimensions{}
	}
	if c.visible {
		// 绘制全屏半透明遮罩层
		paint.Fill(gtx.Ops, c.theme.MaskLayerBgColor)
	}

	for c.cancelClickable.Clicked(gtx) {
		c.visible = false
		if c.cancelFunc != nil {
			c.cancelFunc()
		}
	}

	for c.confirmClickable.Clicked(gtx) {
		c.visible = false
		if c.confirmFunc != nil {
			c.confirmFunc()
		}
	}

	width := gtx.Dp(unit.Dp(c.width))
	height := gtx.Dp(unit.Dp(c.height))
	return new(widget.Clickable).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{
			Top: unit.Dp(0),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Set the size of the Confirm
				gtx.Constraints = layout.Exact(image.Point{X: width, Y: height})
				rc := clip.RRect{
					Rect: image.Rectangle{Max: image.Point{
						X: gtx.Constraints.Min.X,
						Y: gtx.Constraints.Min.Y,
					}},
					NW: 10, NE: 10, SE: 10, SW: 10,
				}
				paint.FillShape(gtx.Ops, c.theme.Palette.Fg, rc.Op(gtx.Ops))
				// Center the text inside the Confirm
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: 0, Right: 10, Bottom: 10, Top: 10}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Stack{Alignment: layout.N}.Layout(gtx,
									layout.Stacked(func(gtx layout.Context) layout.Dimensions {
										return Label(c.theme, c.title).Layout(gtx)
									}),
								)
							})
						})
					}),
					DrawLineFlex(c.theme.Palette.Bg, unit.Dp(1), unit.Dp(c.width)),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(50))
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: 5, Right: 5, Bottom: 2, Top: 2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								label := material.Label(c.theme.Material(), c.theme.TextSize, c.message)
								label.Color = theme2.LightBlue
								label.Font.Weight = font.Medium
								return label.Layout(gtx)
							})
						})
					}),
					DrawLineFlex(c.theme.SeparatorColor, unit.Dp(1), unit.Dp(c.width)),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								but := Button(c.theme, &c.cancelClickable, "取消", unit.Dp(50))
								but.Background = c.theme.ConfirmButtonColor
								but.width = unit.Dp(150)
								return but.Layout(gtx, c.theme)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return DrawLine(gtx, c.theme.SeparatorColor, unit.Dp(35), unit.Dp(1))
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								but := Button(c.theme, &c.confirmClickable, "确认", unit.Dp(50))
								but.Background = c.theme.ConfirmButtonColor
								but.width = unit.Dp(150)
								return but.Layout(gtx, c.theme)
							}),
						)
					}),
				)

			})
		})
	})
}
