package main

import (
	"Hook_TCL/gui"
	"Hook_TCL/internal/tclient"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {

	fmt.Println("start")
	// Creating a themed application with a modern feel
	myApp := app.NewWithID("HookU TG")
	myApp.Settings() // You could allow user theme switching
	myWindow := myApp.NewWindow("HookU TG - Automate Your Telegram")

	cfg := tclient.NewConfig()

	// Telegram Client setup
	tcl, err := tclient.NewTelegramClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("second")

	// Main Tabs
	usersTabContent := gui.UserTab(tcl, myWindow)
	chatsTabContent := gui.ChatTab(tcl, myWindow)
	messagesTabContent := gui.MessagesTab(tcl, myWindow)

	// Creating a modern styled tab container
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Users", theme.AccountIcon(), usersTabContent),
		container.NewTabItemWithIcon("Chats", theme.ComputerIcon(), chatsTabContent),
		container.NewTabItemWithIcon("Messages", theme.MailSendIcon(), messagesTabContent),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	contentContainer := container.NewStack(tabs)

	fmt.Println("third")

	// Navigation Bar for switching between Profile and Main Tabs
	navBar := container.NewHBox(
		widget.NewButtonWithIcon("Profile", theme.MenuIcon(), func() {
			profileContent := gui.ProfileTab(tcl, myWindow)
			contentContainer.Objects = []fyne.CanvasObject{profileContent}
			contentContainer.Refresh()
		}),
		widget.NewButtonWithIcon("Back to Tabs", theme.NavigateBackIcon(), func() {
			contentContainer.Objects = []fyne.CanvasObject{tabs}
			contentContainer.Refresh()
		}),
	)

	mainContainer := container.NewBorder(navBar, nil, nil, nil, contentContainer)

	// Adding window settings
	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(900, 700))
	myWindow.SetFixedSize(true)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

	fmt.Println("last")

}
