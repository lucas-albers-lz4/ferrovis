# TODO.md - Ferrovis MVP Implementation Plan

# Iron Strength - Fitness Accountability App with Full Weasel Mode™ Psychological Manipulation
## Implementation Sequence Notes
- Phase 0 (Infrastructure) must be completed before all other phases
- Phase 1 (Core Auth) should be completed before Phase 2 (Workout System)
- Phase 2 (Workout System) should be completed before Phase 3 (Weasel Psychology Engine)
- Phase 3 (Weasel Engine) should be completed before Phase 4 (Social Features)
- Phase 4 (Social Features) enhances all previous phases with accountability mechanisms
- Phase 5 (Gamification) builds on Phase 3 and 4 for maximum psychological impact
- Phase 6 (Deployment) prepares for production release

## Phase 0: Infrastructure & Database Setup (P0: Critical Foundation) [COMPLETED] ✅

### Goal: Establish core infrastructure for mobile app + Go backend + PostgreSQL database
### Scope: Database schema, API foundation, mobile app structure, development tooling

**🎉 Phase 0 Summary - COMPLETED:**
- ✅ Local PostgreSQL database running in Docker (`liftbuddy-db`)
- ✅ Go backend API with GORM + PostgreSQL integration
- ✅ Complete database schema with 10 tables (users, workouts, achievements, weasel_messages, etc.)
- ✅ Seed data loaded (Starting Strength/5x5 programs, exercises, achievements, fake social activity)
- ✅ Structured logging with slog throughout backend
- ✅ Development scripts (setup, verification) working perfectly
- ✅ Mobile app structure ready for development
- ✅ Git repository properly configured (embedded repo issue resolved)

#### Database Schema Implementation ✅
- [x] **Create core database migrations:**
  - [x] Users table with weasel mode preferences
  - [x] Workouts table with fake progress boost tracking
  - [x] User achievements table (including fake achievements)
  - [x] Buddy relationships table
  - [x] Weasel messages table with user reaction tracking
  - [x] Workout programs table (Starting Strength, 5x5, etc.)
  - [x] Exercises table with metadata for weasel calculations

- [x] **Backend API Foundation:**
  - [x] Set up GORM with PostgreSQL connection
  - [x] Implement database migration system
  - [x] Create basic CRUD operations for all entities
  - [ ] Add JWT authentication middleware
  - [x] Implement CORS and security middleware
  - [ ] Add request validation and error handling

- [ ] **Mobile App Foundation:**
  - [ ] Set up navigation system (React Navigation)
  - [ ] Create reusable UI components library
  - [ ] Implement state management (Context API or Zustand)
  - [ ] Add secure token storage
  - [ ] Set up push notification framework (Expo Notifications)

## Phase 1: User Authentication & Weasel Onboarding (P0: Critical User Experience) [PENDING]

### Goal: Implement user registration/login with psychological manipulation from first interaction
### Scope: Authentication flow, profile setup, initial weasel mode calibration

#### Core Authentication System
- [ ] **Backend Authentication:**
  - [ ] User registration endpoint with validation
  - [ ] User login endpoint with JWT token generation
  - [ ] Password hashing and security best practices
  - [ ] User profile CRUD endpoints
  - [ ] Email verification system (optional for MVP)

- [ ] **Mobile Authentication Flow:**
  - [ ] Registration screen with progressive disclosure
  - [ ] Login screen with persistent session
  - [ ] Profile setup wizard
  - [ ] Secure token storage and automatic refresh

#### Weasel Mode Onboarding
- [ ] **Psychological Onboarding Flow:**
  - [ ] "Quick 30-second setup" (actually 5 minutes) with progress bar manipulation
  - [ ] Fake social proof during loading screens ("247 people joined today!")
  - [ ] Goal selection with artificially inflated success statistics
  - [ ] Commitment escalation questioning ("Just one more quick question...")
  - [ ] Weasel mode intensity selection with humorous descriptions

