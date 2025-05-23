# 🏋️ LiftBuddy MVP Feature Plan

**Core functionality + "Full Weasel Mode" psychological manipulation techniques**

## 🎯 **MVP Core Features (Must-Have)**

### **1. User Authentication & Onboarding**
```typescript
// Basic user flow
- Email/password registration
- Simple profile setup (name, goals)
- "Buddy vs Coach" selection
- First workout program selection
```

**Weasel Mode Elements:**
- ✅ **Fake urgency**: "423 people started their fitness journey today!"
- ✅ **Social proof**: "Users who set up their profile are 73% more likely to stick with it"
- ✅ **Commitment escalation**: "Just one quick question..." (then 5 questions)

### **2. Workout Logging System**
```typescript
// Core functionality
- Pre-built programs (Starting Strength, 5x5)
- Set/rep/weight tracking
- Rest timer with notifications
- Workout completion marking
```

**Weasel Mode Elements:**
- ✅ **Variable rewards**: Random motivational messages ("Beast mode activated!" vs "Meh, you showed up")
- ✅ **Progress illusion**: "You're 12% stronger than last week!" (even if marginal)
- ✅ **Sunk cost fallacy**: "You've already invested 47 minutes this week..."

### **3. Buddy/Coach System**
```typescript
// Social accountability  
- Invite buddy via email/phone
- Share workout completion status
- Weekly progress summaries
- Basic messaging/comments
```

**Weasel Mode Elements:**
- ✅ **Peer pressure**: "Your buddy completed their workout. What's your excuse?"
- ✅ **Fake competition**: "You're behind Alex by 2 workouts this week"
- ✅ **Disappointment guilt**: "Alex is probably wondering where you are..."

## 🎮 **Gamification System (The Hook)**

### **4. Streak & Achievement System**
```typescript
interface Streak {
  current: number;
  longest: number;
  lastWorkout: Date;
  streakType: 'workout' | 'weekly' | 'monthly';
}

interface Achievement {
  id: string;
  title: string;
  description: string;
  unlockedAt?: Date;
  progress: number;
  target: number;
}
```

**Achievement Categories:**
- **Consistency**: "Show Up Samurai" (7-day streak)
- **Strength**: "Weight Warrior" (first PR)
- **Social**: "Team Player" (helped buddy 5 times)
- **Funny**: "Sweat Sommelier" (worked out 3 different times of day)

**Weasel Mode Elements:**
- ✅ **Near-miss psychology**: "You're 1 workout away from 'Iron Will' badge!"
- ✅ **Loss aversion**: "Don't lose your 6-day streak!"
- ✅ **Artificial scarcity**: "Only 12% of users earn this badge"

### **5. Progress Visualization**
```typescript
// Visual progress tricks
- Strength progression charts (amplified scales)
- "Fitness score" calculation
- Before/after comparison tools
- Weekly "gains" summaries
```

**Weasel Mode Elements:**
- ✅ **Cherry-picked stats**: Always find something positive to highlight
- ✅ **Misleading scales**: Charts that make small progress look huge
- ✅ **Projection magic**: "At this rate, you'll bench 200lbs by March!"

## 🐾 **Full Weasel Mode Features (The Fun Part)**

### **6. Motivational Nagging System**
```typescript
interface WeaselMessage {
  type: 'guilt' | 'fomo' | 'pride' | 'urgency' | 'social' | 'funny';
  message: string;
  triggerCondition: string;
  frequency: 'low' | 'medium' | 'aggressive';
}
```

**Message Categories:**

**Guilt Trips:**
- "Your dumbbells are collecting dust and feelings"
- "Even your shadow looks disappointed"
- "Your muscles are sending me strongly worded letters"

**FOMO (Fear of Missing Out):**
- "While you were deciding, Brad already finished his workout"
- "The gym equipment is having a party without you"
- "Your future swole self is tapping their foot impatiently"

**Fake Urgency:**
- "Your workout window closes in 37 minutes!"
- "Motivation levels are at 73% and dropping!"
- "Quick! Your buddy just posted their workout!"

