package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type teacherService struct {
	storage storage.IStorage
}

func NewTeacherService(storage storage.IStorage) teacherService {
	return teacherService{
		storage: storage,
	}
}
func (u teacherService) Create(ctx context.Context, teacher models.Teacher) (models.Teacher, error) {

	pKey, err := u.storage.Teacher().Create(ctx, teacher)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Teacher{}, err
	}

	return pKey, nil
}
func (u teacherService) Update(ctx context.Context, teacher models.Teacher) (models.Teacher, error) {

	pKey, err := u.storage.Teacher().Update(ctx, teacher)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Teacher{}, err
	}

	return pKey, nil
}
func (u teacherService) GetAll(ctx context.Context, teacher models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {

	pKey, err := u.storage.Teacher().GetAll(ctx, teacher)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling car", err.Error())
		return models.GetAllTeachersResponse{}, err
	}

	return pKey, nil
}
func (u teacherService) GetByID(ctx context.Context, id string) (models.Teacher, error) {

	pKey, err := u.storage.Teacher().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID car", err.Error())
		return models.Teacher{}, err
	}

	return pKey, nil
}
func (u teacherService) Delete(ctx context.Context, id string) error {

	err := u.storage.Teacher().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting car", err.Error())
		return err
	}

	return nil
}
