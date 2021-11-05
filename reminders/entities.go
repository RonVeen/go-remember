package reminders

type (
	Reminder struct {
		UID 	  string  `json:"id"`
		Title 	  string  `json:"title"`
		Comment   string  `json:"comment"`
		Completed bool    `json:"completed"`
	}

	Reminders []Reminder
)