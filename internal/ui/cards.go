package ui

import (
	"aegis/internal/queries"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"image/color"
)

// createUserCards creates a slice of Fyne canvas objects representing user cards.
//
// Args:
//
//	allUsers: A slice of maps, where each map represents a user and their data.
//	a: The Fyne application instance.
//
// Returns:
//
//	A slice of Fyne canvas objects.
func createUserCards(allUsers []map[string]string, a fyne.App) []fyne.CanvasObject {
	userCards := []fyne.CanvasObject{}

	for _, user := range allUsers {
		card := createUserCard(user, a)
		userCards = append(userCards, card)

		if len(userCards) < len(allUsers) {
			userCards = append(userCards, widget.NewSeparator())
		}
	}

	return userCards
}

// createUserCard creates a Fyne container representing a user card.
//
// Args:
//
//	user: A map representing a user and their data.
//	a: The Fyne application instance.
//
// Returns:
//
//	A Fyne container representing a user card.
func createUserCard(user map[string]string, a fyne.App) *fyne.Container {
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
		openPasswordUpdateWindow(a, user["username"])
	})

	deleteBtn := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
		queries.DeleteUserByPasswordHash(user["username"])
		refreshUserList(a)
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

// createErrorCard creates a Fyne container representing an error card.
//
// Args:
//
//	message: The error message to display.
//
// Returns:
//
//	A Fyne container representing an error card.
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

// createEmptyStateCard creates a Fyne container representing an empty state card.
//
// Returns:
//
//	A Fyne container representing an empty state card.
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
