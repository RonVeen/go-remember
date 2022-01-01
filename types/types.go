package types

import "go-remember/persistence"

type AppConfig struct {
	Repo persistence.ReminderRepository
	Port int
}
