package handler

import (
	"fmt"
	"clone/rent_car_us/api/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) CreateOrder(c *gin.Context) {
	Order := models.CreateOrder{}
	if err := c.ShouldBindJSON(&Order); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.Order().CreateOrder(Order)
	if err != nil {
		handleResponse(c, "error while creating Order", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

func (h Handler) UpdateOrder(c *gin.Context) {
	Order := models.GetOrder{}

	if err := c.ShouldBindJSON(&Order); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	Order.Id = c.Param("id")

	err := uuid.Validate(Order.Id)
	if err != nil {
		handleResponse(c, "error while validating Order id,id: "+Order.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.Order().UpdateOrder(Order)
	if err != nil {
		handleResponse(c, "error while updating Order", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

func (h Handler) GetAllOrders(c *gin.Context) {
	var (
		request = models.GetAllOrdersRequest{}
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
	Orders, err := h.Store.Order().GetAll(request)
	if err != nil {
		handleResponse(c, "error while gettign Orders", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, "", http.StatusOK, Orders)
}


func (h Handler) GetOne(c *gin.Context) {
 
	id := c.Param("id")
	fmt.Println("id: ", id)
   
	admin, err := h.Store.Order().GetOne(id)
	if err != nil {
	 handleResponse(c, "error while getting admin by id", http.StatusInternalServerError, err)
	 return
	}
	handleResponse(c, "", http.StatusOK, admin)
   }

   func (h Handler) DeleteOrder(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Store.Order().DeleteOrder(id)
	if err != nil {
		handleResponse(c, "error while deleting order", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, id)
}
