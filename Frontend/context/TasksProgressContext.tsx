import React, {createContext, PropsWithChildren, useCallback, useState} from 'react';
import {DailyTasksProgress, UserTasksProgressDto} from '@/models/AllTasksProgress';
import {formatDateToISO} from '@/utils/date';

export interface TasksProgressContextValue {
    cachedTaskProgress: UserTasksProgressDto | null;
    setCachedTaskProgress: React.Dispatch<React.SetStateAction<UserTasksProgressDto | null>>;
    loading: boolean;
    loadRange: (userId: string, startDate: Date, endDate: Date) => Promise<void>;
}

export const TasksProgressContext = createContext<TasksProgressContextValue | undefined>(undefined);

export function TasksProgressProvider({children}: PropsWithChildren) {
    const [cachedTaskProgress, setCachedTaskProgress] = useState<UserTasksProgressDto | null>(null);
    const [loading, setLoading] = useState(false);

    const loadRange = useCallback(async (userId: string, startDate: Date, endDate: Date) => {
        setLoading(true);
        try {
            const progress = await fetchUserTasksProgress(userId, startDate, endDate);
            setCachedTaskProgress(progress);
        } catch (e) {
            console.error('Failed to load tasks:', e);
        } finally {
            setLoading(false);
        }
    }, []);

    return (
        <TasksProgressContext.Provider value={{cachedTaskProgress, setCachedTaskProgress, loading, loadRange}}>
            {children}
        </TasksProgressContext.Provider>
    );
}

// Mock function to simulate backend call - fetches a range of data
async function fetchUserTasksProgress(userId: string, startDate: Date, endDate: Date): Promise<UserTasksProgressDto> {
    console.log(`Fetching data from ${formatDateToISO(startDate)} to ${formatDateToISO(endDate)}`);
    // Simulate network delay
    await new Promise(resolve => setTimeout(resolve, 500));

    // Generate mock data for the date range
    const tasksProgress: DailyTasksProgress[] = [];
    const currentDate = new Date(startDate);

    while (currentDate <= endDate) {
        const dayOfMonth = currentDate.getDate();

        tasksProgress.push({
            date: new Date(currentDate),
            tasks: [
                {
                    taskId: 'a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d',
                    taskName: 'Water Plants',
                    type: 'boolean' as const,
                    progress: dayOfMonth % 3 !== 0,
                },
                {
                    taskId: 'b2c3d4e5-f6a7-4b5c-9d0e-1f2a3b4c5d6e',
                    taskName: 'Go to bed',
                    type: 'time' as const,
                    progress: '23:30:00.000',
                },
                {
                    taskId: 'c3d4e5f6-a7b8-4c5d-0e1f-2a3b4c5d6e7f',
                    taskName: 'Push Ups',
                    type: 'number' as const,
                    progress: 30 + (dayOfMonth % 20),
                },
                {
                    taskId: 'd4e5f6a7-b8c9-4d5e-1f2a-3b4c5d6e7f8a',
                    taskName: 'Meditate',
                    type: 'duration' as const,
                    progress: `00:${10 + (dayOfMonth % 30)}:${(dayOfMonth * 13) % 60}.000`,
                },
            ],
        });

        currentDate.setDate(currentDate.getDate() + 1);
    }

    return {
        userId: userId,
        TasksProgress: tasksProgress,
    };
}
