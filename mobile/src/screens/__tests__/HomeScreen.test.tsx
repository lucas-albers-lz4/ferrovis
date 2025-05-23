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
    expect(screen.getByText('Ferrovis')).toBeOnTheScreen();
    expect(
      screen.getByText('Your Fitness Accountability Partner with Weasel Mode™')
    ).toBeOnTheScreen();
    expect(
      screen.getByText('Built with Expo + React Native + Go + PostgreSQL')
    ).toBeOnTheScreen();
  });

  it('displays now available and coming soon features', () => {
    render(<HomeScreen />);

    // Check "Now Available" section
    expect(screen.getByText('Now Available:')).toBeOnTheScreen();
    expect(screen.getByText('• Backend API with workout programs')).toBeOnTheScreen();
    expect(screen.getByText('• Starting Strength & StrongLifts 5x5 programs')).toBeOnTheScreen();
    expect(screen.getByText('• Next workout calculation')).toBeOnTheScreen();
    expect(screen.getByText('• Exercise instructions & guidance')).toBeOnTheScreen();

    // Check "Coming Soon" section
    expect(screen.getByText('Coming Soon:')).toBeOnTheScreen();
    expect(screen.getByText('• Complete workout tracking')).toBeOnTheScreen();
    expect(screen.getByText('• Buddy accountability system')).toBeOnTheScreen();
    expect(screen.getByText('• Progress tracking & streaks')).toBeOnTheScreen();
    expect(screen.getByText('• Full Weasel Mode™ psychological features')).toBeOnTheScreen();
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

  it('has demo navigation buttons', () => {
    render(<HomeScreen />);

    const programsButton = screen.getByText('View Programs');
    const workoutButton = screen.getByText('Next Workout');

    expect(programsButton).toBeOnTheScreen();
    expect(workoutButton).toBeOnTheScreen();

    // Test button presses (should not throw)
    fireEvent.press(programsButton);
    fireEvent.press(workoutButton);
  });
});
