package main

import (
	"DailyTasks/Progress"
	"DailyTasks/Tasks"
	"fmt"
	"time"
)

func printProgress(tasksWithProgress map[Tasks.Task]Progress.PrintableProgress) {
	dates := getThisWeek()
	printHeader(dates[:])

	for task, progress := range tasksWithProgress {
		fmt.Printf("%15s:", task.Name)
		for _, date := range dates {
			fmt.Printf("%11s", progress.GetPrintableProgressAtDate(date))
		}
		fmt.Println()
	}
}

func printHeader(dates []time.Time) {
	fmt.Printf("%15s:", "Tasks/Dates")

	for _, date := range dates {
		fmt.Printf("%11s", date.Format("02/01/2006"))
	}
	fmt.Println("")
}

func getThisWeek() [7]time.Time {
	today := today()
	var thisWeek [7]time.Time
	for weekdayIndex := 0; weekdayIndex < 7; weekdayIndex++ {
		// Adjust Weekday to make Monday 0 (0 is Sunday by default)
		adjustedWeekday := (int(today.Weekday()) + 6) % 7
		daysToAdd := weekdayIndex - adjustedWeekday
		thisWeek[weekdayIndex] = today.AddDate(0, 0, daysToAdd)
	}

	return thisWeek
}

func today() time.Time {
	return getStartOfDay(time.Now())
}

func getStartOfDay(day time.Time) time.Time {
	return day.Truncate(24 * time.Hour)
}
