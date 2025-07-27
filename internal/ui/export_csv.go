package ui

import (
	"fyne.io/fyne/v2/dialog"
	"log"

	"aegis/internal/pass_export"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// openExportPassToFileWindow opens a new window for exporting passwords to a CSV file.
//
// Args:
//
//	a: The Fyne application instance.
func openExportPassToFileWindow(a fyne.App) {
	updateWindow := a.NewWindow("Export DB To CSV")
	updateWindow.Resize(fyne.NewSize(400, 250))
	updateWindow.CenterOnScreen()

	titleLabel := widget.NewLabel("Export DB")
	titleLabel.TextStyle.Bold = true
	titleLabel.Importance = widget.HighImportance

	selectCsvBtn := widget.NewButton("Select Path", func() {
		dialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				log.Println("File open error:", err)
				return
			}
			if writer == nil {
				return
			}

			defer writer.Close()

			pass_export.ExportPasswordsCsv(writer.URI().Path())

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
	)

	content := container.NewStack(
		windowBg,
		container.NewPadded(form),
	)

	updateWindow.SetContent(content)
	updateWindow.Show()
}
