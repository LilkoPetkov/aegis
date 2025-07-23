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
	db := queries.DB
	defer db.Close()

	queries.CreatePasswordsTable()
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

	importCsvButton := widget.NewButton("Import Passwords", func() {
		openImportPassFromFile(a)
	})
	importCsvButton.Importance = widget.HighImportance
	exportCsvButton := widget.NewButton("Export Passwords", func() {
		openExportPassToFileWindow(a)
	})
	exportCsvButton.Importance = widget.HighImportance
	addButton := widget.NewButton("Add New Password", func() {
		openAddUserWindow(a)
	})
	addButton.Importance = widget.HighImportance

	buttonBar := container.NewHBox(
		importCsvButton,
		exportCsvButton,
		addButton,
	)

	headerContainer := container.NewBorder(
		nil, nil,
		container.NewPadded(titleText),
		buttonBar,
	)

	headerBg := canvas.NewLinearGradient(
		color.NRGBA{R: 35, G: 65, B: 75, A: 255},
		color.NRGBA{R: 25, G: 35, B: 50, A: 255},
		0,
	)
	headerWithBg := container.NewStack(headerBg, container.NewPadded(headerContainer))

	userListContainer = buildUserList(a)

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
