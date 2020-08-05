package model

import "time"

// Member struct
type Member struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Members type list of member
type Members []Member

// InitNewMember call for iniial member data
func InitNewMember() *Member {
	now := time.Now()
	return &Member{
		CreatedAt: now,
		UpdatedAt: now,
	}
}
