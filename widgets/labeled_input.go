package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/chapar-rest/chapar/ui/chapartheme"
)

type LabeledInput struct {
	Label          string
	SpaceBetween   int
	MinEditorWidth unit.Dp
	MinLabelWidth  unit.Dp
	Editor         *widget.Editor
	Hint           string
}

func (l *LabeledInput) Layout(gtx layout.Context, theme *chapartheme.Theme) layout.Dimensions {
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
				Color:        theme.BorderColor,
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
