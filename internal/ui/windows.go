package ui

import (
	"aegis/internal/queries"

	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"image/color"
)

var windowBg = canvas.NewLinearGradient(
	color.NRGBA{R: 189, G: 121, B: 108, A: 255},
	color.NRGBA{R: 108, G: 67, B: 96, A: 255},
	90,
)

func openAddUserWindow(a fyne.App, db *sql.DB) {
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

		queries.AddNewPassword(db, username, password)

		addWindow.Close()
		refreshUserList(db, a)
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

func openNewWindowForPasswordUpdate(a fyne.App, db *sql.DB, username string) {
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

		queries.EditUserPassword(db, password, username)

		updateWindow.Close()
		refreshUserList(db, a)
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
