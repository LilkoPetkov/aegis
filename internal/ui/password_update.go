package ui

import (
	"aegis/internal/queries"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// openPasswordUpdateWindow opens a new window for updating a user's password.
//
// Args:
//
//	a: The Fyne application instance.
//	username: The username of the user to update.
func openPasswordUpdateWindow(a fyne.App, username string) {
	updateWindow := a.NewWindow("Update Password")
	updateWindow.Resize(fyne.NewSize(400, 250))
	updateWindow.CenterOnScreen()

	titleLabel := widget.NewLabel("Update Password for: " + username)
	titleLabel.TextStyle.Bold = true
	titleLabel.Importance = widget.HighImportance

	passwordLabel := widget.NewLabel("New Password:")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter new password")

	statusLabel := widget.NewLabel("")

	submitBtn := widget.NewButton("Update Password", func() {
		password := passwordEntry.Text

		if password == "" {
			statusLabel.SetText("Password field is required")
			statusLabel.Importance = widget.DangerImportance
			statusLabel.Refresh()
			return
		}

		queries.EditUserPassword(password, username)

		updateWindow.Close()
		refreshUserList(a)
	})
	submitBtn.Importance = widget.HighImportance

	cancelBtn := widget.NewButton("Cancel", func() {
		updateWindow.Close()
	})

	buttonContainer := container.NewHBox(
		submitBtn,
		cancelBtn,
	)

	form := container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		passwordLabel,
		passwordEntry,
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
