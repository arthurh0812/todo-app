package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID string `json:"id" firestore:"id"`
	FirstName string `json:"firstName" firestore:"firstName,omitempty"`
	LastName string `json:"lastName" firestore:"lastName,omitempty"`
	Email string `json:"email" firestore:"email"`
	BirthDate time.Time `json:"birthDate" firestore:"birthdate"`
}

func GetUserByID(ctx context.Context, id string) (*User, error) {
	users := store.Collection("users")
	res := users.Where("id", "==", id)
	iter := res.Documents(ctx)
	user, err := iter.Next()
	if err != nil {
		return nil, err
	}
	return convertDocumentToUser(user)
}

func convertDocumentToUser(user *firestore.DocumentSnapshot) (*User, error) {
	id, err := extractDocumentFieldString(user, "id")
	if err != nil {
		return nil, err
	}
	firstName, err := extractDocumentFieldString(user, "firstName")
	if err != nil {
		return nil, err
	}
	lastName, err := extractDocumentFieldString(user, "lastName")
	if err != nil {
		return nil, err
	}
	birthDate, err := extractDocumentFieldTimestamp(user, "birthDate")
	if err != nil {
		return nil, err
	}
	email, err := extractDocumentFieldString(user, "email")
	if err != nil {
		return nil, err
	}
	return &User{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		BirthDate: birthDate,
		Email: email,
	}, nil
}

func extractDocumentFieldString(doc *firestore.DocumentSnapshot, field string) (string, error) {
	fieldValue, err := doc.DataAt(field)
	if err != nil {
		return "", err
	}
	// underlying firestore value should be string
	fieldString, ok := fieldValue.(string)
	if !ok {
		return "", fmt.Errorf("extracting string: field %q is not convertible to string", field)
	}
	return fieldString, nil
}

func extractDocumentFieldInt(doc *firestore.DocumentSnapshot, field string) (int64, error) {
	fieldValue, err := doc.DataAt(field)
	if err != nil {
		return 0, err
	}
	// underlying firestore value should be int64
	fieldInt64, ok := fieldValue.(int64)
	if !ok {
		return 0, fmt.Errorf("extracting int: value at field %q is not convertible to int64", field)
	}
	return fieldInt64, nil
}

func extractDocumentFieldTimestamp(doc *firestore.DocumentSnapshot, field string) (time.Time, error) {
	fieldValue, err := doc.DataAt(field)
	if err != nil {
		return time.Now(), err
	}
	// underlying firestore value should be time.Time
	fieldValueTime, ok := fieldValue.(time.Time)
	if !ok {
		return time.Now(), fmt.Errorf("extracting timestamp: value at field %q is not convertible to time.Time", field)
	}
	return fieldValueTime, nil
}

var ErrNoInputFirstName = errors.New("reading user input object: firstName property is missing")
var ErrNoInputLastName = errors.New("reading user input object: lastName property is missing")
var ErrNoInputEmail = errors.New("reading user input object: email property is missing")
var ErrNoInputBirthDate = errors.New("reading user input object: birthDate")


func convertInputObjectToUser(user map[string]interface{}) (*User, error) {
	firstName, ok := extractObjectFieldString(user, "firstName")
	if !ok {
		return nil, ErrNoInputFirstName
	}
	lastName, ok := extractObjectFieldString(user, "lastName")
	if !ok {
		return nil, ErrNoInputLastName
	}
	birthDateString, ok := extractObjectFieldString(user, "birthDate")
	if !ok {
		return nil, ErrNoInputBirthDate
	}
	birthDate, err := time.Parse(time.RFC3339, birthDateString)
	if err != nil {
		return nil, err
	}
	email, ok := extractObjectFieldString(user, "email")
	if !ok {
		return nil, ErrNoInputEmail
	}
	return &User{
		FirstName: firstName,
		LastName: lastName,
		BirthDate: birthDate,
		Email: email,
	}, nil
}

func extractObjectFieldString(object map[string]interface{}, field string) (string, bool) {
	fieldValue, ok := object[field]
	if !ok {
		return "", false
	}
	fieldString, ok := fieldValue.(string)
	if !ok {
		return "", false
	}
	return fieldString, true
}

func CreateUser(ctx context.Context, user map[string]interface{}) (*User, error) {
	convertedUser, err := convertInputObjectToUser(user)
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