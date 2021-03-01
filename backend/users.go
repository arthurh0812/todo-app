package main

type user struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Episode Episode `json:"episode"`
}

// makeshift database
var data map[string]*user

func getUserByID(id string) *user {
	return data[id]
}
