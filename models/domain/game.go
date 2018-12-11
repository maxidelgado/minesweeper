package domain

import "time"

type Game struct {
	Email       string        `json:"email"`
	Board       Board         `json:"board"`
	Over        bool          `json:"over"`
	ElapsedTime time.Duration `json:"elapsedTime"`
}
