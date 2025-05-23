#!/bin/bash

# LiftBuddy Development Environment Verification Script
# Checks that all development services are running correctly

echo "⚡ Verifying Ferrovis development environment..."
echo ""

# Check if Docker is running
echo "🐳 Checking Docker..."
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running"
    exit 1
else
    echo "✅ Docker is running"
fi

# Check if PostgreSQL container is running
echo ""
echo "🐘 Checking PostgreSQL database..."
if docker ps --format 'table {{.Names}}' | grep -q "ferrovis-db"; then
    echo "✅ PostgreSQL container is running"
    
    # Test database connection
    if docker exec ferrovis-db pg_isready -U postgres -d ferrovis_dev > /dev/null 2>&1; then
        echo "✅ Database is accepting connections"
    else
        echo "❌ Database is not accepting connections"
    fi
else
    echo "❌ PostgreSQL container is not running"
    echo "💡 Run './scripts/dev-setup.sh' to start the database"
fi

# Check backend API
echo ""
echo "🔧 Checking backend API..."
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    response=$(curl -s http://localhost:8080/health)
    if echo "$response" | grep -q "ok"; then
        echo "✅ Backend API is running and healthy"
    else
        echo "❌ Backend API returned unexpected response"
    fi
else
    echo "❌ Backend API is not responding"
    echo "💡 Start with: cd backend && go run cmd/server/main.go"
fi

# Check if mobile development server is running
echo ""
echo "📱 Checking mobile development server..."
if curl -s http://localhost:8081 > /dev/null 2>&1 || \
   curl -s http://localhost:19000 > /dev/null 2>&1 || \
   curl -s http://localhost:19006 > /dev/null 2>&1; then
    echo "✅ Mobile development server is running"
else
    echo "❌ Mobile development server is not running"
    echo "💡 Start with: cd mobile && npm start"
fi

# Check environment files
echo ""
echo "📝 Checking environment configuration..."
if [ -f "backend/.env" ]; then
    echo "✅ Backend .env file exists"
else
    echo "❌ Backend .env file missing"
fi

if [ -f "mobile/.env" ]; then
    echo "✅ Mobile .env file exists"
else
    echo "❌ Mobile .env file missing"
fi

echo ""
echo "🎯 Development Environment Status Summary:"
echo "  Database: $(docker ps --format 'table {{.Names}}' | grep -q 'ferrovis-db' && echo '✅ Running' || echo '❌ Stopped')"
echo "  Backend:  $(curl -s http://localhost:8080/health > /dev/null 2>&1 && echo '✅ Running' || echo '❌ Stopped')"
echo "  Mobile:   $(curl -s http://localhost:8081 > /dev/null 2>&1 || curl -s http://localhost:19000 > /dev/null 2>&1 || curl -s http://localhost:19006 > /dev/null 2>&1 && echo '✅ Running' || echo '❌ Stopped')"

echo ""
echo "📋 Quick Commands:"
echo "  Database:    ./scripts/dev-setup.sh"
echo "  Backend:     cd backend && go run cmd/server/main.go"
echo "  Mobile:      cd mobile && npm start"
echo "  This Check:  ./scripts/verify-setup.sh" 