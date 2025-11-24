import React, {useContext, useState} from 'react';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import {Tabs} from 'expo-router';
import {KeyboardAvoidingView, Modal, Platform, StyleSheet, Text, TextInput, TouchableOpacity, View} from 'react-native';

import Colors from '@/constants/Colors';
import {useColorScheme} from '@/components/useColorScheme';
import {TasksProgressContext, TasksProgressProvider} from '@/context/TasksProgressContext';
import {TaskProgress, TaskType, taskTypes} from "@/models/AllTasksProgress";

export default function TabLayout() {
  return (
    <TasksProgressProvider>
      <TabsLayoutContent/>
    </TasksProgressProvider>
  );
}

function TabsLayoutContent() {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'dark'];
  const {cachedTaskProgress, setCachedTaskProgress} = useContext(TasksProgressContext)!;

  // New Task Modal State
  const [modalVisible, setModalVisible] = useState(false);
  const [newTaskName, setNewTaskName] = useState('');
  const [newTaskType, setNewTaskType] = useState<TaskType>('boolean');

  // Mock user ID
  const userId = 'user-123-abc';

  const handleAddTask = async () => {
    const trimmedTaskName = newTaskName.trim();
    if (!trimmedTaskName) return;

    // Generate a simple id for demo/mock purposes
    // TODO think what should generate uuid, backend or frontend
    const newId = `task-${Math.random().toString(36).slice(2)}-${Date.now()}`;

    // Determine default progress based on type
    const defaultProgressByType: Record<TaskType, boolean | number | string> = {
      boolean: false,
      number: 0,
      time: '00:00:00.000',
      duration: '00:00:00.000',
    };

    let newTask: TaskProgress;
    switch (newTaskType) {
      case 'boolean':
        newTask = {
          taskId: newId,
          taskName: trimmedTaskName,
          type: 'boolean',
          progress: defaultProgressByType['boolean'] as boolean,
        };
        break;
      case 'number':
        newTask = {
          taskId: newId,
          taskName: trimmedTaskName,
          type: 'number',
          progress: defaultProgressByType['number'] as number,
        };
        break;
      case 'time':
        newTask = {
          taskId: newId,
          taskName: trimmedTaskName,
          type: 'time',
          progress: defaultProgressByType['time'] as string,
        };
        break;
      case 'duration':
      default:
        newTask = {
          taskId: newId,
          taskName: trimmedTaskName,
          type: 'duration',
          progress: defaultProgressByType['duration'] as string,
        };
        break;
    }

    if (!cachedTaskProgress) {
      // TODO what if this is the first task?
      // If nothing is cached yet, just close the modal. Views will populate on next load.
      setNewTaskName('');
      setNewTaskType('boolean');
      setModalVisible(false);
      return;
    }

    // TODO extract to a separate function
    // Optimistically update the shared cache so all views reflect the new task immediately
    const prev = cachedTaskProgress;
    const updated = {
      ...cachedTaskProgress,
      TasksProgress: cachedTaskProgress.TasksProgress.map(day => ({
        ...day,
        // append the newly created task to every day currently in cache
        tasks: [...day.tasks, {...newTask}],
      })),
    };

    setCachedTaskProgress(updated);

    // In a real app, call backend to persist the new task, and rollback on failure
    try {
      console.log('Creating task (mock):', {
        userId,
        name: newTask.taskName,
        type: newTask.type,
        createdAt: new Date().toISOString(),
      });
      // await createTaskApi(userId, newTask)
    } catch (e) {
      console.error('Failed to create task, rolling back cache', e);
      setCachedTaskProgress(prev);
      return;
    }

    // Reset modal state
    setNewTaskName('');
    setNewTaskType('boolean');
    setModalVisible(false);
  };

  return (
    <>
      <Tabs
        screenOptions={{
          tabBarActiveTintColor: theme.tabIconSelected,
          tabBarInactiveTintColor: theme.tabIconDefault,
          tabBarStyle: {
            backgroundColor: theme.componentBackground,
            borderTopColor: theme.borderTop,
          },
          headerStyle: {
            backgroundColor: theme.background,
          },
          headerTitleStyle: {
            color: theme.text,
          },
          headerTintColor: theme.tint,
          headerShown: false,
        }}>
        <Tabs.Screen
          name="index"
          options={{
            title: 'Day',
            tabBarIcon: ({color}) => <TabBarIcon name="sun-o" color={color}/>,
          }}
        />
        <Tabs.Screen
          name="week"
          options={{
            title: 'Week',
            tabBarIcon: ({color}) => <TabBarIcon name="calendar" color={color}/>,
          }}
        />
        <Tabs.Screen
          name="month"
          options={{
            title: 'Month',
            tabBarIcon: ({color}) => <TabBarIcon name="calendar-o" color={color}/>,
          }}
        />
        <Tabs.Screen
          name="year"
          options={{
            title: 'Year',
            tabBarIcon: ({color}) => <TabBarIcon name="line-chart" color={color}/>,
          }}
        />
      </Tabs>

      {/* Global Add Task FAB */}
      <TouchableOpacity
        style={[styles.fab, {backgroundColor: theme.tint}]}
        onPress={() => setModalVisible(true)}
      >
        <FontAwesome name="plus" size={24} color="white"/>
      </TouchableOpacity>

      {/* Add Task Modal */}
      <Modal
        animationType="fade"
        transparent={true}
        visible={modalVisible}
        onRequestClose={() => setModalVisible(false)}
      >
        <KeyboardAvoidingView
          behavior={Platform.OS === "ios" ? "padding" : "height"}
          style={styles.centeredView}
        >
          <View style={[styles.modalView, {backgroundColor: theme.componentBackground || '#1a1a1a'}]}>
            <Text style={[styles.modalTitle, {color: theme.text}]}>Add New Task</Text>

            <Text style={[styles.label, {color: theme.tabIconDefault}]}>Task Name</Text>
            <TextInput
              style={[styles.input, {
                color: theme.text,
                backgroundColor: theme.background,
                borderColor: theme.borderTop || '#333'
              }]}
              onChangeText={setNewTaskName}
              value={newTaskName}
              placeholder="Enter task name..."
              placeholderTextColor={theme.tabIconDefault}
              autoFocus={true}
            />

            <Text style={[styles.label, {color: theme.tabIconDefault}]}>Task Type</Text>
            <View style={styles.typeContainer}>
              {taskTypes.map((type) => (
                <TouchableOpacity
                  key={type}
                  style={[
                    styles.typeButton,
                    {borderColor: theme.borderTop || '#333', backgroundColor: theme.background},
                    newTaskType === type && {backgroundColor: theme.tint, borderColor: theme.tint}
                  ]}
                  onPress={() => setNewTaskType(type)}
                >
                  <Text style={[
                    styles.typeButtonText,
                    {color: theme.text},
                    newTaskType === type && {color: 'white', fontWeight: 'bold'}
                  ]}>
                    {type.charAt(0).toUpperCase() + type.slice(1)}
                  </Text>
                </TouchableOpacity>
              ))}
            </View>

            <View style={styles.modalActions}>
              <TouchableOpacity
                style={[styles.button, styles.buttonClose]}
                onPress={() => setModalVisible(false)}
              >
                <Text style={styles.textStyle}>Cancel</Text>
              </TouchableOpacity>
              <TouchableOpacity
                style={[styles.button, {backgroundColor: theme.tint}]}
                onPress={handleAddTask}
              >
                <Text style={styles.textStyle}>Create</Text>
              </TouchableOpacity>
            </View>
          </View>
        </KeyboardAvoidingView>
      </Modal>
    </>
  );
}

