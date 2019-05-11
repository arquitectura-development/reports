package types

type Habit struct {
	ID         int    `json:"ID"`
	UserID     int    `json:"UserID"`
	Name       string `json:"Name"`
	HabitType  int    `json:"HabitType"`
	Difficulty int    `json:"Difficulty"`
	Score      int    `json:"Score"`
}

type Task struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	DueDate     string `json:"dueDate"`
	Reminder    string `json:"reminder"`
	Done        bool   `json:"done"`
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