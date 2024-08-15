package widgets

import (
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/x-module/ui/naive/resource"
	theme2 "github.com/x-module/ui/theme"
	"github.com/x-module/ui/widgets"
	"image"
	"image/color"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/x/component"
)

type DropDown struct {
	dropDown        *widget.Bool
	menuContextArea component.ContextArea
	menu            component.MenuState
	list            *widget.List
	theme           *theme2.Theme

	MinWidth unit.Dp
	width    unit.Dp
	menuInit bool

	isOpen              bool
	selectedOptionIndex int
	lastSelectedIndex   int
	options             []*DropDownOption

	size image.Point

	borderWidth  unit.Dp
	cornerRadius unit.Dp

	onValueChange func(value string)
}

type DropDownOption struct {
	Text       string
	Value      string
	Identifier string
	clickable  widget.Clickable

	Icon      *widget.Icon
	IconColor color.NRGBA

	isDivider bool
	isDefault bool
}

// SetWidth 设置width
func (c *DropDown) SetWidth(width unit.Dp) {
	c.width = width
}

func NewDropDownOption(text string) *DropDownOption {
	return &DropDownOption{
		Text:      text,
		Value:     text,
		isDivider: false,
	}
}

func NewDropDownDivider() *DropDownOption {
	return &DropDownOption{
		isDivider: true,
	}
}

func (o *DropDownOption) WithIdentifier(identifier string) *DropDownOption {
	o.Identifier = identifier
	return o
}

func (o *DropDownOption) WithValue(value string) *DropDownOption {
	o.Value = value
	return o
}

func (o *DropDownOption) WithIcon(icon *widget.Icon, color color.NRGBA) *DropDownOption {
	o.Icon = icon
	o.IconColor = color
	return o
}

func (o *DropDownOption) DefaultSelected() *DropDownOption {
	o.isDefault = true
	return o
}

func (o *DropDownOption) GetText() string {
	if o == nil {
		return ""
	}

	return o.Text
}

func (o *DropDownOption) GetValue() string {
	if o == nil {
		return ""
	}

	return o.Value
}

func (c *DropDown) SetSelected(index int) {
	c.selectedOptionIndex = index
	c.lastSelectedIndex = index
}

func (c *DropDown) SetOnChanged(f func(value string)) {
	c.onValueChange = f
}

func (c *DropDown) SetSelectedByTitle(title string) {
	if len(c.options) == 0 {
		return
	}

	for i, opt := range c.options {
		if opt.Text == title {
			c.selectedOptionIndex = i
			c.lastSelectedIndex = i
			break
		}
	}
}

func (c *DropDown) SetSelectedByIdentifier(identifier string) {
	for i, opt := range c.options {
		if opt.Identifier == identifier {
			c.selectedOptionIndex = i
			c.lastSelectedIndex = i
			break
		}
	}
}

func (c *DropDown) SetSelectedByValue(value string) {
	for i, opt := range c.options {
		if opt.Value == value {
			c.selectedOptionIndex = i
			c.lastSelectedIndex = i
			break
		}
	}
}

func NewDropDown(theme *theme2.Theme, options ...string) *DropDown {
	c := &DropDown{
		dropDown: &widget.Bool{Value: true},
		menuContextArea: component.ContextArea{
			Activation:       pointer.ButtonPrimary,
			AbsolutePosition: true,
		},
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
		borderWidth:  unit.Dp(1),
		cornerRadius: unit.Dp(4),
		theme:        theme,
		menuInit:     true,
	}
	if len(options) > 0 {
		for _, opt := range options {
			c.options = append(c.options, NewDropDownOption(opt))
		}
	}
	return c
}

func NewDropDownWithoutBorder(theme *theme2.Theme, options ...*DropDownOption) *DropDown {
	c := &DropDown{
		menuContextArea: component.ContextArea{
			Activation:       pointer.ButtonPrimary,
			AbsolutePosition: true,
		},
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
		options:  options,
		theme:    theme,
		menuInit: true,
	}

	return c
}

func (c *DropDown) SelectedIndex() int {
	return c.selectedOptionIndex
}

func (c *DropDown) SetOptions(options []*DropDownOption) {
	c.selectedOptionIndex = 0
	c.options = options
	if len(c.options) > 0 {
		c.menuInit = true
	}
}

func (c *DropDown) GetSelected() *DropDownOption {
	if len(c.options) == 0 {
		return nil
	}

	return c.options[c.selectedOptionIndex]
}

