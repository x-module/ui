package widget

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/unit"
	widget2 "gioui.org/widget"
	"github.com/x-module/ui/widget/values"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image/color"
	"reflect"

	"gioui.org/widget/material"
)

var (
	DefaultHover    = 48
	DefaultSelected = 96
)

type Palette struct {
	// Bg is the background color atop which content is currently being
	// drawn.
	Bg color.NRGBA

	// Fg is a color suitable for drawing on top of Bg.
	Fg color.NRGBA

	// ContrastBg is a color used to draw attention to active,
	// important, interactive widgets such as buttons.
	ContrastBg color.NRGBA

	// ContrastFg is a color suitable for content drawn on top of
	// ContrastBg.
	ContrastFg color.NRGBA

	// Bg2 specifies the background color for components like navibar
	Bg2 color.NRGBA

	HoverAlpha, SelectedAlpha uint8
}

type SubThemeID string

type Theme struct {
	*material.Theme

	// Alpha is the set of alpha values to be applied for certain
	// states like hover, selected, etc...
	HoverAlpha, SelectedAlpha uint8

	// Bg2 specifies the background color for components like navibar
	Bg2 color.NRGBA

	// sub theme maps and their type info used to
	// accept dynamic subtheme registration.
	subThemes     map[SubThemeID]interface{}
	subThemeTypes map[SubThemeID]reflect.Type

	// 	===============new =================
	Color    *values.Color
	Styles   *values.WidgetStyles
	Window   *app.Window
	TextSize unit.Sp

	// Alpha is the set of alpha values to be applied for certain
	// states like hover, selected, etc...

	checkBoxCheckedIcon   *widget2.Icon
	checkBoxUncheckedIcon *widget2.Icon
	radioCheckedIcon      *widget2.Icon
	radioUncheckedIcon    *widget2.Icon
	chevronUpIcon         *widget2.Icon
	dropDownIcon          *widget2.Icon
	chevronDownIcon       *widget2.Icon
	navigationCheckIcon   *widget2.Icon
	navMoreIcon           *widget2.Icon

	Icons         *Icons
	dropDownMenus []*DropDown
}

// NewTheme instantiates a theme, extending material theme.
// func NewTheme(fontDir string, embeddedFonts [][]byte, noSystemFonts bool) *Theme {
func NewTheme(window *app.Window) *Theme {
	th := material.NewTheme()
	// var options = []text.ShaperOption{
	// 	text.WithCollection(LoadBuiltin(fontDir, embeddedFonts)),
	// }
	// if noSystemFonts {
	// 	options = append(options, text.NoSystemFonts())
	// }
	// th.Shaper = text.NewShaper(options...)

	color := values.NewColor().DefaultThemeColors()
	color.DarkThemeColors()
	t := &Theme{
		Theme:         th,
		HoverAlpha:    uint8(DefaultHover),
		SelectedAlpha: uint8(DefaultSelected),
		Bg2:           th.Bg,

		// 	===============new =================
		Color:    color,
		Window:   window,
		Styles:   values.DefaultWidgetStyles(),
		TextSize: values.TextSize16,
		Icons:    &Icons{},
	}

	// =================new===================
	t.updateStyles(false)
	t.checkBoxCheckedIcon = MustIcon(widget2.NewIcon(icons.ToggleCheckBox))
	t.checkBoxUncheckedIcon = MustIcon(widget2.NewIcon(icons.ToggleCheckBoxOutlineBlank))
	t.radioCheckedIcon = MustIcon(widget2.NewIcon(icons.ToggleRadioButtonChecked))
	t.radioUncheckedIcon = MustIcon(widget2.NewIcon(icons.ToggleRadioButtonUnchecked))
	t.chevronUpIcon = MustIcon(widget2.NewIcon(icons.NavigationExpandLess))
	t.chevronDownIcon = MustIcon(widget2.NewIcon(icons.NavigationExpandMore))
	t.navMoreIcon = MustIcon(widget2.NewIcon(icons.NavigationMoreHoriz))
	t.navigationCheckIcon = MustIcon(widget2.NewIcon(icons.NavigationCheck))
	t.dropDownIcon = MustIcon(widget2.NewIcon(icons.NavigationArrowDropDown))

	return t
}

func (t *Theme) WithPalette(p Palette) *Theme {
	t.Theme.Palette = material.Palette{
		Bg:         p.Bg,
		Fg:         p.Fg,
		ContrastFg: p.ContrastFg,
		ContrastBg: p.ContrastBg,
	}

	if p.HoverAlpha > 0 {
		t.HoverAlpha = p.HoverAlpha
	}
	if p.SelectedAlpha > 0 {
		t.SelectedAlpha = p.SelectedAlpha
	}

	t.Bg2 = p.Bg2
	return t
}

func (th *Theme) Register(ID SubThemeID, sub interface{}) error {
	if th.subThemes == nil {
		th.subThemes = make(map[SubThemeID]interface{})
		th.subThemeTypes = make(map[SubThemeID]reflect.Type)
	}

	// confliction check
	if t, ok := th.subThemeTypes[ID]; ok {
		if t != reflect.TypeOf(sub) {
			return fmt.Errorf("type %v already registered as %s", ID, t.Name())
		}
	}

	th.subThemes[ID] = sub
	th.subThemeTypes[ID] = reflect.TypeOf(sub)
	return nil
}

func (th *Theme) Get(ID SubThemeID) interface{} {
	if _, exist := th.subThemeTypes[ID]; !exist {
		panic(fmt.Sprintf("%v not registered", ID))
	}

	return th.subThemes[ID]
}

func (t *Theme) closeAllDropdownMenus(group uint) {
	for _, dropDown := range t.dropDownMenus {
		if dropDown.Group == group {
			dropDown.IsOpen = false
		}
	}
}

func (t *Theme) isOpenDropdownGroup(group uint) bool {
	for _, dropDown := range t.dropDownMenus {
		if dropDown.Group == group {
			if dropDown.IsOpen {
				return true
			}
		}
	}
	return false
}

func (t *Theme) updateStyles(isDarkModeOn bool) {
	// update switch style colors
	t.Styles.SwitchStyle.ActiveColor = t.Color.Primary
	t.Styles.SwitchStyle.InactiveColor = t.Color.Gray3
	t.Styles.SwitchStyle.ThumbColor = t.Color.White

	// update icon button style colors
	t.Styles.IconButtonColorStyle.Background = color.NRGBA{}
	t.Styles.IconButtonColorStyle.Foreground = t.Color.Gray1

	// update Collapsible widget style colors
	t.Styles.CollapsibleStyle.Background = t.Color.Surface
	t.Styles.CollapsibleStyle.Foreground = color.NRGBA{}

	// update clickable colors
	t.Styles.ClickableStyle.Color = t.Color.SurfaceHighlight
	t.Styles.ClickableStyle.HoverColor = t.Color.Gray5

	// dropdown clickable colors
	t.Styles.DropdownClickableStyle.Color = t.Color.SurfaceHighlight
	col := t.Color.Gray3
	if isDarkModeOn {
		col = t.Color.Gray5
	}
	t.Styles.DropdownClickableStyle.HoverColor = Hovered(col)
}
func Hovered(c color.NRGBA) (d color.NRGBA) {
	const r = 0x20 // lighten ratio
	return color.NRGBA{
		R: byte(255 - int(255-c.R)*(255-r)/256),
		G: byte(255 - int(255-c.G)*(255-r)/256),
		B: byte(255 - int(255-c.B)*(255-r)/256),
		A: c.A,
	}
}
