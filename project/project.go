package project

import (
	"go-mongo/model"
)

type Store interface {
	Create(*model.Project) error
	Remove(field, value string) error
	Update(old *model.Project, new *model.Project) error
	UpdateProfile(u *model.Project) error

	GetByEmail(string) (*model.Project, error)

	RemoveTweet(u *model.Project, id *string) error
}
