package domain

import "time"

type Game struct {
	Id          int       `json:"id"`
	Email       string    `json:"email"`
	Board       Board     `json:"board"`
	Over        bool      `json:"over"`
	CreatedAt   time.Time `json:"elapsedTime"`
	FinishedAt  time.Time `json:"finishedAt"`
	ElapsedTime int       `json:"elapsedTime"`
}
