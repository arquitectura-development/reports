package reports

import (
	"arqui-reports/types"
	"log"
	"math"
	"time"
)

func ProcessUserReport(tasks []types.Task, habits []types.Habit, response *types.UserReport) {
	goodHabits := []types.Habit{}
	badHabits := []types.Habit{}
	tasksToday := []types.Task{}
	tasksDelayed := []types.Task{}

	for _, habit := range habits {
		if habit.Score >= 50 {
			goodHabits = append(goodHabits, habit)
		} else if habit.Score < 0 {
			badHabits = append(badHabits, habit)
		}
	}

	layout := "02/01/2006"
	todayDateString := time.Now().Format(layout)
	todayDate, err := time.Parse(layout, todayDateString)
	if err != nil {
		log.Panicln(err)
	}

	for _, task := range tasks {
		if !task.Done {
			dueDateString := task.DueDate
			dueDate, err := time.Parse(layout, dueDateString)
			if err != nil {
				log.Panicln(err)
			} else if dueDate.Equal(todayDate) {
				tasksToday = append(tasksToday, task)
			} else if dueDate.Before(todayDate) {
				tasksDelayed = append(tasksDelayed, task)
			}
		}
	}

	response.BadHabits = badHabits
	response.GoodHabits = goodHabits
	response.TodayTasks = tasksToday
	response.DelayedTasks = tasksDelayed
}

func ProcessAdminTasks(tasks []types.Task, response *types.AdminTasksReport) {
	completedTotal := 0
	completedBefore := 0
	completedAfter := 0
	delayed := 0
	availableTotal := 0
	availableRemainig := 0
	availableForToday := 0

	layout := "02/01/2006"
	todayDateString := time.Now().Format(layout)
	todayDate, err := time.Parse(layout, todayDateString)
	if err != nil {
		log.Panicln(err)
	}
	for _, task := range tasks {

		dueDateString := task.DueDate
		dueDate, err := time.Parse(layout, dueDateString)
		if err != nil {
			log.Panicln(err)
		}

		if task.Done {
			completionDateString := task.CompletionDate
			completionDate, err := time.Parse(layout, completionDateString)
			if err != nil {
				log.Panicln(err)
			}

			if completionDate.After(dueDate) || completionDate.Equal(dueDate) {
				completedBefore++
			} else {
				completedAfter++
			}
			completedTotal++
		} else if dueDate.Before(todayDate) {
			delayed++
		} else {
			if dueDate.Equal(todayDate) {
				availableForToday++
			} else {
				availableRemainig++
			}
			availableTotal++
		}
	}

	response.Completed = types.CompletedTasks{
		Total:  completedTotal,
		Before: completedBefore,
		After:  completedAfter,
	}
	response.Delayed = delayed
	response.Available = types.AvailableTasks{
		Total:     availableTotal,
		Remaining: availableRemainig,
		ForToday:  availableForToday,
	}
}

func ProcessAdminHabitsWithoutNames(habits []types.Habit, perRange *types.Ranges, lowestTemp *types.UserHabitTemp, highestTemp *types.UserHabitTemp) {
	countRed := 0
	countOrange := 0
	countYellow := 0
	countGreen := 0
	countBlue := 0
	highestScore := math.Inf(-1)
	lowestScore := math.Inf(1)

	for _, habit := range habits {
		if float64(habit.Score) < lowestScore {
			lowestScore = float64(habit.Score)
			lowestTemp.UserID = habit.UserID
			lowestTemp.HabitName = habit.Name
		}
		if float64(habit.Score) > highestScore {
			highestScore = float64(habit.Score)
			highestTemp.UserID = habit.UserID
			highestTemp.HabitName = habit.Name
		}

		if habit.Score < 0 {
			countRed++
		} else if habit.Score >= 0 && habit.Score < 10 {
			countOrange++
		} else if habit.Score >= 10 && habit.Score < 40 {
			countYellow++
		} else if habit.Score >= 40 && habit.Score < 50 {
			countGreen++
		} else if habit.Score >= 50 {
			countBlue++
		}
	}

	perRange.Red = countRed
	perRange.Orange = countOrange
	perRange.Yellow = countYellow
	perRange.Green = countGreen
	perRange.Blue = countBlue
}
