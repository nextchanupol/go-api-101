package model

import (
	"log"
	"time"

	pgsql12 "github.com/nextchanupol/go-api-101/pkg/database"
)

// Member struct
type Member struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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

// CreateMember creates a new member
func CreateMember(member *Member) (*Member, error) {
	createStatement := `insert into member 
												(id, first_name, last_name, email, password, created_at, updated_at)
											values 
												($1, $2, $3, $4, $5, $6, $7) returning *`

	statement, err := pgsql12.Db.Prepare(createStatement)
	defer statement.Close()

	if err != nil {
		log.Printf("Prepare error: %v", err)
		return nil, err
	}

	log.Printf("Prepare error: %v %v", member.CreatedAt, member.UpdatedAt)

	_, err = statement.Exec(member.ID, member.FirstName, member.LastName, member.Email, member.Password, member.CreatedAt, member.UpdatedAt)

	if err != nil {
		log.Printf("Exec error: %v", err)
		return nil, err
	}
	return member, nil
}
