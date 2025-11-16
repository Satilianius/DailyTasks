import {RefreshControl, ScrollView, StyleSheet} from 'react-native';
import {View} from '@/components/Themed';
import TaskCard from '@/components/TaskCard';
import DateNavigator from "@/components/DateNavigator";
import {useContext, useEffect, useMemo, useState} from "react";
import {isBooleanTask, TaskProgress} from "@/models/AllTasksProgress";
import {TasksProgressContext} from "@/context/TasksProgressContext";

export default function DayScreen() {
    const [currentDate, setCurrentDate] = useState(new Date());
    const {cachedTaskProgress, setCachedTaskProgress, loading, loadRange} = useContext(TasksProgressContext)!;
    const [refreshing, setRefreshing] = useState(false);

    // Mock user ID - replace with actual auth user ID
    const userId = 'user-123-abc';

    const currentDayTasks = useMemo<TaskProgress[]>(() => {
        if (!cachedTaskProgress) return [];

        const dailyProgress = cachedTaskProgress.TasksProgress.find(day =>
            isSameDay(day.date, currentDate)
        );
        return dailyProgress?.tasks || [];
    }, [cachedTaskProgress, currentDate]);

    const cacheRange = useMemo(() => {
        if (!cachedTaskProgress || cachedTaskProgress.TasksProgress.length === 0) {
            return {start: null, end: null};
        }

        const dates = cachedTaskProgress.TasksProgress.map(d => d.date);
        return {
            start: new Date(Math.min(...dates.map(d => d.getTime()))),
            end: new Date(Math.max(...dates.map(d => d.getTime())))
        };
    }, [cachedTaskProgress]);

    useEffect(() => {
        if (!cachedTaskProgress || isCurrentDateNotInCache()) {
            void loadMonthData();
        }
        function isCurrentDateNotInCache() {
            return !cacheRange.start || !cacheRange.end
                || currentDate < cacheRange.start || currentDate > cacheRange.end;
        }
    }, [currentDate]);

    const loadMonthData = async () => {
        // Load 1 month back, 1 week forward (better for historical tracking)
        const startDate = new Date(currentDate);
        startDate.setDate(startDate.getDate() - 30);

        const endDate = new Date(currentDate);
        endDate.setDate(endDate.getDate() + 7);

        await loadRange(userId, startDate, endDate);
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
        const task = currentDayTasks.find(t => t.taskId === taskId);
        if (!task || !isBooleanTask(task) || !cachedTaskProgress) return;

        const updatedProgress = cachedTaskProgress.TasksProgress.map(day =>
            isSameDay(day.date, currentDate)
                ? {
                    ...day,
                    tasks: day.tasks.map(t =>
                        t.taskId === taskId && isBooleanTask(t)
                            ? {...t, progress: !t.progress}
                            : t
                    )
                }
                : day
        );

        setCachedTaskProgress({...cachedTaskProgress, TasksProgress: updatedProgress});

        // TODO: Send update to backend
        // taskService.updateTask(taskId, { progress: !task.progress });
    };

    const handleManualRefresh = async () => {
        setRefreshing(true);
        await loadMonthData();
        setRefreshing(false);
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
                refreshControl={
                    <RefreshControl
                        refreshing={refreshing}
                        onRefresh={handleManualRefresh}
                    />
                }
            >
                {loading && !cachedTaskProgress ? (
                    <View style={styles.taskGrid}>
                        {/* You can add a loading spinner here */}
                    </View>
                ) : (
                    <View style={styles.taskGrid}>
                        {currentDayTasks.map((task) => (
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
