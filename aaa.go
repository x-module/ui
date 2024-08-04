package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			panic(err)
		}
		app.Main()
	}()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var (
		passwordEditor  widget.Editor
		showPasswordBtn widget.Clickable
		showPassword    bool
	)
	passwordEditor.SingleLine = true

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&op.Ops{}, e)
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Label(th, unit.Sp(16), "Password:").Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
								layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
									return PasswordEditor(th, &passwordEditor, showPassword).Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return ShowPasswordButton(th, &showPasswordBtn).Layout(gtx)
								}),
							)
						}),
					)
				})
			})

			if showPasswordBtn.Clicked() {
				showPassword = !showPassword
			}

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}

func PasswordEditor(th *material.Theme, editor *widget.Editor, showPassword bool) material.EditorStyle {
	if !showPassword {
		maskedText := ""
		for range editor.Text() {
			maskedText += "•"
		}
		editor.SetText(maskedText)
	}

	editorStyle := material.Editor(th, editor, "Password")
	editorStyle.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	editorStyle.TextSize = unit.Sp(16)
	editorStyle.Hint = "Enter your password"
	editorStyle.Mask = !showPassword
	return editorStyle
}

func ShowPasswordButton(th *material.Theme, btn *widget.Clickable) material.ButtonStyle {
	icon := material.IconButton(th, btn, &widget.Icon{src: image.NewUniform(color.NRGBA{0, 0, 0, 0})})
	icon.Text = "👁️" // Use an emoji or replace with an actual icon
	icon.Size = unit.Dp(24)
	return icon
}
