package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type adminService struct {
	storage storage.IStorage
}

func NewAdminService(storage storage.IStorage) adminService {
	return adminService{
		storage: storage,
	}
}
func (u adminService) Create(ctx context.Context, admin models.Admin) (models.Admin, error) {

	pKey, err := u.storage.Admin().Create(ctx, admin)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Admin{}, err
	}

	return pKey, nil
}
func (u adminService) Update(ctx context.Context, admin models.Admin) (models.Admin, error) {

	pKey, err := u.storage.Admin().Update(ctx, admin)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return  models.Admin{}, err
	}

	return pKey, nil
}
func (u adminService) GetAll(ctx context.Context, admin models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error) {

	pKey, err := u.storage.Admin().GetAll(ctx, admin)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling car", err.Error())
		return models.GetAllAdminsResponse{}, err
	}

	return pKey, nil
}
func (u adminService) GetByID(ctx context.Context, id string) (models.Admin, error) {

	pKey, err := u.storage.Admin().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID car", err.Error())
		return models.Admin{}, err
	}

	return pKey, nil
}
func (u adminService) Delete(ctx context.Context, id string) error {

	err := u.storage.Admin().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting car", err.Error())
		return err
	}

	return nil
}
