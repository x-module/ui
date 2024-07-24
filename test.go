package main

//
// import (
//	"fmt"
//	tips2 "game.test.client/ui/pages/tips"
//	"game.test.client/widgets"
//	"gioui.org/app"
//	"gioui.org/layout"
//	"gioui.org/op"
//	"gioui.org/unit"
//	"gioui.org/widget"
//	"gioui.org/widget/material"
//	"github.com/x-module/ui/theme"
//	"github.com/chapar-rest/chapar/ui/explorer"
//	widgets2 "github.com/chapar-rest/chapar/ui/widgets"
//	"time"
// )
//
// var modal *widgets.Modal
// var clickable *widget.Clickable
// var th = material.NewTheme()
// var theme = chapartheme.New(th, true)
// var Prompt *widgets.Prompt = widgets.NewPrompt("", "", "")
// var NotificationController = &widgets.Notification{}
// var badgeInput *widgets.BadgeInput
// var tips *tips2.Tips
//
// var binaryFile *widgets2.BinaryFile
// var flatButton *widgets.FlatButton
//
// func main() {
//	modal = widgets.NewModal(theme)
//	modal.SetTitle("玩家登录")
//	modal.SetHeight(200)
//	clickable = new(widget.Clickable)
//	binaryFile = widgets2.NewBinaryFile("/Users/lijin/go/gioui/go.mod")
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
//	tips = tips2.New()
//
//	flatButton = &widgets.FlatButton{
//		Icon:         widgets.ArrowDropDownIcon,
//		IconPosition: 1,
//		SpaceBetween: unit.Dp(10),
//		Text:         "aaa",
//		Clickable:    clickable,
//
//		BackgroundPadding: unit.Dp(1),
//		CornerRadius:      5,
//		MinWidth:          unit.Dp(120),
//		BackgroundColor:   theme.SideBarBgColor,
//		TextColor:         theme.Palette.Fg,
//		ContentPadding:    unit.Dp(5),
//		MarginRight:       10,
//	}
//
//	go func() {
//		w := new(app.Window)
//		var ops op.Ops
//		for {
//			e := w.Event()
//			switch e := e.(type) {
//			case app.DestroyEvent:
//				panic(e.Err)
//			case app.FrameEvent:
//				gtx := app.NewContext(&ops, e)
//				//tips.Layout(gtx, theme)
//				//binaryFile.Layout(gtx, theme)
//				modalLayout(gtx, th)
//				e.Frame(gtx.Ops)
//			}
//		}
//	}()
//	app.Main()
// }
//
// func data(gtx layout.Context, th *material.Theme) layout.Dimensions {
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
// }
//
// func Send(text string, duration time.Duration) {
//	fmt.Println("time.Now():", time.Now().Format(time.DateTime))
//	fmt.Println("time.Now().Add(duration):", time.Now().Add(duration).Format(time.DateTime))
//	NotificationController.EndAt = time.Now().Add(duration)
//	NotificationController.Text = text
// }
//
// func ShowPrompt(title, content, modalType string, onSubmit func(selectedOption string, remember bool), options ...widgets.Option) {
//	Prompt.Type = modalType
//	Prompt.Title = title
//	Prompt.Content = content
//	Prompt.Content = content
//
//	Prompt.SetOptions(options...)
//	Prompt.WithoutRememberBool()
//	Prompt.SetOnSubmit(onSubmit)
//	Prompt.Show()
// }
//
// var checkBox = &widget.Bool{
//	Value: true,
// }
// var explore = &explorer.Explorer{}
//
// func modalLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
//	if clickable.Clicked(gtx) {
//		fmt.Println("asdfasdfasdfasd")
//		fmt.Println("asdfasdfasdfasd")
//		fmt.Println("asdfasdfasdfasd")
//		fmt.Println("asdfasdfasdfasd")
//		//modal.Display(func(gtx layout.Context) layout.Dimensions {
//		//	//return material.Label(th, unit.Sp(16), "modalState.Content").Layout(gtx)
//		//	return data(gtx, th)
//		//})
//		//Send("success", 10*time.Second)
//
//		//options := []widgets.Option{
//		//	{
//		//		Text: "aaa",
//		//		Icon: widgets.ArrowDropDownIcon,
//		//	},
//		//	{
//		//		Text: "bbbb",
//		//		Icon: widgets.ExpandIcon,
//		//	},
//		//}
//		//
//		//ShowPrompt("title", "content", widgets.ModalTypeErr, func(selectedOption string, remember bool) {
//		//}, options...)
//
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
//
//					//return widgets.Button(th, clickable, widgets.ArrowDropDownIcon, 1, "click").Layout(gtx, theme)
//					//func RadioButton(th *material.Theme, group *widget.Enum, key, label string) RadioButtonStyle {
//					//e := &widget.Enum{
//					//	Value: "test-value",
//					//}
//					//return widgets.RadioButton(th, e, "key", "value").Layout(gtx)
//
//					//return widgets.CheckBox(th, checkBox, "Remember me").Layout(gtx)
//
//					//return material.Button(th, clickable, "Close").Layout(gtx)
//					//return data(gtx, th)
//					//return badgeInput.Layout(gtx, theme)
//					//return widgets.NewFileSelector("aaa.txt", explore, "*.go").Layout(gtx, theme)
//
//					//return flatButton.Layout(gtx, theme)
//
//					return widgets2.SaveButtonLayout(gtx, theme, clickable)
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
// }
