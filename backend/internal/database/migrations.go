package database

import (
	"log/slog"
	"time"
)

// RunMigrations executes all database migrations
func RunMigrations() error {
	slog.Info("Running database migrations...")

	// Auto-migrate all models
	err := DB.AutoMigrate(
		&User{},
		&Workout{},
		&Program{},
		&Exercise{},
		&Achievement{},
		&UserAchievement{},
		&BuddyRelationship{},
		&WeaselMessage{},
		&Streak{},
		&FakeSocialActivity{},
	)
	if err != nil {
		return err
	}

	slog.Info("Database migrations completed successfully")

	// Seed initial data
	if err := SeedInitialData(); err != nil {
		slog.Warn("Failed to seed initial data", "error", err)
		return err
	}

	return nil
}

// SeedInitialData populates the database with initial programs, exercises, and achievements
func SeedInitialData() error {
	slog.Info("Seeding initial data...")

	// Check if data already exists
	var programCount int64
	DB.Model(&Program{}).Count(&programCount)
	if programCount > 0 {
		slog.Info("Initial data already exists, skipping seed")
		return nil
	}

	// Seed workout programs
	programs := []Program{
		{
			Name:        "Starting Strength",
			Description: "A beginner-friendly program focusing on compound movements",
			Difficulty:  "beginner",
			Duration:    12, // 12 weeks
			Structure: `{
				"schedule": "3x per week",
				"exercises": ["squat", "deadlift", "bench_press", "overhead_press", "barbell_row"],
				"progression": "linear"
			}`,
		},
		{
			Name:        "StrongLifts 5x5",
			Description: "Simple 5x5 compound movement program for strength building",
			Difficulty:  "beginner",
			Duration:    12,
			Structure: `{
				"schedule": "3x per week",
				"exercises": ["squat", "deadlift", "bench_press", "overhead_press", "barbell_row"],
				"sets": 5,
				"reps": 5,
				"progression": "linear"
			}`,
		},
	}

	if err := DB.Create(&programs).Error; err != nil {
		return err
	}

	// Seed exercises
	exercises := []Exercise{
		{
			Name:               "Squat",
			Category:           "compound",
			MuscleGroups:       `["quadriceps", "glutes", "hamstrings", "core"]`,
			Instructions:       "Stand with feet shoulder-width apart, lower body by bending knees and hips, then return to standing position.",
			ProgressMultiplier: 1.2, // Slightly inflated for weasel mode
		},
		{
			Name:               "Deadlift",
			Category:           "compound",
			MuscleGroups:       `["hamstrings", "glutes", "back", "traps", "core"]`,
			Instructions:       "Lift barbell from floor to hip level by extending hips and knees, then lower back down.",
			ProgressMultiplier: 1.3, // Most impressive exercise for weasel stats
		},
		{
			Name:               "Bench Press",
			Category:           "compound",
			MuscleGroups:       `["chest", "shoulders", "triceps"]`,
			Instructions:       "Lie on bench, lower barbell to chest, then press back up to arms length.",
			ProgressMultiplier: 1.1,
		},
		{
			Name:               "Overhead Press",
			Category:           "compound",
			MuscleGroups:       `["shoulders", "triceps", "core"]`,
			Instructions:       "Press barbell from shoulder level to overhead, then lower back down.",
			ProgressMultiplier: 1.15,
		},
		{
			Name:               "Barbell Row",
			Category:           "compound",
			MuscleGroups:       `["back", "biceps", "rear_delts"]`,
			Instructions:       "Bend at hips, pull barbell from arm's length to lower chest, then lower back down.",
			ProgressMultiplier: 1.1,
		},
	}

	if err := DB.Create(&exercises).Error; err != nil {
		return err
	}

	// Seed achievements
	achievements := []Achievement{
		// Consistency achievements
		{
			Name:              "Show Up Samurai",
			Description:       "Complete 7 workouts in a row",
			Category:          "consistency",
			Icon:              "🥋",
			Target:            7,
			IsFakeAchievement: false,
			RarityPercent:     35,
			WeaselMessage:     "You're becoming unstoppable! Your dedication is inspiring!",
		},
		{
			Name:              "Iron Will",
			Description:       "Maintain a 30-day workout streak",
			Category:          "consistency",
			Icon:              "⚡",
			Target:            30,
			IsFakeAchievement: false,
			RarityPercent:     8,
			WeaselMessage:     "You're in the elite 8%! Your willpower is legendary!",
		},
		// Strength achievements
		{
			Name:              "Weight Warrior",
			Description:       "Achieve your first personal record",
			Category:          "strength",
			Icon:              "💪",
			Target:            1,
			IsFakeAchievement: false,
			RarityPercent:     60,
			WeaselMessage:     "Progress detected! Your muscles are literally growing as we speak!",
		},
		{
			Name:              "Century Club",
			Description:       "Deadlift 100lbs or more",
			Category:          "strength",
			Icon:              "🏋️",
			Target:            100,
			IsFakeAchievement: false,
			RarityPercent:     25,
			WeaselMessage:     "Welcome to the big leagues! You're stronger than 75% of humans!",
		},
		// Social achievements
		{
			Name:              "Team Player",
			Description:       "Help your buddy complete 5 workouts",
			Category:          "social",
			Icon:              "🤝",
			Target:            5,
			IsFakeAchievement: false,
			RarityPercent:     40,
			WeaselMessage:     "You're not just getting swole, you're helping others get swole too!",
		},
		// Funny/Creative achievements
		{
			Name:              "Sweat Sommelier",
			Description:       "Work out at 3 different times of day",
			Category:          "funny",
			Icon:              "🍷",
			Target:            3,
			IsFakeAchievement: false,
			RarityPercent:     45,
			WeaselMessage:     "You've mastered the art of perspiration timing!",
		},
		{
			Name:              "Gym Whisperer",
			Description:       "Complete 100 total workouts",
			Category:          "consistency",
			Icon:              "🗣️",
			Target:            100,
			IsFakeAchievement: true, // This one's a bit creative
			RarityPercent:     5,
			WeaselMessage:     "Legend has it that barbells now listen to your commands...",
		},
	}

	if err := DB.Create(&achievements).Error; err != nil {
		return err
	}

	// Seed some fake social activity
	fakeSocialActivities := []FakeSocialActivity{
		{
			ActivityType:     "workout_completed",
			FakeUserName:     "Mike_Fitness",
			Details:          "crushed their bench press PR!",
			Timestamp:        time.Now().Add(-2 * time.Hour),
			TargetUserGroups: `["beginners", "strength_focused"]`,
		},
		{
			ActivityType:     "streak_extended",
			FakeUserName:     "Sarah_Strong",
			Details:          "extended their streak to 12 days!",
			Timestamp:        time.Now().Add(-4 * time.Hour),
			TargetUserGroups: `["consistency_focused"]`,
		},
		{
			ActivityType:     "pr_achieved",
			FakeUserName:     "FitnessFanatic22",
			Details:          "just deadlifted 225lbs for the first time!",
			Timestamp:        time.Now().Add(-6 * time.Hour),
			TargetUserGroups: `["beginners", "strength_focused"]`,
		},
	}

	if err := DB.Create(&fakeSocialActivities).Error; err != nil {
		return err
	}

	slog.Info("Initial data seeded successfully")
	return nil
}
