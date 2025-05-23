#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');

const processes = [];

function runCommand(name, command, cwd = process.cwd(), color = '36') {
  console.log(`\x1b[${color}m[${name}]\x1b[0m Starting: ${command}`);

  const child = spawn('sh', ['-c', command], {
    cwd,
    stdio: 'pipe'
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
  console.log('⚡ Starting Ferrovis Development Environment');
  console.log('==========================================\n');

  // Start Docker services
  console.log('🐳 Starting Docker services...');
  await new Promise((resolve) => {
    const docker = spawn('make', ['docker-up'], { stdio: 'inherit' });
    docker.on('close', resolve);
  });

  // Wait a bit for services to be ready
  await new Promise(resolve => setTimeout(resolve, 2000));

  console.log('\n🚀 Starting development servers...\n');

  // Start backend
  runCommand('Backend', 'go run cmd/server/main.go', 'backend', '33');

  // Start mobile (in background, will wait for user to choose platform)
  runCommand('Mobile', 'npm start', 'mobile', '35');

  console.log('\n✅ Development environment is running!');
  console.log('=====================================');
  console.log('🔧 Backend API: http://localhost:8080');
  console.log('📱 Mobile app: Follow Expo CLI prompts above');
  console.log('🗄️  Database: PostgreSQL on localhost:5432');
  console.log('\nPress Ctrl+C to stop all services\n');

  // Keep the process alive
  setInterval(() => {}, 1000);
}

main().catch(console.error);
