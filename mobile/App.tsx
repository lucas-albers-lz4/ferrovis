import { StatusBar } from 'expo-status-bar';
import { useState } from 'react';
import { SafeAreaView, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import HomeScreen from './src/screens/HomeScreen';
import ProgramsScreen from './src/screens/ProgramsScreen';
import WorkoutScreen from './src/screens/WorkoutScreen';

type Screen = 'home' | 'programs' | 'workout';

export default function App() {
  const [currentScreen, setCurrentScreen] = useState<Screen>('home');

  const renderScreen = () => {
    switch (currentScreen) {
      case 'home':
        return <HomeScreen />;
      case 'programs':
        return <ProgramsScreen />;
      case 'workout':
        return <WorkoutScreen />;
      default:
        return <HomeScreen />;
    }
  };

  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.tabBar}>
        <TouchableOpacity
          style={[styles.tab, currentScreen === 'home' && styles.activeTab]}
          onPress={() => setCurrentScreen('home')}
        >
          <Text style={[styles.tabText, currentScreen === 'home' && styles.activeTabText]}>
            Home
          </Text>
        </TouchableOpacity>

        <TouchableOpacity
          style={[styles.tab, currentScreen === 'programs' && styles.activeTab]}
          onPress={() => setCurrentScreen('programs')}
        >
          <Text style={[styles.tabText, currentScreen === 'programs' && styles.activeTabText]}>
            Programs
          </Text>
        </TouchableOpacity>

        <TouchableOpacity
          style={[styles.tab, currentScreen === 'workout' && styles.activeTab]}
          onPress={() => setCurrentScreen('workout')}
        >
          <Text style={[styles.tabText, currentScreen === 'workout' && styles.activeTabText]}>
            Workout
          </Text>
        </TouchableOpacity>
      </View>

      <View style={styles.screenContainer}>
        {renderScreen()}
      </View>

      <StatusBar style='auto' />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  tabBar: {
    flexDirection: 'row',
    backgroundColor: '#f8f8f8',
    borderBottomWidth: 1,
    borderBottomColor: '#e0e0e0',
    paddingTop: 8,
  },
  tab: {
    flex: 1,
    paddingVertical: 12,
    alignItems: 'center',
  },
  activeTab: {
    borderBottomWidth: 2,
    borderBottomColor: '#007AFF',
  },
  tabText: {
    fontSize: 16,
    color: '#666',
  },
  activeTabText: {
    color: '#007AFF',
    fontWeight: '600',
  },
  screenContainer: {
    flex: 1,
  },
});
