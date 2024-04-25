package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type taskService struct {
	storage storage.IStorage
}

func NewTaskService(storage storage.IStorage) taskService {
	return taskService{
		storage: storage,
	}
}
func (u taskService) Create(ctx context.Context, task models.Task) (models.Task, error) {

	pKey, err := u.storage.Task().Create(ctx, task)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Task{}, err
	}

	return pKey, nil
}
func (u taskService) Update(ctx context.Context, task models.Task) (models.Task, error) {

	pKey, err := u.storage.Task().Update(ctx, task)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Task{}, err
	}

	return pKey, nil
}
func (u taskService) GetAll(ctx context.Context, task models.GetAllTasksRequest) (models.GetAllTasksResponse, error) {

	pKey, err := u.storage.Task().GetAll(ctx, task)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling car", err.Error())
		return models.GetAllTasksResponse{}, err
	}

	return pKey, nil
}
func (u taskService) GetByID(ctx context.Context, id string) (models.Task, error) {

	pKey, err := u.storage.Task().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID car", err.Error())
		return models.Task{}, err
	}

	return pKey, nil
}
func (u taskService) Delete(ctx context.Context, id string) error {

	err := u.storage.Task().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting car", err.Error())
		return err
	}

	return nil
}
