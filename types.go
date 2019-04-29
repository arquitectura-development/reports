package reports

type Habit struct {
	ID         int    `json:"ID"`
	UserID     int    `json:"UserID"`
	Name       string `json:"Name"`
	HabitType  int    `json:"HabitType"`
	Difficulty int    `json:"Difficulty"`
	Score      int    `json:"Score"`
}

type Task struct {
	ID          int    `json:"ID"`
	UserID      int    `json:"UserID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Difficulty  string `json:"Difficulty"`
	DueDate     string `json:"DueDate"`
	Reminder    string `json:"Reminder"`
	Done        bool   `json:"Done"`
}

type UserReport struct {
	GoodHabits   []Habit `json:"GoodHabits"`
	BadHabits    []Habit `json:"BadHabits"`
	TodayTasks   []Task  `json:"TodayTasks"`
	DelayedTasks []Task  `json:"DelayedTasks"`
}

type CompletedTasks struct {
	Total  int `json:"Total"`
	Before int `json:"Before"`
	After  int `json:"After"`
}

type AvailableTasks struct {
	Total     int `json:"Total"`
	Remaining int `json:"Remaining"`
	ForToday  int `json:"ForToday"`
}

type AdminTasksReport struct {
	Completed CompletedTasks `json:"Completed"`
	Delayed   int            `json:"Delayed"`
	Available AvailableTasks `json:"Available"`
}

type Ranges struct {
	Red    int `json:"Red"`
	Orange int `json:"Orange"`
	Yellow int `json:"Yellow"`
	Green  int `json:"Green"`
	Blue   int `json:"Blue"`
}

type HabitOwner struct {
	Name     string `json:"Name"`
	Username string `json:"Username"`
}

type AdminHabitsReport struct {
	PerRange   Ranges     `json:"PerRange"`
	WorstHabit HabitOwner `json:"WorstHabit"`
	BestHabit  HabitOwner `json:"BestHabit"`
}

var sampleUserReport = UserReport{
	GoodHabits: []Habit{
		{
			ID:         0,
			UserID:     0,
			Name:       "Good habit name",
			HabitType:  0,
			Difficulty: 0,
			Score:      0,
		},
	},
	BadHabits: []Habit{
		{
			ID:         0,
			UserID:     0,
			Name:       "Bad habit name",
			HabitType:  0,
			Difficulty: 0,
			Score:      0,
		},
	},
	TodayTasks: []Task{
		{
			ID:          0,
			UserID:      0,
			Title:       "Today task name",
			Description: "Today task description",
			Difficulty:  "Today task difficulty",
			DueDate:     "Today task due date",
			Reminder:    "Today task reminder",
			Done:        true,
		},
	},
	DelayedTasks: []Task{
		{
			ID:          0,
			UserID:      0,
			Title:       "Delayed task name",
			Description: "Delayed task description",
			Difficulty:  "Delayed task difficulty",
			DueDate:     "Delayed task due date",
			Reminder:    "Delayed task reminder",
			Done:        true,
		},
	},
}

var sampleAdminTasksReport = AdminTasksReport{
	Completed: CompletedTasks{
		Total:  0,
		Before: 0,
		After:  0,
	},
	Delayed: 0,
	Available: AvailableTasks{
		Total:     0,
		Remaining: 0,
		ForToday:  0,
	},
}

var sampleAdminHabitsReport = AdminHabitsReport{
	PerRange: Ranges{
		Red:    0,
		Orange: 0,
		Yellow: 0,
		Green:  0,
		Blue:   0,
	},
	WorstHabit: HabitOwner{
		Name:     "Name Worst",
		Username: "username_worst",
	},
	BestHabit: HabitOwner{
		Name:     "Name Best",
		Username: "username_best",
	},
}
