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

func ChatTab(c *tclient.TelegramClient, window fyne.Window) fyne.CanvasObject {
	chatNameEntry := widget.NewEntry()
	chatNameEntry.SetPlaceHolder("Please provide chat username !(without @)")

	results := widget.NewMultiLineEntry()
	results.SetPlaceHolder("Results: ")
	results.Wrapping = fyne.TextWrapWord
	results.SetMinRowsVisible(20)

	search := func() {
		chatName := chatNameEntry.Text

		if !ValidateTelegramUsername(chatName) {
			results.SetText("Please enter a valid username")
			return
		}

		users, err := c.UsersFromSuper(chatName)
		if err != nil {
			results.SetText("Error: " + err.Error())
			return
		}

		formated := formatChatMembers(users)

		results.SetText(formated)
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

			_, err = writer.Write([]byte(results.Text))
			if err != nil {
				dialog.ShowError(err, window)
			}
		}, window)

		saveDialog.SetFileName("chat_members.txt")
		saveDialog.Show()
	})

	clearButton := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		chatNameEntry.SetText("")
		results.SetText("")
	})

	return container.NewVBox(
		widget.NewLabelWithStyle("Chat members search", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		chatNameEntry,
		container.NewHBox(searchButton, saveButton, clearButton),
		results,
	)
}

func formatChatMembers(members []string) string {
	if len(members) == 0 {
		return "No members found"
	}

	result := "Chat Members:\n"
	for _, member := range members {
		result += fmt.Sprintf("%s\n", member)
	}
	return result
}
