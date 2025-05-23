#!/bin/bash

# Ferrovis GitHub Actions Local Testing Script
# This script runs GitHub Actions locally using act with proper configuration

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Check if act is installed
if ! command -v act &> /dev/null; then
    print_error "act is not installed. Install it with: brew install act"
    exit 1
fi

# Check if Docker is running
if ! docker info &> /dev/null; then
    print_error "Docker is not running. Please start Docker Desktop."
    exit 1
fi

echo "🚀 Ferrovis GitHub Actions Local Testing"
echo "========================================"
print_info "Running GitHub Actions locally with act"
echo ""

# Check for GitHub token
GITHUB_TOKEN_FOR_ACT=""
if [ -n "$GITHUB_TOKEN" ]; then
    GITHUB_TOKEN_FOR_ACT="$GITHUB_TOKEN"
    print_info "Using GITHUB_TOKEN from environment"
elif [ -f ~/.config/gh/hosts.yml ]; then
    # Try to extract token from GitHub CLI config
    TOKEN_FROM_GH=$(gh auth status --show-token 2>/dev/null | grep -o 'Token: gho_[^[:space:]]*' | cut -d' ' -f2 || echo "")
    if [ -n "$TOKEN_FROM_GH" ]; then
        GITHUB_TOKEN_FOR_ACT="$TOKEN_FROM_GH"
        print_info "Using GitHub token from gh CLI"
    fi
fi

if [ -z "$GITHUB_TOKEN_FOR_ACT" ]; then
    print_warning "No GitHub token found. GitHub Actions may fail to clone actions."
    print_info "Set GITHUB_TOKEN environment variable or run 'gh auth login'"
    print_info "You can still test with cached actions or by skipping steps that need GitHub API access"
    GITHUB_TOKEN_FOR_ACT="fake_token_for_local_testing"
fi

# Show available options
if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    echo "Usage: $0 [workflow] [job]"
    echo ""
    echo "Available workflows:"
    echo "  ci        - Run CI pipeline (lint, test, build, security)"
    echo "  release   - Run release pipeline"
    echo "  lint      - Run only linting jobs"
    echo "  test      - Run only testing jobs"
    echo "  build     - Run only build jobs"
    echo "  security  - Run only security jobs"
    echo ""
    echo "Examples:"
    echo "  $0                    # Run CI workflow"
    echo "  $0 ci                 # Run CI workflow"
    echo "  $0 release            # Run release workflow"
    echo "  $0 ci lint            # Run only lint job from CI"
    echo "  $0 ci test            # Run only test job from CI"
    echo ""
    echo "Environment Variables:"
    echo "  GITHUB_TOKEN         - GitHub token for authentication"
    echo "  ACT_VERBOSE          - Set to '1' to enable verbose output"
    echo ""
    exit 0
fi

# Set workflow and job
WORKFLOW=${1:-ci}
JOB=${2:-}

# Workflow file mapping
case $WORKFLOW in
    "ci")
        WORKFLOW_FILE=".github/workflows/ci.yml"
        ;;
    "release")
        WORKFLOW_FILE=".github/workflows/release.yml"
        ;;
    "lint")
        WORKFLOW_FILE=".github/workflows/ci.yml"
        JOB="lint"
        ;;
    "test")
        WORKFLOW_FILE=".github/workflows/ci.yml"
        JOB="test"
        ;;
    "build")
        WORKFLOW_FILE=".github/workflows/ci.yml"
        JOB="build"
        ;;
    "security")
        WORKFLOW_FILE=".github/workflows/ci.yml"
        JOB="security"
        ;;
    *)
        print_error "Unknown workflow: $WORKFLOW"
        print_info "Run '$0 --help' for available options"
        exit 1
        ;;
esac

# Check if workflow file exists
if [ ! -f "$WORKFLOW_FILE" ]; then
    print_error "Workflow file not found: $WORKFLOW_FILE"
    exit 1
fi

print_info "Running workflow: $WORKFLOW_FILE"
if [ -n "$JOB" ]; then
    print_info "Specific job: $JOB"
fi

# Create temporary directory for artifacts
mkdir -p /tmp/act-artifacts

# Create network for services
docker network create act-network 2>/dev/null || true

print_info "Starting local GitHub Actions execution..."
echo ""

# Build act command using array to properly handle arguments
ACT_ARGS=()

# Add job if specified
if [ -n "$JOB" ]; then
    ACT_ARGS+=("-j" "$JOB")
fi

# Add workflow file
ACT_ARGS+=("-W" "$WORKFLOW_FILE")

# Add platform mappings for better compatibility
ACT_ARGS+=("-P" "ubuntu-latest=catthehacker/ubuntu:act-latest")
ACT_ARGS+=("-P" "ubuntu-22.04=catthehacker/ubuntu:act-22.04")
ACT_ARGS+=("-P" "ubuntu-20.04=catthehacker/ubuntu:act-20.04")

# Use linux/amd64 architecture for M-series compatibility
ACT_ARGS+=("--container-architecture" "linux/amd64")

# Add secrets for local testing
ACT_ARGS+=("-s" "GITHUB_TOKEN=$GITHUB_TOKEN_FOR_ACT")

# Artifact server
ACT_ARGS+=("--artifact-server-path" "/tmp/act-artifacts")

# Bind mount the workspace
ACT_ARGS+=("--bind")

# For Go and Node.js caching, create cache directories
mkdir -p ~/.cache/act/{go,npm}

# Add verbose flag if requested
if [ "$ACT_VERBOSE" = "1" ]; then
    ACT_ARGS+=("--verbose")
fi

print_info "Executing: act ${ACT_ARGS[*]}"
echo ""

# Run act with error handling using the array
if act "${ACT_ARGS[@]}"; then
    echo ""
    print_status "GitHub Actions completed successfully!"

    # Show artifacts if any were created
    if [ -d "/tmp/act-artifacts" ] && [ "$(ls -A /tmp/act-artifacts)" ]; then
        echo ""
        print_info "Artifacts created:"
        ls -la /tmp/act-artifacts/
    fi
else
    EXIT_CODE=$?
    echo ""
    if [ $EXIT_CODE -eq 1 ]; then
        print_warning "GitHub Actions failed - this may be due to linting issues or test failures"
        print_info "Check the output above to see specific issues found"
    else
        print_error "GitHub Actions failed with exit code: $EXIT_CODE"
    fi

    echo ""
    print_info "Common issues and solutions:"
    echo "1. GitHub authentication errors: Set GITHUB_TOKEN environment variable"
    echo "   - Get a token from https://github.com/settings/tokens"
    echo "   - Or run: gh auth login"
    echo "2. Linting failures: Fix code quality issues shown in the output"
    echo "3. TLS certificate errors: Normal in local environment, can be ignored"
    echo "4. Cache warnings: Expected for first run, will improve on subsequent runs"
    echo "5. Network timeouts: Try running again or check Docker/internet connection"
    echo ""
    print_info "For more verbose output, set ACT_VERBOSE=1"

    exit $EXIT_CODE
fi

# Cleanup
print_info "Cleaning up..."
docker network rm act-network 2>/dev/null || true

echo ""
print_status "Local GitHub Actions testing complete!"
print_info "Check the output above for any issues or warnings"
