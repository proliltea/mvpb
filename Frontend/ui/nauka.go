package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildNaukaView(win fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Научная деятельность"),
		widget.NewButton("Зарегистрировать публикацию", func() {
			dialog.ShowInformation("Публикация", "Публикация успешно зарегистрирована (имитация)", win)
		}),
		widget.NewButton("Регистрация участия в конференции", func() {
			dialog.ShowInformation("Конференция", "Участие в конференции зарегистрировано (имитация)", win)
		}),
	)
}
