package storage

import (
	"context"
	"rent-car/api/models"
)

type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	Order() IOrderStorage
}

type ICarStorage interface {
	Create(context.Context,models.Car) (string, error)
	GetByID(ctx context.Context,id string) (models.Car, error)
	GetAvaibleCars(ctx context.Context,req models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	GetAll(context.Context,models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	Update(context.Context,models.Car) (string, error)
	Delete(ctx context.Context,id string) error
}

type ICustomerStorage interface {
	Create(context.Context,models.Customer) (string, error)
	GetByID(ctx context.Context,id string) (models.Customer, error)
	GetAllCustomer(ctx context.Context,req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	UpdateCustomer(context.Context,models.Customer) (string, error)
	Delete(ctx context.Context,id string) error
}

type IOrderStorage interface {
	Create(context.Context,models.CreateOrder) (string, error)
	GetByID(ctx context.Context,id string) (models.OrderAll, error)
	GetAll(ctx context.Context,request models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error)
	Update(context.Context,models.UpdateOrder) (string, error)
	Delete(ctx context.Context,id string) error
}

