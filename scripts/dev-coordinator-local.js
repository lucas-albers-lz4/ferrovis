#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');

const processes = [];

function runCommand(name, command, cwd = process.cwd(), color = '36', env = {}) {
  console.log(`\x1b[${color}m[${name}]\x1b[0m Starting: ${command}`);

  const child = spawn('sh', ['-c', command], {
    cwd,
    stdio: 'pipe',
    env: { ...process.env, ...env }
  });

  child.stdout.on('data', (data) => {
    const lines = data.toString().split('\n').filter(line => line.trim());
    lines.forEach(line => {
      console.log(`\x1b[${color}m[${name}]\x1b[0m ${line}`);
    });
  });

  child.stderr.on('data', (data) => {
    const lines = data.toString().split('\n').filter(line => line.trim());
    lines.forEach(line => {
      console.log(`\x1b[${color}m[${name}]\x1b[0m \x1b[31m${line}\x1b[0m`);
    });
  });

  child.on('close', (code) => {
    console.log(`\x1b[${color}m[${name}]\x1b[0m Process exited with code ${code}`);
  });

  processes.push({ name, child });
  return child;
}

function cleanup() {
  console.log('\n🧹 Shutting down development processes...');
  processes.forEach(({ name, child }) => {
    console.log(`⏹️  Stopping ${name}...`);
    child.kill('SIGTERM');
  });
  process.exit(0);
}

// Handle cleanup on exit
process.on('SIGINT', cleanup);
process.on('SIGTERM', cleanup);

// Main development setup
async function main() {
  console.log('⚡ Starting Ferrovis Development Environment (Local Backend)');
  console.log('========================================================\n');

  // Start only database with Docker
  console.log('🐳 Starting PostgreSQL database with Docker...');
  await new Promise((resolve) => {
    const docker = spawn('docker', ['compose', 'up', '-d', 'database', '--wait'], { stdio: 'inherit' });
    docker.on('close', resolve);
  });

  // Wait a bit for database to be ready
  await new Promise(resolve => setTimeout(resolve, 3000));

  console.log('\n🚀 Starting development servers...\n');

  // Environment variables for local backend
  const backendEnv = {
    DB_HOST: 'localhost',
    DB_PORT: '5432',
    DB_NAME: 'ferrovis_dev',
    DB_USER: 'postgres',
    DB_PASSWORD: 'ferrovis123',
    DB_SSLMODE: 'disable',
    PORT: '8080',
    ENVIRONMENT: 'development',
    JWT_SECRET: 'dev_jwt_secret_change_in_production',
    GIN_MODE: 'debug'
  };

  // Start backend locally (not in Docker)
  runCommand('Backend', 'go run cmd/server/main.go', 'backend', '33', backendEnv);

  // Start mobile (in background, will wait for user to choose platform)
  runCommand('Mobile', 'npm start', 'mobile', '35');

  console.log('\n✅ Development environment is running!');
  console.log('=====================================');
  console.log('🔧 Backend API: http://localhost:8080 (local Go process)');
  console.log('📱 Mobile app: Follow Expo CLI prompts above');
  console.log('🗄️  Database: PostgreSQL on localhost:5432 (Docker)');
  console.log('\nPress Ctrl+C to stop all services\n');

  // Keep the process alive
  setInterval(() => {}, 1000);
}

main().catch(console.error);
