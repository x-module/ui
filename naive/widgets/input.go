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
	icon      *widget.Icon
	iconClick widget.Clickable

	theme *theme.Theme
	// size image.Point

	width unit.Dp

	showPassword bool
	onIconClick  func()
}

func NewInput(theme *theme.Theme, hint string, text ...string) *Input {
	t := &Input{
		theme:  theme,
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
func NewTextArea(theme *theme.Theme, hint string, text ...string) *Input {
	t := &Input{
		theme:  theme,
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

func (i *Input) SetWidth(width unit.Dp) {
	i.width = width
}

func (i *Input) SetOnIconClick(f func()) {
	i.onIconClick = f
}
func (i *Input) Password() {
	i.editor.Mask = '*'
	i.icon, _ = widget.NewIcon(icons.ActionVisibilityOff)
	// t.IconPosition = IconPositionEnd
	i.showPassword = false
}
func (i *Input) SetIcon(icon *widget.Icon) {
	i.icon = icon
}

// SetRadius 设置radius
func (i *Input) SetRadius(radius unit.Dp) {
	i.radius = radius
}
func (i *Input) ReadOnly() {
	i.editor.ReadOnly = true
}
func (i *Input) SetBefore(before layout.Widget) {
	i.before = before
}
func (i *Input) SetAfter(after layout.Widget) {
	i.after = after
}

func (i *Input) SetSize(size resource.Size) {
	i.size = size
}

func (i *Input) SetText(text string) {
	i.editor.SetText(text)
}
func (i *Input) GetText() string {
	return i.editor.Text()
}
func (i *Input) update(gtx layout.Context, th *theme.Theme) {
	disabled := gtx.Source == (input.Source{})
	for {
		ev, ok := i.click.Update(gtx.Source)
		if !ok {
			break
		}
		switch ev.Kind {
		case gesture.KindPress:
			gtx.Execute(key.FocusCmd{Tag: &i.editor})
		default:

		}
	}
	i.state = inactive
	if i.click.Hovered() && !disabled {
		i.state = hovered
	}
	// if t.editor.Len() > 0 {
	// 	t.state = activated
	// }
	if gtx.Source.Focused(&i.editor) && !disabled {
		i.state = focused
	}

	i.bgColor = resource.DefaultBgGrayColor

	if i.editor.ReadOnly {
		return
	}

	switch i.state {
	case inactive:
		i.borderColor = resource.DefaultBorderGrayColor
	case hovered:
		i.borderColor = resource.HoveredBorderBlueColor
	case focused:
		i.bgColor = resource.FocusedBgColor
		i.borderColor = resource.FocusedBorderBlueColor
	case activated:
		i.borderColor = resource.DefaultBorderGrayColor
	}
}

func (i *Input) Layout(gtx layout.Context) layout.Dimensions {
	if i.width > 0 {
		gtx.Constraints.Max.X = gtx.Dp(i.width)
	} else {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
	}
	i.update(gtx, i.theme)
	gtx.Constraints.Min.X = gtx.Constraints.Max.X
	gtx.Constraints.Min.Y = 0
	macro := op.Record(gtx.Ops)
	dims := i.layout(gtx, i.theme)
	call := macro.Stop()
	defer pointer.PassOp{}.Push(gtx.Ops).Pop()
	defer clip.Rect(image.Rectangle{Max: dims.Size}).Push(gtx.Ops).Pop()
	i.click.Add(gtx.Ops)
	event.Op(gtx.Ops, &i.editor)
	call.Add(gtx.Ops)
	return dims
}

func (i *Input) layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	border := widget.Border{
		Color:        i.borderColor,
		Width:        unit.Dp(1),
		CornerRadius: i.radius,
	}
	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rr := gtx.Dp(i.radius)
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, rr).Push(gtx.Ops).Pop()
				paint.Fill(gtx.Ops, i.bgColor)
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			func(gtx layout.Context) layout.Dimensions {
				return i.size.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					inputLayout := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {

						gtx.Constraints.Max.X = gtx.Dp(i.width)
						editor := material.Editor(th.Material(), &i.editor, i.hint)
						editor.HintColor = resource.HintTextColor
						editor.SelectionColor = resource.TextSelectionColor

						gtx.Constraints.Min.Y = gtx.Dp(i.size.Height) // 设置最小高度为 100dp
						gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
						editor.TextSize = i.size.TextSize

						if i.height > 0 {
							gtx.Constraints.Min.Y = gtx.Dp(i.height)      // 设置最小高度为 100dp
							gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
						}
						if i.editor.ReadOnly {
							editor.Color = resource.HintTextColor
						} else {
							editor.Color = resource.DefaultTextWhiteColor
						}
						return editor.Layout(gtx)
					})

					var widgets []layout.FlexChild
					if i.before != nil {
						widgets = append(widgets, layout.Rigid(i.before))
					}
					widgets = append(widgets, inputLayout)
					if i.icon != nil {
						iconLayout := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							i.icon = widgets2.ActionVisibilityIcon
							if i.iconClick.Clicked(gtx) {
								if i.onIconClick != nil {
									i.onIconClick()
								}
								if !i.showPassword {
									i.editor.Mask = 0
									i.icon = widgets2.ActionVisibilityIcon
									i.showPassword = true
								} else {
									i.editor.Mask = '*'
									i.icon = widgets2.ActionVisibilityOffIcon
									i.showPassword = false
								}
							}
							return i.iconClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return i.icon.Layout(gtx, resource.IconGrayColor)
							})
						})
						widgets = append(widgets, iconLayout)
					} else {
						if i.after != nil {
							widgets = append(widgets, layout.Rigid(i.after))
						}
					}
					spacing := layout.SpaceBetween
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: spacing}.Layout(gtx, widgets...)
				})
			},
		)
	})
}
