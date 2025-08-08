import React from 'react';
import { StyleSheet } from 'react-native';
import { DataTable } from 'react-native-paper';
import { ThemedView } from '@/components/ThemedView';

type DayValue = boolean | number;

interface TaskProgress {
    Task: string;
    Monday: DayValue;
    Tuesday: DayValue;
    Wednesday: DayValue;
    Thursday: DayValue;
    Friday: DayValue;
    Saturday: DayValue;
    Sunday: DayValue;
}

const formatValue = (value: DayValue, taskName: string): string => {
    if (typeof value === 'boolean') {
        return value ? '✓' : '−';
    }
    if (taskName.includes('(m)')) {
        return value === 0 ? '−' : `${value}m`;
    }
    if (taskName.includes('(ms)')) {
        return value === 0 ? '−' : `${(value / 1000).toFixed(1)}s`;
    }
    return value.toString();
};

const ProgressTable = () => {
    const data: TaskProgress[] = [
        {
            Task: 'Water Plants',
            Monday: true,
            Tuesday: true,
            Wednesday: false,
            Thursday: false,
            Friday: false,
            Saturday: false,
            Sunday: false,
        },
        {
            Task: 'Running(m)',
            Monday: 1100,
            Tuesday: 1542,
            Wednesday: 0,
            Thursday: 0,
            Friday: 0,
            Saturday: 0,
            Sunday: 0,
        },
        {
            Task: 'Speed Climb(ms)',
            Monday: 30000,
            Tuesday: 28263,
            Wednesday: 0,
            Thursday: 0,
            Friday: 0,
            Saturday: 0,
            Sunday: 0,
        },
    ];

    return (
        <ThemedView style={styles.container}>
            <DataTable>
                <DataTable.Header>
                    <DataTable.Title>Task</DataTable.Title>
                    <DataTable.Title numeric>Mon</DataTable.Title>
                    <DataTable.Title numeric>Tue</DataTable.Title>
                    <DataTable.Title numeric>Wed</DataTable.Title>
                    <DataTable.Title numeric>Thu</DataTable.Title>
                    <DataTable.Title numeric>Fri</DataTable.Title>
                    <DataTable.Title numeric>Sat</DataTable.Title>
                    <DataTable.Title numeric>Sun</DataTable.Title>
                </DataTable.Header>

                {data.map((item) => (
                    <DataTable.Row key={item.Task}>
                        <DataTable.Cell>{item.Task}</DataTable.Cell>
                        <DataTable.Cell numeric>{formatValue(item.Monday, item.Task)}</DataTable.Cell>
                        <DataTable.Cell numeric>{formatValue(item.Tuesday, item.Task)}</DataTable.Cell>
                        <DataTable.Cell numeric>{formatValue(item.Wednesday, item.Task)}</DataTable.Cell>
                        <DataTable.Cell numeric>{formatValue(item.Thursday, item.Task)}</DataTable.Cell>
                        <DataTable.Cell numeric>{formatValue(item.Friday, item.Task)}</DataTable.Cell>
                        <DataTable.Cell numeric>{formatValue(item.Saturday, item.Task)}</DataTable.Cell>
                        <DataTable.Cell numeric>{formatValue(item.Sunday, item.Task)}</DataTable.Cell>
                    </DataTable.Row>
                ))}
            </DataTable>
        </ThemedView>
    );
};

const styles = StyleSheet.create({
    container: {
        padding: 8,
    },
});

export default ProgressTable;
