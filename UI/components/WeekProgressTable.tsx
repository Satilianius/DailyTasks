import React from 'react';
import { FlatList, StyleSheet } from 'react-native';
import { ThemedText } from '@/components/ThemedText';
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

    const days = ['Task', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];

    const renderHeader = () => (
        <ThemedView style={styles.headerRow}>
            {days.map((day, index) => (
                <ThemedText
                    key={day}
                    style={[styles.headerCell, index === 0 ? styles.taskColumn : styles.dayColumn]}
                    type="defaultSemiBold">
                    {day}
                </ThemedText>
            ))}
        </ThemedView>
    );

    const renderRow = ({ item }: { item: TaskProgress }) => (
        <ThemedView style={styles.row}>
            <ThemedText style={[styles.cell, styles.taskColumn]} numberOfLines={2}>
                {item.Task}
            </ThemedText>
            <ThemedText style={[styles.cell, styles.dayColumn]}>
                {formatValue(item.Monday, item.Task)}
            </ThemedText>
            <ThemedText style={[styles.cell, styles.dayColumn]}>
                {formatValue(item.Tuesday, item.Task)}
            </ThemedText>
            <ThemedText style={[styles.cell, styles.dayColumn]}>
                {formatValue(item.Wednesday, item.Task)}
            </ThemedText>
            <ThemedText style={[styles.cell, styles.dayColumn]}>
                {formatValue(item.Thursday, item.Task)}
            </ThemedText>
            <ThemedText style={[styles.cell, styles.dayColumn]}>
                {formatValue(item.Friday, item.Task)}
            </ThemedText>
            <ThemedText style={[styles.cell, styles.dayColumn]}>
                {formatValue(item.Saturday, item.Task)}
            </ThemedText>
            <ThemedText style={[styles.cell, styles.dayColumn]}>
                {formatValue(item.Sunday, item.Task)}
            </ThemedText>
        </ThemedView>
    );

    return (
        <ThemedView style={styles.container}>
            {renderHeader()}
            <FlatList
                data={data}
                renderItem={renderRow}
                keyExtractor={(item) => item.Task}
                scrollEnabled={false}
            />
        </ThemedView>
    );
};

const styles = StyleSheet.create({
    container: {
        borderRadius: 8,
        overflow: 'hidden',
        borderWidth: 1,
        borderColor: '#e5e7eb',
    },
    headerRow: {
        flexDirection: 'row',
        padding: 8,
        borderBottomWidth: 1,
        borderBottomColor: '#e5e7eb',
    },
    row: {
        flexDirection: 'row',
        padding: 8,
        borderBottomWidth: 1,
        borderBottomColor: '#e5e7eb',
    },
    headerCell: {
        fontSize: 14,
        textAlign: 'center',
    },
    cell: {
        fontSize: 14,
        textAlign: 'center',
    },
    taskColumn: {
        flex: 2,
        textAlign: 'left',
    },
    dayColumn: {
        flex: 1,
    },
});

export default ProgressTable;
