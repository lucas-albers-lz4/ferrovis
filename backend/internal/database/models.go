package database

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system with weasel mode preferences
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Basic user info
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"-"` // Never include in JSON responses

	// Weasel mode configuration
	WeaselModeEnabled   bool   `gorm:"default:true" json:"weasel_mode_enabled"`
	WeaselIntensity     string `gorm:"default:medium" json:"weasel_intensity"` // gentle, medium, aggressive, full_chaos
	AllowGuiltTrips     bool   `gorm:"default:true" json:"allow_guilt_trips"`
	AllowFakeStats      bool   `gorm:"default:true" json:"allow_fake_stats"`
	AllowSocialPressure bool   `gorm:"default:true" json:"allow_social_pressure"`

	// User preferences
	PreferredWorkoutTime string `json:"preferred_workout_time"` // morning, afternoon, evening
	FitnessGoal          string `json:"fitness_goal"`           // strength, endurance, weight_loss, general

	// Relationships
	Workouts        []Workout           `gorm:"foreignKey:UserID" json:"workouts,omitempty"`
	Achievements    []UserAchievement   `gorm:"foreignKey:UserID" json:"achievements,omitempty"`
	WeaselMessages  []WeaselMessage     `gorm:"foreignKey:UserID" json:"weasel_messages,omitempty"`
	BuddyRequests   []BuddyRelationship `gorm:"foreignKey:UserID" json:"buddy_requests,omitempty"`
	CoachingClients []BuddyRelationship `gorm:"foreignKey:BuddyID" json:"coaching_clients,omitempty"`
}

// Workout represents a completed workout session
type Workout struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`

	// Workout details
	ProgramID   uint      `json:"program_id"`
	Program     Program   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"program,omitempty"`
	CompletedAt time.Time `json:"completed_at"`
	Duration    int       `json:"duration"` // Duration in minutes

	// Exercise data (JSON field for flexibility)
	Exercises string `gorm:"type:jsonb" json:"exercises"` // JSON array of exercises with sets/reps/weight

	// Weasel mode enhancements
	FakeProgressBoost int  `gorm:"default:0" json:"fake_progress_boost"` // Artificial stat inflation percentage
	IsPersonalRecord  bool `gorm:"default:false" json:"is_personal_record"`

	// Session notes
	Notes string `json:"notes"`
}

// Program represents a workout program (Starting Strength, 5x5, etc.)
type Program struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"` // beginner, intermediate, advanced
	Duration    int    `json:"duration"`   // Program duration in weeks

	// Program structure (JSON field for flexibility)
	Structure string `gorm:"type:jsonb" json:"structure"` // JSON defining exercises, progression, etc.

	// Relationships
	Workouts []Workout `gorm:"foreignKey:ProgramID" json:"workouts,omitempty"`
}

// Exercise represents an individual exercise
type Exercise struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name         string `gorm:"uniqueIndex;not null" json:"name"`
	Category     string `json:"category"`      // compound, isolation, cardio
	MuscleGroups string `json:"muscle_groups"` // JSON array of muscle groups
	Instructions string `gorm:"type:text" json:"instructions"`

	// Weasel mode metadata for fake progress calculations
	ProgressMultiplier float64 `gorm:"default:1.0" json:"progress_multiplier"` // For inflating progress stats
}

// Achievement represents gamification achievements
type Achievement struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"` // consistency, strength, social, funny
	Icon        string `json:"icon"`
	Target      int    `json:"target"` // Target value to unlock achievement

	// Weasel mode features
	IsFakeAchievement bool   `gorm:"default:false" json:"is_fake_achievement"` // Some achievements are creative fiction
	RarityPercent     int    `gorm:"default:50" json:"rarity_percent"`         // "Only X% of users earn this!"
	WeaselMessage     string `json:"weasel_message"`                           // Custom message when unlocked

	// Relationships
	UserAchievements []UserAchievement `gorm:"foreignKey:AchievementID" json:"user_achievements,omitempty"`
}

// UserAchievement represents the join table for users and achievements
type UserAchievement struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID        uint        `gorm:"not null" json:"user_id"`
	User          User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	AchievementID uint        `gorm:"not null" json:"achievement_id"`
	Achievement   Achievement `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"achievement,omitempty"`

	UnlockedAt time.Time `json:"unlocked_at"`
	Progress   int       `json:"progress"` // Current progress toward achievement
}

// BuddyRelationship represents the buddy/coach system
type BuddyRelationship struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID  uint `gorm:"not null" json:"user_id"`
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	BuddyID uint `gorm:"not null" json:"buddy_id"`
	Buddy   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"buddy,omitempty"`

	RelationshipType string `gorm:"not null" json:"relationship_type"` // peer, coach
	Status           string `gorm:"default:pending" json:"status"`     // pending, active, paused

	// Invitation details
	InvitedAt  time.Time  `json:"invited_at"`
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
}

// WeaselMessage represents psychological manipulation messages sent to users
type WeaselMessage struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`

	// Message details
	MessageType string `gorm:"not null" json:"message_type"` // guilt, fomo, urgency, social, funny
	Content     string `gorm:"type:text;not null" json:"content"`
	Intensity   string `gorm:"not null" json:"intensity"` // gentle, medium, aggressive

	// Delivery and tracking
	SentAt       time.Time  `json:"sent_at"`
	ReadAt       *time.Time `json:"read_at,omitempty"`
	UserReaction string     `json:"user_reaction"` // ignored, annoyed, motivated, worked_out

	// Effectiveness tracking
	TriggeredWorkout bool `gorm:"default:false" json:"triggered_workout"` // Did user work out within 24h?
}

// Streak represents user workout streaks
type Streak struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`

	StreakType  string    `gorm:"not null" json:"streak_type"` // workout, weekly, monthly
	Current     int       `gorm:"default:0" json:"current"`
	Longest     int       `gorm:"default:0" json:"longest"`
	LastWorkout time.Time `json:"last_workout"`
	StreakStart time.Time `json:"streak_start"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
}

// FakeSocialActivity represents algorithmic fake social activity for pressure
type FakeSocialActivity struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	ActivityType     string    `gorm:"not null" json:"activity_type"` // workout_completed, pr_achieved, streak_extended
	FakeUserName     string    `gorm:"not null" json:"fake_user_name"`
	Details          string    `json:"details"`
	Timestamp        time.Time `json:"timestamp"`
	TargetUserGroups string    `json:"target_user_groups"` // JSON array of user groups to show this to
}
