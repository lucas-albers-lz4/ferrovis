# Ferrovis - Iron Strength Fitness App
# Makefile for development, testing, and deployment commands

.PHONY: help setup dev clean test build lint docker-up docker-down docker-logs mobile-install mobile-start backend-test backend-build backend-run db-reset db-shell git-setup

# Default target
help: ## Show this help message
	@echo "⚡ Ferrovis - Iron Strength Development Commands"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

## 🚀 Quick Start Commands
setup: ## Complete development environment setup
	@echo "⚡ Setting up Ferrovis development environment..."
	@make docker-up
	@make mobile-install
	@echo "✅ Setup complete! Run 'make dev' to start development"

dev: ## Start full development environment
	@echo "⚡ Starting Ferrovis development environment..."
	@make docker-up
	@echo "🔧 Backend API: http://localhost:8080"
	@echo "📱 Run 'make mobile-start' in another terminal for mobile app"
	@echo "🎯 Run 'make verify' to check all services"

verify: ## Verify all services are running correctly
	@echo "⚡ Verifying Ferrovis services..."
	@./scripts/verify-setup.sh

clean: ## Clean up development environment
	@echo "🧹 Cleaning up Ferrovis environment..."
	@make docker-down
	@docker system prune -f
	@echo "✅ Cleanup complete"

## 🐳 Docker Commands
docker-up: ## Start all services with Docker Compose
	@echo "🐳 Starting Ferrovis services..."
	@docker compose up -d database
	@echo "⏳ Waiting for database to be ready..."
	@docker compose up -d --wait

docker-down: ## Stop all Docker services
	@docker compose down

docker-logs: ## View logs from all services
	@docker compose logs -f

docker-build: ## Rebuild all Docker images
	@docker compose build --no-cache

## 🔧 Backend Commands
backend-run: ## Run backend server locally
	@echo "🔧 Starting Ferrovis backend server..."
	@cd backend && go run cmd/server/main.go

backend-build: ## Build backend binary
	@echo "🔨 Building Ferrovis backend..."
	@cd backend && go build -o bin/ferrovis cmd/server/main.go
	@echo "✅ Backend built: backend/bin/ferrovis"

backend-test: ## Run backend tests
	@echo "🧪 Running backend tests..."
	@cd backend && go test ./...

backend-lint: ## Run backend linting
	@echo "🔍 Linting backend code..."
	@cd backend && golangci-lint run

backend-deps: ## Update backend dependencies
	@echo "📦 Updating backend dependencies..."
	@cd backend && go mod tidy && go mod download

## 📱 Mobile Commands
mobile-install: ## Install mobile app dependencies
	@echo "📱 Installing mobile dependencies..."
	@cd mobile && npm install

mobile-start: ## Start mobile development server
	@echo "📱 Starting Ferrovis mobile app..."
	@cd mobile && npm start

mobile-build: ## Build mobile app for production
	@echo "📱 Building mobile app..."
	@cd mobile && npm run build

mobile-test: ## Run mobile app tests
	@echo "🧪 Running mobile tests..."
	@cd mobile && npm test

mobile-lint: ## Lint mobile app code
	@echo "🔍 Linting mobile code..."
	@cd mobile && npm run lint

mobile-type-check: ## Type check mobile app
	@echo "🔍 Type checking mobile app..."
	@cd mobile && npm run type-check

## 🗄️ Database Commands
db-reset: ## Reset database with fresh schema and seed data
	@echo "🗄️ Resetting Ferrovis database..."
	@docker compose down database
	@docker volume rm ferrovis_postgres_data 2>/dev/null || true
	@docker compose up -d database --wait
	@echo "✅ Database reset complete"

db-shell: ## Open PostgreSQL shell
	@echo "🗄️ Opening Ferrovis database shell..."
	@docker exec -it ferrovis-db psql -U postgres -d ferrovis_dev

db-backup: ## Backup database to file
	@echo "💾 Backing up Ferrovis database..."
	@mkdir -p backups
	@docker exec ferrovis-db pg_dump -U postgres ferrovis_dev > backups/ferrovis_$(shell date +%Y%m%d_%H%M%S).sql
	@echo "✅ Database backed up to backups/"

