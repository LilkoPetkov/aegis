package ui

import (
	"aegis/internal/queries"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"log"
)

func Copy(app fyne.App, password string) func(...int) {
	return func(...int) {
		app.Clipboard().SetContent(password)
	}
}
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

func refreshUserList(a fyne.App) {
	newContent := buildUserList(a)
	userListContainer.Objects = newContent.Objects
	scrollContainer.Content = userListContainer
	scrollContainer.Refresh()
}
