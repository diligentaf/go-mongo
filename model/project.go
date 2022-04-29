package model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Name        string `json:"name" bson:"name"`
	Projectname string `json:"projectname" bson:"_id"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password" bson:"password"`

	// Fluff shown to project as profile
	Bio            string `json:"bio" bson:"bio"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
	HeaderPicture  string `json:"header_picture" bson:"header_picture"`

	Tweets *[]primitive.ObjectID `json:"tweets" bson:"tweets"`
}

func NewProject() *Project {
	var u Project
	u.Name = "Twitter Project "
	u.Tweets = &[]primitive.ObjectID{}
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
