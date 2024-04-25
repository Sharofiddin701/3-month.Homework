package handler

import (
	"context"
	_ "lms_back/api/docs"
	"lms_back/api/models"
	"lms_back/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReportAdminByID godoc
// @Router /report_admin/{id} [GET]
// @Summary get a report for admin by ID
// @Description This API retrieves a report for admin by its ID
// @Tags report
// @Accept json
// @Produce json
// @Param id path string true "Report ID"
// @Success 200 {object} models.ReportAdmin
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
func (h Handler) GetReportAdminByID(c *gin.Context) {
	admin := models.AdminPrimaryKey{}
	admin.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(c, config.CtxTime)
	defer cancel()

	report, err := h.Service.Report().GetReportAdmin(ctx, admin)
	if err != nil {
		handleResponse(c, "error while getting report by id", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "", http.StatusOK, report)
}
