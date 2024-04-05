package service

import (
	"context"
	"fmt"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/storage"
)

type carService struct {
	storage storage.IStorage
}

func NewCarService(storage storage.IStorage) carService {
	return carService{
		storage: storage,
	}
}
func (u carService) Create(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Create(ctx, car)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u carService) Update(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Update(ctx, car)
	if err != nil {
		fmt.Println("ERROR in service layer while updating car", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u carService) Delete(ctx context.Context,id string) error {

	err := u.storage.Car().Delete(ctx, id)
	if err != nil {
		fmt.Println("error service delete car", err.Error())
		return err
	}

	return nil
}

func (u carService) GetAllCars(ctx context.Context, car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {

	pKey, err := u.storage.Car().GetAllCars(ctx, car)
	if err != nil {
	 fmt.Println("ERROR in service layer while getalling car", err.Error())
	 return models.GetAllCarsResponse{}, err
	}
   
	return pKey, nil
   }

   func (u carService) GetByIDCar(ctx context.Context,Id string) (models.Car, error) {
	pKey, err := u.storage.Car().GetByID(ctx,Id)
	if err != nil {
	 fmt.Println("ERROR in service layer while getbyID car", err.Error())
	 return models.Car{}, err
	}
   
	return pKey, nil
   }



func (u carService) GetAllCarsFree(ctx context.Context, car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {

	pKey, err := u.storage.Car().GetAllCarsFree(ctx, car)
	if err != nil {
	 fmt.Println("ERROR in service layer while getalling car", err.Error())
	 return models.GetAllCarsResponse{}, err
	}
   
	return pKey, nil
   }
