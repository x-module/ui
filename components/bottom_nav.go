package components

import (
	"game.test.client/widget"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/planetdecred/godcr/ui/decredmaterial"
	"github.com/planetdecred/godcr/ui/values"
)

var (
	C layout.Context
	D layout.Dimensions
)
var (
	bottomNavigationBarHeight = values.MarginPadding100
)

type BottomNavigationBarHandler struct {
	Clickable     *decredmaterial.Clickable
	Image         *decredmaterial.Image
	ImageInactive *decredmaterial.Image
	Title         string
	PageID        string
}

type BottomNavigationBar struct {
	// *load.Load
	Theme widget.Theme

	FloatingActionButton  []BottomNavigationBarHandler
	BottomNavigationItems []BottomNavigationBarHandler
	CurrentPage           string

	axis        layout.Axis
	textSize    unit.Sp
	bottomInset unit.Dp
	height      unit.Dp
	alignment   layout.Alignment
	direction   layout.Direction
}

func (bottomNavigationBar *BottomNavigationBar) LayoutBottomNavigationBar(gtx layout.Context) layout.Dimensions {
	return layout.Stack{Alignment: layout.S}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return decredmaterial.LinearLayout{
				Width:       decredmaterial.WrapContent,
				Height:      decredmaterial.WrapContent,
				Orientation: layout.Horizontal,
				Background:  bottomNavigationBar.Theme.Color.Surface,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					list := layout.List{Axis: layout.Horizontal}
					return list.Layout(gtx, len(bottomNavigationBar.BottomNavigationItems), func(gtx layout.Context, i int) layout.Dimensions {
						background := bottomNavigationBar.Theme.Color.Surface
						if bottomNavigationBar.BottomNavigationItems[i].PageID == bottomNavigationBar.CurrentPage {
							background = bottomNavigationBar.Theme.Color.Gray5
						}
						return decredmaterial.LinearLayout{
							Orientation: bottomNavigationBar.axis,
							Width:       (gtx.Constraints.Max.X * 100 / len(bottomNavigationBar.BottomNavigationItems)) / 100, // Divide each cell equally
							Height:      decredmaterial.WrapContent,
							Padding:     layout.UniformInset(values.MarginPadding10),
							Alignment:   bottomNavigationBar.alignment,
							Direction:   bottomNavigationBar.direction,
							Background:  background,
							Clickable:   bottomNavigationBar.BottomNavigationItems[i].Clickable,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								img := bottomNavigationBar.BottomNavigationItems[i].ImageInactive

								if bottomNavigationBar.BottomNavigationItems[i].PageID == bottomNavigationBar.CurrentPage {
									img = bottomNavigationBar.BottomNavigationItems[i].Image
								}

								return img.Layout24dp(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Bottom: bottomNavigationBar.bottomInset,
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									textColor := bottomNavigationBar.Theme.Color.GrayText1
									if bottomNavigationBar.BottomNavigationItems[i].PageID == bottomNavigationBar.CurrentPage {
										textColor = bottomNavigationBar.Theme.Color.DeepBlue
									}
									txt := bottomNavigationBar.Theme.Label(bottomNavigationBar.textSize, bottomNavigationBar.BottomNavigationItems[i].Title)
									txt.Color = textColor
									return txt.Layout(gtx)
								})
							}),
						)
					})
				}),
			)
		}),
	)
}

func (bottomNavigationBar *BottomNavigationBar) LayoutSendReceive(gtx layout.Context) layout.Dimensions {
	gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
	if bottomNavigationBar.CurrentPage == "Overview" || bottomNavigationBar.CurrentPage == "Transactions" {
		return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Constraints.Max.X
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return decredmaterial.LinearLayout{
							Width:       decredmaterial.WrapContent,
							Height:      decredmaterial.WrapContent,
							Orientation: layout.Horizontal,
							Background:  bottomNavigationBar.Theme.Color.Primary,
							Border:      decredmaterial.Border{Radius: decredmaterial.Radius(20)},
							Padding:     layout.UniformInset(values.MarginPadding8),
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return decredmaterial.LinearLayout{
											Width:       decredmaterial.WrapContent,
											Height:      decredmaterial.WrapContent,
											Orientation: layout.Horizontal,
											Padding: layout.Inset{
												Right: values.MarginPadding16,
												Left:  values.MarginPadding16,
											},
											Clickable: bottomNavigationBar.FloatingActionButton[0].Clickable,
										}.Layout(gtx,
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return layout.Inset{Right: values.MarginPadding8}.Layout(gtx,
													func(gtx layout.Context) layout.Dimensions {
														return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
															return bottomNavigationBar.FloatingActionButton[0].Image.Layout24dp(gtx)
														})
													})
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return layout.Inset{
													Left: values.MarginPadding0,
												}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
													return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
														txt := bottomNavigationBar.Theme.Label(values.TextSize16, bottomNavigationBar.FloatingActionButton[0].Title)
														txt.Color = bottomNavigationBar.Theme.Color.DefaultThemeColors().White
														return txt.Layout(gtx)
													})
												})
											}),
										)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										verticalSeparator := bottomNavigationBar.Theme.SeparatorVertical(50, 1)
										verticalSeparator.Color = bottomNavigationBar.Theme.Color.DefaultThemeColors().White
										return verticalSeparator.Layout(gtx)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return decredmaterial.LinearLayout{
											Width:       decredmaterial.WrapContent,
											Height:      decredmaterial.WrapContent,
											Orientation: layout.Horizontal,
											Padding: layout.Inset{
												Right: values.MarginPadding16,
												Left:  values.MarginPadding16,
											},
											Clickable: bottomNavigationBar.FloatingActionButton[1].Clickable,
										}.Layout(gtx,
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return layout.Inset{Right: values.MarginPadding8}.Layout(gtx,
													func(gtx layout.Context) layout.Dimensions {
														return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
															return bottomNavigationBar.FloatingActionButton[1].Image.Layout24dp(gtx)
														})
													})
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return layout.Inset{
													Left: values.MarginPadding0,
												}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
													return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
														txt := bottomNavigationBar.Theme.Label(values.TextSize16, bottomNavigationBar.FloatingActionButton[1].Title)
														txt.Color = bottomNavigationBar.Theme.Color.DefaultThemeColors().White
														return txt.Layout(gtx)
													})
												})
											}),
										)
									}),
								)
							}),
						)
					})
				}),
			)
		})
	}
	return layout.Dimensions{}
}

func (bottomNavigationBar *BottomNavigationBar) OnViewCreated() {
	bottomNavigationBar.axis = layout.Vertical
	bottomNavigationBar.textSize = values.TextSize12
	bottomNavigationBar.bottomInset = values.MarginPadding0
	bottomNavigationBar.height = bottomNavigationBarHeight
	bottomNavigationBar.alignment = layout.Middle
	bottomNavigationBar.direction = layout.Center
}
