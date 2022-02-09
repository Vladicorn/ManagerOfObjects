package repo

import "time"

type Object struct {
	Id        uint
	Name      string
	Start     *time.Time
	Url       string
	PriceKeep int
	PriceNeed int
	PriceSum  int
	//Tasks     Task
}
