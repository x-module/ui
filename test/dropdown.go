package test

import (
	"fmt"
	widget2 "game.test.client/widget"
	"game.test.client/widget/assets"
	"game.test.client/widget/values"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"os"
)

func init() {
	decredIcons, err := assets.Icons()
	if err != nil {
		fmt.Println("Error loading icons.", err.Error())
		panic("Error loading icons")
	}
	assets.DecredIcons = decredIcons
}

func main() {
	go func() {
		window := new(app.Window)
		err := runDropDown(window)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runDropDown(window *app.Window) error {
	var ops op.Ops
	theme := widget2.NewTheme(window)

	fmt.Println("Styles.ClickableStyle:", theme.Styles.ClickableStyle)
	fmt.Println("Styles.ClickableStyle:", theme.Styles.ClickableStyle)
	fmt.Println("Styles.ClickableStyle:", theme.Styles.ClickableStyle)

	var DropDownItems = []widget2.DropDownItem{
		{
			Text: "option-one",
			Icon: theme.Icons.WalletIcon,
		},
		{
			Text: "option-two",
			Icon: theme.Icons.WalletIcon,
		},
		{
			Text: "option-three",
			Icon: theme.Icons.WalletIcon,
		},
	}
	var (
		editor1 = widget.Editor{SingleLine: true}
		editor2 = widget.Editor{SingleLine: true}
	)

	dropDown := theme.DropDown(DropDownItems, 2, 0)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			widget2.DisplayOneDropdown(gtx, dropDown)
			UniformPadding(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{Alignment: layout.N}.Layout(gtx,
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						return theme.Editor(&editor1, "editor is focused").Layout(gtx)
					}),

					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top: values.MarginPadding60 * 2,
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return theme.EditorPassword(&editor2, "ssssssssss").Layout(gtx)
						})
					}),
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top: values.MarginPadding60,
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return dropDown.Layout(gtx, dropDown.Width, true)
						})
					}),
				)
			})

			// layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			// 	layout.Rigid(layout.Spacer{Height: 10}.Layout),
			// 	layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			//
			// 		return theme.Editor(&editor1, "editor is focused").Layout(gtx)
			// 	}),
			// 	layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// 		return dropDown.Layout(gtx, dropDown.Width, true)
			// 	}),
			// 	layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// 		// return theme.Editor(&editor2, "editor is focused").Layout(gtx, theme.Base)
			// 		return theme.EditorPassword(&editor2, "ssssssssss").Layout(gtx)
			// 	}),
			// )
			// theme.Editor(new(widget.Editor), "editor is focused").Layout(gtx, theme.Base)
			e.Frame(gtx.Ops)
		}
	}
}

var MaxWidth = unit.Dp(800)

func UniformPadding(gtx layout.Context, body layout.Widget) layout.Dimensions {
	width := gtx.Constraints.Max.X

	padding := values.MarginPadding24

	if (width - 2*gtx.Dp(padding)) > gtx.Dp(MaxWidth) {
		paddingValue := float32(width-gtx.Dp(MaxWidth)) / 2
		padding = unit.Dp(paddingValue)
	}

	return layout.Inset{
		Top:    values.MarginPadding24,
		Right:  padding,
		Bottom: values.MarginPadding24,
		Left:   padding,
	}.Layout(gtx, body)
}
func runDropDown22(window *app.Window) error {
	var ops op.Ops
	theme := widget2.NewTheme(window)

	var DropDownItems = []widget2.DropDownItem{
		{
			Text: "option-one",
			Icon: theme.Icons.WalletIcon,
		},
		{
			Text: "option-two",
			Icon: theme.Icons.WalletIcon,
		},
		{
			Text: "option-three",
			Icon: theme.Icons.WalletIcon,
		},
	}
	var (
		editor1 = widget.Editor{SingleLine: true}
		editor2 = widget.Editor{SingleLine: true}
	)

	dropDown := theme.DropDown(DropDownItems, 0, 2)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(layout.Spacer{Height: 10}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return theme.Editor(&editor1, "editor is focused").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return dropDown.Layout(gtx, dropDown.Width, true)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// return theme.Editor(&editor2, "editor is focused").Layout(gtx, theme.Base)
					return theme.EditorPassword(&editor2, "ssssssssss").Layout(gtx)
				}),
			)
			// theme.Editor(new(widget.Editor), "editor is focused").Layout(gtx, theme.Base)
			e.Frame(gtx.Ops)
		}
	}
}
