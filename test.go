package main

//
//import (
//	"fmt"
//	"game.test.client/widgets"
//	"gioui.org/app"
//	"gioui.org/layout"
//	"gioui.org/op"
//	"gioui.org/unit"
//	"gioui.org/widget"
//	"gioui.org/widget/material"
//	"github.com/chapar-rest/chapar/ui/chapartheme"
//	"time"
//)
//
//var modal *widgets.Modal
//var clickable *widget.Clickable
//
//var theme = chapartheme.New(material.NewTheme(), true)
//var Prompt *widgets.Prompt = widgets.NewPrompt("", "", "")
//var NotificationController = &widgets.Notification{}
//var badgeInput *widgets.BadgeInput
//
//func main() {
//	modal = widgets.NewModal(theme)
//	modal.SetTitle("玩家登录")
//	modal.SetHeight(200)
//	clickable = new(widget.Clickable)
//
//	items := []*widgets.BadgeInputItem{
//		{
//			Identifier: "bbb",
//			Value:      "bbb",
//		},
//		{
//			Identifier: "aaa",
//			Value:      "aaa",
//		},
//	}
//	badgeInput = widgets.NewBadgeInput(items...)
//
//	go func() {
//		w := new(app.Window)
//		th := material.NewTheme()
//		var ops op.Ops
//		for {
//			e := w.Event()
//			switch e := e.(type) {
//			case app.DestroyEvent:
//				panic(e.Err)
//			case app.FrameEvent:
//				gtx := app.NewContext(&ops, e)
//				modalLayout(gtx, th)
//				e.Frame(gtx.Ops)
//			}
//		}
//	}()
//	app.Main()
//}
//
//func data(gtx layout.Context, th *material.Theme) layout.Dimensions {
//	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
//		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
//			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
//				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
//					label := material.Label(th, unit.Sp(16), "用户名:")
//					label.Color = theme.Palette.Fg
//					return label.Layout(gtx)
//				}),
//				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
//					return widgets.NewTextField("用户名", "请输入密码").Layout(gtx, theme)
//				}),
//			)
//		}),
//		layout.Rigid(layout.Spacer{Height: 20}.Layout),
//		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
//			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
//				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
//					//return widgets.NewTextField("用户名", "请输入密码").Layout(gtx, theme)
//					input := widgets.LabeledInput{
//						Label:        "用户名:",
//						SpaceBetween: 10,
//						Editor:       new(widget.Editor),
//						Hint:         "请输入密码",
//					}
//					return input.Layout(gtx, theme)
//				}),
//			)
//		}),
//	)
//}
//
//func Send(text string, duration time.Duration) {
//	fmt.Println("time.Now():", time.Now().Format(time.DateTime))
//	fmt.Println("time.Now().Add(duration):", time.Now().Add(duration).Format(time.DateTime))
//	NotificationController.EndAt = time.Now().Add(duration)
//	NotificationController.Text = text
//}
//
//func ShowPrompt(title, content, modalType string, onSubmit func(selectedOption string, remember bool), options ...widgets.Option) {
//	Prompt.Type = modalType
//	Prompt.Title = title
//	Prompt.Content = content
//	Prompt.Content = content
//
//	Prompt.SetOptions(options...)
//	Prompt.WithoutRememberBool()
//	Prompt.SetOnSubmit(onSubmit)
//	Prompt.Show()
//}
//
//func modalLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
//	if clickable.Clicked(gtx) {
//		//modal.Display(func(gtx layout.Context) layout.Dimensions {
//		//	//return material.Label(th, unit.Sp(16), "modalState.Content").Layout(gtx)
//		//	return data(gtx, th)
//		//})
//		//Send("success", 10*time.Second)
//
//		options := []widgets.Option{
//			{
//				Text: "aaa",
//				Icon: widgets.ArrowDropDownIcon,
//			},
//			{
//				Text: "bbbb",
//				Icon: widgets.ExpandIcon,
//			},
//		}
//
//		ShowPrompt("title", "content", widgets.ModalTypeErr, func(selectedOption string, remember bool) {
//		}, options...)
//	}
//
//	//notification := widgets.Notification{
//	//	Text:  "test",
//	//	EndAt: time.Now().Add(180),
//	//}
//	//
//	//return notification.Layout(gtx, theme)
//
//	return layout.Stack{}.Layout(gtx,
//		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
//			gtx.Constraints.Min = gtx.Constraints.Max
//			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
//				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
//					//return material.Button(th, clickable, "Close").Layout(gtx)
//					//return data(gtx, th)
//					return badgeInput.Layout(gtx, theme)
//				}),
//			)
//		}),
//		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
//			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
//				//return modal.Layout(gtx)
//				//return NotificationController.Layout(gtx, theme)
//				return Prompt.Layout(gtx, theme)
//			})
//		}),
//	)
//}
