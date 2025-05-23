import React, { useState, useEffect } from 'react';
import { View, Text, StyleSheet, TouchableOpacity, Alert } from 'react-native';
import { checkApiHealth } from '../services/api';

export default function HomeScreen() {
  const [apiStatus, setApiStatus] = useState<'checking' | 'connected' | 'disconnected'>('checking');

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
      'Welcome to LiftBuddy! 🏋️',
      'Your fitness accountability companion is ready to help you build lasting lifting habits.',
      [{ text: 'Let\'s Go!', style: 'default' }]
    );
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>LiftBuddy</Text>
      <Text style={styles.subtitle}>Your Fitness Accountability Partner</Text>
      
      <View style={styles.statusContainer}>
        <Text style={[styles.statusText, { color: getStatusColor() }]}>
          {getStatusText()}
        </Text>
        <TouchableOpacity style={styles.refreshButton} onPress={checkConnection}>
          <Text style={styles.refreshButtonText}>Refresh</Text>
        </TouchableOpacity>
      </View>

      <View style={styles.featuresContainer}>
        <Text style={styles.sectionTitle}>Coming Soon:</Text>
        <Text style={styles.featureItem}>• Simple workout tracking</Text>
        <Text style={styles.featureItem}>• Buddy accountability system</Text>
        <Text style={styles.featureItem}>• Progress tracking & streaks</Text>
        <Text style={styles.featureItem}>• Beginner-friendly programs</Text>
      </View>

      <TouchableOpacity style={styles.welcomeButton} onPress={showWelcomeMessage}>
        <Text style={styles.welcomeButtonText}>Welcome Message</Text>
      </TouchableOpacity>

      <Text style={styles.footerText}>
        Built with Expo + React Native + Go
      </Text>
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
}); 