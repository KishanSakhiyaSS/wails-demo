package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kishansakhiya/wails-demo/backend/app/models"
	"github.com/kishansakhiya/wails-demo/backend/app/services"

	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-io-go/ip2locationio"
)

// SystemController handles HTTP requests for system information
type SystemController struct {
	systemService *services.SystemService
}

// NewSystemController creates a new instance of SystemController
func NewSystemController() *SystemController {
	return &SystemController{
		systemService: services.NewSystemService(),
	}
}

// GetAllSystemInfo handles GET request for all system information
// @Summary Get all system information
// @Description Retrieve comprehensive system information including CPU, GPU, memory, disk, and hardware details
// @Tags system
// @Accept json
// @Produce json
// @Success 200 {object} models.SystemInfo
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/system [get]
func (c *SystemController) GetAllSystemInfo(ctx *gin.Context) {
	// Create context with timeout
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 30*time.Second)
	defer cancel()

	// Add timeout to gin context
	ctx.Request = ctx.Request.WithContext(reqCtx)

	data, err := c.systemService.GetAllSystemInfo()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get system information", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetCPUInfo handles GET request for CPU information
// @Summary Get CPU information
// @Description Retrieve detailed CPU information including cores, frequency, and usage
// @Tags cpu
// @Accept json
// @Produce json
// @Success 200 {object} models.CPUInfo
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/cpu [get]
func (c *SystemController) GetCPUInfo(ctx *gin.Context) {
	data, err := c.systemService.GetCPUInfo()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get CPU information", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetGPUInfo handles GET request for GPU information
// @Summary Get GPU information
// @Description Retrieve GPU information including model, memory, and driver details
// @Tags gpu
// @Accept json
// @Produce json
// @Success 200 {object} models.GPUInfo
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/gpu [get]
func (c *SystemController) GetGPUInfo(ctx *gin.Context) {
	data, err := c.systemService.GetGPUInfo()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get GPU information", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetOSInfo handles GET request for OS information
// @Summary Get OS information
// @Description Retrieve operating system information including name, version, and architecture
// @Tags os
// @Accept json
// @Produce json
// @Success 200 {object} models.OSInfo
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/os [get]
func (c *SystemController) GetOSInfo(ctx *gin.Context) {
	data, err := c.systemService.GetOSInfo()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get OS information", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetLocationInfo handles GET request for location information
// @Summary Get location information
// @Description Retrieve system location information including timezone and locale
// @Tags location
// @Accept json
// @Produce json
// @Success 200 {object} models.LocationInfo
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/location [get]
func (c *SystemController) GetLocationInfo(ctx *gin.Context) {
	data, err := c.systemService.GetLocationInfo()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get location information", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetMemoryInfo handles GET request for memory information
// @Summary Get memory information
// @Description Retrieve memory information including total, used, and available memory
// @Tags memory
// @Accept json
// @Produce json
// @Success 200 {object} models.MemoryInfo
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/memory [get]
func (c *SystemController) GetMemoryInfo(ctx *gin.Context) {
	data, err := c.systemService.GetMemoryInfo()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get memory information", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetDiskInfo handles GET request for disk information
// @Summary Get disk information
// @Description Retrieve disk information including partitions, usage, and I/O statistics
// @Tags disk
// @Accept json
// @Produce json
// @Success 200 {object} models.DiskInfo
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/disk [get]
func (c *SystemController) GetDiskInfo(ctx *gin.Context) {
	data, err := c.systemService.GetDiskInfo()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get disk information", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetHardwareInfo handles GET request for hardware information
// @Summary Get hardware information
// @Description Retrieve hardware information including motherboard, BIOS, and device details
// @Tags hardware
// @Accept json
// @Produce json
// @Success 200 {object} models.HardwareInfo
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/hardware [get]
func (c *SystemController) GetHardwareInfo(ctx *gin.Context) {
	data, err := c.systemService.GetHardwareInfo()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get hardware information", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetUsagePercentages handles GET request for usage percentages
// @Summary Get usage percentages
// @Description Retrieve usage percentages for CPU, GPU, memory, and disk
// @Tags usage
// @Accept json
// @Produce json
// @Success 200 {object} models.UsagePercentages
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/usage [get]
func (c *SystemController) GetUsagePercentages(ctx *gin.Context) {
	data, err := c.systemService.GetUsagePercentages()
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get usage percentages", err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// HealthCheck handles GET request for health check
// @Summary Health check
// @Description Check if the API is running and healthy
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Router /health [get]
func (c *SystemController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "System benchmark API is running",
	})
}

// TestRoute handles GET request for testing personal methods
// @Summary Test route for personal methods
// @Description A flexible endpoint for testing custom methods and logic
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/test [get]
func (c *SystemController) TestRoute(ctx *gin.Context) {

	apikey := "0940FABBC3BBC0E2FAF3CD2BB700E193"

	config, err := ip2locationio.OpenConfiguration(apikey)
	if err != nil {
		fmt.Print(err)
		return
	}

	ipl, err := ip2locationio.OpenIPGeolocation(config)
	if err != nil {
		fmt.Print(err)
		return
	}

	// location, err := c.systemService.GetLocationInfo()
	// if err != nil {
	// 	c.sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to get location information", err)
	// 	return
	// }

	// ip := location.Ip
	ip := "171.25.193.38"
	res, err := ipl.LookUp(ip, "")
	if err != nil {
		fmt.Print(err)
		return
	}

	testData := map[string]any{
		"response": res,
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Status:  "ok",
		Message: "Test route executed successfully",
		Data:    testData,
	})
}

// sendErrorResponse sends a standardized error response
func (c *SystemController) sendErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	errorResponse := models.ErrorResponse{
		Error:   message,
		Details: err.Error(),
	}

	ctx.JSON(statusCode, errorResponse)
}
