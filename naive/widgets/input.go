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
	widgets2 "github.com/x-module/ui/widgets"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image"
)

type Input struct {
	CommonWidget
	editor    widget.Editor
	height    unit.Dp
	before    layout.Widget
	after     layout.Widget
	Icon      *widget.Icon
	iconClick widget.Clickable

	showPassword bool
}

func NewInput(hint string, text ...string) *Input {
	t := &Input{
		editor: widget.Editor{},
	}
	t.size = resource.Medium
	t.hint = hint
	t.radius = resource.DefaultRadiusSize
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
	t.size = resource.Medium
	t.hint = hint
	t.radius = resource.DefaultRadiusSize
	if len(text) > 0 {
		t.editor.SetText(text[0])
	}
	t.editor.SingleLine = false
	return t
}

func (t *Input) Password() {
	t.editor.Mask = '*'
	t.Icon, _ = widget.NewIcon(icons.ActionVisibilityOff)
	// t.IconPosition = IconPositionEnd
	t.showPassword = false
}

// SetRadius 设置radius
func (t *Input) SetRadius(radius unit.Dp) {
	t.radius = radius
}
func (t *Input) ReadOnly() {
	t.editor.ReadOnly = true
}
func (t *Input) SetBefore(before layout.Widget) {
	t.before = before
}
func (t *Input) SetAfter(after layout.Widget) {
	t.after = after
}

func (t *Input) SetSize(size resource.Size) {
	t.size = size
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

	t.bgColor = resource.DefaultBgBlueColor

	if t.editor.ReadOnly {
		return
	}

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
				return t.size.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					inputLayout := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						editor := material.Editor(th.Material(), &t.editor, t.hint)
						editor.HintColor = resource.HintTextColor
						editor.SelectionColor = resource.TextSelectionColor

						gtx.Constraints.Min.Y = gtx.Dp(t.size.Height) // 设置最小高度为 100dp
						gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
						editor.TextSize = t.size.TextSize

						if t.height > 0 {
							gtx.Constraints.Min.Y = gtx.Dp(t.height)      // 设置最小高度为 100dp
							gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
						}
						if t.editor.ReadOnly {
							editor.Color = resource.HintTextColor
						} else {
							editor.Color = resource.DefaultTextWhiteColor
						}
						return editor.Layout(gtx)
					})

					var widgets []layout.FlexChild
					if t.before != nil {
						widgets = append(widgets, layout.Rigid(t.before))
					}
					widgets = append(widgets, inputLayout)

					if t.Icon != nil {
						iconLayout := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							clk := &widget.Clickable{}
							clk = &t.iconClick
							if t.iconClick.Clicked(gtx) {
								if !t.showPassword {
									t.editor.Mask = 0
									t.Icon = widgets2.ActionVisibilityIcon
									t.showPassword = true
								} else {
									t.editor.Mask = '*'
									t.Icon = widgets2.ActionVisibilityOffIcon
									t.showPassword = false
								}
							}
							b := DefaultButton(th, clk, "", unit.Dp(30))
							b.SetIcon(t.Icon, widgets2.IconPositionStart)
							b.Inset = layout.Inset{Left: unit.Dp(8), Right: unit.Dp(2), Top: unit.Dp(2), Bottom: unit.Dp(2)}
							return b.Layout(gtx)
						})
						widgets = append(widgets, iconLayout)
					} else {
						if t.after != nil {
							widgets = append(widgets, layout.Rigid(t.after))
						}
					}
					spacing := layout.SpaceBetween
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: spacing}.Layout(gtx, widgets...)
				})
			},
		)
	})
}
