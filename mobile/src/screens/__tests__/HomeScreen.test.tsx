// HomeScreen.test.tsx - Test for HomeScreen component
import { fireEvent, render, screen } from '@testing-library/react-native';
import HomeScreen from '../HomeScreen';

// Mock the API service
jest.mock('../../services/api', () => ({
  checkApiHealth: jest.fn(),
}));

describe('HomeScreen', () => {
  it('renders correctly', () => {
    render(<HomeScreen />);

    // Check if main elements are present
    expect(screen.getByText('LiftBuddy')).toBeOnTheScreen();
    expect(
      screen.getByText('Your Fitness Accountability Partner')
    ).toBeOnTheScreen();
    expect(
      screen.getByText('Built with Expo + React Native + Go')
    ).toBeOnTheScreen();
  });

  it('displays coming soon features', () => {
    render(<HomeScreen />);

    expect(screen.getByText('Coming Soon:')).toBeOnTheScreen();
    expect(screen.getByText('• Simple workout tracking')).toBeOnTheScreen();
    expect(screen.getByText('• Buddy accountability system')).toBeOnTheScreen();
    expect(screen.getByText('• Progress tracking & streaks')).toBeOnTheScreen();
    expect(screen.getByText('• Beginner-friendly programs')).toBeOnTheScreen();
  });

  it('has functional refresh button', () => {
    render(<HomeScreen />);

    const refreshButton = screen.getByText('Refresh');
    expect(refreshButton).toBeOnTheScreen();

    // Test button press (should not throw)
    fireEvent.press(refreshButton);
  });

  it('has functional welcome button', () => {
    render(<HomeScreen />);

    const welcomeButton = screen.getByText('Welcome Message');
    expect(welcomeButton).toBeOnTheScreen();

    // Test button press (should not throw)
    fireEvent.press(welcomeButton);
  });
});
