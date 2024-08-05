package widgets

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gen2brain/dlgs"
	"github.com/x-module/ui/theme"
)

type DirSelector struct {
	textField *TextField
	DirName   string

	explorer    *Explorer
	onSelectDir func(dir string)

	changed bool
	width   unit.Dp
}

func NewDirSelector(DirName string, explorer *Explorer, placeholder string) *DirSelector {
	bf := &DirSelector{
		DirName:   DirName,
		textField: NewTextField(DirName, placeholder),
		explorer:  explorer,
		width:     unit.Dp(200),
	}

	bf.textField.SetText(DirName)
	bf.textField.IconPosition = IconPositionEnd
	bf.textField.SetMinWidth(200)
	bf.updateIcon()
	bf.setOnSelectDir(bf.handleExplorerSelect)
	return bf
}

// 设置width
func (b *DirSelector) SetWidth(width unit.Dp) {
	b.width = width
}

func (b *DirSelector) SetExplorer(explorer *Explorer) {
	b.explorer = explorer
}

func (b *DirSelector) handleExplorerSelect() {
	dir, _, err := dlgs.File("Select Directory", "", true)
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

func (b *DirSelector) SetOnSelectDir(f func(dir string)) {
	b.onSelectDir = f
}

func (b *DirSelector) setOnSelectDir(f func()) {
	b.textField.SetOnIconClick(func() {
		if b.DirName != "" {
			b.RemoveDir()
			b.changed = true
			return
		} else {
			f()
		}
	})
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
