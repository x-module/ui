package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
	theme2 "github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
)

//type MenuState struct {
//	OptionList layout.List
//	Options    []func(gtx C) D
//}

func main() {
	//w := new(app.Window)
	var th = material.NewTheme()
	theme := theme2.New(th, true)
	menuItemOptions := []*widgets.ListMenuOption{
		{
			Text:  "属性",
			Value: "shuxing",
		},
		{
			Text:  "退出",
			Value: "exit",
		},
	}
	menu := widgets.NewListMenu(theme, "头 大", menuItemOptions)
	menu.Clicked(func(key int, menu string) {
		fmt.Println("---click,key:", key, " menu:", menu)
		fmt.Println("---click,key:", key, " menu:", menu)
	})
	var ops op.Ops
	confirm := widgets.NewConfirm(theme)
	confirm.Confirm(func() {
		fmt.Println("confirm")
	})

	confirm.Cancel(func() {
		fmt.Println("cancel")
	})
	go func() {
		w := new(app.Window)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				//label := material.Label(th, unit.Sp(20), "adfasdfasd").Layout(gtx)
				confirm.Message("确定退出当前账号吗?")
				confirm.Layout(gtx)
				//layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				//	layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				//		return menu.Layout(gtx)
				//	}),
				//	layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				//		return material.Label(th, unit.Sp(20), "adfasdfasd").Layout(gtx)
				//	}), layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				//		return material.Label(th, unit.Sp(20), "adfasdfasd").Layout(gtx)
				//	}), layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				//		return material.Label(th, unit.Sp(20), "adfasdfasd").Layout(gtx)
				//	}),
				//)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
