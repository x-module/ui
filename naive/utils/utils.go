/**
 * Created by Goland
 * @file   utils.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/16 13:08
 * @desc   utils.go
 */

package utils

import (
	"github.com/x-module/ui/naive/widgets"
	"time"
)

var NotificationController = widgets.NewNotification()
var SystemNoticeController = &widgets.SystemNotice{}

func AppNotice(text string, duration ...time.Duration) {
	dru := time.Second * 3
	if len(duration) > 0 {
		dru = duration[0]
	}
	NotificationController.EndAt = time.Now().Add(dru)
	NotificationController.Text = text
}

func SystemNotice(message string) {
	_ = SystemNoticeController.Notice(message)
}
