export enum TaskType {
  Boolean = 'boolean',
  Number = 'number',
  Time = 'time',
  Duration = 'duration',
}

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
  type: TaskType.Boolean;
  progress: boolean;
}

export interface NumberTaskProgress extends BaseTaskProgress {
  type: TaskType.Number;
  progress: number;
}

export interface TimeTaskProgress extends BaseTaskProgress {
  type: TaskType.Time;
  progress: string; // HH:mm:ss.mmm format
}

export interface DurationTaskProgress extends BaseTaskProgress {
  type: TaskType.Duration;
  progress: string; // HH:mm:ss.mmm format
}

// Type guard functions for discriminating task types
export function isBooleanTask(task: TaskProgress): task is BooleanTaskProgress {
  return task.type === TaskType.Boolean;
}

export function isNumberTask(task: TaskProgress): task is NumberTaskProgress {
  return task.type === TaskType.Number;
}

export function isTimeTask(task: TaskProgress): task is TimeTaskProgress {
  return task.type === TaskType.Time;
}

export function isDurationTask(task: TaskProgress): task is DurationTaskProgress {
  return task.type === TaskType.Duration;
}
