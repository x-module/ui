package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
)

type LabeledInput struct {
	Label          string
	SpaceBetween   int
	MinEditorWidth unit.Dp
	MinLabelWidth  unit.Dp
	Editor         *widget.Editor
	Hint           string
}

func NewLabeledInput(label string, hint string) *LabeledInput {
	return &LabeledInput{
		Label: label,
		Editor: &widget.Editor{
			SingleLine: true,
			Submit:     true,
		},
		Hint:           hint,
		SpaceBetween:   3,
		MinEditorWidth: unit.Dp(300),
		MinLabelWidth:  unit.Dp(50),
	}
}

func (l *LabeledInput) SetText(text string) *LabeledInput {
	l.Editor.SetText(text)
	return l

}

func (l *LabeledInput) SetSpaceBetween(space int) *LabeledInput {
	l.SpaceBetween = space
	return l
}
func (l *LabeledInput) SetMinEditorWidth(width unit.Dp) *LabeledInput {
	l.MinEditorWidth = width
	return l
}
func (l *LabeledInput) SetMinLabelWidth(width unit.Dp) *LabeledInput {
	l.MinLabelWidth = width
	return l
}

func (l *LabeledInput) SetEditor(editor *widget.Editor) *LabeledInput {
	l.Editor = editor
	return l
}

// 获取editer的当前值
func (l *LabeledInput) GetText() string {
	return l.Editor.Text()
}
func (l *LabeledInput) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	borderColor := theme.BorderColor
	if gtx.Source.Focused(l.Editor) {
		borderColor = theme.BorderColorFocused
	}
	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Right: unit.Dp(l.SpaceBetween)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = gtx.Dp(l.MinLabelWidth)
				return material.Label(theme.Material(), theme.TextSize, l.Label).Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if l.MinEditorWidth > 0 {
				gtx.Constraints.Min.X = gtx.Dp(l.MinEditorWidth)
			} else {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
			}
			return widget.Border{
				Color:        borderColor,
				Width:        unit.Dp(1),
				CornerRadius: unit.Dp(4),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					editor := material.Editor(theme.Material(), l.Editor, l.Hint)
					editor.SelectionColor = theme.TextSelectionColor
					return editor.Layout(gtx)
				})
			})
		}),
	)
}
