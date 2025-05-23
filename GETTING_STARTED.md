# 🚀 LiftBuddy - Getting Started

**Quick setup guide to get LiftBuddy running in 5 minutes!**

## ⚡ Quick Start

### 1. **Prerequisites**
```bash
# Ensure you have these installed:
node --version    # Should be 18+
go version        # Should be 1.21+
docker --version  # Any recent version
```

### 2. **Automated Setup**
```bash
# Run the setup script (handles everything!)
./scripts/dev-setup.sh
```

### 3. **Manual Setup** (if script doesn't work)
```bash
# Start database
docker run --name liftbuddy-db \
  -e POSTGRES_DB=liftbuddy_dev \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=yourpassword \
  -p 5432:5432 -d postgres:15

# Setup backend
cd backend
go mod tidy
go run cmd/server/main.go

# Setup mobile (in new terminal)
cd mobile
npm install
npm start
```

## 📱 **Testing the App**

### **Backend API**
- Visit: http://localhost:8080/health
- Should see: `{"status":"ok","message":"LiftBuddy API is running"}`

### **Mobile App**
1. Install **Expo Go** app on your phone
2. Run `npm start` in the `mobile/` directory
3. Scan the QR code with your phone
4. The app should show "API Connected ✅"

## 🎯 **What You'll See**

### **Mobile App Features**
- ✅ Clean, modern UI
- ✅ API connection status
- ✅ Welcome message and feature preview
- ✅ Real-time updates with hot reload

### **Backend API Features**
- ✅ Health check endpoint
- ✅ Placeholder routes for all planned features
- ✅ CORS configuration for mobile app
- ✅ Ready for database integration

## 🛠️ **Next Development Steps**

1. **Database Integration**
   - Add PostgreSQL models (GORM)
   - Create database migrations
   - Implement actual endpoints

2. **Authentication**
   - JWT token system
   - User registration/login
   - Protected routes

3. **Core Features**
   - Workout logging
   - Buddy system
   - Progress tracking

4. **Mobile Screens**
   - Login/Register screens
   - Workout logging interface
   - Progress charts

## 🐛 **Troubleshooting**

### **Backend won't start**
```bash
# Check if port 8080 is in use
lsof -ti:8080 | xargs kill -9

# Reinstall dependencies
cd backend && go mod tidy
```

### **Mobile app won't connect**
```bash
# Clear Metro cache
cd mobile && npx expo start --clear

# Reinstall dependencies
rm -rf node_modules && npm install
```

### **Database issues**
```bash
# Remove and restart database
docker rm -f liftbuddy-db
./scripts/dev-setup.sh
```

## 📚 **Learn More**

- **Full Documentation**: [DEVELOPMENT.md](DEVELOPMENT.md)
- **Project Overview**: [README.md](README.md)
- **Architecture Details**: See README.md Architecture section

## 🏆 **Success Criteria**

You're ready to develop when:
- ✅ Backend API responds at http://localhost:8080/health
- ✅ Mobile app shows "API Connected ✅"
- ✅ Hot reload works (edit code, see instant changes)
- ✅ No error messages in terminals

**Happy coding! 🎉** 