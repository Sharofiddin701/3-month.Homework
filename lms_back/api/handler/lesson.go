package handler

import (
	"context"
	"fmt"
	_ "lms_back/api/docs"
	"lms_back/api/models"
	"lms_back/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// CreateLesson godoc
// @Router 		   /lesson [POST]
// @Summary 	   create a lesson
// @Description    This api is creates a new lesson and returns its id
// @Tags 		   lesson
// @Accept		   json
// @Produce		   json
// @Param		   lesson body    models.CreateLesson true "lesson"
// @Success		   200  {object}  models.Lesson
// @Failure		   400  {object}  models.Response
// @Failure		   404  {object}  models.Response
// @Failure		   500  {object}  models.Response
func (h Handler) CreateLesson(c *gin.Context) {
	lesson := models.Lesson{}

	if err := c.ShouldBindJSON(&lesson); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.CtxTime)
	defer cancel()

	id, err := h.Service.Lesson().Create(ctx, lesson)
	if err != nil {
		handleResponse(c, "error while creating lesson", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// UpdateLesson godoc
// @Router                /lesson/{id} [PUT]
// @Summary 			  update a lesson
// @Description:          this api updates lesson information
// @Tags 			      lesson
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Lesson ID"
// @Param       		  car body models.UpdateLesson true "Lesson"
// @Success 		      200 {object} models.Lesson
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateLesson(c *gin.Context) {

	lesson := models.Lesson{}
	if err := c.ShouldBindJSON(&lesson); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	lesson.Id = c.Param("id")
	err := uuid.Validate(lesson.Id)

	if err != nil {
		handleResponse(c, "error while validating", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.CtxTime)
	defer cancel()

	id, err := h.Service.Lesson().Update(ctx, lesson)
	if err != nil {
		handleResponse(c, "error while updating lesson", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GetAllLessons godoc
// @Router 			/lesson [GET]
// @Summary 		Get All Lessons
// @Description 	This API returns lesson list
// @Tags 			lesson
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllLessonsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllLessons(c *gin.Context) {

	request := models.GetAllLessonsRequest{}

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

	ctx, cancel := context.WithTimeout(c, config.CtxTime)
	defer cancel()

	branches, err := h.Service.Lesson().GetAll(ctx, request)
	if err != nil {
		handleResponse(c, "error while getting branches", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, branches)
}

// @Security ApiKeyAuth
// GetByIDLesson godoc
// @Router       /lesson/{id} [GET]
// @Summary      return a lesson by ID
// @Description  Retrieves a lesson by its ID
// @Tags         lesson
// @Accept       json
// @Produce      json
// @Param        id path string true "Lesson ID"
// @Success      200 {object} models.GetLesson
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDLesson(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	ctx, cancel := context.WithTimeout(c, config.CtxTime)
	defer cancel()

	Branch, err := h.Service.Lesson().GetByID(ctx, id)
	if err != nil {
		handleResponse(c, "error while getting lesson by id", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "", http.StatusOK, Branch)
}

// @Security ApiKeyAuth
// DeleteLessson godoc
// @Router          /lesson/{id} [DELETE]
// @Summary         Deletes a lesson by ID
// @Description     Deletes a lesson by its ID
// @Tags            lesson
// @Accept          json
// @Produce         json
// @Param           id path string true "Lesson ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteLessson(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.CtxTime)
	defer cancel()

	err = h.Service.Lesson().Delete(ctx, id)
	if err != nil {
		handleResponse(c, "error while deleting lesson", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "", http.StatusOK, id)
}
