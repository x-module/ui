/**
 * Created by Goland
 * @file   Size.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/13 16:19
 * @desc   Size.go
 */

package resource

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

type Size struct {
	TextSize unit.Sp
	Height   unit.Dp
	Inset    layout.Inset
	IconSize unit.Dp
}

var (
	Tiny   = Size{TextSize: unit.Sp(9), Height: unit.Dp(10), Inset: layout.UniformInset(unit.Dp(4)), IconSize: unit.Dp(14)}
	Small  = Size{TextSize: unit.Sp(12), Height: unit.Dp(15), Inset: layout.UniformInset(unit.Dp(6)), IconSize: unit.Dp(18)}
	Medium = Size{TextSize: unit.Sp(14), Height: unit.Dp(20), Inset: layout.UniformInset(unit.Dp(8)), IconSize: unit.Dp(24)}
	Large  = Size{TextSize: unit.Sp(20), Height: unit.Dp(25), Inset: layout.UniformInset(unit.Dp(10)), IconSize: unit.Dp(30)}
)

var (
	DefaultTextSize   = unit.Sp(14)
	DefaultIconSize   = unit.Dp(24)
	DefaultRadiusSize = unit.Dp(4)
)
