package reports

import (
	"arqui-reports/types"
	"testing"
	"time"
)

var layout = "02/01/2006"
var todayDateString = time.Now().Format(layout)
var pastDateString = time.Now().AddDate(0, 0, -1).Format(layout)
var tomorrowDateString = time.Now().AddDate(0, 0, 1).Format(layout)
var futureDateString = time.Now().AddDate(1, 0, 1).Format(layout)

var habitsSample = []types.Habit{
	types.Habit{
		ID:         1,
		UserID:     1,
		Name:       "Program in Kotlin",
		HabitType:  1,
		Difficulty: 3,
		Score:      15,
		Color:      3,
	},
	types.Habit{
		ID:         2,
		UserID:     1,
		Name:       "Do Clean Code",
		HabitType:  1,
		Difficulty: 2,
		Score:      140,
		Color:      5,
	},
	types.Habit{
		ID:         7,
		UserID:     1,
		Name:       "Bully kids",
		HabitType:  2,
		Difficulty: 1,
		Score:      -10,
		Color:      1,
	},
	types.Habit{
		ID:         141,
		UserID:     1,
		Name:       "Do things",
		HabitType:  3,
		Difficulty: 1,
		Score:      70,
		Color:      5,
	},
}

var habitsSampleAdmin = []types.Habit{
	types.Habit{
		ID:         1,
		UserID:     1,
		Name:       "Program in Kotlin",
		HabitType:  1,
		Difficulty: 3,
		Score:      15,
		Color:      3,
	},
	types.Habit{
		ID:         10,
		UserID:     1,
		Name:       "Program in Kotlin2",
		HabitType:  1,
		Difficulty: 3,
		Score:      9,
		Color:      2,
	},
	types.Habit{
		ID:         2,
		UserID:     1,
		Name:       "Do Clean Code",
		HabitType:  1,
		Difficulty: 2,
		Score:      140,
		Color:      5,
	},
	types.Habit{
		ID:         7,
		UserID:     2,
		Name:       "Bully kids",
		HabitType:  2,
		Difficulty: 1,
		Score:      -10,
		Color:      1,
	},
	types.Habit{
		ID:         141,
		UserID:     1,
		Name:       "Do things",
		HabitType:  3,
		Difficulty: 1,
		Score:      70,
		Color:      5,
	},
}

var tasksSample = []types.Task{
	types.Task{
		ID:             8,
		UserID:         1,
		Title:          "task done",
		Description:    "desc",
		DueDate:        todayDateString,
		Reminder:       "",
		Done:           true,
		CompletionDate: pastDateString,
	},
	types.Task{
		ID:             9,
		UserID:         1,
		Title:          "task for today",
		Description:    "do stuff today",
		DueDate:        todayDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             9,
		UserID:         1,
		Title:          "task for tomorrow",
		Description:    "",
		DueDate:        tomorrowDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             9,
		UserID:         1,
		Title:          "delayed task",
		Description:    "",
		DueDate:        pastDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             9,
		UserID:         1,
		Title:          "very distant future task",
		Description:    "",
		DueDate:        futureDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
}

var tasksSampleAdmin = []types.Task{
	types.Task{
		ID:             1,
		UserID:         1,
		Title:          "task done",
		Description:    "desc",
		DueDate:        todayDateString,
		Reminder:       "",
		Done:           true,
		CompletionDate: pastDateString,
	},
	types.Task{
		ID:             2,
		UserID:         1,
		Title:          "task for today",
		Description:    "do stuff today",
		DueDate:        todayDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             3,
		UserID:         1,
		Title:          "task for tomorrow",
		Description:    "",
		DueDate:        tomorrowDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             4,
		UserID:         1,
		Title:          "delayed task",
		Description:    "",
		DueDate:        pastDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             5,
		UserID:         1,
		Title:          "very distant future task",
		Description:    "",
		DueDate:        futureDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             15,
		UserID:         2,
		Title:          "very distant future task",
		Description:    "",
		DueDate:        futureDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             25,
		UserID:         3,
		Title:          "very distant future task",
		Description:    "",
		DueDate:        futureDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             12,
		UserID:         2,
		Title:          "task for today",
		Description:    "do stuff today",
		DueDate:        todayDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             14,
		UserID:         2,
		Title:          "delayed task",
		Description:    "",
		DueDate:        pastDateString,
		Reminder:       "",
		Done:           false,
		CompletionDate: "",
	},
	types.Task{
		ID:             54,
		UserID:         2,
		Title:          "completed delayed task",
		Description:    "",
		DueDate:        pastDateString,
		Reminder:       "",
		Done:           true,
		CompletionDate: todayDateString,
	},
}

func TestProcessUserReport(t *testing.T) {

	response := types.UserReport{}

	ProcessUserReport(tasksSample, habitsSample, &response)

	if len(response.BadHabits) != 1 {
		t.Error("Error on bad habits")
	}
	if len(response.GoodHabits) != 2 {
		t.Error("Error on good habits")
	}
	if len(response.TodayTasks) != 1 {
		t.Error("Error on tasks for today")
	}
	if len(response.DelayedTasks) != 1 {
		t.Error("Error on delayed tasks")
	}
}

func TestProcessAdminTasks(t *testing.T) {

	response := types.AdminTasksReport{}

	ProcessAdminTasks(tasksSampleAdmin, &response)

	if response.Completed.Total != 2 {
		t.Error("Error on completed tasks total")
	}
	if response.Completed.Before != 1 {
		t.Error("Error on completed tasks before")
	}
	if response.Completed.After != 1 {
		t.Error("Error on completed tasks after")
	}
	if response.Available.Total != 6 {
		t.Error("Error on available tasks total")
	}
	if response.Available.Remaining != 4 {
		t.Error("Error on completed tasks before")
	}
	if response.Available.ForToday != 2 {
		t.Error("Error on completed tasks after")
	}
}

func TestProcessAdminHabitsWithoutNames(t *testing.T) {
	ranges := types.Ranges{}
	lowestTemp := types.UserHabitTemp{UserID: -1}
	highestTemp := types.UserHabitTemp{UserID: -1}

	ProcessAdminHabitsWithoutNames(habitsSampleAdmin, &ranges, &lowestTemp, &highestTemp)

	if ranges.Red != 1 {
		t.Error("Error on red habits qty")
	}
	if ranges.Orange != 1 {
		t.Error("Error on orange habits qty")
	}
	if ranges.Yellow != 1 {
		t.Error("Error on yellow habits qty")
	}
	if ranges.Green != 0 {
		t.Error("Error on green habits qty")
	}
	if ranges.Blue != 2 {
		t.Error("Error on blue habits qty")
	}
	if lowestTemp.HabitName != "Bully kids" && lowestTemp.UserID != 2 {
		t.Error("Error obtaining lowest score habit")
	}
	if highestTemp.HabitName != "Do Clean Code" && highestTemp.UserID != 1 {
		t.Error("Error obtaining highest score habit")
	}
}
