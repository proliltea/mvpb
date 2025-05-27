package ui

import (
	"strconv"

	"github.com/antonminin/borodyansky/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildEducationView() fyne.CanvasObject {
	studentsBox := container.NewVBox()
	refreshBtn := widget.NewButton("Загрузить студентов", func() {
		studentsBox.Objects = nil
		students, err := api.GetAllStudents()
		if err != nil {
			studentsBox.Add(widget.NewLabel("Ошибка загрузки"))
			return
		}
		for _, s := range students {
			studentsBox.Add(widget.NewLabel(s.FullName + " (" + s.Group + ")"))
		}
		studentsBox.Refresh()
	})

	addGradeBtn := widget.NewButton("Добавить оценку", func() {
		studentID := widget.NewEntry()
		subject := widget.NewEntry()
		score := widget.NewEntry()

		form := dialog.NewForm("Добавить оценку", "Отправить", "Отмена",
			[]*widget.FormItem{
				widget.NewFormItem("ID студента", studentID),
				widget.NewFormItem("Предмет", subject),
				widget.NewFormItem("Оценка", score),
			},
			func(b bool) {
				if b {
					id, _ := strconv.Atoi(studentID.Text)
					sc, _ := strconv.ParseFloat(score.Text, 64)
					api.AddGrade(api.Grade{
						StudentID: uint(id),
						Subject:   subject.Text,
						Score:     sc,
						Passed:    sc >= 3.0,
					})
				}
			}, nil,
		)

		form.Show()
	})

	return container.NewVBox(
		widget.NewLabel("Учебная деятельность"),
		refreshBtn,
		addGradeBtn,
		studentsBox,
	)
}
