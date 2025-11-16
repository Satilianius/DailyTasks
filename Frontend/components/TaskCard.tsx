import {Pressable, StyleSheet} from "react-native";
import {Text, View} from '@/components/Themed';
import {isBooleanTask, isDurationTask, isNumberTask, isTimeTask, TaskProgress} from "@/models/AllTasksProgress";
import {useColorScheme} from "@/components/useColorScheme";
import Colors from "@/constants/Colors";
import BooleanProgressEditor from '@/components/BooleanProgressEditor';
import {useContext} from 'react';
import {TasksProgressContext} from '@/context/TasksProgressContext';

interface TaskCardProps {
  task: TaskProgress;
  date: Date;
  userId: string;
  onPress?: () => void; // Reserved for opening task details in the future
}

export default function TaskCard({task, date, userId, onPress}: TaskCardProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'dark'];
  const {updateTaskProgress} = useContext(TasksProgressContext)!;

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
      style={[
        styles.card,
        isCompleted
          ? {backgroundColor: theme.success}
          : {backgroundColor: theme.componentBackground}]}>

      <Text style={[styles.cardTitle, {color: theme.text}]}>
        {task.taskName}
      </Text>

      <View style={styles.progressContainer}>
        {isBooleanTask(task)
          ? (
            <BooleanProgressEditor
              value={task.progress}
              onChange={(next) => {
                // Update via shared context (optimistic cache + mock backend)
                void updateTaskProgress(userId, date, task.taskId, next);
              }}
            />)
          : (
            <Text style={styles.progressValue}>
              {getProgressDisplay()}
            </Text>)}
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
