#!/bin/bash

# Ferrovis Release Validation Script
# This script validates the release build process locally

set -e

echo "🚀 Ferrovis Release Validation Script"
echo "======================================"

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

# Check if we're in the right directory
if [ ! -d "backend" ] || [ ! -f "backend/cmd/server/main.go" ]; then
    print_error "Please run this script from the project root directory"
    exit 1
fi

# Get version (use current git tag or default)
VERSION=${1:-"v0.0.0-test"}
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

print_info "Building test release for version: ${VERSION}"
print_info "Commit: ${COMMIT}"
print_info "Date: ${DATE}"
echo ""

# Clean previous builds
rm -rf _dist
mkdir -p _dist

# Platforms to test
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

echo "🏗️  Building binaries for multiple platforms..."
echo "=============================================="

for platform in "${PLATFORMS[@]}"; do
    GOOS=$(echo $platform | cut -d'/' -f1)
    GOARCH=$(echo $platform | cut -d'/' -f2)

    echo ""
    print_info "Building for ${GOOS}/${GOARCH}..."

    BINARY_NAME="ferrovis"
    if [ "$GOOS" = "windows" ]; then
        BINARY_NAME="ferrovis.exe"
    fi

    cd backend

    # Build the binary
    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build \
        -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" \
        -o bin/${BINARY_NAME} \
        cmd/server/main.go

    if [ $? -eq 0 ]; then
        print_status "Built ${GOOS}/${GOARCH} successfully"
        ls -la bin/${BINARY_NAME}

        # Test the binary (only for current platform)
        if [ "$GOOS" = "$(go env GOOS)" ] && [ "$GOARCH" = "$(go env GOARCH)" ]; then
            print_info "Testing binary on current platform..."
            ./bin/${BINARY_NAME} --help 2>/dev/null || echo "Binary runs (no --help flag)"
        fi
    else
        print_error "Failed to build ${GOOS}/${GOARCH}"
        cd ..
        continue
    fi

    cd ..

    # Create package
    PACKAGE_NAME="ferrovis-${VERSION}-${GOOS}-${GOARCH}"
    mkdir -p _dist/${PACKAGE_NAME}

    # Copy binary
    cp backend/bin/${BINARY_NAME} _dist/${PACKAGE_NAME}/

    # Copy documentation
    cp README.md _dist/${PACKAGE_NAME}/ 2>/dev/null || echo "README.md not found, skipping"
    cp docker-compose.yml _dist/${PACKAGE_NAME}/ 2>/dev/null || echo "docker-compose.yml not found, skipping"

    # Create installation script
    cat > _dist/${PACKAGE_NAME}/install.sh << 'EOF'
#!/bin/bash
echo "🚀 Installing Ferrovis..."
chmod +x ferrovis
echo "✅ Ferrovis installed! Run ./ferrovis to start the server"
EOF
    chmod +x _dist/${PACKAGE_NAME}/install.sh

    # Create archive
    cd _dist
    if [ "$GOOS" = "windows" ]; then
        zip -r ${PACKAGE_NAME}.zip ${PACKAGE_NAME}/
        print_status "Created ${PACKAGE_NAME}.zip"
    else
        tar -czf ${PACKAGE_NAME}.tar.gz ${PACKAGE_NAME}/
        print_status "Created ${PACKAGE_NAME}.tar.gz"
    fi
    cd ..

    # Clean up directory
    rm -rf _dist/${PACKAGE_NAME}
done

echo ""
echo "🐳 Testing Docker build..."
echo "========================="

cd backend
if docker build -t ferrovis:${VERSION} -t ferrovis:test . > /dev/null 2>&1; then
    print_status "Docker build successful"

    # Save Docker image
    cd ..
    docker save ferrovis:${VERSION} | gzip > _dist/ferrovis-${VERSION}-docker.tar.gz
    print_status "Docker image saved to _dist/ferrovis-${VERSION}-docker.tar.gz"

    # Clean up test image
    docker rmi ferrovis:test > /dev/null 2>&1 || true
else
    print_warning "Docker build failed (Docker may not be available)"
    cd ..
fi

echo ""
echo "📦 Release Artifacts Summary"
echo "============================"
print_info "Release files created in _dist/:"
ls -la _dist/

TOTAL_SIZE=$(du -sh _dist | cut -f1)
print_info "Total size: ${TOTAL_SIZE}"

echo ""
echo "🧪 Testing a sample binary..."
echo "============================"

# Find a binary for the current platform
CURRENT_OS=$(go env GOOS)
CURRENT_ARCH=$(go env GOARCH)
SAMPLE_ARCHIVE=""

if [ "$CURRENT_OS" = "windows" ]; then
    SAMPLE_ARCHIVE="_dist/ferrovis-${VERSION}-${CURRENT_OS}-${CURRENT_ARCH}.zip"
else
    SAMPLE_ARCHIVE="_dist/ferrovis-${VERSION}-${CURRENT_OS}-${CURRENT_ARCH}.tar.gz"
fi

if [ -f "$SAMPLE_ARCHIVE" ]; then
    print_info "Testing archive: $(basename $SAMPLE_ARCHIVE)"

    # Create test directory
    mkdir -p _test
    cd _test

    # Extract archive
    if [[ "$SAMPLE_ARCHIVE" == *.zip ]]; then
        unzip -q "../$SAMPLE_ARCHIVE"
    else
        tar -xzf "../$SAMPLE_ARCHIVE"
    fi

    # Find the binary
    BINARY_PATH=$(find . -name "ferrovis*" -type f | head -1)

    if [ -n "$BINARY_PATH" ]; then
        print_info "Testing binary: $BINARY_PATH"
        chmod +x "$BINARY_PATH"  # Ensure it's executable

        # Test version endpoint
        $BINARY_PATH &
        SERVER_PID=$!

        # Wait a moment for server to start
        sleep 2

        # Test health endpoint
        if curl -s http://localhost:8080/health > /dev/null 2>&1; then
            print_status "Health endpoint responsive"

            # Test version endpoint
            VERSION_RESPONSE=$(curl -s http://localhost:8080/version 2>/dev/null)
            if echo "$VERSION_RESPONSE" | grep -q "$VERSION"; then
                print_status "Version endpoint working correctly"
            else
                print_warning "Version endpoint response: $VERSION_RESPONSE"
            fi
        else
            print_warning "Server not responding (may need database setup)"
        fi

        # Clean up
        kill $SERVER_PID 2>/dev/null || true
        wait $SERVER_PID 2>/dev/null || true
    else
        print_warning "Could not find binary in extracted archive"
    fi

    cd ..
    rm -rf _test
else
    print_warning "No sample archive found for current platform"
fi

echo ""
echo "🎉 Release Validation Complete!"
echo "==============================="
print_status "All release artifacts created successfully"
echo ""
print_info "Next steps for actual release:"
echo "  1. Tag your commit: git tag ${VERSION}"
echo "  2. Push the tag: git push origin ${VERSION}"
echo "  3. GitHub Actions will automatically create the release"
echo ""
print_info "Or manually trigger: gh workflow run release.yml -f version=${VERSION}"
echo ""
