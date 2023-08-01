package services

import (
	"context"
	"errors"
	"ginMongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection //Poineter to collection
	ctx context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService  {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx: ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	// Get the user in the request and add it to mongodb

	// Initialed in the constractor
	_ , err := u.usercollection.InsertOne(u.ctx , user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error)  {
	var user *models.User

	// Write the mongo query
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx , query).Decode(&user)
	return user , err
}

func (u *UserServiceImpl) GetAll() ([]*models.User , error) {
	// create slice object 
	var users []*models.User 
	cursor, err := u.usercollection.Find(u.ctx , bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
		  return nil, err	
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err	
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name" , Value: user.Name} , bson.E{Key: "user_address" , Value: user.Address} , bson.E{Key: "user_age" , Value: user.Age}}}}
	
	result, _ := u.usercollection.UpdateOne(u.ctx , filter , update)
	if result.MatchedCount != 1 {
		return errors.New("no match documents found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	// Write the mongo query
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx , filter)
	if result.DeletedCount != 1 {
		return errors.New("no match documents found for delete")
	}
	return nil
}