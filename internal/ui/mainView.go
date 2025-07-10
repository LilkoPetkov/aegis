package ui

import (
	"aegis/internal/queries"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"image/color"
)

var userListContainer *fyne.Container
var scrollContainer *container.Scroll

func RunUI() {
	db := queries.InitDB()
	queries.CreatePasswordsTable(db)
	a := app.New()
	w := a.NewWindow("Aegis Password Manager")
	w.CenterOnScreen()
	w.SetTitle("Aegis")
	w.Resize(fyne.Size{Width: 800, Height: 800})
	w.SetIcon(theme.AccountIcon())

	bg := canvas.NewLinearGradient(
		color.NRGBA{R: 169, G: 142, B: 101, A: 255},
		color.NRGBA{R: 101, G: 67, B: 56, A: 255},
		90,
	)

	titleText := canvas.NewText("Aegis Password Manager", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	titleText.Alignment = fyne.TextAlignCenter
	titleText.TextSize = 24
	titleText.TextStyle.Bold = true

	addButton := widget.NewButton("Add New Password", func() {
		openAddUserWindow(a, db)
	})
	addButton.Importance = widget.HighImportance

	headerContainer := container.NewBorder(
		nil, nil,
		container.NewPadded(titleText),
		container.NewPadded(addButton),
	)

	headerBg := canvas.NewLinearGradient(
		color.NRGBA{R: 35, G: 65, B: 75, A: 255},
		color.NRGBA{R: 25, G: 35, B: 50, A: 255},
		0,
	)
	headerWithBg := container.NewStack(headerBg, container.NewPadded(headerContainer))

	userListContainer = buildUserList(db, a)

	scrollContainer = container.NewScroll(userListContainer)
	scrollContainer.SetMinSize(fyne.NewSize(780, 450))

	content := container.NewBorder(
		container.NewVBox(
			headerWithBg,
			widget.NewSeparator(),
		),
		nil, nil, nil,
		container.NewPadded(scrollContainer),
	)

	w.SetContent(container.NewStack(bg, content))
	w.ShowAndRun()
}
