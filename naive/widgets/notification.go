package widgets

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/utils"
	"github.com/x-module/ui/widgets"
	"image"
	"time"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type Notification struct {
	Text      string
	EndAt     time.Time
	closeIcon *IconButton
}

func NewNotification() *Notification {
	notice := &Notification{
		closeIcon: &IconButton{
			Icon:                 widgets.CloseIcon,
			Color:                resource.IconGrayColor,
			BackgroundColor:      resource.NotificationBgColor,
			BackgroundColorHover: utils.Hovered(resource.ModalBgGrayColor),
			Size:                 resource.DefaultIconSize,
			Clickable:            &widget.Clickable{},
		},
	}
	notice.closeIcon.OnClick = func() {
		notice.Text = ""
	}
	return notice
}

func (n *Notification) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	if n.Text == "" || n.EndAt == (time.Time{}) || time.Now().After(n.EndAt) {
		return layout.Dimensions{}
	}

	// set max width for the notification
	gtx.Constraints.Max.X = gtx.Dp(300)
	// // set max height for the notification
	gtx.Constraints.Max.Y = gtx.Dp(40)

	// utils.ColorBackground(gtx, gtx.Constraints.Max, resource.GreenColor)

	macro := op.Record(gtx.Ops)
	dim := layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 8).Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, resource.NotificationBgColor)
			return layout.Dimensions{Size: gtx.Constraints.Min}
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						bd := material.Body1(theme.Material(), n.Text)
						bd.Color = resource.NotificationTextWhiteColor
						return bd.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return n.closeIcon.Layout(gtx)
					}),
				)
			})
		},
	)
	call := macro.Stop()
	return layout.NE.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Top: unit.Dp(20), Right: unit.Dp(20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			call.Add(gtx.Ops)
			return dim
		})
	})
}
