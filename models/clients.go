package models

import "time"

type Client struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
}