func (c *DropDown) box(gtx layout.Context, theme *theme2.Theme, text string, maxWidth unit.Dp) layout.Dimensions {
	borderColor := theme.BorderColor
	if c.isOpen {
		borderColor = resource.HoveredBorderBlueColor
	}

	border := widget.Border{
		Color:        borderColor,
		Width:        c.borderWidth,
		CornerRadius: c.cornerRadius,
	}

	if maxWidth == 0 {
		maxWidth = unit.Dp(gtx.Constraints.Max.X)
	}

	c.size.X = gtx.Dp(maxWidth)
	return c.dropDown.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		if c.dropDown.Hovered() {
			border.Color = resource.HoveredBorderBlueColor
		}
		return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// calculate the minimum width of the box, considering icon and padding
			return layout.Background{}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, gtx.Dp(resource.DefaultRadiusSize)).Push(gtx.Ops).Pop()
					paint.Fill(gtx.Ops, resource.DefaultBgGrayColor)
					return layout.Dimensions{Size: gtx.Constraints.Min}
				},
				func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Dp(maxWidth) - gtx.Dp(8)
					return layout.Inset{
						Top:    4,
						Bottom: 4,
						Left:   8,
						Right:  4,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: layout.SpaceBetween}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return Label(theme, text).Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								gtx.Constraints.Min.X = gtx.Dp(16)
								return widgets.ExpandIcon.Layout(gtx, theme.TextColor)
							}),
						)
					})
				},
			)
		})
	})
}

func (c *DropDown) SetSize(size image.Point) {
	c.size = size
}

// Layout the DropDown.
func (c *DropDown) Layout(gtx layout.Context, theme *theme2.Theme) layout.Dimensions {
	c.isOpen = c.menuContextArea.Active()

	for i, opt := range c.options {
		if opt.isDefault {
			c.selectedOptionIndex = i
		}

		for opt.clickable.Clicked(gtx) {
			c.isOpen = false
			c.selectedOptionIndex = i
		}
	}

	if c.selectedOptionIndex != c.lastSelectedIndex {
		if c.onValueChange != nil {
			go c.onValueChange(c.options[c.selectedOptionIndex].Value)
		}
		c.lastSelectedIndex = c.selectedOptionIndex
	}

	// Update menu items only if options change
	if c.menuInit {
		c.menuInit = false
		c.updateMenuItems(theme)
	}

	if c.MinWidth == 0 {
		c.MinWidth = unit.Dp(150)
	}

	text := ""
	if c.selectedOptionIndex >= 0 && c.selectedOptionIndex < len(c.options) {
		text = c.options[c.selectedOptionIndex].Text
	}

	box := c.box(gtx, theme, text, c.width)
	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return box
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return c.menuContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				offset := layout.Inset{
					Top:  unit.Dp(float32(box.Size.Y)/gtx.Metric.PxPerDp + 1),
					Left: unit.Dp(4),
				}
				return offset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Dp(c.MinWidth)
					if c.width != 0 {
						gtx.Constraints.Max.X = gtx.Dp(c.width)
					}
					m := component.Menu(theme.Material(), &c.menu)
					m.SurfaceStyle.Fill = resource.DropDownBgGrayColor
					return m.Layout(gtx)
				})
			})
		}),
	)
}

// updateMenuItems creates or updates menu items based on options and calculates minWidth.
func (c *DropDown) updateMenuItems(th *theme2.Theme) {
	c.menu.Options = c.menu.Options[:0]
	for _, opt := range c.options {
		opt := opt
		c.menu.Options = append(c.menu.Options, func(gtx layout.Context) layout.Dimensions {
			if opt.isDivider {
				dv := component.Divider(th.Material())
				dv.Fill = th.BorderColor
				return dv.Layout(gtx)
			}
			itm := component.MenuItem(th.Material(), &opt.clickable, opt.Text)
			itm.HoverColor = resource.DropDownItemHoveredGrayColor
			if opt.Icon != nil {
				itm.Icon = opt.Icon
				itm.IconColor = opt.IconColor
			}

			if c.GetSelected().Text == opt.Text {
				itm.Label.Color = resource.GreenColor
				// itm.Icon = widgets.ActionStarRateIcon
				// itm.IconSize = unit.Dp(16)
				// itm.IconInset = outlay.Inset{}
				// itm.IconColor = opt.IconColor
			} else {
				itm.Label.Color = resource.DefaultTextWhiteColor
			}
			return itm.Layout(gtx)
		})
	}
}
