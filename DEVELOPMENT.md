# 🛠️ LiftBuddy Development Guide

Complete development setup and workflow documentation for the LiftBuddy fitness accountability app.

## 📋 **Prerequisites**

### **Required Software**
```bash
# Node.js (v18 or later)
node --version  # Should be 18+

# npm or yarn
npm --version

# Git
git --version

# Go (v1.21 or later)
go version

# Docker
docker --version

# AWS CLI
aws --version
```

### **Development Accounts**
- **Expo Account**: [expo.dev](https://expo.dev) (free)
- **AWS Account**: [aws.amazon.com](https://aws.amazon.com)
- **Apple Developer**: [developer.apple.com](https://developer.apple.com) ($99/year, for iOS)
- **Google Play Console**: [play.google.com/console](https://play.google.com/console) ($25 one-time)

### **Recommended Tools**
- **VS Code** with extensions:
  - React Native Tools
  - Go extension
  - ES7+ React/Redux/React-Native snippets
  - Prettier
  - Thunder Client (API testing)

## 🏗️ **Project Structure**

```
liftbuddy/
├── mobile/                 # Expo React Native app
│   ├── src/
│   │   ├── components/     # Reusable UI components
│   │   ├── screens/        # Screen components
│   │   ├── services/       # API calls and business logic
│   │   ├── types/          # TypeScript type definitions
│   │   └── utils/          # Helper functions
│   ├── app.json           # Expo configuration
│   ├── package.json
│   └── tsconfig.json
├── backend/                # Go API server
│   ├── cmd/
│   │   └── server/         # Main application entry
│   ├── internal/
│   │   ├── auth/           # Authentication logic
│   │   ├── database/       # Database models and migrations
│   │   ├── handlers/       # HTTP route handlers
│   │   └── middleware/     # HTTP middleware
│   ├── pkg/                # Shared packages
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── docs/                   # Documentation
├── scripts/                # Build and deployment scripts
├── README.md
└── DEVELOPMENT.md
```

## 🚀 **Quick Start**

### **1. Clone and Setup**
```bash
# Clone repository
git clone https://github.com/yourusername/liftbuddy.git
cd liftbuddy

# Setup mobile app
cd mobile
npm install
npx expo install

# Setup backend
cd ../backend
go mod download
```

### **2. Environment Configuration**
```bash
# Mobile app (.env in mobile/)
EXPO_PUBLIC_API_URL=http://localhost:8080
EXPO_PUBLIC_APP_ENV=development

# Backend (.env in backend/)
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_NAME=liftbuddy_dev
DB_USER=postgres
DB_PASSWORD=yourpassword
JWT_SECRET=your-secret-key
AWS_REGION=us-east-1
```

### **3. Local Database Setup**
```bash
# Using Docker for local PostgreSQL
docker run --name liftbuddy-db \
  -e POSTGRES_DB=liftbuddy_dev \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=yourpassword \
  -p 5432:5432 \
  -d postgres:15

# Run database migrations
cd backend
go run cmd/migrate/main.go
```

### **4. Start Development Servers**
```bash
# Terminal 1: Backend
cd backend
go run cmd/server/main.go

# Terminal 2: Mobile app
cd mobile
npm start
# Scan QR code with Expo Go app on your phone
```

## 📱 **Mobile Development Workflow**

### **Expo Development Cycle**
```bash
# Start development server
npm start

# Run on specific platforms
npm run ios       # iOS simulator
npm run android   # Android emulator
npm run web       # Web browser

# Reset cache (when things break)
npx expo start --clear
```

### **Testing on Physical Devices**
1. **Install Expo Go** app from App Store/Play Store
2. **Scan QR code** from terminal or browser
3. **Instant updates** when you save files
4. **Shake device** to open developer menu

### **Key Development Commands**
```bash
# Install new dependency
npx expo install package-name

# Type checking
npm run type-check

# Linting
npm run lint

# Build for testing
eas build --platform ios --profile preview
```

## 🔧 **Backend Development Workflow**

### **Go Development Setup**
```bash
# Install Air for hot reloading
go install github.com/cosmtrek/air@latest

# Start with hot reload
air

# Without hot reload
go run cmd/server/main.go

# Run tests
go test ./...

# Build binary
go build -o bin/server cmd/server/main.go
```

### **Database Operations**
```bash
# Create new migration
go run cmd/migrate/main.go create add_workouts_table

# Run migrations
go run cmd/migrate/main.go up

# Rollback migration
go run cmd/migrate/main.go down

# Reset database
go run cmd/migrate/main.go reset
```

### **API Testing**
```bash
# Test endpoints with curl
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Or use Thunder Client in VS Code
# Import postman collection from docs/api/
```

## 🌐 **AWS Deployment Workflow**

### **Backend Deployment (App Runner)**
```bash
# Build and push Docker image
docker build -t liftbuddy-backend .
docker tag liftbuddy-backend:latest 123456789.dkr.ecr.us-east-1.amazonaws.com/liftbuddy:latest
docker push 123456789.dkr.ecr.us-east-1.amazonaws.com/liftbuddy:latest

# Deploy via AWS CLI (or use console)
aws apprunner create-service \
  --service-name liftbuddy-api \
  --source-configuration file://apprunner.yaml
```

### **Database Setup (RDS)**
```bash
# Create RDS instance
aws rds create-db-instance \
  --db-instance-identifier liftbuddy-db \
  --db-instance-class db.t3.micro \
  --engine postgres \
  --master-username postgres \
  --master-user-password yourpassword \
  --allocated-storage 20
```

### **Mobile App Deployment (EAS)**
```bash
# Configure EAS
eas login
eas build:configure

# Build for iOS TestFlight
eas build --platform ios --profile preview

# Build for Google Play
eas build --platform android --profile preview

# Submit to app stores
eas submit --platform ios
eas submit --platform android
```

## 🧪 **Testing Strategy**

### **Mobile App Testing**
```bash
# Unit tests with Jest
npm test

# E2E tests with Detox (optional)
npm run test:e2e

# Type checking
npm run type-check
```

### **Backend Testing**
```bash
# Unit tests
go test ./internal/...

# Integration tests
go test ./test/integration/...

# Load testing with k6
k6 run scripts/load-test.js
```

## 🔐 **Security Considerations**

### **Environment Variables**
- **Never commit** `.env` files
- **Use different secrets** for each environment
- **Rotate JWT secrets** regularly

### **API Security**
- **Rate limiting** on all endpoints
- **Input validation** for all user data
- **SQL injection prevention** with parameterized queries
- **CORS configuration** for mobile app

### **Mobile Security**
- **Secure storage** for auth tokens
- **Certificate pinning** for API calls
- **Biometric authentication** (future feature)

## 📊 **Monitoring and Debugging**

### **Backend Monitoring**
```bash
# CloudWatch logs
aws logs tail /aws/apprunner/liftbuddy-api --follow

# Local debugging
go run cmd/server/main.go --debug

# Profiling
go tool pprof http://localhost:8080/debug/pprof/profile
```

### **Mobile Debugging**
- **Flipper** for advanced debugging
- **React Native Debugger**
- **Expo Dev Tools** in browser
- **Device logs** via Expo Go app

## 🚀 **Performance Optimization**

### **Mobile Performance**
- **Image optimization** (WebP, compression)
- **Bundle size analysis** with `npx expo-bundle-analyzer`
- **Memory profiling** with Flipper
- **Lazy loading** for screens

### **Backend Performance**
- **Database indexing** on frequently queried columns
- **Connection pooling** for PostgreSQL
- **Response caching** for static data
- **Gzip compression** for API responses

## 🐛 **Common Issues and Solutions**

### **Mobile Development**
```bash
# Metro bundler cache issues
npx expo start --clear

# iOS build issues
cd ios && pod install

# Android issues
npx expo run:android --variant debug

# Node modules issues
rm -rf node_modules && npm install
```

### **Backend Issues**
```bash
# Module not found
go mod tidy

# Database connection issues
docker ps  # Check if PostgreSQL is running

# Port already in use
lsof -ti:8080 | xargs kill -9
```

## 📚 **Learning Resources**

### **React Native / Expo**
- [Expo Documentation](https://docs.expo.dev/)
- [React Native Documentation](https://reactnative.dev/)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)

### **Go Development**
- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [GORM ORM](https://gorm.io/docs/)

### **AWS Services**
- [App Runner Documentation](https://docs.aws.amazon.com/apprunner/)
- [RDS Documentation](https://docs.aws.amazon.com/rds/)
- [AWS CLI Reference](https://docs.aws.amazon.com/cli/)

## 🤝 **Contributing Guidelines**

### **Code Style**
- **TypeScript**: Use Prettier + ESLint
- **Go**: Use `gofmt` and `golint`
- **Commit messages**: Follow conventional commits

### **Pull Request Process**
1. Create feature branch from `main`
2. Implement feature with tests
3. Update documentation if needed
4. Create pull request with description
5. Request code review

### **Release Process**
1. Update version numbers
2. Create release notes
3. Tag release in Git
4. Deploy to staging
5. Test thoroughly
6. Deploy to production

---

**Happy coding! 🚀** 