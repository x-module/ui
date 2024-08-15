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

type FileSelector struct {
	input    *Input
	fileName string

	actionClick  widget.Clickable
	windowTitle  string
	onSelectFile func(fileName string)

	filter  string
	changed bool
	width   unit.Dp
}

func NewFileSelector(hint string, fileName ...string) *FileSelector {
	bf := &FileSelector{
		input:       NewInput(hint, fileName...),
		width:       unit.Dp(200),
		windowTitle: "Select file",
	}
	if len(fileName) > 0 {
		bf.fileName = fileName[0]
		bf.input.SetText(fileName[0])
	}
	bf.updateIcon()
	return bf
}

func (b *FileSelector) SetFilter(filter string) {
	b.filter = filter
}

// SetWidth 设置width
func (b *FileSelector) SetWidth(width unit.Dp) {
	b.width = width
}

// SetWindowTitle 设置windowTitle
func (b *FileSelector) SetWindowTitle(title string) {
	b.windowTitle = title
}

func (b *FileSelector) action(gtx layout.Context) {
	if b.actionClick.Clicked(gtx) {
		if b.fileName != "" {
			b.RemoveFile()
			b.changed = true
			return
		} else {
			file, _, err := dlgs.File(b.windowTitle, b.filter, false)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Selected file:", file)
			if file == "" {
				return
			}
			b.setFileName(file)
			b.changed = true
			if b.onSelectFile != nil {
				b.onSelectFile(file)
			}
		}
	}
}
func (b *FileSelector) SetOnSelectFile(f func(fileName string)) {
	b.onSelectFile = f
}

func (b *FileSelector) setFileName(name string) {
	b.fileName = name
	b.input.SetText(name)
	b.updateIcon()
	b.changed = true
}

func (b *FileSelector) Changed() bool {
	out := b.changed
	b.changed = false
	return out
}

func (b *FileSelector) RemoveFile() {
	b.fileName = ""
	b.input.SetText("")
	b.updateIcon()
	b.changed = true
}

func (b *FileSelector) GetFileName() string {
	return b.fileName
}

func (b *FileSelector) updateIcon() {
	if b.fileName != "" {
		b.input.SetAfter(func(gtx layout.Context) layout.Dimensions {
			return b.actionClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = gtx.Dp(resource.DefaultIconSize)
				return widgets.DeleteIcon.Layout(gtx, resource.IconGrayColor)
			})
		})
	} else {
		b.input.SetAfter(func(gtx layout.Context) layout.Dimensions {
			return b.actionClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = gtx.Dp(resource.DefaultIconSize)
				return widgets.UploadIcon.Layout(gtx, resource.IconGrayColor)
			})
		})
	}
}

func (b *FileSelector) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	// gtx.Constraints.Max.Y = gtx.Dp(42)
	b.action(gtx)
	gtx.Constraints.Max.X = gtx.Dp(b.width)
	return b.input.Layout(gtx, theme)
}
