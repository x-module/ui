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
}

var (
	Tiny   = Size{TextSize: unit.Sp(9), Height: unit.Dp(10), Inset: layout.UniformInset(unit.Dp(4))}
	Small  = Size{TextSize: unit.Sp(12), Height: unit.Dp(15), Inset: layout.UniformInset(unit.Dp(6))}
	Medium = Size{TextSize: unit.Sp(15), Height: unit.Dp(20), Inset: layout.UniformInset(unit.Dp(8))}
	Large  = Size{TextSize: unit.Sp(20), Height: unit.Dp(25), Inset: layout.UniformInset(unit.Dp(10))}
)
