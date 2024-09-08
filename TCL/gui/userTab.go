package gui

import (
	"Hook_TCL/internal/tclient"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func UserTab(c *tclient.TelegramClient, window fyne.Window) fyne.CanvasObject {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Please provide user name")

	results := widget.NewMultiLineEntry()
	results.SetPlaceHolder("Result: ")
	results.Wrapping = fyne.TextWrapWord
	results.SetMinRowsVisible(20)

	search := func() {
		username := usernameEntry.Text
		if !ValidateTelegramUsername(username) {
			results.SetText("Please enter a valid Telegram username")
			return
		}

		info, err := c.UserInfo(username)
		if err != nil {
			results.SetText("Error: " + err.Error())
			return
		}

		res := formatUserInfo(info)

		results.SetText(res)
	}

	searchButton := widget.NewButton("Search", search)

	saveButton := widget.NewButton("Save result into a file", func() {
		if results.Text == "" {
			results.SetText("No data to save.")
			return
		}

		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if writer == nil {
				return
			}

			defer writer.Close()

			// Записати дані у файл
			_, err = writer.Write([]byte(results.Text))
			if err != nil {
				dialog.ShowError(err, window)
			}
		}, window)

		saveDialog.SetFileName("user_info.txt")
		saveDialog.Show()
	})

	clearButton := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		usernameEntry.SetText("")
		results.SetText("")
	})

	return container.NewVBox(
		widget.NewLabelWithStyle("User search", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		usernameEntry,
		container.NewHBox(searchButton, saveButton, clearButton),
		results,
	)
}

func formatUserInfo(user *tclient.Usr) string {
	return fmt.Sprintf(
		"ID: %d\nFirst Name: %s\nLast Name: %s\nUsernames: %v\nPhone Number: %s\nVerified: %s\nPremium: %s\nSupport: %s\nCan Be Called: %s\nSupports Video Calls: %s\nHas Private Calls: %s\nHas Private Forwards: %s\nRestricted Voice/Video Notes: %s\nNeeds Phone Privacy Exception: %s\n",
		user.ID,
		user.FirstName,
		user.LastName,
		user.Usernames,
		user.PhoneNumber,
		user.IsVerified,
		user.IsPremium,
		user.IsSupport,
		user.CanBeCalled,
		user.SupportVideoCalls,
		user.HasPrivateCalls,
		user.HasPrivateForwards,
		user.HasRestrictedVoiceAndVideoNoteMessages,
		user.NeedPhoneNumberPrivacyException,
	)
}
