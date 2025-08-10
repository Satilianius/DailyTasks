import React, { useMemo } from 'react';
import { ScrollView, StyleSheet, View as RNView } from 'react-native';
import { Text, View } from '@/components/Themed';

import { getWeekDays } from '@/utils/date';

// Data model for tasks (example)
type TaskType = 'bool' | 'int' | 'float' | 'time' | 'duration';

type Task = {
  id: string;
  name: string;
  type: TaskType;
  // values per weekday index 0..6 (Mon..Sun)
  values: (boolean | number | string | null)[];
};

const exampleTasks: Task[] = [
  {
    id: 't1',
    name: 'Water plants',
    type: 'bool',
    values: [true, true, false, true, false, false, true],
  },
  {
    id: 't2',
    name: 'Push ups',
    type: 'int',
    values: [20, 0, 30, 10, 0, 0, 25],
  },
  {
    id: 't3',
    name: 'Go to bed',
    type: 'time',
    values: ['23:15', '22:45', '00:10', '23:50', '23:30', null, '22:55'],
  },
  {
    id: 't4',
    name: 'Spin bike',
    type: 'duration',
    values: [30, 0, 45, 20, 0, 60, 0], // minutes
  },
  {
    id: 't5',
    name: 'Run',
    type: 'float',
    values: [3.5, 0, 5.2, 0, 2.1, 0, 7.0], // km
  },
];

function renderCellValue(type: TaskType, value: boolean | number | string | null): string {
  if (value === null || value === undefined) return '-';
  switch (type) {
    case 'bool':
      return value ? '✅' : '⬜';
    case 'int':
      return String(value);
    case 'float':
      return `${(value as number).toFixed(1)} km`;
    case 'time':
      return String(value);
    case 'duration':
      return `${value} min`;
  }
}

export default function WeekScreen() {
  const days = useMemo(() => getWeekDays(new Date()), []);

  return (
    <View style={styles.container}>
      <Text style={styles.title}>This Week</Text>
      <RNView style={styles.tableWrapper}>
        <ScrollView horizontal contentContainerStyle={styles.scrollContainer}>
          <RNView style={styles.table}>
            {/* Header Row */}
            <RNView style={[styles.row, styles.headerRow]}>
              <RNView style={[styles.cell, styles.taskHeaderCell]}>
                <Text style={[styles.headerText, styles.taskHeaderText]}>Task</Text>
              </RNView>
              {days.map((d) => (
                <RNView key={d.label} style={[styles.cell, styles.dayHeaderCell]}>
                  <Text style={styles.headerText}>{d.label}</Text>
                </RNView>
              ))}
            </RNView>

            {/* Data Rows */}
            {exampleTasks.map((task) => (
              <RNView key={task.id} style={styles.row}>
                <RNView style={[styles.cell, styles.taskNameCell]}>
                  <Text style={styles.taskNameText}>{task.name}</Text>
                </RNView>
                {days.map((_, idx) => (
                  <RNView key={idx} style={[styles.cell, styles.valueCell]}>
                    <Text style={styles.valueText}>{renderCellValue(task.type, task.values[idx] ?? null)}</Text>
                  </RNView>
                ))}
              </RNView>
            ))}
          </RNView>
        </ScrollView>
      </RNView>
    </View>
  );
}

const TASK_COL_WIDTH = 140;
const DAY_COL_WIDTH = 90;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 16,
  },
  title: {
    fontSize: 22,
    fontWeight: '600',
    marginHorizontal: 12,
    marginBottom: 8,
  },
  tableWrapper: {
    flex: 1,
  },
  scrollContainer: {
    paddingBottom: 24,
  },
  table: {
    flexDirection: 'column',
    alignSelf: 'flex-start',
  },
  row: {
    flexDirection: 'row',
    alignItems: 'stretch',
  },
  headerRow: {
    borderBottomWidth: StyleSheet.hairlineWidth,
    borderColor: '#8884',
  },
  cell: {
    paddingVertical: 10,
    paddingHorizontal: 8,
    borderRightWidth: StyleSheet.hairlineWidth,
    borderColor: '#8884',
    minWidth: 80,
    justifyContent: 'center',
  },
  taskHeaderCell: {
    minWidth: TASK_COL_WIDTH,
    borderLeftWidth: StyleSheet.hairlineWidth,
    borderColor: '#8884',
  },
  taskNameCell: {
    minWidth: TASK_COL_WIDTH,
    borderLeftWidth: StyleSheet.hairlineWidth,
    borderColor: '#8884',
  },
  dayHeaderCell: {
    backgroundColor: 'rgba(127,127,127,0.08)',
    minWidth: DAY_COL_WIDTH,
  },
  headerText: {
    fontWeight: '600',
    textAlign: 'center',
  },
  taskHeaderText: {
    textAlign: 'left',
  },
  taskNameText: {
    fontWeight: '500',
  },
  valueCell: {
    minWidth: DAY_COL_WIDTH,
  },
  valueText: {
    textAlign: 'center',
  },
});
