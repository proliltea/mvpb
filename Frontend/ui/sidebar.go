package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildSidebar(content *fyne.Container) *fyne.Container {
	sidebar := container.NewVBox(
		widget.NewLabel("Разделы"),
		widget.NewButton("Учебная деятельность", func() {
			content.Objects = []fyne.CanvasObject{BuildUchebna()}
			content.Refresh()
		}),
		widget.NewButton("Учебно-методическая", func() {
			content.Objects = []fyne.CanvasObject{BuildMetodichka()}
			content.Refresh()
		}),
		widget.NewButton("Научная деятельность", func() {
			content.Objects = []fyne.CanvasObject{BuildNauka()}
			content.Refresh()
		}),
		widget.NewButton("ГИА", func() {
			content.Objects = []fyne.CanvasObject{BuildGIA()}
			content.Refresh()
		}),
		widget.NewButton("Пользователи и доступ", func() {
			content.Objects = []fyne.CanvasObject{BuildUsers()}
			content.Refresh()
		}),
	)
	return sidebar
}
