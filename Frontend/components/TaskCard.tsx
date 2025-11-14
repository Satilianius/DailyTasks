import {Pressable, StyleSheet} from "react-native";
import {Text, View} from '@/components/Themed';
import {isBooleanTask, isDurationTask, isNumberTask, isTimeTask, TaskProgress} from "@/models/AllTasksProgress";
import {useColorScheme} from "@/components/useColorScheme";
import Colors from "@/constants/Colors";

interface TaskCardProps {
  task: TaskProgress;
  onPress: () => void;
}

export default function TaskCard({task, onPress}: TaskCardProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'dark'];

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

  const isCompleted = isBooleanTask(task) && task.progress;

  return (
    <Pressable
      onPress={onPress}
      style={[styles.card, { backgroundColor: theme.componentBackground }, isCompleted && styles.cardCompleted]}>

      <Text style={[styles.cardTitle, { color: theme.text }]}>
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
