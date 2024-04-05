package service

import (
	"context"
	"fmt"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/storage"
)

type OrderService struct {
	storage storage.IStorage
}

func NewOrderService(storage storage.IStorage) OrderService {
	return OrderService{
		storage: storage,
	}
}
func (u OrderService) Create(ctx context.Context, Order models.CreateOrder) (string, error) {

	pKey, err := u.storage.Order().CreateOrder(ctx, Order)
	if err != nil {
		fmt.Println("ERROR in service layer while creating Order", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u OrderService) Update(ctx context.Context, Order models.GetOrder) (string, error) {

	pKey, err := u.storage.Order().UpdateOrder(ctx, Order)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Order", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u OrderService) Delete(ctx context.Context,id string) error {

	err := u.storage.Order().DeleteOrder(ctx, id)
	if err != nil {
		fmt.Println("error service delete Order", err.Error())
		return err
	}

	return nil
}

func (u OrderService) GetAllOrders(ctx context.Context, Order models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {

	pKey, err := u.storage.Order().GetAll(ctx, Order)
	if err != nil {
	 fmt.Println("ERROR in service layer while getalling Order", err.Error())
	 return models.GetAllOrdersResponse{}, err
	}
   
	return pKey, nil
   }

   func (u OrderService) GetByIDOrder(ctx context.Context,Id string) (models.GetOrder, error) {
	pKey, err := u.storage.Order().GetOne(ctx,Id)
	if err != nil {
	 fmt.Println("ERROR in service layer while getbyID Order", err.Error())
	 return models.GetOrder{}, err
	}
   
	return pKey, nil
   }