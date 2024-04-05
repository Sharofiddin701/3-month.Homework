package service

import (
	"clone/rent_car_us/storage"
)

type IServiceManager interface {
	Car() carService
	Customer() CustomerService
	Order() OrderService
}

type Service struct {
	carService carService
	CustomerService CustomerService
	OrderService OrderService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.carService = NewCarService(storage)
	services.CustomerService = NewCustomerService(storage)
	services.OrderService = NewOrderService(storage)

	return services
}

func (s Service) Car() carService {
	return s.carService
}

func (s Service) Customer() CustomerService {
	return s.CustomerService
}


func (s Service) Order() OrderService {
	return s.OrderService
}
