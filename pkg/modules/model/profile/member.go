package model

import (
	"fmt"
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

type MemberData struct {
	List []*Member
}

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

	_, err = statement.Exec(member.ID, member.FirstName, member.LastName, member.Email, member.Password, member.CreatedAt, member.UpdatedAt)

	if err != nil {
		log.Printf("Exec error: %v", err)
		return nil, err
	}
	return member, nil
}

func GetMembers() ([]*Member, error) {
	selectStatement := `SELECT id, first_name, last_name, email, created_at, updated_at FROM member`
	statement, err := pgsql12.Db.Prepare(selectStatement)
	defer statement.Close()

	if err != nil {
		log.Printf("Prepare error: %v", err)
		return nil, err
	}

	rows, err := statement.Query()
	defer rows.Close()

	if err != nil {
		log.Printf("Query error: %v", err)
		return nil, err
	}

	var members []*Member
	for rows.Next() {

		var (
			id        string
			firstName string
			lastName  string
			email     string
			createdAt time.Time
			updatedAt time.Time
		)

		err = rows.Scan(&id, &firstName, &lastName, &email, &createdAt, &updatedAt)
		if err != nil {
			fmt.Printf("rows.Scan error: %v\n", err)
			return nil, err
		}

		m := &Member{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		members = append(members, m)

	}
	return members, nil
}
