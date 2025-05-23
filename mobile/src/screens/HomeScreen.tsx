import { useEffect, useState } from 'react';
import { Alert, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { checkApiHealth } from '../services/api';

export default function HomeScreen() {
  const [apiStatus, setApiStatus] = useState<
    'checking' | 'connected' | 'disconnected'
  >('checking');

  useEffect(() => {
    checkConnection();
  }, []);

  const checkConnection = async () => {
    setApiStatus('checking');
    const isHealthy = await checkApiHealth();
    setApiStatus(isHealthy ? 'connected' : 'disconnected');
  };

  const getStatusColor = () => {
    switch (apiStatus) {
      case 'connected':
        return '#4CAF50';
      case 'disconnected':
        return '#F44336';
      default:
        return '#FF9800';
    }
  };

  const getStatusText = () => {
    switch (apiStatus) {
      case 'connected':
        return 'API Connected ✅';
      case 'disconnected':
        return 'API Disconnected ❌';
      default:
        return 'Checking connection...';
    }
  };

  const showWelcomeMessage = () => {
    Alert.alert(
      'Welcome to Ferrovis! 🏋️',
      'Your fitness accountability companion with Weasel Mode™ is ready to help you build lasting lifting habits.',
      [{ text: "Let's Go!", style: 'default' }]
    );
  };

  const showProgramsDemo = () => {
    Alert.alert(
      'Programs Screen Demo',
      'This would normally navigate to the Programs screen. For now, this is a placeholder to show the upcoming feature.',
      [{ text: 'OK', style: 'default' }]
    );
  };

  const showWorkoutDemo = () => {
    Alert.alert(
      'Workout Screen Demo',
      'This would normally navigate to your next workout. For now, this is a placeholder to show the upcoming feature.',
      [{ text: 'OK', style: 'default' }]
    );
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Ferrovis</Text>
      <Text style={styles.subtitle}>Your Fitness Accountability Partner with Weasel Mode™</Text>

      <View style={styles.statusContainer}>
        <Text style={[styles.statusText, { color: getStatusColor() }]}>
          {getStatusText()}
        </Text>
        <TouchableOpacity
          style={styles.refreshButton}
          onPress={checkConnection}
        >
          <Text style={styles.refreshButtonText}>Refresh</Text>
        </TouchableOpacity>
      </View>

      <View style={styles.featuresContainer}>
        <Text style={styles.sectionTitle}>Now Available:</Text>
        <Text style={styles.featureItem}>• Backend API with workout programs</Text>
        <Text style={styles.featureItem}>• Starting Strength & StrongLifts 5x5 programs</Text>
        <Text style={styles.featureItem}>• Next workout calculation</Text>
        <Text style={styles.featureItem}>• Exercise instructions & guidance</Text>

        <Text style={[styles.sectionTitle, { marginTop: 16 }]}>Coming Soon:</Text>
        <Text style={styles.featureItem}>• Complete workout tracking</Text>
        <Text style={styles.featureItem}>• Buddy accountability system</Text>
        <Text style={styles.featureItem}>• Progress tracking & streaks</Text>
        <Text style={styles.featureItem}>• Full Weasel Mode™ psychological features</Text>
      </View>

      <View style={styles.demoButtonsContainer}>
        <TouchableOpacity
          style={styles.demoButton}
          onPress={showProgramsDemo}
        >
          <Text style={styles.demoButtonText}>View Programs</Text>
        </TouchableOpacity>

        <TouchableOpacity
          style={[styles.demoButton, styles.workoutButton]}
          onPress={showWorkoutDemo}
        >
          <Text style={styles.demoButtonText}>Next Workout</Text>
        </TouchableOpacity>
      </View>

      <TouchableOpacity
        style={styles.welcomeButton}
        onPress={showWelcomeMessage}
      >
        <Text style={styles.welcomeButtonText}>Welcome Message</Text>
      </TouchableOpacity>

      <Text style={styles.footerText}>Built with Expo + React Native + Go + PostgreSQL</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 20,
  },
  title: {
    fontSize: 36,
    fontWeight: 'bold',
    color: '#333',
    marginBottom: 8,
  },
  subtitle: {
    fontSize: 18,
    color: '#666',
    marginBottom: 40,
    textAlign: 'center',
  },
  statusContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 40,
    padding: 16,
    backgroundColor: 'white',
    borderRadius: 8,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  statusText: {
    fontSize: 16,
    fontWeight: '600',
    marginRight: 12,
  },
  refreshButton: {
    backgroundColor: '#007AFF',
    paddingHorizontal: 12,
    paddingVertical: 6,
    borderRadius: 4,
  },
  refreshButtonText: {
    color: 'white',
    fontSize: 14,
    fontWeight: '600',
  },
  featuresContainer: {
    backgroundColor: 'white',
    padding: 20,
    borderRadius: 8,
    width: '100%',
    marginBottom: 30,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  sectionTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#333',
    marginBottom: 12,
  },
  featureItem: {
    fontSize: 16,
    color: '#666',
    marginBottom: 8,
    lineHeight: 24,
  },
  welcomeButton: {
    backgroundColor: '#4CAF50',
    paddingHorizontal: 24,
    paddingVertical: 12,
    borderRadius: 8,
    marginBottom: 20,
  },
  welcomeButtonText: {
    color: 'white',
    fontSize: 16,
    fontWeight: '600',
  },
  footerText: {
    fontSize: 12,
    color: '#999',
    marginTop: 20,
  },
  demoButtonsContainer: {
    flexDirection: 'row',
    width: '100%',
    marginBottom: 20,
    gap: 12,
  },
  demoButton: {
    backgroundColor: '#007AFF',
    flex: 1,
    paddingVertical: 12,
    borderRadius: 8,
    alignItems: 'center',
  },
  workoutButton: {
    backgroundColor: '#4CAF50',
  },
  demoButtonText: {
    color: 'white',
    fontSize: 16,
    fontWeight: '600',
  },
});