function TabBarIcon(props: {
  name: React.ComponentProps<typeof FontAwesome>['name'];
  color: string;
}) {
  return <FontAwesome size={24} style={{marginBottom: -2}} {...props} />;
}

const styles = StyleSheet.create({
  fab: {
    position: 'absolute',
    width: 56,
    height: 56,
    alignItems: 'center',
    justifyContent: 'center',
    right: 20,
    bottom: 120, // Adjusted to sit above the tab bar roughly
    borderRadius: 28,
    elevation: 8,
    shadowColor: '#000',
    shadowOffset: {width: 0, height: 2},
    shadowOpacity: 0.25,
    shadowRadius: 3.84,
    zIndex: 1000,
  },
  centeredView: {
    flex: 1,
    justifyContent: 'flex-end',
    backgroundColor: 'rgba(0,0,0,0.7)', // Darker overlay
  },
  modalView: {
    width: '100%',
    borderTopLeftRadius: 20,
    borderTopRightRadius: 20,
    padding: 20,
    paddingBottom: 40,
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: -2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5,
  },
  modalTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    marginBottom: 20,
    textAlign: 'center',
  },
  label: {
    fontSize: 14,
    fontWeight: '600',
    marginBottom: 8,
    marginTop: 10,
  },
  input: {
    height: 50,
    borderWidth: 1,
    borderRadius: 12,
    paddingHorizontal: 15,
    fontSize: 16,
  },
  typeContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    gap: 10,
    marginBottom: 30,
    marginTop: 5,
  },
  typeButton: {
    paddingVertical: 8,
    paddingHorizontal: 16,
    borderRadius: 20,
    borderWidth: 1,
  },
  typeButtonText: {
    fontSize: 14,
  },
  modalActions: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    gap: 15,
    marginTop: 10,
  },
  button: {
    borderRadius: 12,
    padding: 15,
    elevation: 2,
    flex: 1,
    alignItems: 'center',
  },
  buttonClose: {
    backgroundColor: '#666',
  },
  textStyle: {
    color: 'white',
    fontWeight: 'bold',
    textAlign: 'center',
    fontSize: 16,
  },
});