**Social Pressure:**
- "Your buddy thinks you forgot about leg day"
- "3 out of 4 dentists recommend working out today"
- "Your mom asked if you're still exercising..." (even if they didn't)

### **7. Psychological Manipulation Notifications**

**Variable Ratio Schedule:**
```typescript
// Random notification timing (most addictive)
const notificationSchedule = [
  { timeAfterMissed: '2 hours', probability: 0.3 },
  { timeAfterMissed: '6 hours', probability: 0.7 },
  { timeAfterMissed: '1 day', probability: 0.9 },
  { timeAfterMissed: '2 days', probability: 1.0, intensity: 'weasel' }
];
```

**Escalating Intensity:**
- **Day 1**: Gentle reminder
- **Day 2**: Disappointed friend
- **Day 3**: Full weasel mode activated
- **Day 4**: "Intervention" from app

### **8. Fake Social Features**
```typescript
// Create illusion of active community
interface FakeActivity {
  type: 'workout_completed' | 'pr_achieved' | 'streak_extended';
  fakeUserName: string;
  timestamp: Date;
  details: string;
}
```

**Examples:**
- "💪 Mike_Fitness just crushed their bench press PR!"
- "🔥 Sarah_Strong extended their streak to 12 days!"
- "⚡ FitnessFanatic22 just completed their 3rd workout this week!"

*(All algorithmically generated to create social pressure)*

## 📅 **Implementation Roadmap**

### **Week 1-2: Foundation**
**Priority 1 (Core):**
- User registration/login
- Basic profile setup
- Simple workout logging
- Database schema design

**Priority 2 (Weasel):**
- Motivational message system
- Basic notification framework

### **Week 3-4: Social & Gamification**
**Priority 1 (Core):**
- Buddy invitation system
- Workout sharing/visibility
- Progress tracking

**Priority 2 (Weasel):**
- Achievement system
- Streak tracking
- Variable reward messages

### **Week 5-6: Polish & Psychology**
**Priority 1 (Core):**
- Rest timer with notifications
- Weekly progress reports
- Bug fixes and optimization

**Priority 2 (Weasel):**
- Full weasel mode implementation
- Fake social activity
- Advanced psychological triggers

## 🛠️ **Technical Architecture**

### **Database Schema**
```sql
-- User system
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR UNIQUE,
  name VARCHAR,
  weasel_mode_enabled BOOLEAN DEFAULT true,
  weasel_intensity ENUM('gentle', 'medium', 'aggressive') DEFAULT 'medium'
);

-- Workout tracking
CREATE TABLE workouts (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  exercises JSONB,
  completed_at TIMESTAMP,
  fake_progress_boost INTEGER DEFAULT 0 -- Weasel mode stat inflation
);

-- Achievements & gamification
CREATE TABLE user_achievements (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  achievement_id VARCHAR,
  unlocked_at TIMESTAMP,
  fake_achievement BOOLEAN DEFAULT false -- Some achievements are... creative
);

-- Buddy system
CREATE TABLE buddy_relationships (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  buddy_id UUID REFERENCES users(id),
  relationship_type ENUM('peer', 'coach'),
  status ENUM('pending', 'active', 'paused')
);

-- Weasel mode messages
CREATE TABLE weasel_messages (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  message_type VARCHAR,
  content TEXT,
  sent_at TIMESTAMP,
  user_reaction ENUM('ignored', 'annoyed', 'motivated', 'worked_out')
);
```

### **Psychological Engine**
```typescript
class WeaselEngine {
  // Analyze user behavior and trigger appropriate manipulation
  async analyzeUserState(userId: string): Promise<WeaselStrategy> {
    const user = await getUserWithHistory(userId);
    const daysSinceLastWorkout = getDaysSince(user.lastWorkout);
    const streakRisk = calculateStreakRisk(user);
    const buddyActivity = await getBuddyActivity(user.buddyId);
    
    return {
      guilt_level: this.calculateGuiltLevel(daysSinceLastWorkout),
      fomo_triggers: this.generateFOMOTriggers(buddyActivity),
      urgency_level: this.calculateUrgencyLevel(streakRisk),
      social_pressure: this.generateSocialPressure(user, buddyActivity)
    };
  }
  
  async triggerWeaselMode(userId: string, intensity: 'gentle' | 'aggressive') {
    const strategy = await this.analyzeUserState(userId);
    const message = this.selectOptimalMessage(strategy, intensity);
    await this.sendNotification(userId, message);
    await this.logPsychologicalIntervention(userId, message);
  }
}
```

### **Mobile App Features**
```typescript
// Streak Counter with anxiety-inducing countdown
const StreakCounter = () => {
  const timeUntilStreakLoss = calculateTimeUntilStreakLoss();
  
  return (
    <View style={getStreakStyle(timeUntilStreakLoss)}>
      <Text>🔥 {currentStreak} day streak</Text>
      {timeUntilStreakLoss < 6 && (
        <Text style={styles.urgentWarning}>
          ⚠️ Streak expires in {timeUntilStreakLoss} hours!
        </Text>
      )}
    </View>
  );
};

// Fake progress celebration
const WorkoutComplete = () => {
  const fakeBoosts = generateFakeProgressStats();
  
  return (
    <View>
      <Text>🎉 Workout Complete!</Text>
      <Text>💪 Strength increased by {fakeBoosts.strength}%!</Text>
      <Text>🧠 Discipline increased by {fakeBoosts.discipline}%!</Text>
      <Text>😎 Attractiveness increased by {fakeBoosts.attractiveness}%!</Text>
    </View>
  );
};
```

## 🎭 **Weasel Mode Settings**

**User Control Panel:**
```typescript
interface WeaselSettings {
  enabled: boolean;
  intensity: 'gentle' | 'medium' | 'aggressive' | 'full_chaos';
  allow_guilt_trips: boolean;
  allow_fake_stats: boolean;
  allow_social_pressure: boolean;
  buddy_notification_enabled: boolean;
}
```

**Intensity Levels:**
- **Gentle**: Encouraging reminders, subtle nudges
- **Medium**: Regular notifications, mild social pressure
- **Aggressive**: Daily guilt trips, fake urgency, peer comparison
- **Full Chaos**: Everything enabled, maximum psychological manipulation

## 📊 **Success Metrics**

**Core Metrics:**
- Workout completion rate
- User retention (7-day, 30-day)
- Buddy invitation acceptance rate

**Weasel Mode Effectiveness:**
- Notification response rate by message type
- Workout completion after weasel intervention
- User tolerance levels (complaint rate)
- "Stockholm Syndrome Score" (users who start enjoying the nagging)

---

**Remember**: We're building an app that's **psychologically manipulative but ultimately helpful**. The goal is to trick users into building healthy habits while keeping it fun and self-aware about the manipulation! 🐾

Want me to start implementing any specific feature first? 