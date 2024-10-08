/**
 * Created by Goland
 * @file   color.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/13 13:58
 * @desc   color.go
 */

package resource

import "image/color"

var (
	DefaultWindowBgGrayColor  = color.NRGBA{R: 17, G: 15, B: 20, A: 255}
	DefaultContentBgGrayColor = color.NRGBA{R: 24, G: 24, B: 28, A: 255}

	DefaultBgGrayColor     = color.NRGBA{R: 53, G: 54, B: 56, A: 255}
	DefaultTextWhiteColor  = color.NRGBA{R: 223, G: 223, B: 224, A: 255}
	DefaultBorderGrayColor = color.NRGBA{R: 53, G: 54, B: 56, A: 255}
	DefaultBorderBlueColor = color.NRGBA{R: 127, G: 231, B: 196, A: 255}

	DefaultLineColor   = color.NRGBA{R: 44, G: 44, B: 47, A: 255}
	DefaultMaskBgColor = color.NRGBA{R: 10, G: 10, B: 12, A: 230}

	IconGrayColor            = color.NRGBA{R: 136, G: 136, B: 137, A: 255}
	BorderBlueColor          = color.NRGBA{R: 127, G: 231, B: 196, A: 255}
	BorderLightGrayColor     = color.NRGBA{R: 65, G: 65, B: 68, A: 255}
	HoveredBorderBlueColor   = color.NRGBA{R: 127, G: 231, B: 196, A: 255}
	FocusedBorderBlueColor   = color.NRGBA{R: 127, G: 231, B: 196, A: 255}
	ActivatedBorderBlueColor = color.NRGBA{R: 127, G: 231, B: 196, A: 255}
	FocusedBgColor           = color.NRGBA{R: 33, G: 50, B: 46, A: 255}
	TextSelectionColor       = color.NRGBA{R: 92, G: 136, B: 177, A: 255}
	HintTextColor            = color.NRGBA{R: 136, G: 136, B: 137, A: 255}

	DropDownBgGrayColor          = color.NRGBA{R: 72, G: 72, B: 77, A: 255}
	DropDownItemHoveredGrayColor = color.NRGBA{R: 90, G: 90, B: 96, A: 255}

	GreenColor   = color.NRGBA{R: 101, G: 231, B: 188, A: 255}
	ErrorColor   = color.NRGBA{R: 232, G: 127, B: 127, A: 255}
	WarningColor = color.NRGBA{R: 242, G: 201, B: 126, A: 255}
	SuccessColor = color.NRGBA{R: 99, G: 226, B: 184, A: 255}
	InfoColor    = color.NRGBA{R: 113, G: 192, B: 231, A: 255}

	ActionTipsBgGrayColor = color.NRGBA{A: 255, R: 48, G: 48, B: 51}
	ProgressBarColor      = color.NRGBA{R: 127, G: 200, B: 235, A: 255}

	MenuHoveredBgColor  = color.NRGBA{R: 45, G: 45, B: 48, A: 255}
	MenuSelectedBgColor = color.NRGBA{R: 35, G: 54, B: 51, A: 255}
	LogTextWhiteColor   = color.NRGBA{R: 202, G: 202, B: 203, A: 255}

	NotificationBgColor        = color.NRGBA{R: 72, G: 72, B: 77, A: 255}
	NotificationTextWhiteColor = color.NRGBA{R: 219, G: 219, B: 220, A: 255}
	ModalBgGrayColor           = color.NRGBA{R: 44, G: 44, B: 50, A: 255}

	DropdownMenuBgColor = color.NRGBA{R: 72, G: 72, B: 77, A: 255}
	DropdownTextColor   = color.NRGBA{R: 212, G: 212, B: 213, A: 255}
)
