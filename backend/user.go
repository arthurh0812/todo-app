package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"time"
)

type User struct {
	ID string `json:"id" firestore:"id"`
	FirstName string `json:"firstName" firestore:"firstName,omitempty"`
	LastName string `json:"lastName" firestore:"lastName,omitempty"`
	Email string `json:"email" firestore:"email"`
	BirthDate time.Time `json:"birthDate" firestore:"birthdate"`
}

func GetUserByID(ctx context.Context, id string) *User {
	users := store.Collection("users")
	res := users.Where("id", "==", id)
	iter := res.Documents(ctx)
	user, err := iter.Next()
	if err != nil {
		return nil
	}
	return convertDocumentToUser(user)
}

func convertDocumentToUser(user *firestore.DocumentSnapshot) *User {
	id, _ := user.DataAt("id")
	firstName, _ := user.DataAt("firstName")
	lastName, _ := user.DataAt("lastName")
	birthDate, _ := user.DataAt("birthdate")
	email, _ := user.DataAt("email")
	return &User{
		ID: id.(string),
		FirstName: firstName.(string),
		LastName: lastName.(string),
		BirthDate: birthDate.(time.Time),
		Email: email.(string),
	}
}

var ErrNoInputFirstName = errors.New("reading user input object: error: firstName property is missing")
var ErrNoInputLastName = errors.New("reading user input object: error: lastName property is missing")
var ErrNoInputEmail = errors.New("reading user input object: error: email property is missing")
var ErrNoInputBirthDate = errors.New("reading user input object: error: birthDate")


func convertMapToUser(user map[string]interface{}) (*User, error) {
	firstName, _ := user["firstName"]
	if firstName == nil {
		return nil, ErrNoInputFirstName
	}
	lastName, _ := user["lastName"]
	if lastName == nil {
		return nil, ErrNoInputLastName
	}
	birthDateString, _ := user["birthDate"]
	if birthDateString == nil {
		return nil, ErrNoInputBirthDate
	}
	birthDate, err := time.Parse(time.RFC3339, birthDateString.(string))
	if err != nil {
		return nil, err
	}
	email, _ := user["email"]
	if email == nil {
		return nil, ErrNoInputEmail
	}
	return &User{
		FirstName: firstName.(string),
		LastName: lastName.(string),
		BirthDate: birthDate,
		Email: email.(string),
	}, nil
}

func CreateUser(ctx context.Context, user map[string]interface{}) (*User, error) {
	convertedUser, err := convertMapToUser(user)
	if err != nil {
		return nil, err
	}
	users := store.Collection("users")
	ref, _, err := users.Add(ctx, convertedUser)
	if err != nil {
		return nil, err
	}
	_, err = ref.Update(ctx, []firestore.Update{
		firestore.Update{
			Path: "id",
			Value: ref.ID,
		},
	})
	convertedUser.ID = ref.ID
	return convertedUser, err
}