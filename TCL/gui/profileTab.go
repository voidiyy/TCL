package gui

import (
	"Hook_TCL/internal/tclient"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ProfileTab(c *tclient.TelegramClient, window fyne.Window) fyne.CanvasObject {

	nameLabel := widget.NewLabel("Name: ")
	usernameLabel := widget.NewLabel("Username: ")
	phoneLabel := widget.NewLabel("Phone number: ")
	statusLabel := widget.NewLabel("Status: ")

	user, err := c.Cl.GetMe()
	if err != nil {
		statusLabel.SetText(err.Error())
		return nil
	}

	nameLabel.SetText(fmt.Sprintf("Name: %s", user.FirstName))
	usernameLabel.SetText(fmt.Sprintf("Username: %v", user.Usernames.ActiveUsernames[0]))
	phoneLabel.SetText(fmt.Sprintf("Phone number: %s", user.PhoneNumber))
	statusLabel.SetText(fmt.Sprintf("Status: %s", user.Status.UserStatusType()))

	// Logout button
	logoutButton := widget.NewButton("Logout", func() {
		_, err := c.Cl.LogOut()
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		dialog.ShowInformation("Logout", "Successfully logged out", window)
	})

	// Buttons for additional user actions
	addUserButton := widget.NewButton("Add User", func() {
		dialog.ShowInformation("Add User", "User added successfully", window)
	})

	switchUserButton := widget.NewButton("Switch User", func() {
		dialog.ShowInformation("Switch User", "Switched to another user", window)
	})

	// Layout organization
	return container.NewVBox(
		widget.NewLabelWithStyle("Profile Details", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewGridWithColumns(2, nameLabel, usernameLabel),
		phoneLabel,
		statusLabel,
		container.NewGridWithColumns(2, logoutButton, switchUserButton),
		addUserButton,
	)
}
