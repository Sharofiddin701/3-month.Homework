package service

import (
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/storage"
	"context"
	"fmt"
)

type CustomerService struct {
	storage storage.IStorage
}

func NewCustomerService(storage storage.IStorage) CustomerService {
	return CustomerService{
		storage: storage,
	}
}
func (u CustomerService) Create(ctx context.Context, Customer models.CreateCustomer) (string, error) {

	pKey, err := u.storage.Customer().Create(ctx, Customer)
	if err != nil {
		fmt.Println("ERROR in service layer while creating Customer", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u CustomerService) Update(ctx context.Context, Customer models.GetCustomer) (string, error) {

	pKey, err := u.storage.Customer().Update(ctx, Customer)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Customer", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u CustomerService) Delete(ctx context.Context, id string) error {

	err := u.storage.Customer().Delete(ctx, id)
	if err != nil {
		fmt.Println("error service delete Customer", err.Error())
		return err
	}

	return nil
}

func (u CustomerService) GetAllCustomers(ctx context.Context, Customer models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {

	pKey, err := u.storage.Customer().GetAllCustomers(ctx, Customer)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling Customer", err.Error())
		return models.GetAllCustomersResponse{}, err
	}

	return pKey, nil
}

func (u CustomerService) GetByIDCustomer(ctx context.Context, Id string) (models.GetCustomer, error) {
	pKey, err := u.storage.Customer().GetByID(ctx, Id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyID Customer", err.Error())
		return models.GetCustomer{}, err
	}

	return pKey, nil
}

func (u CustomerService) GetAllCustomerCars(ctx context.Context,Customer models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error) {
	pKey, err := u.storage.Customer().GetAllCustomerCars(ctx, Customer)
	if err != nil {
		fmt.Println("ERROR in service layer while getalling Customer", err.Error())
		return models.GetAllCustomerCarsResponse{}, err
	}

	return pKey, nil
}


func (u CustomerService) UpdatePassword(ctx context.Context, Customer models.PasswordCustomer) (string, error) {

	pKey, err := u.storage.Customer().UpdatePassword(ctx, Customer)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Customer", err.Error())
		return "", err
	}

	return pKey, nil
}