## 🧪 Testing Commands
test: ## Run all tests (backend + mobile)
	@echo "🧪 Running all Ferrovis tests..."
	@make backend-test
	@make mobile-test

test-integration: ## Run integration tests
	@echo "🧪 Running integration tests..."
	@cd backend && go test ./test/integration/...

lint: ## Run all linting (backend + mobile)
	@echo "🔍 Linting all Ferrovis code..."
	@make backend-lint
	@make mobile-lint

## 🏗️ Build Commands
build: ## Build all components (backend + mobile)
	@echo "🏗️ Building all Ferrovis components..."
	@make backend-build
	@make mobile-build

build-docker: ## Build production Docker image
	@echo "🐳 Building Ferrovis production Docker image..."
	@docker build -t ferrovis:latest -f backend/Dockerfile backend/

## 📚 Documentation Commands
docs: ## Generate and serve documentation
	@echo "📚 Generating Ferrovis documentation..."
	@echo "📖 README.md - Project overview"
	@echo "🔧 DEVELOPMENT.md - Development setup"
	@echo "📋 TODO.md - Implementation plan"

## 🌐 Git & Deployment Commands  
git-setup: ## Initialize git repository and create GitHub repo
	@echo "🌐 Setting up Ferrovis git repository..."
	@git add .
	@git commit -m "feat: initial Ferrovis (Iron Strength) project setup with weasel mode"
	@gh repo create ferrovis --public --description "Iron Strength - Fitness accountability app with Full Weasel Mode™ psychological manipulation"
	@git remote add origin https://github.com/lucas-albers-lz4/ferrovis.git
	@git push -u origin main
	@echo "✅ Repository created at https://github.com/lucas-albers-lz4/ferrovis"

deploy-staging: ## Deploy to staging environment
	@echo "🚀 Deploying Ferrovis to staging..."
	@echo "TODO: Implement staging deployment"

deploy-prod: ## Deploy to production environment
	@echo "🚀 Deploying Ferrovis to production..."
	@echo "TODO: Implement production deployment"

## 📊 Monitoring Commands
health: ## Check health of all services
	@echo "💓 Checking Ferrovis service health..."
	@curl -s http://localhost:8080/health | jq || echo "❌ Backend not responding"
	@curl -s http://localhost:8080/schema-test | jq '.weasel_features.schema' || echo "❌ Database not responding"

logs: ## View application logs
	@echo "📋 Viewing Ferrovis logs..."
	@make docker-logs

## 🔧 Development Utilities
update: ## Update all dependencies
	@echo "📦 Updating all Ferrovis dependencies..."
	@make backend-deps
	@cd mobile && npm update
	@echo "✅ Dependencies updated"

format: ## Format all code
	@echo "✨ Formatting all Ferrovis code..."
	@cd backend && go fmt ./...
	@cd mobile && npm run format 2>/dev/null || echo "Mobile formatting skipped"

security: ## Run security checks
	@echo "🔒 Running Ferrovis security checks..."
	@cd backend && gosec ./... 2>/dev/null || echo "⚠️ Install gosec for security scanning"
	@cd mobile && npm audit || echo "⚠️ Mobile security audit complete"

## 📈 Analytics Commands
stats: ## Show project statistics
	@echo "📈 Ferrovis Project Statistics:"
	@echo "Backend Go files: $(shell find backend -name '*.go' | wc -l)"
	@echo "Mobile TypeScript files: $(shell find mobile -name '*.ts' -o -name '*.tsx' | wc -l)"
	@echo "Database tables: $(shell docker exec ferrovis-db psql -U postgres -d ferrovis_dev -t -c 'SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = '\''public'\'';' 2>/dev/null || echo '0')"
	@echo "Total lines of code: $(shell find backend mobile -name '*.go' -o -name '*.ts' -o -name '*.tsx' | xargs wc -l | tail -1 | awk '{print $$1}')"

# Include environment-specific makefiles if they exist
-include Makefile.local 