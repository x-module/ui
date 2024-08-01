package widgets

import (
	"github.com/x-module/ui/theme"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type TextArea struct {
	textEditor  widget.Editor
	Placeholder string
	size        image.Point
	borderColor color.NRGBA
}

func NewTextArea(text, placeholder string) *TextArea {
	t := &TextArea{
		textEditor:  widget.Editor{},
		Placeholder: placeholder,
	}

	t.textEditor.SetText(text)
	t.textEditor.SingleLine = true
	return t
}

// 可能存在bug，弃用，如需要可程序new一个对象
// func (t *TextArea) SetText(text string) {
// 	t.textEditor.SetText(text)
// }

func (t *TextArea) SetMinWidth(width int) {
	t.size.X = width
}
func (t *TextArea) Text() string {
	return t.textEditor.Text()
}

func (t *TextArea) SetBorderColor(color color.NRGBA) {
	t.borderColor = color
}

func (t *TextArea) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
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

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		if t.size.X == 0 {
			t.size.X = gtx.Constraints.Min.X
		}

		gtx.Constraints.Min = t.size
		return layout.Inset{
			Top:    4,
			Bottom: 4,
			Left:   leftPadding,
			Right:  4,
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			inputLayout := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(100))  // 设置最小高度为 100dp
				gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
				editor := material.Editor(th.Material(), &t.textEditor, t.Placeholder)
				editor.Color = th.TextColor
				editor.HintColor = theme.LightBlue
				editor.SelectionColor = th.TextSelectionColor
				return editor.Layout(gtx)
			})
			widgets := []layout.FlexChild{inputLayout}
			spacing := layout.SpaceBetween
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: spacing}.Layout(gtx, widgets...)
		})
	})
}
