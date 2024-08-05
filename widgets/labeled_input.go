package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
)

type LabeledInput struct {
	Label        string
	theme        *theme.Theme
	SpaceBetween int
	editorWidth  unit.Dp
	labelWidth   unit.Dp
	Editor       *widget.Editor
	Hint         string
}

func NewLabeledInput(theme *theme.Theme, label string, hint string) *LabeledInput {
	return &LabeledInput{
		theme: theme,
		Label: label,
		Editor: &widget.Editor{
			SingleLine: true,
			Submit:     true,
		},
		Hint:         hint,
		SpaceBetween: 3,
		editorWidth:  unit.Dp(300),
		labelWidth:   unit.Dp(50),
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
func (l *LabeledInput) SetEditorWidth(width unit.Dp) *LabeledInput {
	l.editorWidth = width
	return l
}
func (l *LabeledInput) SetLabelWidth(width unit.Dp) *LabeledInput {
	l.labelWidth = width
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
func (l *LabeledInput) Layout(gtx layout.Context) layout.Dimensions {
	borderColor := l.theme.BorderColor
	if gtx.Source.Focused(l.Editor) {
		borderColor = l.theme.BorderColorFocused
	}
	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Right: unit.Dp(l.SpaceBetween)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = gtx.Dp(l.labelWidth)
				return Label(l.theme, l.Label).Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if l.editorWidth > 0 {
				gtx.Constraints.Min.X = gtx.Dp(l.editorWidth)
			} else {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
			}
			return widget.Border{
				Color:        borderColor,
				Width:        unit.Dp(1),
				CornerRadius: unit.Dp(4),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					editor := material.Editor(l.theme.Material(), l.Editor, l.Hint)
					editor.Color = l.theme.TextColor
					editor.HintColor = theme.LightBlue
					editor.SelectionColor = l.theme.TextSelectionColor
					return editor.Layout(gtx)
				})
			})
		}),
	)
}
