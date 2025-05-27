package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildUsers() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Управление пользователями"),
		widget.NewButton("Администратор — полный доступ", func() {}),
		widget.NewButton("Преподаватель — доступ к своим данным", func() {}),
		widget.NewButton("Зав. кафедрой — анализ всех данных", func() {}),
	)
}
