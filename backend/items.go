package main

import "time"

type Item struct {
	ID int64 `json:"id"`
	Description string `json:"description"`
	Done bool `json:"done"`
	Author *User `json:"author"`
}

var items map[int64]*Item

func getItemByID(id int) *Item {
	item, ok := items[int64(id)]
	if !ok {
		return nil
	}
	return item
}

func createItem(desc string, authorID int) *Item {
	author := getUserByID(authorID)
	if author == nil {
		return nil
	}
	return &Item{
		Description: desc,
		Author: author,
	}
}

type User struct {
	ID int64
	Name string `json:"name"`
	Email string `json:"email"`
	BirthDate time.Time `json:"birthDate"`
}

var users map[int64]*User

func getUserByID(id int) *User {
	user, ok := users[int64(id)]
	if !ok {
		return nil
	}
	return user
}