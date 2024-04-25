package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type groupService struct {
	storage storage.IStorage
}

func NewGroupService(storage storage.IStorage) groupService {
	return groupService{
		storage: storage,
	}
}
func (u groupService) Create(ctx context.Context, group models.Group) (models.Group, error) {

	pKey, err := u.storage.Group().Create(ctx, group)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Group{}, err
	}

	return pKey, nil
}
func (u groupService) Update(ctx context.Context, group models.Group) (models.Group, error) {

	pKey, err := u.storage.Group().Update(ctx, group)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Group{}, err
	}

	return pKey, nil
}
func (u groupService) GetAll(ctx context.Context, group models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error) {

	pKey, err := u.storage.Group().GetAll(ctx, group)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling car", err.Error())
		return models.GetAllGroupsResponse{}, err
	}

	return pKey, nil
}
func (u groupService) GetByID(ctx context.Context, id string) (models.Group, error) {

	pKey, err := u.storage.Group().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID car", err.Error())
		return models.Group{}, err
	}

	return pKey, nil
}
func (u groupService) Delete(ctx context.Context, id string) error {

	err := u.storage.Admin().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting car", err.Error())
		return err
	}

	return nil
}
