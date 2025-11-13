export type TaskType = 'boolean' | 'number' | 'time' | 'duration';

// Main DTO
export interface UserTasksProgressDto {
    userId: string;
    TasksProgress: DailyTasksProgress[];
}

export interface DailyTasksProgress {
    date: Date;
    tasks: TaskProgress[];
}

// Union type for all task progress types
export type TaskProgress =
    | BooleanTaskProgress
    | NumberTaskProgress
    | TimeTaskProgress
    | DurationTaskProgress;

// Base task progress interface
interface BaseTaskProgress {
    taskId: string;
    taskName: string;
    type: TaskType;
    progress: any;
}

export interface BooleanTaskProgress extends BaseTaskProgress {
    type: 'boolean';
    progress: boolean;
}

export interface NumberTaskProgress extends BaseTaskProgress {
    type: 'number';
    progress: number;
}

export interface TimeTaskProgress extends BaseTaskProgress {
    type: 'time';
    progress: string; // HH:mm:ss.mmm format
}

export interface DurationTaskProgress extends BaseTaskProgress {
    type: 'duration';
    progress: string; // HH:mm:ss.mmm format
}

// Type guard functions for discriminating task types
export function isBooleanTask(task: TaskProgress): task is BooleanTaskProgress {
    return task.type === 'boolean';
}

export function isNumberTask(task: TaskProgress): task is NumberTaskProgress {
    return task.type === 'number';
}

export function isTimeTask(task: TaskProgress): task is TimeTaskProgress {
    return task.type === 'time';
}

export function isDurationTask(task: TaskProgress): task is DurationTaskProgress {
    return task.type === 'duration';
}
