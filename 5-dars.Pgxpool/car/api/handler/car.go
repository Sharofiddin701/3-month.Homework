package handler

import (
	"fmt"
	"net/http"
	_ "rent-car/api/docs"
	"rent-car/api/models"
	"rent-car/pkg/check"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCar godoc
// @Router 		/car [POST]
// @Summary 	create a car
// @Description This api is creates a new car and returns its id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car body models.CreateCar true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCar(c *gin.Context) {
	car := models.CreateCar{}

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
// @Router                /car/{id} [PUT]
// @Summary 			  update a car
// @Description:          this api updates car information
// @Tags 			      car
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Car ID"
// @Param       		  car body models.Car true "car"
// @Success 		      200 {object} models.Car
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
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
// @Router 			/car [GET]
// @Summary 		get all cars
// @Description 	This API returns car list
// @Tags 			car
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllCarsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
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
	cars, err := h.Store.Car().GetAll(request)
	if err != nil {
		handleResponse(c, "error while gettign cars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, "", http.StatusOK, cars)
}

// GetByIDCar godoc
// @Router       /car/{id} [GET]
// @Summary      return a car by ID
// @Description  Retrieves a car by its ID
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        id path string true "Car ID"
// @Success      200 {object} models.Car
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	car, err := h.Store.Car().GetByID(id)
	if err != nil {
		handleResponse(c, "error while getting car by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, car)
}

// DeleteCar godoc
// @Router          /car/{id} [DELETE]
// @Summary         delete a car by ID
// @Description     Deletes a car by its ID
// @Tags            car
// @Accept          json
// @Produce         json
// @Param           id path string true "Car ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
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

// func (h Handler) GetAllOrdersCars() {
// 	var (
// 		values  = r.URL.Query()
// 		search  string
// 		request = models.GetAllCarsRequest{}
// 	)
// 	if _, ok := values["search"]; ok {
// 		search = values["search"][0]
// 	}

// 	request.Search = search

// 	page, err := ParsePageQueryParam(r)
// 	if err != nil {
// 		fmt.Println("error while parsing page, err: ", err)
// 		handleResponse(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	limit, err := ParseLimitQueryParam(r)
// 	if err != nil {
// 		fmt.Println("error while parsing limit, err: ", err)
// 		handleResponse(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	fmt.Println("page: ", page)
// 	fmt.Println("limit: ", limit)

// 	request.Page = page
// 	request.Limit = limit
// 	cars, err := c.Store.Car().GetAllOrdersCars(request)
// 	if err != nil {
// 		fmt.Println("error while getting carsOrder, err: ", err)
// 		handleResponse(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	handleResponse(w, http.StatusOK, cars)
// }

// func (c Controller) GetByIDCustomeCar(w http.ResponseWriter, r *http.Request) {
// 	values := r.URL.Query()
// 	id := values["id"][0]

// 	customer, err := c.Store.Car().GetByIDCustomerCarr(id)
// 	if err != nil {
// 		fmt.Println("error while getting customerCar by id")
// 		handleResponse(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	handleResponse(w, http.StatusOK, customer)
// }
