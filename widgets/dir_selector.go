package widgets

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gen2brain/dlgs"
	"github.com/x-module/ui/theme"
)

type DirSelector struct {
	textField *Input
	DirName   string

	windowTitle string
	onSelectDir func(dir string)

	changed bool
	width   unit.Dp
}

func NewDirSelector(dirName string, placeholder string) *DirSelector {
	bf := &DirSelector{
		DirName:     dirName,
		textField:   NewInput(dirName, placeholder),
		width:       unit.Dp(200),
		windowTitle: "Select Directory",
	}
	bf.textField.SetText(dirName)
	bf.textField.IconPosition = IconPositionEnd
	bf.textField.SetWidth(200)
	bf.updateIcon()
	bf.setOnSelectDir()
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

func (b *DirSelector) setOnSelectDir() {
	b.textField.SetOnIconClick(func() {
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
	})
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
		b.textField.SetIcon(DeleteIcon, IconPositionEnd)
	} else {
		b.textField.SetIcon(UploadIcon, IconPositionEnd)
	}
}

func (b *DirSelector) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	// gtx.Constraints.Max.Y = gtx.Dp(42)
	gtx.Constraints.Max.X = gtx.Dp(b.width)
	return b.textField.Layout(gtx, theme)
}
