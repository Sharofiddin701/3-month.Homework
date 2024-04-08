package handler

import (
	_ "clone/rent_car_us/api/docs"
	"clone/rent_car_us/api/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// CreateOrder godoc
// @Router 		/order [POST]
// @Summary 	create a order
// @Description This api is creates a new order and returns it's id
// @Tags 		order
// @Accept		json
// @Produce		json
// @Param		order body models.CreateOrder true "order"
// @Success		200  {object}  models.CreateOrder
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateOrder(c *gin.Context) {
	Order := models.CreateOrder{}
	if err := c.ShouldBindJSON(&Order); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Order().Create(context.Background(),Order)
	if err != nil {
		handleResponse(c, h.Log, "error while creating Order", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c,  h.Log,"Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// UpdateOrder godoc
// @Router 		/order/{id} [PUT]
// @Summary 	update a order
// @Description This api is update a  order and returns it's id
// @Tags 		order
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Param		order body models.GetOrder true "order"
// @Success		200  {object}  models.GetOrder
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateOrder(c *gin.Context) {
	Order := models.GetOrder{}

	if err := c.ShouldBindJSON(&Order); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	Order.Id = c.Param("id")

	err := uuid.Validate(Order.Id)
	if err != nil {
		handleResponse(c, h.Log, "error while validating Order id,id: "+Order.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Order().Update(context.Background(),Order)
	if err != nil {
		handleResponse(c, h.Log, "error while updating Order", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GETALLOrders godoc
// @Router 		/order [GET]
// @Summary 	Get order list
// @Description Get order list
// @Tags 		order
// @Accept		json
// @Produce		json
// @Param		page path string false "page"
// @Param		limit path string false "limit"
// @Param		search path string false "search"
// @Success		200  {object}  models.GetOrder
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllOrders(c *gin.Context) {
	var (
		request = models.GetAllOrdersRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c,  h.Log,"error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	Orders, err := h.Services.Order().GetAllOrders(context.Background(),request)
	if err != nil {
		handleResponse(c,  h.Log,"error while gettign Orders", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, h.Log, "", http.StatusOK, Orders)
}

// @Security ApiKeyAuth
// GETBYIDORDER godoc
// @Router 		/order/{id} [GET]
// @Summary 	Get order 
// @Description Get order
// @Tags 		order
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.GetOrder
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetOne(c *gin.Context) {
 
	id := c.Param("id")
	fmt.Println("id: ", id)
   
	admin, err := h.Services.Order().GetByIDOrder(context.Background(),id)
	if err != nil {
	 handleResponse(c, h.Log, "error while getting admin by id", http.StatusInternalServerError, err)
	 return
	}
	handleResponse(c, h.Log, "", http.StatusOK, admin)
   }

   func (h Handler) DeleteOrder(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Services.Order().Delete(context.Background(),id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting order", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "", http.StatusOK, id)
}
