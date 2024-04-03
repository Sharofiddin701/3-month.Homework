package service

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"rent-car/storage"
)


type orderService struct {
	storage storage.IStorage
}

func NewOrderService(storage storage.IStorage) orderService {
	return orderService{
		storage: storage,
	}
}

func (os orderService) Create(ctx context.Context, order models.CreateOrder) (string,error) {
	pkey,err := os.storage.Order().Create(ctx,order)
	if err != nil {
		fmt.Println("ERROR in service layer while creating customers", err.Error())
		return "", err
	}
	return pkey,nil
}

func (os orderService) Update(ctx context.Context, order models.UpdateOrder) (string,error) {
	pkey, err := os.storage.Order().Update(ctx,order)
	if err != nil {
		fmt.Println("ERROR in service layer while updating customers",err.Error())
		return "",err
	}
	return pkey,nil
}

func (os orderService) GetByIDOrder(ctx context.Context, id string) (models.OrderAll,error) {
	order, err := os.storage.Order().GetByID(ctx,id)
	if err != nil {
		fmt.Println("ERROR in service layer while getting by id customers",err.Error())
		return models.OrderAll{},err
	}
	return order,nil
}

func (os orderService) Delete(ctx context.Context, id string) (error) {
	err := os.storage.Order().Delete(ctx,id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting customers",err.Error())
		return err
	}
	return nil
}

func (os orderService) GetOrderAll(ctx context.Context,order models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {
	orders, err := os.storage.Order().GetAll(ctx,order)
	if err != nil {
		fmt.Println("ERROR in service layer while getting all customers",err.Error())
		return orders,err
	}
	return orders,nil
}