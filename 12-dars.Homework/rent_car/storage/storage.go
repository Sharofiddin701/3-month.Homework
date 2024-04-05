package storage

import (
	"clone/rent_car_us/api/models"
	"context"
)
type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	Order() IOrderStorage
}

type ICarStorage interface {
	Create(context.Context,models.Car) (string, error)
	Update(context.Context,models.Car) (string, error)
	Delete(context.Context,string) error
	GetByID(context.Context,string) (models.Car, error)
	GetAllCars(context.Context, models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	GetAllCarsFree(context.Context, models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
}

type ICustomerStorage interface {
	Create(context.Context,models.CreateCustomer) (string, error)
	GetByID(context.Context,string) (models.GetCustomer, error)
	GetAllCustomers(context.Context,models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	GetAllCustomerCars(context.Context,models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error)
	Update(context.Context,models.GetCustomer) (string, error)
	Delete(context.Context,string) error

	UpdatePassword(context.Context,models.PasswordCustomer) (string, error)

	///GetCustomer(request models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)

}

type IOrderStorage interface {
	CreateOrder(context.Context,models.CreateOrder) (string, error)
	UpdateOrder(context.Context,models.GetOrder) (string, error)
	GetOne(context.Context,string) (models.GetOrder, error)
	GetAll(context.Context,models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error)
	DeleteOrder(context.Context,string) error
}