- [ ] **Onboarding Weasel Messages:**
  - [ ] Dynamic fake urgency counters
  - [ ] Social proof testimonials (algorithmically generated)
  - [ ] Progress anchoring ("You're already 23% more committed than yesterday!")

## Phase 2: Core Workout System & Basic Weasel Rewards (P0: Core Functionality) [PENDING]

### Goal: Implement essential workout logging with immediate psychological gratification
### Scope: Workout programs, exercise tracking, completion rewards, basic progress inflation

#### Workout Program System
- [ ] **Pre-built Workout Programs:**
  - [ ] Starting Strength program implementation
  - [ ] StrongLifts 5x5 program implementation
  - [ ] Program selection with "personalized" recommendations (rigged)
  - [ ] Exercise database with instructions and form tips
  - [ ] Progressive overload calculation

- [ ] **Workout Logging Interface:**
  - [ ] Exercise selection with search and favorites
  - [ ] Set/rep/weight tracking with easy increment buttons
  - [ ] Rest timer with motivational countdown messages
  - [ ] Workout completion marking with celebration
  - [ ] Session notes and comments

#### Basic Weasel Reward System
- [ ] **Immediate Gratification Engine:**
  - [ ] Variable reward messages on workout completion
  - [ ] Fake progress statistics generation (strength boost %, etc.)
  - [ ] Progress inflation algorithms (make small gains look impressive)
  - [ ] Sunk cost fallacy messaging ("You've invested X minutes this week...")
  - [ ] Random "achievement unlocked" notifications

- [ ] **Progress Visualization Tricks:**
  - [ ] Charts with misleading scales to amplify progress
  - [ ] "Fitness score" calculation with arbitrary but impressive numbers
  - [ ] Projection magic ("At this rate, you'll bench 200lbs by March!")
  - [ ] Cherry-picked statistics highlighting

## Phase 3: Weasel Psychology Engine & Notification System (P1: Psychological Manipulation) [PENDING]

### Goal: Implement sophisticated psychological manipulation system for habit formation
### Scope: Message engine, notification scheduling, user behavior analysis, intervention triggers

#### Weasel Message Classification System
- [ ] **Message Categories Implementation:**
  - [ ] Guilt trip messages ("Your dumbbells are collecting dust and feelings")
  - [ ] FOMO triggers ("The gym equipment is having a party without you")
  - [ ] Fake urgency ("Your workout window closes in 37 minutes!")
  - [ ] Social pressure ("3 out of 4 dentists recommend working out today")
  - [ ] Humorous nagging ("Your future swole self is tapping their foot")

- [ ] **Psychological Trigger Engine:**
  - [ ] User behavior analysis (days since last workout, streak risk, etc.)
  - [ ] Optimal timing calculations for maximum impact
  - [ ] Variable ratio scheduling for addiction-level engagement
  - [ ] Escalating intensity based on user resistance
  - [ ] A/B testing framework for message effectiveness

#### Advanced Notification System
- [ ] **Smart Notification Scheduling:**
  - [ ] Variable ratio schedule implementation (most psychologically addictive)
  - [ ] Time-of-day optimization based on user patterns
  - [ ] Escalating intensity over missed workout days
  - [ ] Context-aware messaging (weather, day of week, etc.)
  - [ ] User tolerance monitoring and adjustment

- [ ] **Weasel Mode Intensity Controls:**
  - [ ] Gentle mode: Encouraging reminders only
  - [ ] Medium mode: Regular notifications + mild social pressure
  - [ ] Aggressive mode: Daily guilt trips + fake urgency + peer comparison
  - [ ] Full Chaos mode: Maximum psychological manipulation enabled

## Phase 4: Buddy System & Social Accountability (P1: Social Psychology) [PENDING]

### Goal: Implement social accountability features with peer pressure mechanisms
### Scope: Buddy invitations, workout sharing, social pressure generation, coach mode

#### Buddy Invitation System
- [ ] **Core Buddy Features:**
  - [ ] Email/SMS buddy invitation system
  - [ ] Buddy vs Coach role selection and permissions
  - [ ] Workout completion visibility and notifications
  - [ ] Weekly progress report sharing
  - [ ] Basic in-app messaging/comments

