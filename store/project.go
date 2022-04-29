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

func (us *ProjectStore) AddFollower(u *model.Project, follower *model.Project) error {
	*u.Followers = append(*u.Followers, *model.NewOwner(follower.Projectname, follower.ProfilePicture, follower.Name, follower.Bio))
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"_id": u.Projectname}, bson.M{"$set": bson.M{"followers": u.Followers}})
	if err != nil {
		return err
	}
	*follower.Followings = append(*follower.Followings, *model.NewOwner(u.Projectname, u.ProfilePicture, u.Name, u.Bio))
	_, err = us.db.UpdateOne(context.TODO(), bson.M{"_id": follower.Projectname}, bson.M{"$set": bson.M{"followings": follower.Followings}})
	if err != nil {
		return err
	}
	return nil
}

func (us *ProjectStore) RemoveFollower(u *model.Project, follower *model.Project) error {
	newFollowers := &[]model.Owner{}
	for _, o := range *u.Followers {
		if o.Projectname != follower.Projectname {
			*newFollowers = append(*newFollowers, o)
		}
	}
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"_id": u.Projectname}, bson.M{"$set": bson.M{"followers": newFollowers}})
	if err != nil {
		return err
	}
	u.Followers = newFollowers

	newFollowings := &[]model.Owner{}
	for _, o := range *follower.Followings {
		if o.Projectname != u.Projectname {
			*newFollowings = append(*newFollowings, o)
		}
	}
	_, err = us.db.UpdateOne(context.TODO(), bson.M{"_id": follower.Projectname}, bson.M{"$set": bson.M{"followings": newFollowings}})
	if err != nil {
		return err
	}
	follower.Followings = newFollowings
	return nil
}

func (us *ProjectStore) IsFollower(username, followerProjectname string) (bool, error) {
	u, err := us.GetByProjectname(username)
	if err != nil {
		return false, err
	}
	follower, err := us.GetByProjectname(followerProjectname)
	if err != nil {
		return false, nil
	}
	doesFollow := false
	for _, o := range *u.Followers {
		if o.Projectname == follower.Projectname {
			doesFollow = true
			break
		}
	}
	hasInFollowings := false
	for _, o := range *follower.Followings {
		if o.Projectname == u.Projectname {
			hasInFollowings = true
			break
		}
	}
	return doesFollow && hasInFollowings, nil
}

func (us *ProjectStore) AddTweet(u *model.Project, t *model.Tweet) error {
	*u.Tweets = append(*u.Tweets, t.ID)
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"_id": u.Projectname}, bson.M{"$set": bson.M{"tweets": u.Tweets}})
	return err
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

func (us *ProjectStore) AddLog(u *model.Project, e *model.Event) error {
	*u.Logs = append(*u.Logs, *e)
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"_id": u.Projectname}, bson.M{"$set": bson.M{"logs": u.Logs}})
	if err != nil {
		return err
	}
	return nil
}

func (us *ProjectStore) AddNotification(u *model.Project, e *model.Event) error {
	*u.Notifications = append(*u.Notifications, *e)
	_, err := us.db.UpdateOne(context.TODO(), bson.M{"_id": u.Projectname}, bson.M{"$set": bson.M{"notifications": u.Notifications}})
	if err != nil {
		return err
	}
	return nil
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

func (us *ProjectStore) GetProjectnameSearchResult(username string) (*[]model.Owner, error) {
	var users []model.Project
	reg := "^" + username // usernames that starts with "query"
	query := bson.M{"_id": bson.M{"$regex": reg}}
	res, err := us.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	var result []model.Owner
	for _, user := range users {
		result = append(result, model.Owner{
			Projectname:    user.Projectname,
			ProfilePicture: user.ProfilePicture,
			Name:           user.Name,
			Bio:            user.Bio,
		})
	}
	return &result, err
}
