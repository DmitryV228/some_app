package model

import "time"

type User struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	IsAdmin      bool      `json:"isAdmin"`
	LastViewedAt time.Time `json:"lastViewedAt"`
}

func (u *User) SetLastViewedAt(t time.Time) {
	u.LastViewedAt = t
}
