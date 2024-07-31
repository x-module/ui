package widgets

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gen2brain/dlgs"
	"github.com/x-module/ui/theme"
)

// DirSelector is a widget that allows the user to select a file. it handles the file selection and display the file name.
// TODO replace binary file with this widget
type DirSelector struct {
	textField *TextField
	FileName  string

	extensions []string

	explorer     *Explorer
	onSelectFile func()

	changed bool
	width   unit.Dp
}

func NewDirSelector(filename string, explorer *Explorer, extensions ...string) *DirSelector {
	bf := &DirSelector{
		FileName:   filename,
		textField:  NewTextField(filename, "File"),
		explorer:   explorer,
		extensions: extensions,
		width:      unit.Dp(200),
	}

	bf.textField.SetText(filename)
	bf.textField.IconPosition = IconPositionEnd
	bf.textField.SetMinWidth(200)
	bf.updateIcon()
	bf.SetOnSelectFile(bf.handleExplorerSelect)
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
	b.SetFileName(dir)
	b.changed = true
}

func (b *DirSelector) SetOnSelectFile(f func()) {
	b.onSelectFile = f
	b.textField.SetOnIconClick(func() {
		if b.FileName != "" {
			b.RemoveFile()
			b.changed = true
			return
		} else {
			// Select file
			f()
		}
	})
}

func (b *DirSelector) SetFileName(name string) {
	b.FileName = name
	b.textField.SetText(name)
	b.updateIcon()
	b.changed = true
}

func (b *DirSelector) Changed() bool {
	out := b.changed
	b.changed = false
	return out
}

func (b *DirSelector) RemoveFile() {
	b.FileName = ""
	b.textField.SetText("")
	b.updateIcon()
	b.changed = true
}

func (b *DirSelector) GetFilePath() string {
	return b.FileName
}

func (b *DirSelector) updateIcon() {
	if b.FileName != "" {
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
