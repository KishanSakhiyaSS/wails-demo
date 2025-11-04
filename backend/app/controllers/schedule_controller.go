package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kishansakhiya/wails-demo/backend/app/database"
	"github.com/kishansakhiya/wails-demo/backend/app/models"
	"github.com/kishansakhiya/wails-demo/backend/app/scheduler"

	"github.com/gin-gonic/gin"
)

// ScheduleController handles HTTP requests for schedule operations
type ScheduleController struct {
	schedulerService *scheduler.SchedulerService
}

// NewScheduleController creates a new instance of ScheduleController
func NewScheduleController(schedulerService *scheduler.SchedulerService) *ScheduleController {
	return &ScheduleController{
		schedulerService: schedulerService,
	}
}

// AddSchedule handles POST request to add a new schedule
// @Summary Add a new schedule
// @Description Create a new schedule with the provided details
// @Tags schedule
// @Accept json
// @Produce json
// @Param schedule body database.Schedule true "Schedule details"
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/schedules [post]
func (c *ScheduleController) AddSchedule(ctx *gin.Context) {
	var schedule database.Schedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := c.schedulerService.AddSchedule(&schedule); err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to add schedule", err)
		return
	}

	ctx.JSON(http.StatusCreated, models.APIResponse{
		Status:  "ok",
		Message: "Schedule added successfully",
		Data:    schedule,
	})
}

// ListSchedules handles GET request to retrieve all schedules
// @Summary List all schedules
// @Description Retrieve a list of all schedules
// @Tags schedule
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/schedules [get]
func (c *ScheduleController) ListSchedules(ctx *gin.Context) {
	schedules, err := c.schedulerService.ListSchedules()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to list schedules", err)
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status: "ok",
		Data:   schedules,
	})
}

// UpdateSchedule handles PUT request to update an existing schedule
// @Summary Update a schedule
// @Description Update an existing schedule by ID
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Param schedule body database.Schedule true "Updated schedule details"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/schedules/{id} [put]
func (c *ScheduleController) UpdateSchedule(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid schedule ID", err)
		return
	}

	var schedule database.Schedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	schedule.ID = id
	if err := c.schedulerService.UpdateSchedule(&schedule); err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update schedule", err)
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "Schedule updated successfully",
		Data:    schedule,
	})
}

// DeleteSchedule handles DELETE request to remove a schedule
// @Summary Delete a schedule
// @Description Delete a schedule by ID
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/schedules/{id} [delete]
func (c *ScheduleController) DeleteSchedule(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid schedule ID", err)
		return
	}

	if err := c.schedulerService.DeleteSchedule(id); err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete schedule", err)
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "Schedule deleted successfully",
	})
}

// ToggleSchedule handles PATCH request to enable/disable a schedule
// @Summary Toggle a schedule
// @Description Enable or disable a schedule by ID
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Param enabled body object{enabled=bool} true "Enable/disable flag"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/schedules/{id}/toggle [patch]
func (c *ScheduleController) ToggleSchedule(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid schedule ID", err)
		return
	}

	var request struct {
		Enabled bool `json:"enabled"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := c.schedulerService.ToggleSchedule(id, request.Enabled); err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to toggle schedule", err)
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "Schedule toggled successfully",
	})
}

// SyncWithSystem handles POST request to sync with system scheduler
// @Summary Sync with system scheduler
// @Description Manually trigger synchronization with the system scheduler
// @Tags schedule
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/schedules/sync [post]
func (c *ScheduleController) SyncWithSystem(ctx *gin.Context) {
	if err := c.schedulerService.SyncWithSystem(); err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to sync with system", err)
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "Synchronization completed successfully",
	})
}

// GetSchedule handles GET request to retrieve a single schedule
// @Summary Get a schedule
// @Description Retrieve a specific schedule by ID
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/schedules/{id} [get]
func (c *ScheduleController) GetSchedule(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "Invalid schedule ID", err)
		return
	}

	// Get schedule from database directly
	schedules, err := c.schedulerService.ListSchedules()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get schedules", err)
		return
	}

	// Find the schedule with matching ID
	var foundSchedule *database.Schedule
	for _, s := range schedules {
		if s.ID == id {
			foundSchedule = s
			break
		}
	}

	if foundSchedule == nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "Schedule not found",
			Details: fmt.Sprintf("No schedule found with ID %d", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status: "ok",
		Data:   foundSchedule,
	})
}

// sendErrorResponse sends a standardized error response
func (c *ScheduleController) sendErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	errorDetails := ""
	if err != nil {
		if err == sql.ErrNoRows {
			errorDetails = "Resource not found"
		} else {
			errorDetails = err.Error()
		}
	}

	errorResponse := models.ErrorResponse{
		Error:   message,
		Details: errorDetails,
	}

	ctx.JSON(statusCode, errorResponse)
}

