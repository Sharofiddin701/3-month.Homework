package service

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"rent-car/storage"
)


type customerService struct {
	storage storage.IStorage
}

func NewCustomerService(storage storage.IStorage) customerService {
	return customerService{
		storage: storage,
	}
}

func (cs customerService) Create(ctx context.Context, customer models.Customer) (string,error) {
	pkey,err := cs.storage.Customer().Create(ctx,customer)
	if err != nil {
		fmt.Println("ERROR in service layer while creating customer", err.Error())
		return "", err
	}
	return pkey,nil
}

func (cs customerService) Update(ctx context.Context, customer models.Customer) (string,error) {
	pkey, err := cs.storage.Customer().UpdateCustomer(ctx,customer)
	if err != nil {
		fmt.Println("ERROR in service layer while updating customer",err.Error())
		return "",err
	}
	return pkey,nil
}

func (cs customerService) GetByIDCar(ctx context.Context, id string) (models.Customer,error) {
	customer, err := cs.storage.Customer().GetByID(ctx,id)
	if err != nil {
		fmt.Println("ERROR in service layer while getting by id customer",err.Error())
		return models.Customer{},err
	}
	return customer,nil
}

func (cs customerService) Delete(ctx context.Context, id string) (error) {
	err := cs.storage.Customer().Delete(ctx,id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting customer",err.Error())
		return err
	}
	return nil
}

func (cs customerService) GetCarAll(ctx context.Context,customer models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	customers, err := cs.storage.Customer().GetAllCustomer(ctx,customer)
	if err != nil {
		fmt.Println("ERROR in service layer while getting all customers",err.Error())
		return customers,err
	}
	return customers,nil
}