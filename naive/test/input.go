package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/naive/widgets"
	"github.com/x-module/ui/theme"
	widgets2 "github.com/x-module/ui/widgets"
)

func main() {
	var username *widgets.Input
	var password *widgets.Input
	var password2 *widgets.Input
	var age *widgets.Input
	var profile *widgets.Input
	// var clickable widget.Clickable
	var th = theme.New(material.NewTheme(), true)

	// w := new(app.Window)
	var ops op.Ops
	username = widgets.NewInput("请输入名称...")
	age = widgets.NewInput("请输入年龄...", "á3452345234523452345")
	password = widgets.NewInput("请输入密码...")
	password2 = widgets.NewInput("请输入确认密码===...")
	profile = widgets.NewTextArea("请输入属性...")

	username.SetSize(resource.Tiny)
	password.SetSize(resource.Small)
	password.SetRadius(unit.Dp(8))
	password2.SetSize(resource.Medium)
	age.SetSize(resource.Large)

	age.ReadOnly()
	// password2.SetAfter(func(gtx layout.Context) layout.Dimensions {
	// 	return widgets2.NavigationSubdirectoryArrowRightIcon.Layout(gtx, resource.IconColor)
	// })
	password2.SetBefore(func(gtx layout.Context) layout.Dimensions {
		return widgets2.ActionPermIdentityIcon.Layout(gtx, resource.IconGrayColor)
	})

	password2.Password()
	go func() {
		w := new(app.Window)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, resource.WindowBgColor, rect.Op())
				// =============================================
				// ==============================================
				layout.Stack{Alignment: layout.Center}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return username.Layout(gtx, th)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return password.Layout(gtx, th)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return password2.Layout(gtx, th)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return age.Layout(gtx, th)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return profile.Layout(gtx, th)
							}),
						)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
