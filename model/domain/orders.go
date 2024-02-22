package domain

import "time"

type Orders struct {
	OrderId   int
	Total     int
	OrderDate time.Time
}
