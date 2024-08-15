/**
 * Created by Goland
 * @file   progressbar.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/15 23:52
 * @desc   progressbar.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/x-module/ui/naive/resource"
	"github.com/x-module/ui/theme"
)

type ProgressBar struct {
	progressBar     material.ProgressBarStyle
	theme           *theme.Theme
	initProgress    float32
	currentProgress float32
}

func NewProgressBar(theme *theme.Theme, initProgress float32) *ProgressBar {
	progressBar := &ProgressBar{
		theme:           theme,
		initProgress:    initProgress,
		currentProgress: initProgress,
		progressBar:     material.ProgressBar(theme.Material(), initProgress),
	}
	progressBar.progressBar.Color = resource.ProgressBarColor
	return progressBar
}
func (p *ProgressBar) SetProgress(progress float32) {
	p.currentProgress = progress
	p.progressBar.Progress = p.currentProgress
}

func (p *ProgressBar) Layout(gtx layout.Context) layout.Dimensions {
	return p.progressBar.Layout(gtx)
}
