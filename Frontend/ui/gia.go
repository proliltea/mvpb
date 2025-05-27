package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildGIA() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Государственная Итоговая Аттестация"),
		widget.NewButton("Регистрация тем дипломов", func() {}),
		widget.NewButton("Привязать тему к научному руководителю", func() {}),
	)
}