- [ ] **Social Pressure Generation:**
  - [ ] Buddy comparison dashboards ("You're behind Alex by 2 workouts")
  - [ ] Competitive streaks and challenges
  - [ ] Disappointment guilt messaging ("Alex is probably wondering where you are...")
  - [ ] Peer notification when someone completes a workout
  - [ ] Social proof generation based on buddy activity

#### Coach Mode Features
- [ ] **Coach-Specific Functionality:**
  - [ ] Multiple client management dashboard
  - [ ] Program assignment and customization
  - [ ] Progress monitoring and reporting tools
  - [ ] Custom motivational message creation
  - [ ] Client accountability score tracking

- [ ] **Advanced Social Features:**
  - [ ] Group challenges and competitions
  - [ ] Social feed of workout completions (real + fake)
  - [ ] Buddy workout synchronization suggestions
  - [ ] Accountability partner matching algorithm

## Phase 5: Gamification & Achievement System (P1: Engagement Optimization) [PENDING]

### Goal: Implement comprehensive gamification with psychological hooks
### Scope: Achievements, streaks, badges, leaderboards, fake social activity

#### Achievement System
- [ ] **Achievement Categories:**
  - [ ] Consistency achievements ("Show Up Samurai" - 7-day streak)
  - [ ] Strength achievements ("Weight Warrior" - first PR)
  - [ ] Social achievements ("Team Player" - helped buddy 5 times)
  - [ ] Funny achievements ("Sweat Sommelier" - 3 different workout times)
  - [ ] Creative fake achievements for psychological boost

- [ ] **Psychological Achievement Mechanics:**
  - [ ] Near-miss psychology ("1 workout away from 'Iron Will' badge!")
  - [ ] Loss aversion messaging ("Don't lose your 6-day streak!")
  - [ ] Artificial scarcity ("Only 12% of users earn this badge")
  - [ ] Progress anchoring and goal post manipulation

#### Streak System & Fake Social Activity
- [ ] **Streak Mechanics:**
  - [ ] Workout streaks with anxiety-inducing countdown timers
  - [ ] Streak protection mechanisms and "grace periods"
  - [ ] Multiple streak types (daily, weekly, monthly)
  - [ ] Streak recovery options with psychological cost

- [ ] **Fake Social Community:**
  - [ ] Algorithmic generation of fake user activities
  - [ ] Realistic fake usernames and workout achievements
  - [ ] Timed fake notifications for social pressure
  - [ ] Community leaderboards (real + fake users for motivation)

## Phase 6: Production Deployment & Monitoring (P0: Release Preparation) [PENDING]

### Goal: Deploy to AWS infrastructure with monitoring and analytics
### Scope: AWS App Runner deployment, database setup, monitoring, app store preparation

#### AWS Infrastructure Setup
- [ ] **Backend Deployment:**
  - [ ] AWS RDS PostgreSQL setup with backups
  - [ ] AWS App Runner service configuration
  - [ ] Environment variable and secrets management
  - [ ] SSL/TLS certificate setup
  - [ ] CloudWatch logging and monitoring

- [ ] **Mobile App Deployment:**
  - [ ] Expo Application Services (EAS) setup
  - [ ] iOS TestFlight deployment
  - [ ] Google Play Console setup and APK generation
  - [ ] Push notification service configuration (APNs, FCM)

#### Analytics & Monitoring
- [ ] **Weasel Mode Analytics:**
  - [ ] User behavior tracking and analysis
  - [ ] Message effectiveness measurement
  - [ ] Psychological intervention success rates
  - [ ] User tolerance and complaint monitoring
  - [ ] "Stockholm Syndrome Score" calculation

- [ ] **Performance Monitoring:**
  - [ ] API response time monitoring
  - [ ] Database performance optimization
  - [ ] Mobile app crash reporting
  - [ ] User retention and engagement metrics

