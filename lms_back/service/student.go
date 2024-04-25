package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type studentService struct {
	storage storage.IStorage
}

func NewStudentService(storage storage.IStorage) studentService {
	return studentService{
		storage: storage,
	}
}
func (u studentService) Create(ctx context.Context, student models.Student) (models.Student, error) {

	pKey, err := u.storage.Student().Create(ctx, student)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Student{}, err
	}

	return pKey, nil
}
func (u studentService) Update(ctx context.Context, student models.Student) (models.Student, error) {

	pKey, err := u.storage.Student().Update(ctx, student)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Student{}, err
	}

	return pKey, nil
}
func (u studentService) GetAll(ctx context.Context, student models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {

	pKey, err := u.storage.Student().GetAll(ctx, student)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling car", err.Error())
		return models.GetAllStudentsResponse{}, err
	}

	return pKey, nil
}
func (u studentService) GetByID(ctx context.Context, id string) (models.Student, error) {

	pKey, err := u.storage.Student().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID car", err.Error())
		return models.Student{}, err
	}

	return pKey, nil
}
func (u studentService) Delete(ctx context.Context, id string) error {

	err := u.storage.Student().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting car", err.Error())
		return err
	}

	return nil
}
