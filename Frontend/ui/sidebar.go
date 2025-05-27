package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BuildSidebar(content *fyne.Container) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Модули"),
		widget.NewButton("Учебная деятельность", func() {
			content.Objects = []fyne.CanvasObject{BuildEducationView()}
			content.Refresh()
		}),
		widget.NewButton("Учебно-методическая", func() {
			content.Objects = []fyne.CanvasObject{widget.NewLabel("Методическая часть — без API")}
			content.Refresh()
		}),
		widget.NewButton("Научная деятельность", func() {
			content.Objects = []fyne.CanvasObject{widget.NewLabel("Кнопка участия в конференции")}
			content.Refresh()
		}),
		widget.NewButton("ГИА", func() {
			content.Objects = []fyne.CanvasObject{widget.NewLabel("Регистрация темы диплома")}
			content.Refresh()
		}),
		widget.NewButton("Пользователи", func() {
			content.Objects = []fyne.CanvasObject{widget.NewLabel("Управление ролями")}
			content.Refresh()
		}),
		layout.NewSpacer(),
	)
}
