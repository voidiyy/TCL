package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateSuperGroupsTab(output *widget.Label) fyne.CanvasObject {
	// Поля вводу
	groupIDEntry := widget.NewEntry()
	groupIDEntry.SetPlaceHolder("Введіть ID супер групи")

	// Кнопки
	getSuperGroupUsersButton := widget.NewButton("Отримати користувачів супер групи", func() {
		groupID := groupIDEntry.Text
		if groupID == "" {
			output.SetText("Будь ласка, введіть ID супер групи.")
			return
		}
		// Тут можна додати логіку отримання користувачів супер групи
		output.SetText(fmt.Sprintf("Отримано користувачів супер групи з ID: %s", groupID))
	})

	// Вертикальний контейнер для вкладки "Супергрупи"
	superGroupsContainer := container.NewVBox(
		widget.NewLabelWithStyle("ID супер групи:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		groupIDEntry,
		container.NewHBox(layout.NewSpacer(), getSuperGroupUsersButton),
	)

	return container.NewScroll(superGroupsContainer)
}
