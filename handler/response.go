package handler

import (
	"go-mongo/model"
)

type projectResponse struct {
	Project struct {
		Name  string `json:"username" bson:"_id"`
		Email string `json:"email"`
		Bio   string `json:"bio"`
		Token string `json:"token"`
	} `json:"user"`
}

func newProjectResponse(u *model.Project) *projectResponse {
	r := new(projectResponse)
	r.Project.Name = u.Name
	r.Project.Email = u.Email
	r.Project.Name = u.Name
	return r
}
