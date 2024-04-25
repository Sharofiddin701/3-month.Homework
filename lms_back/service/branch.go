package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type branchService struct {
	storage storage.IStorage
}

func NewBranchService(storage storage.IStorage) branchService {
	return branchService{
		storage: storage,
	}
}
func (u branchService) Create(ctx context.Context, branch models.Branch) (models.Branch, error) {

	pKey, err := u.storage.Branches().Create(ctx, branch)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Branch{}, err
	}

	return pKey, nil
}
func (u branchService) Update(ctx context.Context, branch models.Branch) (models.Branch, error) {

	pKey, err := u.storage.Branches().Update(ctx, branch)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Branch{}, err
	}

	return pKey, nil
}
func (u branchService) GetAll(ctx context.Context, branch models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error) {

	pKey, err := u.storage.Branches().GetAll(ctx, branch)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling car", err.Error())
		return models.GetAllBranchesResponse{}, err
	}

	return pKey, nil
}
func (u branchService) GetByID(ctx context.Context, id string) (models.Branch, error) {

	pKey, err := u.storage.Branches().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID car", err.Error())
		return models.Branch{}, err
	}

	return pKey, nil
}
func (u branchService) Delete(ctx context.Context, id string) error {

	err := u.storage.Admin().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting car", err.Error())
		return err
	}

	return nil
}
