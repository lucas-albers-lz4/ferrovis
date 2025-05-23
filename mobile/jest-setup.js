// jest-setup.js - Minimal Jest configuration for React Native with Expo
import '@testing-library/react-native/extend-expect';

// Mock Expo modules that aren't available in test environment
jest.mock('expo-status-bar', () => ({
  StatusBar: 'StatusBar',
}));

// Mock global fetch for API tests
global.fetch = jest.fn();
