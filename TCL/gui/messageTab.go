package gui

import (
	"Hook_TCL/internal/tclient"
	"bufio"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
	"strings"
)

func MessagesTab(c *tclient.TelegramClient, window fyne.Window) fyne.CanvasObject {
	messageEntry := widget.NewMultiLineEntry()
	messageEntry.SetPlaceHolder("Enter the message you want to send")
	messageEntry.Wrapping = fyne.TextWrapWord

	recipientsEntry := widget.NewMultiLineEntry()
	recipientsEntry.SetPlaceHolder("Enter recipients, one per line")
	recipientsEntry.Wrapping = fyne.TextWrapWord

	results := widget.NewMultiLineEntry()
	results.SetPlaceHolder("Results will appear here")
	results.Wrapping = fyne.TextWrapWord

	send := func() {
		message := messageEntry.Text
		receiversText := recipientsEntry.Text

		if message == "" || receiversText == "" {
			results.SetText("Message or recipients cannot be empty.")
			return
		}

		receivers := strings.Split(receiversText, "\n")
		results.SetText("Sending messages...\n")

		resultsMap := c.SendMessage(message, receivers)
		var resultLines []string
		for receiver, result := range resultsMap {
			resultLines = append(resultLines, receiver+": "+result)
		}
		results.SetText(strings.Join(resultLines, "\n"))
	}

	loadButton := widget.NewButton("Load recipients from file", func() {
		dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				results.SetText("Error opening file: " + err.Error())
				return
			}
			if r == nil {
				return
			}
			defer r.Close()

			filePath := r.URI().Path()
			recipients, err := loadRecipientsFromFile(filePath)
			if err != nil {
				results.SetText("Error loading file: " + err.Error())
				return
			}

			recipientsEntry.SetText(strings.Join(recipients, "\n"))
		}, window).Show()
	})

	sendButton := widget.NewButton("Send", send)

	clearButton := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		messageEntry.SetText("")
		recipientsEntry.SetText("")
		results.SetText("")
	})

	return container.NewVBox(
		widget.NewLabelWithStyle("Send Messages", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		messageEntry,
		recipientsEntry,
		container.NewGridWithColumns(3, loadButton, sendButton, clearButton),
		results,
	)
}

func loadRecipientsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var recipients []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			recipients = append(recipients, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return recipients, nil
}
