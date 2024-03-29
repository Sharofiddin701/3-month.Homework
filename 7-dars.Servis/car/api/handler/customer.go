package handler

import (
	"context"
	"fmt"
	"net/http"
	_ "rent-car/api/docs"
	"rent-car/api/models"
	"rent-car/pkg/check"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCustomer godoc
// @Router 		/customer [POST]
// @Summary 	create a customer
// @Description This api is creates a new customer and returns its id
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		customer body  models.CreateCustomer true "car"
// @Success		200  {object}  models.Customer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCustomer(c *gin.Context) {
	customer := models.CreateCustomer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateGmailAddress(customer.Gmail); err != nil {
		handleResponse(c, "error while validating email address", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePhoneNumber(customer.Phone); err != nil {
		handleResponse(c, "error while validating phoneNumber", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Customer().Create(context.Background(),customer)
	if err != nil {
		handleResponse(c, "error while creating customer", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "Created Succesfully", http.StatusOK, id)
}

// UpdateCustomer godoc
// @Router                /customer/{id} [PUT]
// @Summary 			  update a customer
// @Description:          this api updates customer information
// @Tags 			      customer
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Customer ID"
// @Param       		  car body models.Customer true "car"
// @Success 		      200 {object} models.Customer
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateCustomer(c *gin.Context) {
	customer := models.Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}
	customer.Id = c.Param("id")
	err := uuid.Validate(customer.Id)
	if err != nil {
		handleResponse(c, "error while validating", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Services.Customer().Update(context.Background(), customer)
	if err != nil {
		handleResponse(c, "error while updating customer", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "Updated Successfully", http.StatusOK, id)
}

// GetAllCustomer godoc
// @Router 			/customer [GET]
// @Summary 		get all customer
// @Description 	This API returns customer list
// @Tags 			customer
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllCustomersResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllCustomer(c *gin.Context) {
	var (
		request = models.GetAllCustomersRequest{}
	)
	
	
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {

		handleResponse(c, "error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	customers, err := h.Services.Customer().GetAll(context.Background(), request)
	if err != nil {
		handleResponse(c, "error while getting customers,", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, customers)
}

// GetByIDCustomer godoc
// @Router       /customer/{id} [GET]
// @Summary      return a customer by ID
// @Description  Retrieves a customer by its ID
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        id path string true "Customer ID"
// @Success      200 {object} models.Customer
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDCustomer(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id", id)

	customer, err := h.Services.Customer().GetByID(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while getting customer by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, customer)
}

// DeleteCustomer godoc
// @Router          /customer/{id} [DELETE]
// @Summary         delete a customer by ID
// @Description     Deletes a customer by its ID
// @Tags            customer
// @Accept          json
// @Produce         json
// @Param           id path string true "Customer ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	err = h.Services.Customer().Delete(context.Background(), id)
	if err != nil {
		handleResponse(c, "Error while deleting customer", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "Deleted successfully", http.StatusOK, id)
}

func (h Handler) GetCustomerCars(c *gin.Context)  {
	var (
		request = models.GetAllCustomersRequest{}
	)

	request.Search = c.Param("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	customers, err := h.Store.Customer().GetCustomerCars(context.Background(), request)
	if err != nil {
		handleResponse(c, "error while getting customers", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, customers)
}

func (h Handler) GetByIDCustomeCarr(c *gin.Context) {
	id := c.Param("id")

	customer, err := h.Services.Customer().GetByIDCustomerCar(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while getting customerCar by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, customer)
}

// func (h Handler) GetCustomerCarss(c *gin.Context){
// 	id := c.Param("id")

// 	customer, err := h.Services.Customer().GetCustomerCarss(context.Background(), id)
// 	if err != nil {
// 		handleResponse(c, "error while getting customerCar by id", http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	handleResponse(c, "", http.StatusOK, customer)
// }