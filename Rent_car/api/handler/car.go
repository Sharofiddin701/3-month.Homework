package handler

import (
	"strconv"
	"fmt"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/pkg/check"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) CreateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		handleResponse(c, "error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.Store.Car().Create(car)
	if err != nil {
		handleResponse(c, "error while creating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

func (h Handler) UpdateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		handleResponse(c, "error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())
		return
	}
	car.Id = c.Param("id")

	err := uuid.Validate(car.Id)
	if err != nil {
		handleResponse(c, "error while validating car id,id: "+car.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.Car().Update(car)
	if err != nil {
		handleResponse(c, "error while updating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

func (h Handler) GetAllCars(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	cars, err := h.Store.Car().GetAllCars(request)
	if err != nil {
		handleResponse(c, "error while gettign cars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, "", http.StatusOK, cars)
}

func (h Handler) DeleteCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Store.Car().Delete(id)
	if err != nil {
		handleResponse(c, "error while deleting car", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, id)
}

func (h Handler) GetByIDCar(c *gin.Context) {
 
	id := c.Param("id")
	fmt.Println("id: ", id)
   
	admin, err := h.Store.Car().GetByID(id)
	if err != nil {
	 handleResponse(c, "error while getting admin by id", http.StatusInternalServerError, err)
	 return
	}
	handleResponse(c, "", http.StatusOK, admin)
   }