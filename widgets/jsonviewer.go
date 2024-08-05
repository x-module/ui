package widgets

import (
	"encoding/json"
	"fmt"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/theme"
	"strings"
)

type JsonViewer struct {
	data string

	lines []string

	selectables []*widget.Selectable

	list *widget.List
}

func NewJsonViewer() *JsonViewer {
	return &JsonViewer{
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

func (j *JsonViewer) SetData(data any) {

	// 使用MarshalIndent序列化map，生成格式化的JSON字符串
	// 第二个参数是每一行输出的前缀（通常为空）
	// 第三个参数是每一级缩进的字符串，这里使用4个空格作为缩进
	formattedJSON, _ := json.MarshalIndent(data, "", "    ")
	jsonData := string(formattedJSON)
	j.data = jsonData
	j.lines = strings.Split(jsonData, "\n")
	j.selectables = make([]*widget.Selectable, len(j.lines))
	for i := range j.selectables {
		j.selectables[i] = &widget.Selectable{}
	}
}

func (j *JsonViewer) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	border := widget.Border{
		Color:        theme.BorderColor,
		Width:        unit.Dp(1),
		CornerRadius: unit.Dp(4),
	}

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(3).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return material.List(theme.Material(), j.list).Layout(gtx, len(j.lines), func(gtx layout.Context, i int) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							l := material.Label(theme.Material(), theme.TextSize, fmt.Sprintf("%d", i+1))
							l.Font.Weight = font.Medium
							l.Color = theme.TextColor
							l.SelectionColor = theme.TextSelectionColor
							l.Alignment = text.End
							return l.Layout(gtx)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							l := material.Label(theme.Material(), theme.TextSize, j.lines[i])
							l.State = j.selectables[i]
							l.SelectionColor = theme.TextSelectionColor
							l.TextSize = unit.Sp(14)
							l.Font.Weight = font.Medium
							l.Color = theme.TextColor
							l.Alignment = text.End

							return l.Layout(gtx)
						})
					}),
				)
			})
		})
	})
}
