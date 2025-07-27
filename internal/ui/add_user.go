package ui

import (
	"aegis/internal/queries"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// openAddUserWindow opens a new window for adding a new password.
//
// Args:
//
//	a: The Fyne application instance.
func openAddUserWindow(a fyne.App) {
	addWindow := a.NewWindow("Add New Password")
	addWindow.Resize(fyne.NewSize(400, 300))
	addWindow.CenterOnScreen()

	titleLabel := widget.NewLabel("Add New Password")
	titleLabel.TextStyle.Bold = true
	titleLabel.Importance = widget.HighImportance

	usernameLabel := widget.NewLabel("Username:")
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter username")

	passwordLabel := widget.NewLabel("Password:")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	statusLabel := widget.NewLabel("")

	submitBtn := widget.NewButton("Add Password", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text

		if username == "" || password == "" {
			statusLabel.SetText("All fields are required")
			statusLabel.Importance = widget.DangerImportance
			statusLabel.Refresh()
			return
		}

		queries.AddNewPassword(username, password)

		addWindow.Close()
		refreshUserList(a)
	})
	submitBtn.Importance = widget.HighImportance

	cancelBtn := widget.NewButton("Cancel", func() {
		addWindow.Close()
	})

	buttonContainer := container.NewHBox(
		submitBtn,
		cancelBtn,
	)

	form := container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		usernameLabel,
		usernameEntry,
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

	addWindow.SetContent(content)
	addWindow.Show()
}
