import { useEffect, useState } from 'react';
import {
    ActivityIndicator,
    Alert,
    ScrollView,
    StyleSheet,
    Text,
    TouchableOpacity,
    View,
} from 'react-native';
import { apiClient, NextWorkout, WorkoutExercise } from '../services/api';

interface WorkoutScreenProps {
    route?: {
        params: {
            programId: number;
            programName?: string;
        };
    };
    navigation?: any;
}

export default function WorkoutScreen({ route, navigation }: WorkoutScreenProps) {
    const [workout, setWorkout] = useState<NextWorkout | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    // For testing, default to program ID 1 if no route params
    const programId = route?.params?.programId || 1;
    const programName = route?.params?.programName || 'Starting Strength';

    useEffect(() => {
        loadNextWorkout();
    }, [programId]);

    const loadNextWorkout = async () => {
        setLoading(true);
        setError(null);

        try {
            const response = await apiClient.getNextWorkout(programId);

            if (response.error) {
                setError(response.error);
            } else if (response.data) {
                setWorkout(response.data.next_workout);
            }
        } catch (err) {
            setError('Failed to load workout');
        } finally {
            setLoading(false);
        }
    };

    const handleStartWorkout = () => {
        Alert.alert(
            'Start Workout',
            'This will start your workout session. Ready to begin?',
            [
                { text: 'Not Yet', style: 'cancel' },
                {
                    text: 'Let\'s Go!',
                    onPress: () => {
                        Alert.alert('Coming Soon', 'Workout tracking will be implemented in the next phase!');
                    }
                },
            ]
        );
    };

    const renderExercise = (exercise: WorkoutExercise, index: number) => (
        <View key={index} style={styles.exerciseCard}>
            <View style={styles.exerciseHeader}>
                <Text style={styles.exerciseName}>{exercise.name}</Text>
                <Text style={styles.exerciseDetails}>
                    {exercise.sets} × {exercise.reps} reps
                </Text>
            </View>

            <Text style={styles.weightText}>Weight: {exercise.weight}</Text>
            <Text style={styles.instructionsText}>{exercise.instructions}</Text>
        </View>
    );

    if (loading) {
        return (
            <View style={styles.centerContainer}>
                <ActivityIndicator size="large" color="#007AFF" />
                <Text style={styles.loadingText}>Loading your next workout...</Text>
            </View>
        );
    }

    if (error) {
        return (
            <View style={styles.centerContainer}>
                <Text style={styles.errorText}>Error: {error}</Text>
                <TouchableOpacity style={styles.retryButton} onPress={loadNextWorkout}>
                    <Text style={styles.retryButtonText}>Try Again</Text>
                </TouchableOpacity>
            </View>
        );
    }

    if (!workout) {
        return (
            <View style={styles.centerContainer}>
                <Text style={styles.errorText}>No workout data available</Text>
            </View>
        );
    }

    return (
        <ScrollView style={styles.container}>
            <View style={styles.header}>
                <Text style={styles.title}>{workout.program_name}</Text>
                <Text style={styles.workoutDay}>Workout {workout.workout_day}</Text>
                <Text style={styles.duration}>Estimated Duration: {workout.estimated_duration}</Text>
            </View>

            <View style={styles.exercisesSection}>
                <Text style={styles.sectionTitle}>Today's Exercises</Text>
                {workout.exercises.map((exercise, index) => renderExercise(exercise, index))}
            </View>

            <View style={styles.tipsSection}>
                <Text style={styles.sectionTitle}>Tips</Text>
                <View style={styles.tipCard}>
                    <Text style={styles.tipText}>• {workout.rest_between_sets}</Text>
                    <Text style={styles.tipText}>• Focus on proper form over heavy weight</Text>
                    <Text style={styles.tipText}>• Warm up with lighter weights before your working sets</Text>
                    <Text style={styles.tipText}>• Track your weights to ensure progressive overload</Text>
                </View>
            </View>

            <TouchableOpacity style={styles.startButton} onPress={handleStartWorkout}>
                <Text style={styles.startButtonText}>Start Workout</Text>
            </TouchableOpacity>

            <View style={styles.bottomPadding} />
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#f5f5f5',
    },
    centerContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#f5f5f5',
        padding: 20,
    },
    header: {
        backgroundColor: 'white',
        padding: 20,
        borderBottomWidth: 1,
        borderBottomColor: '#e0e0e0',
    },
    title: {
        fontSize: 28,
        fontWeight: 'bold',
        color: '#333',
        marginBottom: 8,
    },
    workoutDay: {
        fontSize: 18,
        color: '#007AFF',
        fontWeight: '600',
        marginBottom: 4,
    },
    duration: {
        fontSize: 16,
        color: '#666',
    },
    loadingText: {
        marginTop: 12,
        fontSize: 16,
        color: '#666',
        textAlign: 'center',
    },
    errorText: {
        fontSize: 16,
        color: '#F44336',
        textAlign: 'center',
        marginBottom: 20,
    },
    retryButton: {
        backgroundColor: '#007AFF',
        paddingHorizontal: 20,
        paddingVertical: 10,
        borderRadius: 8,
    },
    retryButtonText: {
        color: 'white',
        fontSize: 16,
        fontWeight: '600',
    },
    exercisesSection: {
        margin: 16,
    },
    sectionTitle: {
        fontSize: 20,
        fontWeight: 'bold',
        color: '#333',
        marginBottom: 16,
    },
    exerciseCard: {
        backgroundColor: 'white',
        padding: 16,
        borderRadius: 12,
        marginBottom: 12,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
        elevation: 3,
    },
    exerciseHeader: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginBottom: 8,
    },
    exerciseName: {
        fontSize: 18,
        fontWeight: 'bold',
        color: '#333',
        flex: 1,
    },
    exerciseDetails: {
        fontSize: 16,
        color: '#007AFF',
        fontWeight: '600',
    },
    weightText: {
        fontSize: 16,
        color: '#666',
        marginBottom: 8,
    },
    instructionsText: {
        fontSize: 14,
        color: '#888',
        lineHeight: 20,
        fontStyle: 'italic',
    },
    tipsSection: {
        margin: 16,
    },
    tipCard: {
        backgroundColor: 'white',
        padding: 16,
        borderRadius: 12,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
        elevation: 3,
    },
    tipText: {
        fontSize: 14,
        color: '#666',
        lineHeight: 22,
        marginBottom: 4,
    },
    startButton: {
        backgroundColor: '#4CAF50',
        margin: 16,
        padding: 16,
        borderRadius: 12,
        alignItems: 'center',
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
        elevation: 3,
    },
    startButtonText: {
        color: 'white',
        fontSize: 18,
        fontWeight: 'bold',
    },
    bottomPadding: {
        height: 20,
    },
});
