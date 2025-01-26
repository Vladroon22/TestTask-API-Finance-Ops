package entity

import "time"

type User struct {
	ID      int
	Balance float64
}

type Tx struct {
	Sender_name   string
	Receiver_name string
	Amount        float64
	Type          string
	CreatedAt     time.Time
}
