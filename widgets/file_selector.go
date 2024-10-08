package widgets

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/x-module/ui/theme"
)

// FileSelector is a widget that allows the user to select a file. it handles the file selection and display the file name.
// TODO replace binary file with this widget
type FileSelector struct {
	textField *Input
	FileName  string

	extensions []string

	explorer     *Explorer
	onSelectFile func(file string)

	changed bool
	width   unit.Dp
}

func NewFileSelector(filename string, explorer *Explorer, placeholder string, extensions ...string) *FileSelector {
	bf := &FileSelector{
		FileName:   filename,
		textField:  NewInput(filename, placeholder),
		explorer:   explorer,
		extensions: extensions,
		width:      unit.Dp(200),
	}

	bf.textField.SetText(filename)
	bf.textField.IconPosition = IconPositionEnd
	bf.textField.SetWidth(200)
	bf.updateIcon()
	bf.setOnSelectFile()
	return bf
}

// 设置width
func (b *FileSelector) SetWidth(width unit.Dp) {
	b.width = width
}

func (b *FileSelector) SetExplorer(explorer *Explorer) {
	b.explorer = explorer
}

func (b *FileSelector) handleExplorerSelect() {

}

func (b *FileSelector) setOnSelectFile() {
	b.textField.SetOnIconClick(func() {
		if b.FileName != "" {
			b.RemoveFile()
			b.changed = true
			return
		} else {
			// Select file
			if b.explorer == nil {
				return
			}
			b.explorer.ChoseFile(func(result Result) {
				if result.Error != nil {
					fmt.Println("failed to get file", result.Error)
					return
				}
				if result.FilePath == "" {
					return
				}
				b.SetFileName(result.FilePath)
				if b.onSelectFile != nil {
					b.onSelectFile(result.FilePath)
				}
				b.changed = true
			}, b.extensions...)
		}
	})
}
func (b *FileSelector) SetOnSelectFile(f func(file string)) {
	b.onSelectFile = f
}

func (b *FileSelector) SetFileName(name string) {
	b.FileName = name
	b.textField.SetText(name)
	b.updateIcon()
	b.changed = true
}

func (b *FileSelector) Changed() bool {
	out := b.changed
	b.changed = false
	return out
}

func (b *FileSelector) RemoveFile() {
	b.FileName = ""
	b.textField.SetText("")
	b.updateIcon()
	b.changed = true
}

func (b *FileSelector) GetFilePath() string {
	return b.FileName
}

func (b *FileSelector) updateIcon() {
	if b.FileName != "" {
		b.textField.SetIcon(DeleteIcon, IconPositionEnd)
	} else {
		b.textField.SetIcon(UploadIcon, IconPositionEnd)
	}
}

func (b *FileSelector) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	// gtx.Constraints.Max.Y = gtx.Dp(32)
	gtx.Constraints.Max.X = gtx.Dp(b.width)
	return b.textField.Layout(gtx, theme)
}
