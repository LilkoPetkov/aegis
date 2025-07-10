package ui

import (
	"aegis/internal/queries"

	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"image/color"
)

func createUserCards(allUsers []map[string]string, a fyne.App, db *sql.DB) []fyne.CanvasObject {
	userCards := []fyne.CanvasObject{}

	for _, user := range allUsers {
		card := createUserCard(user, a, db)
		userCards = append(userCards, card)

		if len(userCards) < len(allUsers) {
			userCards = append(userCards, widget.NewSeparator())
		}
	}

	return userCards
}

func createUserCard(user map[string]string, a fyne.App, db *sql.DB) *fyne.Container {
	cardBg := canvas.NewLinearGradient(
		color.NRGBA{R: 80, G: 132, B: 152, A: 255},
		color.NRGBA{R: 102, G: 38, B: 75, A: 255},
		45,
	)

	usernameLabel := widget.NewLabel(user["username"])
	usernameLabel.TextStyle.Bold = true
	usernameLabel.Importance = widget.MediumImportance

	usernameIcon := widget.NewIcon(theme.AccountIcon())
	usernameContainer := container.NewBorder(nil, nil, usernameIcon, nil, usernameLabel)

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetText(user["password_ciphertext"])
	passwordEntry.Disable()

	passwordIcon := widget.NewIcon(theme.VisibilityOffIcon())
	passwordContainer := container.NewBorder(nil, nil, passwordIcon, nil, passwordEntry)

	copyBtn := widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {
		Copy(a, user["password_ciphertext"])()
	})
	copyBtn.Importance = widget.MediumImportance

	editBtn := widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
		openNewWindowForPasswordUpdate(a, db, user["username"])
	})

	deleteBtn := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
		queries.DeleteUserByPasswordHash(db, user["username"])
		refreshUserList(db, a)
	})
	deleteBtn.Importance = widget.DangerImportance

	buttonContainer := container.NewGridWithColumns(3,
		copyBtn,
		editBtn,
		deleteBtn,
	)

	cardContent := container.NewVBox(
		usernameContainer,
		container.NewPadded(widget.NewSeparator()),
		passwordContainer,
		container.NewPadded(widget.NewSeparator()),
		buttonContainer,
	)

	card := container.NewStack(
		cardBg,
		container.NewPadded(cardContent),
	)

	return card
}

func createErrorCard(message string) *fyne.Container {
	errorBg := canvas.NewLinearGradient(
		color.NRGBA{R: 140, G: 60, B: 60, A: 255},
		color.NRGBA{R: 120, G: 50, B: 50, A: 255},
		45,
	)

	errorIcon := widget.NewIcon(theme.ErrorIcon())
	errorLabel := widget.NewLabel(message)
	errorLabel.Alignment = fyne.TextAlignCenter

	errorContent := container.NewBorder(
		nil, nil,
		errorIcon,
		nil,
		errorLabel,
	)

	return container.NewStack(
		errorBg,
		container.NewPadded(errorContent),
	)
}

func createEmptyStateCard() *fyne.Container {
	emptyBg := canvas.NewLinearGradient(
		color.NRGBA{R: 70, G: 76, B: 90, A: 255},
		color.NRGBA{R: 60, G: 66, B: 80, A: 255},
		45,
	)

	emptyIcon := widget.NewIcon(theme.InfoIcon())
	emptyLabel := widget.NewLabel("No passwords stored yet")
	emptyLabel.Alignment = fyne.TextAlignCenter

	hintLabel := widget.NewLabel("Click 'Add New Password' to get started")
	hintLabel.Alignment = fyne.TextAlignCenter

	emptyContent := container.NewVBox(
		container.NewBorder(nil, nil, emptyIcon, nil, emptyLabel),
		hintLabel,
	)

	return container.NewStack(
		emptyBg,
		container.NewPadded(emptyContent),
	)
}
