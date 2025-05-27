package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildMetodichka() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Учебно-методическая деятельность"),
		widget.NewButton("Загрузить файл дисциплины", func() {}),
		widget.NewButton("Посмотреть файл", func() {}),
	)
}
