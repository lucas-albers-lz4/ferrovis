#!/bin/bash

# LiftBuddy Development Setup Script
# Sets up local development environment with PostgreSQL database

set -e  # Exit on any error

echo "⚡ Setting up Ferrovis development environment..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker Desktop and try again."
    exit 1
fi

# Database configuration
DB_NAME="ferrovis_dev"
DB_USER="postgres"
DB_PASSWORD="ferrovis123"
DB_PORT="5432"
CONTAINER_NAME="ferrovis-db"

echo "📦 Setting up PostgreSQL database container..."

# Stop and remove existing container if it exists
if docker ps -a --format 'table {{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    echo "🗑️  Removing existing database container..."
    docker stop $CONTAINER_NAME || true
    docker rm $CONTAINER_NAME || true
fi

# Create and start PostgreSQL container
echo "🚀 Starting PostgreSQL container..."
docker run --name $CONTAINER_NAME \
    -e POSTGRES_DB=$DB_NAME \
    -e POSTGRES_USER=$DB_USER \
    -e POSTGRES_PASSWORD=$DB_PASSWORD \
    -p $DB_PORT:5432 \
    -d postgres:15

# Wait for PostgreSQL to be ready
echo "⏳ Waiting for PostgreSQL to be ready..."
sleep 5

# Test database connection
echo "🔗 Testing database connection..."
until docker exec $CONTAINER_NAME pg_isready -U $DB_USER -d $DB_NAME; do
    echo "Waiting for PostgreSQL..."
    sleep 2
done

echo "✅ PostgreSQL is ready!"

# Create .env file for backend
echo "📝 Creating backend .env file..."
cat > backend/.env << EOF
# Database configuration
DB_HOST=localhost
DB_PORT=$DB_PORT
DB_NAME=$DB_NAME
DB_USER=$DB_USER
DB_PASSWORD=$DB_PASSWORD
DB_SSLMODE=disable

# JWT configuration
JWT_SECRET=your_super_secret_jwt_key_change_in_production

# Server configuration
PORT=8080

# Development mode
ENVIRONMENT=development
EOF

# Create environment config for mobile app
echo "📱 Creating mobile app environment config..."
cat > mobile/.env << EOF
# Backend API URL (local development)
EXPO_PUBLIC_API_URL=http://localhost:8080

# Environment
EXPO_PUBLIC_ENVIRONMENT=development
EOF

echo ""
echo "🎉 Development environment setup complete!"
echo ""
echo "📋 Next steps:"
echo "1. Start the backend server:"
echo "   cd backend && go run cmd/server/main.go"
echo ""
echo "2. Start the mobile app:"
echo "   cd mobile && npm start"
echo ""
echo "3. Database connection details:"
echo "   Host: localhost"
echo "   Port: $DB_PORT"
echo "   Database: $DB_NAME"
echo "   Username: $DB_USER"
echo "   Password: $DB_PASSWORD"
echo ""
echo "📊 Database management:"
echo "   Connect with psql: docker exec -it $CONTAINER_NAME psql -U $DB_USER -d $DB_NAME"
echo "   Stop database: docker stop $CONTAINER_NAME"
echo "   Start database: docker start $CONTAINER_NAME"
echo "" 