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
	textEditor  *widget.Editor
	theme       *material.Theme
	Text        string
	Placeholder string

	size image.Point

	onTextChange func(text string)
	borderColor  color.NRGBA
}

func NewTextArea(text, placeholder string) *TextArea {
	// clickable = new(widget.Clickable)

	t := &TextArea{
		textEditor:  &widget.Editor{},
		Text:        text,
		Placeholder: placeholder,
	}
	t.textEditor.SetText(text)
	t.textEditor.SingleLine = false
	return t
}

func (t *TextArea) SetText(text string) {
	t.textEditor.SetText(text)
}

func (t *TextArea) SetMinWidth(width int) {
	t.size.X = width
}

func (t *TextArea) SetBorderColor(color color.NRGBA) {
	t.borderColor = color
}

func (t *TextArea) SetOnTextChange(f func(text string)) {
	t.onTextChange = f
}

func (t *TextArea) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	borderColor := theme.BorderColor
	if gtx.Source.Focused(&t.textEditor) {
		borderColor = theme.BorderColorFocused
	}

	cornerRadius := unit.Dp(4)
	border := widget.Border{
		Color:        borderColor,
		Width:        unit.Dp(1),
		CornerRadius: cornerRadius,
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
			t.size.X = gtx.Constraints.Max.X
		}
		gtx.Constraints.Min = t.size
		return layout.UniformInset(4).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(8).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// 设置 Editor 的高度
				gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(100))  // 设置最小高度为 100dp
				gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
				return material.Editor(theme.Material(), t.textEditor, t.Placeholder).Layout(gtx)
			})
		})
	})
}
