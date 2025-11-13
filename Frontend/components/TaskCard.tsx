
import {Pressable, StyleSheet} from "react-native";
import {Text, View} from '@/components/Themed';
import {TaskProgress, isBooleanTask, isNumberTask, isTimeTask, isDurationTask} from "@/models/AllTasksProgress";

interface TaskCardProps {
    task: TaskProgress;
    onPress: () => void;
}

export default function TaskCard({task, onPress}: TaskCardProps) {
    const getProgressDisplay = (): string => {
        if (isBooleanTask(task)) {
            return task.progress ? '✓' : '✗';
        } else if (isNumberTask(task)) {
            return `${task.progress}`;
        } else if (isTimeTask(task)) {
            return task.progress.substring(0, 5); // HH:mm
        } else if (isDurationTask(task)) {
            return task.progress.substring(0, 8); // HH:mm:ss
        }
        return '';
    };

    const getTaskTypeLabel = (): string => {
        switch (task.type) {
            case 'boolean':
                return task.progress ? 'Done' : 'Not Done';
            case 'number':
                return 'Count';
            case 'time':
                return 'Time';
            case 'duration':
                return 'Duration';
            default:
                return '';
        }
    };

    const isCompleted = isBooleanTask(task) && task.progress;

    return (
        <Pressable
            onPress={onPress}
            style={[styles.card, isCompleted && styles.cardCompleted]}>
            <Text style={styles.cardTitle}>
                {task.taskName}
            </Text>
            <View style={styles.progressContainer}>
                <Text style={styles.progressValue}>
                    {getProgressDisplay()}
                </Text>
            </View>
        </Pressable>
    )
}

const styles = StyleSheet.create({
    card: {
        backgroundColor: '#6d6d6d',
        borderRadius: 8,
        padding: 16,
        margin: 8,
        minHeight: 180,
        minWidth: 180,
        justifyContent: 'space-between'
    },
    cardCompleted: {
        backgroundColor: '#4a7c59',
    },
    cardTitle: {
        fontSize: 20,
        fontWeight: 'bold',
        textAlign: 'center',
    },
    progressContainer: {
        gap: 4,
        backgroundColor: 'transparent',
    },
    progressValue: {
        fontSize: 20,
        fontWeight: 'bold',
        textAlign: 'center',
        textAlignVertical: 'center',
    },
})
