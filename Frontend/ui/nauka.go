package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildNauka() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Научная деятельность"),
		widget.NewButton("Добавить публикацию", func() {}),
		widget.NewButton("Зарегистрировать участие в конференции", func() {}),
	)
}
