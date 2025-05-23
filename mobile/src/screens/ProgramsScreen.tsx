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
import { apiClient, Program } from '../services/api';

interface ProgramsScreenProps {
    navigation?: any; // Will add proper navigation types later
}

export default function ProgramsScreen({ navigation }: ProgramsScreenProps) {
    const [programs, setPrograms] = useState<Program[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        loadPrograms();
    }, []);

    const loadPrograms = async () => {
        setLoading(true);
        setError(null);

        try {
            const response = await apiClient.getPrograms();

            if (response.error) {
                setError(response.error);
            } else if (response.data) {
                setPrograms(response.data.programs);
            }
        } catch (err) {
            setError('Failed to load programs');
        } finally {
            setLoading(false);
        }
    };

    const handleProgramSelect = (program: Program) => {
        // For now, show an alert - later we'll navigate to workout screen
        Alert.alert(
            program.name,
            `${program.description}\n\nDifficulty: ${program.difficulty}\nDuration: ${program.duration} weeks`,
            [
                { text: 'Cancel', style: 'cancel' },
                {
                    text: 'View Next Workout',
                    onPress: () => {
                        // TODO: Navigate to workout screen
                        // navigation?.navigate('Workout', { programId: program.id });
                        Alert.alert('Coming Soon', 'Navigation to workout screen will be implemented next!');
                    }
                },
            ]
        );
    };

    const getDifficultyColor = (difficulty: string) => {
        switch (difficulty.toLowerCase()) {
            case 'beginner':
                return '#4CAF50';
            case 'intermediate':
                return '#FF9800';
            case 'advanced':
                return '#F44336';
            default:
                return '#2196F3';
        }
    };

    if (loading) {
        return (
            <View style={styles.centerContainer}>
                <ActivityIndicator size="large" color="#007AFF" />
                <Text style={styles.loadingText}>Loading programs...</Text>
            </View>
        );
    }

    if (error) {
        return (
            <View style={styles.centerContainer}>
                <Text style={styles.errorText}>Error: {error}</Text>
                <TouchableOpacity style={styles.retryButton} onPress={loadPrograms}>
                    <Text style={styles.retryButtonText}>Try Again</Text>
                </TouchableOpacity>
            </View>
        );
    }

    return (
        <ScrollView style={styles.container}>
            <View style={styles.header}>
                <Text style={styles.title}>Workout Programs</Text>
                <Text style={styles.subtitle}>Choose a program to start your fitness journey</Text>
            </View>

            {programs.map((program) => (
                <TouchableOpacity
                    key={program.id}
                    style={styles.programCard}
                    onPress={() => handleProgramSelect(program)}
                >
                    <View style={styles.programHeader}>
                        <Text style={styles.programName}>{program.name}</Text>
                        <View style={[styles.difficultyBadge, { backgroundColor: getDifficultyColor(program.difficulty) }]}>
                            <Text style={styles.difficultyText}>{program.difficulty}</Text>
                        </View>
                    </View>

                    <Text style={styles.programDescription}>{program.description}</Text>

                    <View style={styles.programDetails}>
                        <Text style={styles.detailText}>Duration: {program.duration} weeks</Text>
                    </View>
                </TouchableOpacity>
            ))}

            {programs.length === 0 && (
                <View style={styles.emptyState}>
                    <Text style={styles.emptyText}>No programs available</Text>
                    <Text style={styles.emptySubtext}>Check your connection and try again</Text>
                </View>
            )}
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
    },
    header: {
        padding: 20,
        backgroundColor: 'white',
        borderBottomWidth: 1,
        borderBottomColor: '#e0e0e0',
    },
    title: {
        fontSize: 28,
        fontWeight: 'bold',
        color: '#333',
        marginBottom: 8,
    },
    subtitle: {
        fontSize: 16,
        color: '#666',
    },
    loadingText: {
        marginTop: 12,
        fontSize: 16,
        color: '#666',
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
    programCard: {
        backgroundColor: 'white',
        margin: 16,
        padding: 20,
        borderRadius: 12,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
        elevation: 3,
    },
    programHeader: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginBottom: 12,
    },
    programName: {
        fontSize: 20,
        fontWeight: 'bold',
        color: '#333',
        flex: 1,
    },
    difficultyBadge: {
        paddingHorizontal: 12,
        paddingVertical: 4,
        borderRadius: 12,
    },
    difficultyText: {
        color: 'white',
        fontSize: 12,
        fontWeight: '600',
        textTransform: 'capitalize',
    },
    programDescription: {
        fontSize: 16,
        color: '#666',
        lineHeight: 24,
        marginBottom: 16,
    },
    programDetails: {
        borderTopWidth: 1,
        borderTopColor: '#f0f0f0',
        paddingTop: 12,
    },
    detailText: {
        fontSize: 14,
        color: '#888',
    },
    emptyState: {
        padding: 40,
        alignItems: 'center',
    },
    emptyText: {
        fontSize: 18,
        color: '#666',
        marginBottom: 8,
    },
    emptySubtext: {
        fontSize: 14,
        color: '#999',
    },
});
