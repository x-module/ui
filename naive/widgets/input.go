package widgets

import (
	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/input"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
	"image"
)

type Input struct {
	CommonWidget
	editor widget.Editor
	height unit.Dp
}

func NewInput(hint string, text ...string) *Input {
	t := &Input{
		editor: widget.Editor{},
	}
	t.hint = hint
	t.radius = 4
	if len(text) > 0 {
		t.editor.SetText(text[0])
	}
	t.editor.SingleLine = true
	return t
}
func NewTextArea(hint string, text ...string) *Input {
	t := &Input{
		editor: widget.Editor{},
		height: unit.Dp(100),
	}
	t.hint = hint
	t.radius = 4
	if len(text) > 0 {
		t.editor.SetText(text[0])
	}
	t.editor.SingleLine = false
	return t
}

func (t *Input) Password() {
	t.editor.Mask = '*'
}

// SetRadius 设置radius
func (t *Input) SetRadius(radius unit.Dp) {
	t.radius = radius
}

func (t *Input) SetText(text string) {
	t.editor.SetText(text)
}

func (t *Input) update(gtx layout.Context, th *theme.Theme) {
	disabled := gtx.Source == (input.Source{})
	for {
		ev, ok := t.click.Update(gtx.Source)
		if !ok {
			break
		}
		switch ev.Kind {
		case gesture.KindPress:
			gtx.Execute(key.FocusCmd{Tag: &t.editor})
		default:

		}
	}
	t.state = inactive
	if t.click.Hovered() && !disabled {
		t.state = hovered
	}
	// if t.editor.Len() > 0 {
	// 	t.state = activated
	// }
	if gtx.Source.Focused(&t.editor) && !disabled {
		t.state = focused
	}

	t.bgColor = resource.DefaultBgColor
	switch t.state {
	case inactive:
		t.borderColor = resource.DefaultBorderBgColor
	case hovered:
		t.borderColor = resource.HoveredBorderColor
	case focused:
		t.bgColor = resource.FocusedBgColor
		t.borderColor = resource.FocusedBorderColor
	case activated:
		t.borderColor = resource.DefaultBorderBgColor
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
	event.Op(gtx.Ops, &t.editor)
	call.Add(gtx.Ops)
	return dims
}

func (t *Input) layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	border := widget.Border{
		Color:        t.borderColor,
		Width:        unit.Dp(1),
		CornerRadius: t.radius,
	}
	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rr := gtx.Dp(t.radius)
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, rr).Push(gtx.Ops).Pop()
				paint.Fill(gtx.Ops, t.bgColor)
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
						if t.height > 0 {
							gtx.Constraints.Min.Y = gtx.Dp(t.height)      // 设置最小高度为 100dp
							gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
						}
						editor := material.Editor(th.Material(), &t.editor, t.hint)
						editor.Color = resource.TextColor
						editor.HintColor = resource.HintTextColor
						editor.SelectionColor = resource.TextSelectionColor
						return editor.Layout(gtx)
					})
					widgets := []layout.FlexChild{inputLayout}
					spacing := layout.SpaceBetween
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: spacing}.Layout(gtx, widgets...)
				})
			},
		)
	})
}
