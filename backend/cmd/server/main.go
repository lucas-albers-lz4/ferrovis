// Package main provides the entry point for the Ferrovis API server.
// This server handles fitness tracking, workout management, and the infamous Weasel Mode™
// psychological manipulation features for accountability.
package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucas-albers-lz4/ferrovis/internal/database"
)

// Version information (injected at build time)
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

const (
	// HTTP status code constants to avoid magic numbers
	statusOK                  = http.StatusOK
	statusInternalServerError = http.StatusInternalServerError

	// Business logic constants
	defaultTableCount = 10
	weekDuration      = 12 // 12 weeks for workout programs
	defaultPort       = "8080"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, using environment variables")
	}

	// Log version information
	slog.Info("Starting Ferrovis API server",
		"version", version,
		"commit", commit,
		"build_date", date)

	// Connect to database
	if err := database.Connect(); err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := database.Close(); err != nil {
			slog.Error("Failed to close database connection", "error", err)
		}
	}()

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Configure appropriately for production
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(statusOK, gin.H{
			"status":     "ok",
			"message":    "Ferrovis API is running",
			"version":    version,
			"commit":     commit,
			"build_date": date,
		})
	})

	// Version endpoint
	r.GET("/version", func(c *gin.Context) {
		c.JSON(statusOK, gin.H{
			"version":    version,
			"commit":     commit,
			"build_date": date,
			"service":    "ferrovis-api",
		})
	})

	// Database test endpoint
	r.GET("/db-test", func(c *gin.Context) {
		// Test database connection
		sqlDB, err := database.DB.DB()
		if err != nil {
			c.JSON(statusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to get database instance",
				"error":   err.Error(),
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(statusInternalServerError, gin.H{
				"status":  "error",
				"message": "Database ping failed",
				"error":   err.Error(),
			})
			return
		}

		// Get database stats
		stats := sqlDB.Stats()
		c.JSON(statusOK, gin.H{
			"status":  "ok",
			"message": "Database connection is healthy",
			"database": gin.H{
				"open_connections": stats.OpenConnections,
				"in_use":           stats.InUse,
				"idle":             stats.Idle,
			},
		})
	})

	// Weasel mode schema showcase endpoint
	r.GET("/schema-test", func(c *gin.Context) {
		var programCount, exerciseCount, achievementCount, fakeActivityCount int64

		database.DB.Model(&database.Program{}).Count(&programCount)
		database.DB.Model(&database.Exercise{}).Count(&exerciseCount)
		database.DB.Model(&database.Achievement{}).Count(&achievementCount)
		database.DB.Model(&database.FakeSocialActivity{}).Count(&fakeActivityCount)

		// Get a sample achievement with weasel features
		var sampleAchievement database.Achievement
		database.DB.Where("is_fake_achievement = ?", true).First(&sampleAchievement)

		// Get sample fake social activity
		var sampleFakeActivity database.FakeSocialActivity
		database.DB.First(&sampleFakeActivity)

		c.JSON(statusOK, gin.H{
			"status":  "ok",
			"message": "Ferrovis Weasel Mode™ Database Schema",
			"weasel_features": gin.H{
				"schema": gin.H{
					"total_tables":  defaultTableCount,
					"core_entities": []string{"users", "workouts", "achievements", "weasel_messages", "buddy_relationships"},
					"weasel_tables": []string{"fake_social_activities", "streaks", "user_achievements"},
				},
				"seed_data": gin.H{
					"workout_programs":       programCount,
					"exercises":              exerciseCount,
					"achievements":           achievementCount,
					"fake_social_activities": fakeActivityCount,
				},
				"weasel_mode_examples": gin.H{
					"fake_achievement": gin.H{
						"name":           sampleAchievement.Name,
						"description":    sampleAchievement.Description,
						"weasel_message": sampleAchievement.WeaselMessage,
						"rarity_percent": sampleAchievement.RarityPercent,
					},
					"fake_social_activity": gin.H{
						"fake_user": sampleFakeActivity.FakeUserName,
						"activity":  sampleFakeActivity.Details,
						"type":      sampleFakeActivity.ActivityType,
					},
				},
				"psychological_features": []string{
					"Progress inflation algorithms",
					"Variable reward messaging",
					"Fake social pressure generation",
					"Achievement rarity manipulation",
					"Streak anxiety timers",
					"Guilt trip message classification",
				},
			},
		})
	})

	// API routes group
	api := r.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.POST("/register", registerHandler)
	auth.POST("/login", loginHandler)

	// Protected routes (will add JWT middleware later)
	// User routes
	api.GET("/user/profile", getUserProfile)
	api.PUT("/user/profile", updateUserProfile)

	// Workout routes
	api.POST("/workouts", createWorkout)
	api.GET("/workouts", getWorkouts)
	api.GET("/workouts/:id", getWorkout)

	// Program routes - NEW
	api.GET("/programs", getPrograms)
	api.GET("/programs/:id", getProgram)
	api.GET("/programs/:id/next-workout", getNextWorkout)

	// Buddy routes
	api.POST("/buddies/invite", inviteBuddy)
	api.GET("/buddies", getBuddies)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	slog.Info("Starting Ferrovis API server", "port", port)
	if err := r.Run(":" + port); err != nil {
		slog.Error("Failed to start server", "error", err)
		// Don't use os.Exit here since defer won't run
		return
	}
}

