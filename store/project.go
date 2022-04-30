package store

import (
	"context"

	"go-mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectStore struct {
	db *mongo.Collection
}

func NewProjectStore(db *mongo.Collection) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

func (us *ProjectStore) Create(u *model.Project) error {
	_, err := us.db.InsertOne(context.TODO(), u)
	return err
}

func (us *ProjectStore) Remove(field, value string) error {
	_, err := us.db.DeleteOne(context.TODO(), bson.M{field: value})
	return err
}

func (us *ProjectStore) Update(old *model.Project, new *model.Project) error {
	var err error
	if old.Projectname != new.Projectname {
		err = us.Remove("_id", old.Projectname)
	} else if old.Email != new.Email {
		err = us.Remove("email", old.Email)
	} else if old.Password != new.Password {
		err = us.Remove("password", old.Password)
	}
	err = us.Create(new)
	return err
}

func (us *ProjectStore) UpdateProfile(u *model.Project) error {
	_, err := us.db.UpdateOne(context.TODO(),
		bson.M{"_id": u.Projectname},
		bson.M{"$set": bson.M{
			"name":            u.Name,
			"bio":             u.Bio,
			"profile_picture": u.ProfilePicture,
			"header_picture":  u.HeaderPicture,
		},
		})
	return err
}

func (us *ProjectStore) GetByEmail(email string) (*model.Project, error) {
	var u model.Project
	err := us.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&u)
	return &u, err
}

func (us *ProjectStore) GetByProjectname(username string) (*model.Project, error) {
	var u model.Project
	err := us.db.FindOne(context.TODO(), bson.M{"_id": username}).Decode(&u)
	return &u, err
}

func (us *ProjectStore) RemoveTweet(u *model.Project, id *string) error {
	oid, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return err
	}
	newTweets := &[]primitive.ObjectID{}
	for _, tid := range *u.Tweets {
		if tid != oid {
			*newTweets = append(*newTweets, tid)
		}
	}
	_, err = us.db.UpdateOne(context.TODO(), bson.M{"_id": u.Projectname}, bson.M{"$set": bson.M{"tweets": newTweets}})
	return err
}

func (us *ProjectStore) GetProjectListFromProjectnameList(usernames []string) (*[]model.Project, error) {
	var users []model.Project
	query := bson.M{"_id": bson.M{"$in": usernames}}
	res, err := us.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return &users, err
}

func (us *ProjectStore) GetTweetIdListFromProjectnameList(usernames []string) (*[]primitive.ObjectID, error) {
	var users []model.Project
	query := bson.M{"_id": bson.M{"$in": usernames}}
	res, err := us.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	var tweetsId []primitive.ObjectID
	for _, user := range users {
		tweetsId = append(tweetsId, *user.Tweets...)
	}
	return &tweetsId, err
}
