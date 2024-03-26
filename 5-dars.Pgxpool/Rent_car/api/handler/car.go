package handler

import (
	_ "clone/rent_car_us/api/docs"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/pkg/check"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCar godoc
// @Router 		/car [POST]
// @Summary 	create a car
// @Description This api is creates a new car and returns it's id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car body models.Car true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
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

// UpdateCar godoc
// @Router   /car [PUT]
// @Summary  update a car
// @Description This api is creates a new car and returns it's id
// @Tags   car
// @Accept  json
// @Produce  json
// @Param  id path string true "car"
// @Param  car body models.Car true "car"
// @Success  200  {object}  models.Car
// @Failure  400  {object}  models.Response
// @Failure  404  {object}  models.Response
// @Failure  500  {object}  models.Response
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

// GetAllCars godoc
// @Router   /car [Get]
// @Summary     get all a car
// @Description This api is creates a new car and returns it's id
// @Tags   car
// @Accept  json
// @Produce  json
// @Param       id path string true "Car ID"
// @Param  car body models.Car true "car"
// @Success  200  {object}  models.Car
// @Failure  400  {object}  models.Response
// @Failure  404  {object}  models.Response
// @Failure  500  {object}  models.Response
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

// DeleteCar godoc
// @Router   /car [DELETE]
// @Summary  delete a car
// @Description This api is creates a new car and returns it's id
// @Tags   car
// @Accept  json
// @Produce  json
// @Param       id path string true "Car ID"
// @Param  car body models.Car true "car"
// @Success  200  {object}  models.Car
// @Failure  400  {object}  models.Response
// @Failure  404  {object}  models.Response
// @Failure  500  {object}  models.Response
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

// GetByIDCar godoc
// @Router   /car [Get]
// @Summary  get a car by id
// @Description This api is creates a new car and returns it's id
// @Tags   car
// @Accept  json
// @Produce  json
// @Param       id path string true "Car ID"
// @Param  car body models.Car true "car"
// @Success  200  {object}  models.Car
// @Failure  400  {object}  models.Response
// @Failure  404  {object}  models.Response
// @Failure  500  {object}  models.Response
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
