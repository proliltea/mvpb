package services

import (
	"educational-service/models"
	"educational-service/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	return s.repo.GetAllStudents()
}

func (s *StudentService) GetStudentByID(id uint) (models.Student, error) {
	return s.repo.GetStudentByID(id)
}

func (s *StudentService) RegisterStudent(student *models.Student) error {
	return s.repo.CreateStudent(student)
}

func (s *StudentService) RecordGrade(grade *models.Grade) error {
	return s.repo.AddGrade(grade)
}
