package types

type Habit struct {
	ID         int    `json:"id"`
	UserID     int    `json:"userID"`
	Name       string `json:"name"`
	HabitType  int    `json:"habitType"`
	Difficulty int    `json:"difficulty"`
	Score      int    `json:"score"`
	Color      int    `json:"color"`
}

type Task struct {
	ID             int    `json:"id"`
	UserID         int    `json:"userID"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Difficulty     string `json:"difficulty"`
	DueDate        string `json:"dueDate"`
	CompletionDate string `json:"completionDate"`
	Reminder       string `json:"reminder"`
	Done           bool   `json:"done"`
}

type UserReport struct {
	GoodHabits   []Habit `json:"goodHabits"`
	BadHabits    []Habit `json:"badHabits"`
	TodayTasks   []Task  `json:"todayTasks"`
	DelayedTasks []Task  `json:"delayedTasks"`
}

type CompletedTasks struct {
	Total  int `json:"total"`
	Before int `json:"before"`
	After  int `json:"after"`
}

type AvailableTasks struct {
	Total     int `json:"total"`
	Remaining int `json:"remaining"`
	ForToday  int `json:"forToday"`
}

type AdminTasksReport struct {
	Completed CompletedTasks `json:"completed"`
	Delayed   int            `json:"delayed"`
	Available AvailableTasks `json:"available"`
}

type Ranges struct {
	Red    int `json:"red"`
	Orange int `json:"orange"`
	Yellow int `json:"yellow"`
	Green  int `json:"green"`
	Blue   int `json:"blue"`
}

type HabitOwner struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type AdminHabitsReport struct {
	PerRange   Ranges     `json:"perRange"`
	WorstHabit HabitOwner `json:"worstHabit"`
	BestHabit  HabitOwner `json:"bestHabit"`
}

type UserData struct {
	Success bool   `json:"success"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}

type UserHabitTemp struct {
	UserID    int    `json:"userId"`
	HabitName string `json:"habitName"`
}
