/**
 * Created by Goland
 * @file   switch.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/15 23:30
 * @desc   switch.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
)

type Switch struct {
	theme *theme.Theme
	open  material.SwitchStyle
	ok    widget.Bool
}

func NewSwitch(theme *theme.Theme, description string) *Switch {
	ok := widget.Bool{}
	openSwitch := &Switch{
		theme: theme,
		ok:    ok,
		open:  material.Switch(theme.Material(), &ok, description),
	}
	openSwitch.open.Color.Enabled = resource.GreenColor
	openSwitch.open.Color.Disabled = resource.InfoColor
	return openSwitch
}

func (s *Switch) Open() bool {
	return s.ok.Value
}

func (s *Switch) Layout(gtx layout.Context) layout.Dimensions {
	return s.open.Layout(gtx)
}
