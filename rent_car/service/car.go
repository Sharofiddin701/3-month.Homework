package service

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"rent-car/storage"
)


type carService struct {
	storage storage.IStorage
}

func NewCarService(storage storage.IStorage) carService {
	return carService{
		storage: storage,
	}
}

func (u carService) Create(ctx context.Context, car models.Car) (string,error) {
	pkey,err := u.storage.Car().Create(ctx,car)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return "", err
	}
	return pkey,nil
}

func (u carService) Update(ctx context.Context, car models.Car) (string,error) {
	pkey, err := u.storage.Car().Update(ctx,car)
	if err != nil {
		fmt.Println("ERROR in service layer while updating car",err.Error())
		return "",err
	}
	return pkey,nil
}

func (u carService) GetByIDCar(ctx context.Context, id string) (models.Car,error) {
	car, err := u.storage.Car().GetByID(ctx,id)
	if err != nil {
		fmt.Println("ERROR in service layer while getting by id car",err.Error())
		return models.Car{},err
	}
	return car,nil
}

func (u carService) Delete(ctx context.Context, id string) (error) {
	err := u.storage.Car().Delete(ctx,id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting car",err.Error())
		return err
	}
	return nil
}

func (u carService) GetCarAll(ctx context.Context,car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	cars, err := u.storage.Car().GetAll(ctx,car)
	if err != nil {
		fmt.Println("ERROR in service layer while getting all cars",err.Error())
		return cars,err
	}
	return cars,nil
}

func (u carService) GetAvaibleCars(ctx context.Context,car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	cars, err := u.storage.Car().GetAvaibleCars(ctx,car)
	if err != nil {
		fmt.Println("ERROR in service layer while getting all cars with no customers",err.Error())
		return cars,err
	}
	return cars,nil
}

 