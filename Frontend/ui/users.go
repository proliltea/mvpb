package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildUsersView() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Управление пользователями и доступом"),
		widget.NewLabel("Роли:"),
		widget.NewLabel(" - Администратор — полный доступ"),
		widget.NewLabel(" - Преподаватель — доступ к своим данным"),
		widget.NewLabel(" - Заведующий кафедрой — просмотр и анализ всех данных"),
		widget.NewLabel("Разграничение прав имитируется"),
	)
}
