package handler

import (
	"go-mongo/project"
)

type Handler struct {
	projectStore project.Store
}

func NewHandler(pj project.Store) *Handler {
	return &Handler{
		projectStore: pj,
	}
}
