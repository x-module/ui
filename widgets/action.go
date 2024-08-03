/**
 * Created by Goland
 * @file   Action.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/22 18:56
 * @desc   Action.go
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
	"image"
	"image/color"
)

type Action struct {
	visible         bool
	theme           *theme2.Theme
	closeIcon       *IconButton
	title           string
	message         string
	height          int
	width           int
	actionClickable widget.Clickable
	actionFunc      func()
}

func NewAction(th *theme2.Theme) *Action {
	bkColor := color.NRGBA{}
	hoveredColor := Hovered(bkColor)
	iconSize := unit.Dp(6)
	modal := &Action{
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

func (c *Action) Action(fun func()) {
	c.actionFunc = fun
}
func (c *Action) SetWidth(width int) {
	c.width = width
}
func (c *Action) Visible() bool {
	return c.visible
}
func (c *Action) SetTitle(title string) *Action {
	c.title = title
	return c
}

func (c *Action) SetHeight(height int) *Action {
	c.height = height
	return c
}

func (c *Action) Message(message string) {
	c.message = message
	c.visible = true
}
func (c *Action) Close() {
	c.visible = false
}

func (c *Action) Layout(gtx layout.Context) layout.Dimensions {
	if !c.visible {
		return layout.Dimensions{}
	}
	if c.visible {
		// 绘制全屏半透明遮罩层
		paint.Fill(gtx.Ops, c.theme.MaskLayerBgColor)
	}
	for c.actionClickable.Clicked(gtx) {
		c.visible = false
		if c.actionFunc != nil {
			c.actionFunc()
		}
	}
	width := gtx.Dp(unit.Dp(c.width))
	height := gtx.Dp(unit.Dp(c.height))
	return new(widget.Clickable).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{
			Top: unit.Dp(0),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Set the size of the Action
				gtx.Constraints = layout.Exact(image.Point{X: width, Y: height})
				rc := clip.RRect{
					Rect: image.Rectangle{Max: image.Point{
						X: gtx.Constraints.Min.X,
						Y: gtx.Constraints.Min.Y,
					}},
					NW: 10, NE: 10, SE: 10, SW: 10,
				}
				paint.FillShape(gtx.Ops, c.theme.Palette.Fg, rc.Op(gtx.Ops))
				// Center the text inside the Action
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
						but := Button(c.theme, &c.actionClickable, nil, 1, "确 定", unit.Dp(50))
						but.Background = c.theme.ConfirmButtonColor
						but.width = unit.Dp(300)
						return but.Layout(gtx, c.theme)
					}),
				)

			})
		})
	})
}
