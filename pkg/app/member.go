package app

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	pgsql12 "github.com/nextchanupol/go-api-101/pkg/database"
	model "github.com/nextchanupol/go-api-101/pkg/modules/model/profile"
)

// CreateMember create a new member
func CreateMember(c echo.Context) (err error) {

	// m := new(model.Member)
	m := model.InitNewMember()

	if err := c.Bind(m); err != nil {
		return err
	}

	createStatement := `insert into member 
												(id, first_name, last_name, email, password)
											values 
												($1, $2, $3, $4, $5) returning *`

	statement, err := pgsql12.Db.Prepare(createStatement)
	defer statement.Close()

	if err != nil {
		log.Printf("Prepare error: %v", err)
		return err
	}

	log.Printf("Prepare error: %v %v", m.CreatedAt, m.UpdatedAt)

	_, err = statement.Exec(m.ID, m.FirstName, m.LastName, m.Email, m.Password)

	if err != nil {
		log.Printf("Exec error: %v", err)
		return err
	}

	return c.JSON(http.StatusCreated, m)
}
