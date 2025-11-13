import {ScrollView, StyleSheet} from 'react-native';
import {View} from '@/components/Themed';
import TaskCard from '@/components/TaskCard';
import DateNavigator from "@/components/DateNavigator";
import {useEffect, useState} from "react";
import {
    isBooleanTask,
    TaskProgress,
    UserTasksProgressDto
} from "@/models/AllTasksProgress";

// TODO review
export default function DayScreen() {
    const [currentDate, setCurrentDate] = useState(new Date());
    const [tasks, setTasks] = useState<TaskProgress[]>([]);
    const [loading, setLoading] = useState(true);
    const [cachedData, setCachedData] = useState<UserTasksProgressDto | null>(null);
    const [cacheStartDate, setCacheStartDate] = useState<Date | null>(null);
    const [cacheEndDate, setCacheEndDate] = useState<Date | null>(null);

    // Mock user ID - replace with actual auth user ID
    const userId = 'user-123-abc';

    useEffect(() => {
        loadMonthData();
    }, []);

    useEffect(() => {
        if (cachedData) {
            updateTasksForCurrentDate();
            // Check if we need to load more data
            if (isDateOutsideCache(currentDate)) {
                loadMonthData();
            }
        }
    }, [currentDate, cachedData]);

    const loadMonthData = async () => {
        setLoading(true);
        try {
            // Calculate date range: 2 weeks before and 2 weeks after current date
            const startDate = new Date(currentDate);
            startDate.setDate(startDate.getDate() - 14);

            const endDate = new Date(currentDate);
            endDate.setDate(endDate.getDate() + 14);

            const userProgress = await fetchUserTasksProgress(userId, startDate, endDate);

            setCachedData(userProgress);
            setCacheStartDate(startDate);
            setCacheEndDate(endDate);
        } catch (error) {
            console.error('Failed to load tasks:', error);
            setTasks([]);
        } finally {
            setLoading(false);
        }
    };

    const updateTasksForCurrentDate = () => {
        if (!cachedData) return;

        const dateString = currentDate.toISOString().split('T')[0];
        const dailyProgress = cachedData.TasksProgress.find(
            (dailyTaskProgress) => dailyTaskProgress.date === dateString
        );
        setTasks(dailyProgress?.tasks || []);
    };

    const isDateOutsideCache = (date: Date): boolean => {
        if (!cacheStartDate || !cacheEndDate) return true;
        return date < cacheStartDate || date > cacheEndDate;
    };

    const handlePreviousDay = () => {
        const newDate = new Date(currentDate);
        newDate.setDate(newDate.getDate() - 1);
        setCurrentDate(newDate);
    };

    const handleNextDay = () => {
        const newDate = new Date(currentDate);
        newDate.setDate(newDate.getDate() + 1);
        setCurrentDate(newDate);
    };

    const handleTaskPress = (taskId: string) => {
        // For now, only handle boolean tasks
        const task = tasks.find(t => t.taskId === taskId);
        if (task && isBooleanTask(task)) {
            // Update local state
            const updatedTasks = tasks.map(t =>
                t.taskId === taskId && isBooleanTask(t)
                    ? {...t, progress: !t.progress}
                    : t
            );
            setTasks(updatedTasks);

            // Update cache
            if (cachedData) {
                const dateString = currentDate.toISOString().split('T')[0];
                const updatedProgress = cachedData.TasksProgress.map(day =>
                    day.date === dateString
                        ? {...day, tasks: updatedTasks}
                        : day
                );
                setCachedData({...cachedData, TasksProgress: updatedProgress});
            }

            // TODO: Send update to backend
            // taskService.updateTask(taskId, { progress: !task.progress });
        }
    };

    return (
        <View style={styles.container}>
            {/* Date Navigator */}
            <DateNavigator
                currentDate={currentDate}
                onPrevious={handlePreviousDay}
                onNext={handleNextDay}
            />

            {/* Task Grid */}
            <ScrollView
                style={styles.scrollView}
                contentContainerStyle={styles.scrollContent}
            >
                {loading ? (
                    <View style={styles.taskGrid}>
                        {/* You can add a loading spinner here */}
                    </View>
                ) : (
                    <View style={styles.taskGrid}>
                        {tasks.map((task) => (
                            <TaskCard
                                key={task.taskId}
                                task={task}
                                onPress={() => handleTaskPress(task.taskId)}
                            />
                        ))}
                    </View>
                )}
            </ScrollView>
        </View>
    );
}

// Mock function to simulate backend call - fetches a month of data
const fetchUserTasksProgress = async (userId: string, startDate: Date, endDate: Date): Promise<UserTasksProgressDto> => {
    console.log(`Fetching data from ${startDate.toISOString()} to ${endDate.toISOString()}`);
    // Simulate network delay
    await new Promise(resolve => setTimeout(resolve, 500));

    // Generate mock data for the date range
    const tasksProgress = [];
    const currentDate = new Date(startDate);

    while (currentDate <= endDate) {
        const dateString = currentDate.toISOString().split('T')[0];
        const dayOfMonth = currentDate.getDate();

        // Vary the data based on the day for more realistic mock data
        tasksProgress.push({
            date: dateString,
            tasks: [
                {
                    taskId: 'a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d',
                    taskName: 'Water Plants',
                    type: 'boolean' as const,
                    progress: dayOfMonth % 3 !== 0
                },
                {
                    taskId: 'b2c3d4e5-f6a7-4b5c-9d0e-1f2a3b4c5d6e',
                    taskName: 'Go to bed',
                    type: 'time' as const,
                    progress: '23:30'
                },
                {
                    taskId: 'c3d4e5f6-a7b8-4c5d-0e1f-2a3b4c5d6e7f',
                    taskName: 'Push Ups',
                    type: 'number' as const,
                    progress: 30 + (dayOfMonth % 20)
                },
                {
                    taskId: 'd4e5f6a7-b8c9-4d5e-1f2a-3b4c5d6e7f8a',
                    taskName: 'Meditate',
                    type: 'duration' as const,
                    progress: `00:${10 + (dayOfMonth % 30)}:${(dayOfMonth * 13) % 60}.000`
                },
            ]
        });

        currentDate.setDate(currentDate.getDate() + 1);
    }

    return {
        userId: userId,
        TasksProgress: tasksProgress
    };
};

function isSameDay(date1: Date, date2: Date): boolean {
    return date1.getFullYear() === date2.getFullYear() &&
        date1.getMonth() === date2.getMonth() &&
        date1.getDate() === date2.getDate();
}


const styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        gap: 8,
    },
    scrollView: {
        width: '100%',
    },
    scrollContent: {
        alignItems: 'center',
        paddingVertical: 8,
    },
    taskGrid: {
        flexDirection: 'row',
        flexWrap: 'wrap',
        justifyContent: 'center',
        maxWidth: 1200,
        width: '100%',
        paddingHorizontal: 8,
    },
    title: {
        fontSize: 20,
        fontWeight: 'bold',
    },
    caption: {
        opacity: 0.7,
    },
    dateNav: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: 16,
    },
});