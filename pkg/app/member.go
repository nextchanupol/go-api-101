package app

import (
	"net/http"

	"github.com/labstack/echo"
	model "github.com/nextchanupol/go-api-101/pkg/modules/model/profile"
)

// CreateMember create a new member
func CreateMember(c echo.Context) (err error) {

	// m := new(model.Member)
	m := model.InitNewMember()

	if err := c.Bind(m); err != nil {

		return c.JSON(http.StatusInternalServerError, err)
	}

	m, err = model.CreateMember(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &m)
}

func GetMembers(c echo.Context) (err error) {

	list, err := model.GetMembers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &list)
}
