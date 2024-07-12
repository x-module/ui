/**
 * Created by Goland
 * @file   theme.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/11 17:07
 * @desc   theme.go
 */

package widget

import (
	"gioui.org/app"
	"gioui.org/unit"
	"gioui.org/widget/material"
	values2 "github.com/x-module/ui/values"
	"image/color"
)

type Theme struct {
	Color    *values2.Color
	Base     *material.Theme
	Styles   *values2.WidgetStyles
	Window   *app.Window
	TextSize unit.Sp

	Toast *Toast
}

func NewTheme(window *app.Window) *Theme {
	th := &Theme{
		Color:    values2.NewColor().DefaultThemeColors(),
		Base:     material.NewTheme(),
		Window:   window,
		Styles:   values2.DefaultWidgetStyles(),
		TextSize: values2.TextSize16,
	}
	th.Toast = NewToast(th)
	return th
}

// Hovered blends color towards a brighter color.
func Hovered(c color.NRGBA) (d color.NRGBA) {
	const r = 0x20 // lighten ratio
	return color.NRGBA{
		R: byte(255 - int(255-c.R)*(255-r)/256),
		G: byte(255 - int(255-c.G)*(255-r)/256),
		B: byte(255 - int(255-c.B)*(255-r)/256),
		A: c.A,
	}
}