## REMINDER On the Implementation Process: (DONT REMOVE THIS SECTION)
- For each change:
  1. **Baseline Verification:**
     - Run backend tests: `cd backend && go test ./...` ✓
     - Run mobile tests: `cd mobile && npm test` ✓
     - Run linting: `cd backend && golangci-lint run` ✓
     - Check mobile type safety: `cd mobile && npm run type-check` ✓

  2. **Pre-Change Verification:**
     - Run targeted tests relevant to the component being modified ✓
     - Test API endpoints with curl or Thunder Client ✓
     - Test mobile screens in Expo Go app ✓

  3. **Make Required Changes:**
     - Follow KISS and YAGNI principles ✓
     - Maintain consistent code style (Go: gofmt, TypeScript: Prettier) ✓
     - Document weasel mode features with clear psychological rationale ✓
     - **For database changes:**
       - Create proper migrations with rollback capability
       - Test migrations on sample data
       - Update API endpoints and mobile models accordingly
     - **For weasel mode features:**
       - Implement user controls for intensity levels
       - Add metrics tracking for psychological intervention effectiveness
       - Ensure features remain helpful despite manipulation tactics
     - **NEW REMINDER:** Test both backend API and mobile app integration frequently. Don't wait until the end of a feature.

  4. **Post-Change Verification:**
     - Run targeted tests to verify the changes work as expected ✓
     - Test end-to-end workflows (registration → onboarding → workout logging) ✓
     - Verify mobile app hot reload works correctly ✓
     - Run full test suite: `cd backend && go test ./...` ✓
     - Run mobile tests: `cd mobile && npm test` ✓
     - **CRITICAL:** Test weasel mode features don't become genuinely harmful or overly annoying

  5. **Git Commit:**
     - Stop after completing a logical portion of a feature to make well reasoned git commits ✓
     - Request suggested git commands for committing the changes ✓
     - Use conventional commit format: `feat: add user registration with weasel onboarding` ✓
     - Review and execute the git commit commands yourself ✓

  6. **Building and Testing Hints:**
     - Backend: `cd backend && go run cmd/server/main.go` starts development server
     - Mobile: `cd mobile && npm start` starts Expo development server
     - Database: `docker run --name liftbuddy-db -e POSTGRES_DB=liftbuddy_dev -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=yourpassword -p 5432:5432 -d postgres:15`
     - Full setup: `./scripts/dev-setup.sh` runs automated development setup
     - API testing: Use `curl` commands or Thunder Client extension in VS Code
     - Mobile testing: Scan QR code with Expo Go app, shake device for developer menu

##END REMINDER On the Implementation Process:

## Testing Strategy Notes
- **Unit Tests:** Focus on weasel message selection algorithms, progress calculation inflation, and psychological trigger logic
- **Integration Tests:** Test complete user flows from registration through workout completion with weasel interventions
- **Psychological Testing:** A/B test different weasel mode intensities and message types for optimal engagement without annoyance
- **User Acceptance Testing:** Ensure features remain helpful and motivational despite manipulation tactics
- **Performance Testing:** Verify notification systems don't overwhelm users or drain battery life

## Weasel Mode Ethics Guidelines
- **Transparency:** Users should understand they're being psychologically influenced (make it part of the fun)
- **Beneficial Manipulation:** All psychological techniques should ultimately help users build healthy habits
- **User Control:** Always provide granular controls for weasel mode intensity
- **Escape Hatches:** Allow users to disable or reduce psychological pressure when needed
- **Self-Aware Humor:** Keep manipulation tactics lighthearted and obviously intentional (not deceptive)

## Success Metrics
- **Core Metrics:** Workout completion rate, user retention (7-day, 30-day), buddy invitation acceptance rate
- **Weasel Effectiveness:** Notification response rate by message type, workout completion after intervention
- **User Satisfaction:** App store ratings, user complaints vs compliments about weasel mode
- **Engagement:** Time spent in app, feature usage patterns, streak maintenance rates
- **Social Features:** Buddy relationship creation and maintenance, social pressure response rates 