package repo

import (
	"time"
)

type Task struct {
	Id          uint
	Name        string
	Start       *time.Time
	End         *time.Time
	Description string
}
