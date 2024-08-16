package widgets

import (
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
	"image"
	"time"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Notification struct {
	Text  string
	EndAt time.Time
}

type Notif struct {
	// Text is the text to display in the notification.
	Text string
	// Duration is the duration to display the notification.
	Duration time.Duration
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
				bd := material.Body1(theme.Material(), n.Text)
				bd.Color = resource.NotificationTextWhiteColor
				return bd.Layout(gtx)
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
