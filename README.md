# ⚡ Ferrovis

**Iron Strength** - A mobile-first fitness accountability app designed for absolute beginner weight lifters. Track your workouts, connect with a buddy or coach, and build sustainable lifting habits through social accountability and smart motivational techniques.

## 🎯 **Core Problem**

Most fitness apps are complex and intimidating for beginners. Ferrovis focuses on:
- **Simple workout tracking** (sets, reps, weight)
- **Accountability through social connection** (buddy/coach system)
- **Habit formation** using proven psychological techniques
- **Beginner-friendly guidance** to reduce gym anxiety

## 🏗️ **Architecture**

### **Mobile App**
- **Frontend**: Expo + React Native + TypeScript
- **Platform**: iOS and Android (single codebase)
- **Features**: Workout logging, progress tracking, buddy interactions

### **Backend API**
- **Language**: Go
- **Framework**: Gin (HTTP router)
- **Database**: PostgreSQL (AWS RDS)
- **Deployment**: AWS App Runner (containerized)
- **Authentication**: JWT tokens

### **Infrastructure**
- **Hosting**: AWS App Runner
- **Database**: AWS RDS PostgreSQL
- **Storage**: AWS S3 (future: progress photos)
- **Monitoring**: AWS CloudWatch
- **Cost**: ~$25-45/month

## 📱 **Key Features**

### **MVP Features (v1.0)**
- ✅ **Workout Logging**: Simple sets/reps/weight tracking
- ✅ **Buddy System**: Invite friend or coach for accountability
- ✅ **Progress Tracking**: Personal records and consistency streaks
- ✅ **Pre-built Programs**: Starting Strength, StrongLifts 5x5
- ✅ **Rest Timer**: Between-set timing with notifications

### **Planned Features (v2.0)**
- 📱 **Progress Photos**: Before/after visual tracking
- 🏆 **Achievement System**: Badges and milestones
- 📊 **Advanced Analytics**: Strength progression charts
- 💪 **Form Guides**: Exercise instruction videos
- 🔔 **Smart Notifications**: Accountability reminders

## 🧠 **Motivational Psychology**

Ferrovis uses evidence-based psychological techniques to build lasting habits:

### **Social Accountability**
- **Buddy Visibility**: Friends see your workout completion status
- **Weekly Reports**: Progress summaries shared with accountability partner
- **Social Pressure**: Gentle FOMO and peer comparison

### **Habit Formation**
- **Streak Tracking**: "Don't break the chain" mentality
- **Micro-Commitments**: "Just show up for 10 minutes" approach
- **Implementation Intentions**: Scheduled workout reminders

### **Progress Momentum**
- **Visual Progress**: Charts showing strength gains over time
- **Achievement Unlocks**: Celebration of milestones and PRs
- **Loss Aversion**: Protect earned streaks and progress

## 🚀 **Tech Stack Rationale**

### **Why Expo + React Native?**
- ✅ **Cross-platform**: Single codebase for iOS + Android
- ✅ **Fast Development**: Hot reload, great tooling
- ✅ **Large Ecosystem**: 2M+ npm packages available
- ✅ **Developer Experience**: Instant device testing via QR code

### **Why Go Backend?**
- ✅ **Performance**: Excellent for concurrent API requests
- ✅ **Simplicity**: Single binary deployment
- ✅ **JSON APIs**: Perfect match with React Native
- ✅ **AWS Support**: Official SDKs and great container support

### **Why AWS App Runner?**
- ✅ **Learning Goal**: Hands-on AWS experience
- ✅ **Simplicity**: Deploy like Heroku, but on AWS
- ✅ **Cost-Effective**: Auto-scaling containers
- ✅ **Production-Ready**: Can handle real traffic

## 🛠️ **Local Development & Testing**

### **Quick Start**
```bash
# Setup development environment
make setup

# Start development servers
make dev

# Run all tests
make test

# Run linting
make lint
```

### **GitHub Actions Local Testing**
Test your workflows locally before pushing using [act](https://github.com/nektos/act):

```bash
# Install act (macOS)
brew install act

# Test CI pipeline locally
make act-test

# Test only linting
make act-lint

# Test only builds
make act-build

# Test release process
make act-release
```

**Benefits:**
- ✅ Catch CI failures before pushing
- ✅ Save GitHub Actions minutes
- ✅ Faster development feedback loop
- ✅ Test on Apple Silicon with proper configuration

**For detailed troubleshooting and configuration, see [docs/LOCAL_TESTING.md](docs/LOCAL_TESTING.md)**

### **Development Commands**
- `make help` - Show all available commands
- `make docker-up` - Start database and services
- `make backend-run` - Run Go API server locally
- `make mobile-start` - Start React Native development server
- `make db-reset` - Reset database with fresh schema

## 📊 **Target Users**

### **Primary: Absolute Beginner Lifters**
- Never lifted weights before or inconsistent
- Intimidated by gym environment
- Need structure and accountability
- Want simple, non-overwhelming tools

### **Secondary: Accountability Partners**
- **Lifting Buddies**: Peers also starting their fitness journey
- **Coaches/Trainers**: Professional or experienced lifters helping beginners
- **Friends/Family**: Support network invested in user's success

## 🎨 **Design Principles**

### **Simplicity First**
- Minimal cognitive load
- Clear, obvious actions
- No overwhelming features or options

### **Beginner-Focused**
- Reduce gym anxiety through guidance
- Celebrate small wins
- "Good enough" is better than perfect

### **Social by Design**
- Accountability built into core workflows
- Progress sharing is frictionless
- Community support without comparison pressure

## 📈 **Success Metrics**

### **User Engagement**
- Weekly active users (target: 60%+ retention)
- Workout completion rate (target: 70%+)
- Buddy interaction frequency

### **Habit Formation**
- 7-day workout streaks (target: 40% of users)
- 30-day retention (target: 30%)
- Progress photo uploads

### **Social Features**
- Buddy invitations sent/accepted
- Weekly report engagement
- Accountability check-ins

## 🛣️ **Development Roadmap**

### **Phase 1: MVP (Weeks 1-4)**
- Basic workout logging
- User authentication
- Simple buddy invitation system
- Progress tracking

### **Phase 2: Enhancement (Weeks 5-8)**
- Push notifications
- Streak tracking
- Basic analytics dashboard
- App store submission

### **Phase 3: Growth (Months 3-6)**
- Advanced social features
- Coach-specific tools
- Analytics and insights
- User acquisition features

## 🤝 **Contributing**

This is a learning project focused on:
- Modern mobile app development
- AWS cloud infrastructure
- Go backend development
- User experience design for fitness apps

See [DEVELOPMENT.md](DEVELOPMENT.md) for detailed setup instructions and [docs/DEPENDENCY_MANAGEMENT.md](docs/DEPENDENCY_MANAGEMENT.md) for our dependency strategy.

## 📄 **License**

MIT License - feel free to learn from and adapt this code for your own projects.

---

**Built with ❤️ for the fitness community**
