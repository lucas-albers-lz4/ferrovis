# Local GitHub Actions Testing with Act

This document explains how to test GitHub Actions workflows locally using [act](https://github.com/nektos/act).

## Prerequisites

1. **Install act**:
   ```bash
   brew install act
   ```

2. **Docker Desktop** must be running

3. **Project setup**: Ensure you're in the Ferrovis project root

## Quick Start

The easiest way to test workflows locally is using our Makefile targets:

```bash
# Run full CI pipeline
make act-test

# Run only linting
make act-lint

# Run only build jobs
make act-build

# Run only security scan
make act-security

# Run release pipeline (simulated)
make act-release

# Show help for act commands
make act-help
```

## Direct Script Usage

You can also use the testing script directly:

```bash
# Show all available options
./scripts/test-github-actions.sh --help

# Run specific workflows
./scripts/test-github-actions.sh ci
./scripts/test-github-actions.sh release
./scripts/test-github-actions.sh lint
```

## Common Issues and Solutions

### 1. Architecture Warnings (Apple Silicon)

**Issue**: `You are using Apple M-series chip and you have not specified container architecture`

**Solution**: The `.actrc` file automatically configures `--container-architecture linux/amd64`

### 2. Go Setup Failures

**Issue**: `/bin/sh: 1: version: not found`

**Solution**:
- Ensure `backend/go.mod` and `backend/go.sum` exist
- Our script uses better Docker images: `catthehacker/ubuntu:act-latest`

### 3. TLS Certificate Errors

**Issue**: `tls: failed to verify certificate: x509: certificate signed by unknown authority`

**Solution**:
- Normal for local testing environments
- Our configuration disables TLS verification for golangci-lint
- Set `GODEBUG=x509ignoreCN=0` and `GOSUMDB=off`

### 4. Cache Warnings

**Issue**: `Dependencies file is not found` or cache-related warnings

**Solution**:
- Expected on first run
- Cache will improve on subsequent runs
- Our script creates proper cache directories

### 5. Network Connectivity

**Issue**: Timeouts or network errors

**Solution**:
- Check Docker Desktop is running
- Ensure stable internet connection
- Use `--container-architecture linux/amd64` flag

## Configuration Files

### `.actrc`
Global configuration for act with platform mappings and architecture settings.

### `scripts/test-github-actions.sh`
Main testing script with error handling and environment setup.

### `.golangci.yml`
Simplified linter configuration that works with local testing.

## Environment Variables

The testing setup automatically configures:

```bash
GITHUB_REPOSITORY="lucas-albers-lz4/ferrovis"
GITHUB_REF="refs/heads/main"
GITHUB_SHA="<current-git-sha>"
GITHUB_ACTOR="local-tester"
CI=true
GITHUB_ACTIONS=true
GOCACHE="/root/.cache/go-build"
GOMODCACHE="/root/go/pkg/mod"
GODEBUG="x509ignoreCN=0"
GOSUMDB="off"
```

## Workflow-Specific Notes

### CI Workflow (`.github/workflows/ci.yml`)
- **lint**: Go and Node.js linting with TLS workarounds
- **test**: Backend tests with PostgreSQL service
- **build**: Binary and Docker image builds
- **security**: Security scanning with gosec and npm audit

### Release Workflow (`.github/workflows/release.yml`)
- **lint-and-test**: Combined linting and testing
- **build-backend**: Multi-platform binary builds
- **build-docker**: Docker image creation
- **release**: Artifact collection and release notes

## Troubleshooting

### View Detailed Logs
```bash
# Add verbose flag to any act command
act --verbose -W .github/workflows/ci.yml

# Or modify the script to add more debugging
./scripts/test-github-actions.sh ci --verbose
```

### Clean Up Docker Resources
```bash
# Clean up act containers and networks
docker container prune -f
docker network prune -f
docker volume prune -f

# Remove act artifacts
rm -rf /tmp/act-artifacts
```

### Test Individual Jobs
```bash
# Test only the lint job
act -j lint -W .github/workflows/ci.yml

# Test only the build job
act -j build -W .github/workflows/ci.yml
```

## Performance Tips

1. **Use local caches**: Act automatically binds local directories for better performance
2. **Incremental testing**: Test individual jobs instead of full workflows during development
3. **Docker layer caching**: Docker will cache layers between runs for faster subsequent executions

## Integration with Development Workflow

1. **Before pushing changes**:
   ```bash
   make act-lint  # Quick linting check
   ```

2. **Before creating PR**:
   ```bash
   make act-test  # Full CI pipeline
   ```

3. **Before tagging release**:
   ```bash
   make act-release  # Test release process
   ```

This ensures your changes will pass CI/CD before pushing to GitHub, saving time and avoiding failed builds.
