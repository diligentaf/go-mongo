package project

import (
	"go-mongo/model"
)

type Store interface {
	Create(*model.Project) error
	Remove(field, value string) error
	GetByName(string) (*model.Project, error)
}
