package handler

import (
	"go-mongo/model"

	"github.com/labstack/echo/v4"
)

// Registration request
type projectRegisterRequest struct {
	Project struct {
		Name         string `json:"name" validate:"required"`
		Email        string `json:"email" validate:"required,email"`
		Password     string `json:"password" validate:"required"`
		TokenAddress string `json:"token_address" validate:"required"`
	}
}

func (r *projectRegisterRequest) bind(c echo.Context, u *model.Project) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Name = r.Project.Name
	u.Email = r.Project.Email
	u.TokenAddress = r.Project.TokenAddress
	h, err := u.HashPassword(r.Project.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}
