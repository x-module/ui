package widgets

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gen2brain/dlgs"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
)

type DirSelector struct {
	textField *Input
	DirName   string

	actionClick widget.Clickable
	windowTitle string
	onSelectDir func(dir string)

	changed bool
	width   unit.Dp
}

func NewDirSelector(hint string, dirName ...string) *DirSelector {
	bf := &DirSelector{
		textField:   NewInput(hint, dirName...),
		width:       unit.Dp(200),
		windowTitle: "Select Directory",
	}
	if len(dirName) > 0 {
		bf.DirName = dirName[0]
		bf.textField.SetText(dirName[0])
	}
	bf.updateIcon()
	return bf
}

// SetWidth 设置width
func (b *DirSelector) SetWidth(width unit.Dp) {
	b.width = width
}

// SetWindowTitle 设置windowTitle
func (b *DirSelector) SetWindowTitle(title string) {
	b.windowTitle = title
}

func (b *DirSelector) action(gtx layout.Context) {
	if b.actionClick.Clicked(gtx) {
		if b.DirName != "" {
			b.RemoveDir()
			b.changed = true
			return
		} else {
			dir, _, err := dlgs.File(b.windowTitle, "", true)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Selected Directory:", dir)
			if dir == "" {
				return
			}
			b.SetDirName(dir)
			b.changed = true
			if b.onSelectDir != nil {
				b.onSelectDir(dir)
			}
		}
	}
}
func (b *DirSelector) SetOnSelectDir(f func(dir string)) {
	b.onSelectDir = f
}

func (b *DirSelector) SetDirName(name string) {
	b.DirName = name
	b.textField.SetText(name)
	b.updateIcon()
	b.changed = true
}

func (b *DirSelector) Changed() bool {
	out := b.changed
	b.changed = false
	return out
}

func (b *DirSelector) RemoveDir() {
	b.DirName = ""
	b.textField.SetText("")
	b.updateIcon()
	b.changed = true
}

func (b *DirSelector) GetDirPath() string {
	return b.DirName
}

func (b *DirSelector) updateIcon() {
	if b.DirName != "" {
		b.textField.SetAfter(func(gtx layout.Context) layout.Dimensions {
			return b.actionClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = gtx.Dp(resource.DefaultIconSize)
				return widgets.DeleteIcon.Layout(gtx, resource.IconGrayColor)
			})
		})
	} else {
		b.textField.SetAfter(func(gtx layout.Context) layout.Dimensions {
			return b.actionClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = gtx.Dp(resource.DefaultIconSize)
				return widgets.UploadIcon.Layout(gtx, resource.IconGrayColor)
			})
		})
	}
}

func (b *DirSelector) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	// gtx.Constraints.Max.Y = gtx.Dp(42)
	b.action(gtx)
	gtx.Constraints.Max.X = gtx.Dp(b.width)
	return b.textField.Layout(gtx, theme)
}
