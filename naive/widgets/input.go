package widgets

import (
	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/input"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
)

type Input struct {
	textEditor  widget.Editor
	Background  color.NRGBA
	Text        string
	Placeholder string
	click       gesture.Click
	state       state
	border      color.NRGBA
}

func NewInput(text, placeholder string) *Input {
	t := &Input{
		textEditor:  widget.Editor{},
		Text:        text,
		Placeholder: placeholder,
	}
	t.textEditor.SetText(text)
	t.textEditor.SingleLine = true
	return t
}
func (t *Input) Password() {
	t.textEditor.Mask = '*'
}

func (t *Input) SetText(text string) {
	t.textEditor.SetText(text)
}

type state uint8
type LabelAlignment uint8

const (
	inactive state = iota
	hovered
	activated
	focused
)

func (t *Input) update(gtx layout.Context, th *theme.Theme) {
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

func (t *Input) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	t.update(gtx, th)
	gtx.Constraints.Min.X = gtx.Constraints.Max.X
	gtx.Constraints.Min.Y = 0
	macro := op.Record(gtx.Ops)
	dims := t.layout(gtx, th)
	call := macro.Stop()
	defer pointer.PassOp{}.Push(gtx.Ops).Pop()
	defer clip.Rect(image.Rectangle{Max: dims.Size}).Push(gtx.Ops).Pop()
	t.click.Add(gtx.Ops)
	event.Op(gtx.Ops, &t.textEditor)
	call.Add(gtx.Ops)
	return dims
}

func (t *Input) layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	t.Background = th.Palette.Fg
	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			rr := gtx.Dp(0)
			defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, rr).Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, t.border)
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top:    8,
				Bottom: 8,
				Left:   8,
				Right:  4,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				inputLayout := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					editor := material.Editor(th.Material(), &t.textEditor, t.Placeholder)
					editor.Color = theme.LightBlue
					editor.HintColor = theme.LightBlue
					editor.SelectionColor = th.TextSelectionColor
					return editor.Layout(gtx)
				})
				widgets := []layout.FlexChild{inputLayout}
				spacing := layout.SpaceBetween
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: spacing}.Layout(gtx, widgets...)
			})
		},
	)

}
