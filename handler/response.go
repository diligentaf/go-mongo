package handler

import (
	"go-mongo/model"
)

type projectResponse struct {
	Project struct {
		Name         string `json:"username" bson:"_id"`
		Email        string `json:"email"`
		TokenAddress string `json:"token_address"`
	} `json:"user"`
}

func newProjectResponse(u *model.Project) *projectResponse {
	r := new(projectResponse)
	r.Project.Name = u.Name
	r.Project.Email = u.Email
	r.Project.TokenAddress = u.TokenAddress
	return r
}
