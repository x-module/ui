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
	"gioui.org/io/input"
	"gioui.org/io/key"
	"gioui.org/layout"
	"github.com/x-module/ui/utils"
)

type Hover struct {
}

type state uint8
type LabelAlignment uint8

const (
	inactive state = iota
	hovered
	activated
	focused
)

func (t *Hover) update(gtx layout.Context, th *theme.Theme) {
	disabled := gtx.Source == (input.Source{})
	for {
		ev, ok := t.click.Update(gtx.Source)
		if !ok {
			break
		}
		switch ev.Kind {
		case gesture.KindPress:
			gtx.Execute(key.FocusCmd{Tag: &t.textEditor})
		default:
		}
	}

	t.state = inactive
	if t.click.Hovered() && !disabled {
		t.state = hovered
	}
	if t.textEditor.Len() > 0 {
		t.state = activated
	}
	if gtx.Source.Focused(&t.textEditor) && !disabled {
		t.state = focused
	}
	switch t.state {
	case inactive:
		t.border = utils.WithAlpha(th.Fg, 128)
	case hovered:
		t.border = utils.WithAlpha(th.Fg, 221)
	case focused:
		t.border = th.ContrastBg
	case activated:
		t.border = utils.WithAlpha(th.Fg, 221)
	}
}
