package main

import (
	"log/slog"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucas-albers-lz4/ferrovis/internal/database"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, using environment variables")
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer database.Close()

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
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Ferrovis API is running",
		})
	})

	// Database test endpoint
	r.GET("/db-test", func(c *gin.Context) {
		// Test database connection
		sqlDB, err := database.DB.DB()
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to get database instance",
				"error":   err.Error(),
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Database ping failed",
				"error":   err.Error(),
			})
			return
		}

		// Get database stats
		stats := sqlDB.Stats()
		c.JSON(200, gin.H{
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

		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Ferrovis Weasel Mode™ Database Schema",
			"weasel_features": gin.H{
				"schema": gin.H{
					"total_tables":  10,
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
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", registerHandler)
			auth.POST("/login", loginHandler)
		}

		// Protected routes (will add JWT middleware later)
		protected := api.Group("/")
		{
			// User routes
			protected.GET("/user/profile", getUserProfile)
			protected.PUT("/user/profile", updateUserProfile)

			// Workout routes
			protected.POST("/workouts", createWorkout)
			protected.GET("/workouts", getWorkouts)
			protected.GET("/workouts/:id", getWorkout)

			// Buddy routes
			protected.POST("/buddies/invite", inviteBuddy)
			protected.GET("/buddies", getBuddies)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("Starting Ferrovis API server", "port", port)
	if err := r.Run(":" + port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

// Placeholder handlers - will implement in separate files
func registerHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Register endpoint - TODO"})
}

func loginHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Login endpoint - TODO"})
}

func getUserProfile(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get user profile - TODO"})
}

func updateUserProfile(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update user profile - TODO"})
}

func createWorkout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Create workout - TODO"})
}

func getWorkouts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get workouts - TODO"})
}

func getWorkout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get single workout - TODO"})
}

func inviteBuddy(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Invite buddy - TODO"})
}

func getBuddies(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get buddies - TODO"})
}
