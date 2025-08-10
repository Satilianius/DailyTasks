export function getStartOfWeek(date: Date): Date {
  // Start week on Monday
  const d = new Date(date);
  const day = d.getDay(); // 0 (Sun) - 6 (Sat)
  const diffToMonday = (day === 0 ? -6 : 1) - day;
  d.setDate(d.getDate() + diffToMonday);
  d.setHours(0, 0, 0, 0);
  return d;
}

export function formatDayLabel(d: Date): string {
  return d.toLocaleDateString(undefined, { weekday: 'short' });
}

export function getWeekDays(date: Date): { date: Date; label: string }[] {
  const start = getStartOfWeek(date);
  return Array.from({ length: 7 }).map((_, i) => {
    const curr = new Date(start);
    curr.setDate(start.getDate() + i);
    return { date: curr, label: formatDayLabel(curr) };
  });
}
