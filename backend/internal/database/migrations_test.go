package database

import (
	"testing"
)

func TestConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant interface{}
		expected interface{}
	}{
		{"standardProgramDuration", standardProgramDuration, 12},
		{"squatMultiplier", squatMultiplier, 1.2},
		{"deadliftMultiplier", deadliftMultiplier, 1.3},
		{"benchMultiplier", benchMultiplier, 1.1},
		{"ohpMultiplier", ohpMultiplier, 1.15},
		{"rowMultiplier", rowMultiplier, 1.1},
		{"weeklyStreakTarget", weeklyStreakTarget, 7},
		{"monthlyStreakTarget", monthlyStreakTarget, 30},
		{"firstPRTarget", firstPRTarget, 1},
		{"centuryClubTarget", centuryClubTarget, 100},
		{"teamPlayerTarget", teamPlayerTarget, 5},
		{"sweatSommelierTarget", sweatSommelierTarget, 3},
		{"gymWhispererTarget", gymWhispererTarget, 100},
		{"showUpSamuraiRarity", showUpSamuraiRarity, 35},
		{"ironWillRarity", ironWillRarity, 8},
		{"weightWarriorRarity", weightWarriorRarity, 60},
		{"centuryClubRarity", centuryClubRarity, 25},
		{"teamPlayerRarity", teamPlayerRarity, 40},
		{"sweatSommelierRarity", sweatSommelierRarity, 45},
		{"gymWhispererRarity", gymWhispererRarity, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("Expected %s to be %v, got %v", tt.name, tt.expected, tt.constant)
			}
		})
	}
}

func TestProgressMultipliers(t *testing.T) {
	// Test that progress multipliers are reasonable values
	multipliers := map[string]float64{
		"squat":    squatMultiplier,
		"deadlift": deadliftMultiplier,
		"bench":    benchMultiplier,
		"ohp":      ohpMultiplier,
		"row":      rowMultiplier,
	}

	for exercise, multiplier := range multipliers {
		t.Run(exercise, func(t *testing.T) {
			if multiplier < 1.0 || multiplier > 2.0 {
				t.Errorf("Progress multiplier for %s should be between 1.0 and 2.0, got %f", exercise, multiplier)
			}
		})
	}
}

func TestAchievementTargets(t *testing.T) {
	// Test that achievement targets are positive values
	targets := map[string]int{
		"weeklyStreak":   weeklyStreakTarget,
		"monthlyStreak":  monthlyStreakTarget,
		"firstPR":        firstPRTarget,
		"centuryClub":    centuryClubTarget,
		"teamPlayer":     teamPlayerTarget,
		"sweatSommelier": sweatSommelierTarget,
		"gymWhisperer":   gymWhispererTarget,
	}

	for achievement, target := range targets {
		t.Run(achievement, func(t *testing.T) {
			if target <= 0 {
				t.Errorf("Achievement target for %s should be positive, got %d", achievement, target)
			}
		})
	}
}

func TestAchievementRarities(t *testing.T) {
	// Test that achievement rarities are valid percentages
	rarities := map[string]int{
		"showUpSamurai":  showUpSamuraiRarity,
		"ironWill":       ironWillRarity,
		"weightWarrior":  weightWarriorRarity,
		"centuryClub":    centuryClubRarity,
		"teamPlayer":     teamPlayerRarity,
		"sweatSommelier": sweatSommelierRarity,
		"gymWhisperer":   gymWhispererRarity,
	}

	for achievement, rarity := range rarities {
		t.Run(achievement, func(t *testing.T) {
			if rarity < 0 || rarity > 100 {
				t.Errorf("Achievement rarity for %s should be between 0-100, got %d", achievement, rarity)
			}
		})
	}
}
