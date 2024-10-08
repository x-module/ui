package widgets

import (
	"github.com/x-module/ui/theme"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const (
	IconPositionStart = 0
	IconPositionEnd   = 1
)

type Input struct {
	textEditor widget.Editor
	Icon       *widget.Icon
	iconClick  widget.Clickable

	IconPosition int

	Text        string
	Placeholder string

	size image.Point

	onIconClick  func()
	onTextChange func(text string)
	borderColor  color.NRGBA

	showPassword bool
}

func NewInput(text, placeholder string) *Input {
	t := &Input{
		textEditor:  widget.Editor{},
		Text:        text,
		Placeholder: placeholder,
		onIconClick: func() {},
	}
	t.textEditor.SetText(text)
	t.textEditor.SingleLine = true
	return t
}
func (t *Input) Password() {
	t.textEditor.Mask = '*'
	t.Icon, _ = widget.NewIcon(icons.ActionVisibilityOff)
	t.IconPosition = IconPositionEnd
	t.showPassword = false
}

func (t *Input) SetText(text string) {
	t.textEditor.SetText(text)
}

func (t *Input) SetIcon(icon *widget.Icon, position int) {
	t.Icon = icon
	t.IconPosition = position
}

func (t *Input) SetWidth(width int) {
	t.size.X = width
}

func (t *Input) SetBorderColor(color color.NRGBA) {
	t.borderColor = color
}

func (t *Input) SetOnTextChange(f func(text string)) {
	t.onTextChange = f
}

func (t *Input) SetOnIconClick(f func()) {
	t.onIconClick = f
}

func (t *Input) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	borderColor := th.BorderColor
	if gtx.Source.Focused(&t.textEditor) {
		borderColor = th.BorderColorFocused
	}

	cornerRadius := unit.Dp(4)
	border := widget.Border{
		Color:        borderColor,
		Width:        unit.Dp(1),
		CornerRadius: cornerRadius,
	}

	leftPadding := unit.Dp(8)
	if t.Icon != nil && t.IconPosition == IconPositionStart {
		leftPadding = unit.Dp(0)
	}

	for {
		event, ok := t.textEditor.Update(gtx)
		if !ok {
			break
		}
		if _, ok := event.(widget.ChangeEvent); ok {
			if t.onTextChange != nil {
				t.onTextChange(t.textEditor.Text())
			}
		}
	}

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		if t.size.X == 0 {
			t.size.X = gtx.Constraints.Min.X
		}

		gtx.Constraints.Min = t.size
		return layout.Inset{
			Top:    8,
			Bottom: 8,
			Left:   leftPadding,
			Right:  4,
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			inputLayout := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				editor := material.Editor(th.Material(), &t.textEditor, t.Placeholder)
				editor.Color = th.TextColor
				editor.HintColor = theme.LightBlue
				editor.SelectionColor = th.TextSelectionColor
				return editor.Layout(gtx)
			})
			widgets := []layout.FlexChild{inputLayout}
			spacing := layout.SpaceBetween
			if t.Icon != nil {
				iconLayout := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					clk := &widget.Clickable{}
					if t.onIconClick != nil {
						clk = &t.iconClick
						if t.iconClick.Clicked(gtx) {
							t.onIconClick()
							if !t.showPassword {
								t.textEditor.Mask = 0
								t.Icon = ActionVisibilityIcon
								t.showPassword = true
							} else {
								t.textEditor.Mask = '*'
								t.Icon = ActionVisibilityOffIcon
								t.showPassword = false
							}
						}
					}
					b := ButtonWithIcon(th, clk, t.Icon, IconPositionStart, "", 0)
					b.Inset = layout.Inset{Left: unit.Dp(8), Right: unit.Dp(2), Top: unit.Dp(2), Bottom: unit.Dp(2)}
					return b.Layout(gtx)
				})
				if t.IconPosition == IconPositionEnd {
					widgets = []layout.FlexChild{inputLayout, iconLayout}
				} else {
					widgets = []layout.FlexChild{iconLayout, inputLayout}
					spacing = layout.SpaceEnd
				}
			}
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: spacing}.Layout(gtx, widgets...)
		})
	})
}