// Placeholder handlers - will implement in separate files
func registerHandler(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Register endpoint - TODO"})
}

func loginHandler(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Login endpoint - TODO"})
}

func getUserProfile(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Get user profile - TODO"})
}

func updateUserProfile(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Update user profile - TODO"})
}

func createWorkout(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Create workout - TODO"})
}

func getWorkouts(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Get workouts - TODO"})
}

func getWorkout(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Get single workout - TODO"})
}

func inviteBuddy(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Invite buddy - TODO"})
}

func getBuddies(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": "Get buddies - TODO"})
}

// Program handlers
func getPrograms(c *gin.Context) {
	var programs []database.Program
	if err := database.DB.Find(&programs).Error; err != nil {
		c.JSON(statusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch programs",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(statusOK, gin.H{
		"status":   "ok",
		"programs": programs,
	})
}

func getProgram(c *gin.Context) {
	id := c.Param("id")
	var program database.Program

	if err := database.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Program not found",
		})
		return
	}

	c.JSON(statusOK, gin.H{
		"status":  "ok",
		"program": program,
	})
}

func getNextWorkout(c *gin.Context) {
	programID := c.Param("id")

	// For now, return a basic next workout based on Starting Strength
	// This is a simplified version - real implementation would track user progress
	var program database.Program
	if err := database.DB.First(&program, programID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Program not found",
		})
		return
	}

	// Basic workout structure for Starting Strength/5x5 programs
	nextWorkout := gin.H{
		"program_name": program.Name,
		"workout_day":  "A", // Alternate between A/B
		"exercises": []gin.H{
			{
				"name":         "Squat",
				"sets":         3,
				"reps":         5,
				"weight":       "start_weight + progression", // Will be calculated based on user
				"instructions": "Stand with feet shoulder-width apart, lower body by bending knees and hips",
			},
			{
				"name":         "Bench Press",
				"sets":         3,
				"reps":         5,
				"weight":       "start_weight + progression",
				"instructions": "Lie on bench, lower barbell to chest, then press back up",
			},
			{
				"name":         "Barbell Row",
				"sets":         3,
				"reps":         5,
				"weight":       "start_weight + progression",
				"instructions": "Bend at hips, pull barbell from arm's length to lower chest",
			},
		},
		"estimated_duration": "45-60 minutes",
		"rest_between_sets":  "3-5 minutes for compound movements",
	}

	c.JSON(statusOK, gin.H{
		"status":       "ok",
		"next_workout": nextWorkout,
	})
}
