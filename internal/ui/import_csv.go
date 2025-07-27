package ui

import (
	"fyne.io/fyne/v2/dialog"
	"log"

	"aegis/internal/pass_import"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// openImportPassFromFile opens a new window for importing passwords from a CSV file.
//
// Args:
//
//	a: The Fyne application instance.
func openImportPassFromFile(a fyne.App) {
	updateWindow := a.NewWindow("Import CSV")
	updateWindow.Resize(fyne.NewSize(400, 250))
	updateWindow.CenterOnScreen()

	titleLabel := widget.NewLabel("Import CSV")
	titleLabel.TextStyle.Bold = true
	titleLabel.Importance = widget.HighImportance
	statusLabel := widget.NewLabel("")

	selectCsvBtn := widget.NewButton("Select CSV File", func() {
		dialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("File open error:", err)
				return
			}
			if reader == nil {
				return
			}

			defer reader.Close()

			pass_import.ImportPasswordsCsv(reader.URI().Path())

			updateWindow.Close()
			refreshUserList(a)

		}, updateWindow)

		dialog.Show()
	})

	cancelBtn := widget.NewButton("Cancel", func() {
		updateWindow.Close()
	})

	buttonContainer := container.NewHBox(
		selectCsvBtn,
		cancelBtn,
	)

	form := container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		buttonContainer,
		statusLabel,
	)

	content := container.NewStack(
		windowBg,
		container.NewPadded(form),
	)

	updateWindow.SetContent(content)
	updateWindow.Show()
}
