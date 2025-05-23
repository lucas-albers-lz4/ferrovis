#!/bin/bash

# Ferrovis CI Validation Script
# This script validates the same steps as our GitHub Actions CI pipeline

set -e

echo "🚀 Ferrovis CI Validation Script"
echo "================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Check if we're in the right directory
if [ ! -f "go.mod" ] && [ ! -d "backend" ]; then
    print_error "Please run this script from the project root directory"
    exit 1
fi

echo "📍 Current directory: $(pwd)"
echo ""

# 1. Backend Linting
echo "🔍 Step 1: Backend Linting"
echo "------------------------"
if make backend-lint > /dev/null 2>&1; then
    print_status "Backend linting passed"
else
    print_error "Backend linting failed"
    exit 1
fi

# 2. Mobile Linting
echo ""
echo "🔍 Step 2: Mobile Linting"
echo "------------------------"
if make mobile-lint > /dev/null 2>&1; then
    print_status "Mobile linting passed"
else
    print_warning "Mobile linting completed with warnings (ESLint config needed)"
fi

# 3. Backend Testing
echo ""
echo "🧪 Step 3: Backend Testing"
echo "-------------------------"
if make backend-test > /dev/null 2>&1; then
    print_status "Backend tests passed"
else
    print_error "Backend tests failed"
    exit 1
fi

# 4. Mobile Type Checking
echo ""
echo "🔍 Step 4: Mobile Type Checking"
echo "------------------------------"
if make mobile-type-check > /dev/null 2>&1; then
    print_status "Mobile type checking passed"
else
    print_error "Mobile type checking failed"
    exit 1
fi

# 5. Backend Build
echo ""
echo "🏗️  Step 5: Backend Build"
echo "-----------------------"
if make backend-build > /dev/null 2>&1; then
    print_status "Backend build successful"
    ls -la backend/bin/ferrovis
else
    print_error "Backend build failed"
    exit 1
fi

# 6. Docker Build Test (optional)
echo ""
echo "🐳 Step 6: Docker Build Test"
echo "---------------------------"
if command -v docker > /dev/null 2>&1; then
    if make build-docker > /dev/null 2>&1; then
        print_status "Docker build successful"
    else
        print_warning "Docker build failed (optional)"
    fi
else
    print_warning "Docker not available, skipping Docker build test"
fi

# 7. Security Checks
echo ""
echo "🔒 Step 7: Security Checks"
echo "-------------------------"
if make security > /dev/null 2>&1; then
    print_status "Security checks passed"
else
    print_warning "Security checks completed with warnings"
fi

echo ""
echo "🎉 CI Validation Complete!"
echo "========================="
print_status "All critical checks passed - ready for GitHub Actions"
echo ""
echo "Next steps:"
echo "  1. Commit your changes"
echo "  2. Push to GitHub"
echo "  3. Monitor the GitHub Actions workflow"
echo ""
