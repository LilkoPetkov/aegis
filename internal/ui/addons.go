package ui

import (
	"aegis/internal/queries"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"log"
)

// Copy copies the given password to the clipboard.
//
// Args:
//
//	app: The Fyne application instance.
//	password: The password to copy.
//
// Returns:
//
//	A function that can be called to copy the password.
func Copy(app fyne.App, password string) func(...int) {
	return func(...int) {
		app.Clipboard().SetContent(password)
	}
}

// buildUserList builds the list of user cards to be displayed in the UI.
//
// Args:
//
//	a: The Fyne application instance.
//
// Returns:
//
//	A Fyne container with the user cards.
func buildUserList(a fyne.App) *fyne.Container {
	userData, err := queries.FetchUserData()
	if err != nil {
		log.Printf("Could not fetch user data: %s", err)
		errorCard := createErrorCard("Error loading users")
		return container.NewVBox(errorCard)
	}

	if len(userData) == 0 {
		emptyCard := createEmptyStateCard()
		return container.NewVBox(emptyCard)
	}

	for _, user := range userData {
		username := user["username"]
		decryptedPassword := queries.FetchPassword(username)
		user["password_ciphertext"] = decryptedPassword
	}

	userCards := createUserCards(userData, a)
	return container.NewVBox(userCards...)
}

// refreshUserList refreshes the list of user cards in the UI.
//
// Args:
//
//	a: The Fyne application instance.
func refreshUserList(a fyne.App) {
	newContent := buildUserList(a)
	userListContainer.Objects = newContent.Objects
	scrollContainer.Content = userListContainer
	scrollContainer.Refresh()
}
