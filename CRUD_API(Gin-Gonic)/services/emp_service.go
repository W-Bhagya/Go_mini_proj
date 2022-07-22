package services

import "crud.com/api/models"

type EmpService interface {
	CreateUser(*models.Employee) error
	GetUser(*string) (*models.Employee, error)
	GetAll() ([]*models.Employee, error)
	UpdateUser(*models.Employee) error
	DeleteUser(*string) error
}