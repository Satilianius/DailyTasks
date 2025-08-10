import { describe, test, expect } from '@jest/globals';
import { getStartOfWeek, getWeekDays } from './date';

describe('date utils', () => {
  test('getStartOfWeek returns Monday for any date in the week', () => {
    const dates = [
      new Date('2025-08-10T12:00:00Z'), // Sunday
      new Date('2025-08-11T12:00:00Z'), // Monday
      new Date('2025-08-12T12:00:00Z'), // Tuesday
      new Date('2025-08-16T12:00:00Z'), // Saturday
    ];
    for (const d of dates) {
      const start = getStartOfWeek(d);
      expect(start.getDay()).toBe(1); // Monday
      // Ensure time is midnight
      expect(start.getHours()).toBe(0);
      expect(start.getMinutes()).toBe(0);
    }
  });

  test('getWeekDays produces 7 consecutive days starting Monday', () => {
    const base = new Date('2025-08-13T12:00:00Z');
    const week = getWeekDays(base);
    expect(week).toHaveLength(7);
    expect(week[0].date.getDay()).toBe(1);
    for (let i = 1; i < week.length; i++) {
      const prev = week[i - 1].date;
      const curr = week[i].date;
      const diffMs = curr.getTime() - prev.getTime();
      const oneDayMs = 24 * 60 * 60 * 1000;
      // Allow for possible DST shifts by accepting a small range around 24h
      expect(Math.abs(diffMs - oneDayMs)).toBeLessThanOrEqual(60 * 60 * 1000);
    }
  });
});
