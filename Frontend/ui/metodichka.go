package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildMetodichkaView(win fyne.Window) fyne.CanvasObject {
	discipline := widget.NewEntry()
	file := widget.NewEntry()

	uploadBtn := widget.NewButton("Имитация загрузки файла", func() {
		dialog.ShowInformation("Успешно", "Файл загружен для дисциплины: "+discipline.Text, win)
	})

	return container.NewVBox(
		widget.NewLabel("Учебно-методическая деятельность"),
		widget.NewForm(
			widget.NewFormItem("Дисциплина", discipline),
			widget.NewFormItem("Имя файла", file),
		),
		uploadBtn,
	)
}
