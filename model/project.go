package model

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	UID          string `json:"uid" bson:"uid,omitempty"`
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"password"`
	TokenAddress string `json:"token_address" bson:"token_address"`
}

func NewProject() *Project {
	var u Project
	u.UID = uuid.New().String()
	u.Name = "from model new project"
	return &u
}

func (u *Project) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *Project) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
