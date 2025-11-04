package routes

import (
	"github.com/kishansakhiya/wails-demo/backend/app/controllers"
	"github.com/kishansakhiya/wails-demo/backend/app/database"
	"github.com/kishansakhiya/wails-demo/backend/app/middleware"
	"github.com/kishansakhiya/wails-demo/backend/app/scheduler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine) {
	// Create controller instance
	systemController := controllers.NewSystemController()

	// Initialize database and scheduler service for schedule endpoints
	db, err := database.NewDB()
	if err != nil {
		// Log error but continue without schedule endpoints
		// In production, you might want to handle this differently
	} else {
		schedulerService := scheduler.NewSchedulerService(db)
		scheduleController := controllers.NewScheduleController(schedulerService)

		// Schedule endpoints
		r.POST("/api/v1/schedules", scheduleController.AddSchedule)
		r.GET("/api/v1/schedules", scheduleController.ListSchedules)
		r.GET("/api/v1/schedules/:id", scheduleController.GetSchedule)
		r.PUT("/api/v1/schedules/:id", scheduleController.UpdateSchedule)
		r.DELETE("/api/v1/schedules/:id", scheduleController.DeleteSchedule)
		r.PATCH("/api/v1/schedules/:id/toggle", scheduleController.ToggleSchedule)
		r.POST("/api/v1/schedules/sync", scheduleController.SyncWithSystem)
	}

	// Add global middleware
	r.Use(middleware.CORS())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Recovery())

	// Health check endpoint
	r.GET("/health", systemController.HealthCheck)

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// Add rate limiting for API endpoints
		v1.Use(middleware.RateLimit(100, time.Minute)) // 100 requests per minute

		// Get all system information
		v1.GET("/system", systemController.GetAllSystemInfo)

		// Individual module endpoints
		v1.GET("/cpu", systemController.GetCPUInfo)
		v1.GET("/gpu", systemController.GetGPUInfo)
		v1.GET("/os", systemController.GetOSInfo)
		v1.GET("/location", systemController.GetLocationInfo)
		v1.GET("/memory", systemController.GetMemoryInfo)
		v1.GET("/disk", systemController.GetDiskInfo)
		v1.GET("/hardware", systemController.GetHardwareInfo)
		v1.GET("/usage", systemController.GetUsagePercentages)
		v1.GET("/test", systemController.TestRoute)
	}

	// Add 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Endpoint not found",
			"message": "The requested endpoint does not exist",
		})
	})
}
