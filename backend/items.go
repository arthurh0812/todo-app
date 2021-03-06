package main

import "context"

type Item struct {
	ID int64 `json:"id"`
	Description string `json:"description"`
	Done bool `json:"done"`
}

var items map[int64]*Item

func getItemByID(id int) *Item {
	item, ok := items[int64(id)]
	if !ok {
		return nil
	}
	return item
}

func createItem(desc string, authorID string) *Item {
	author := GetUserByID(context.Background(), authorID)
	if author == nil {
		return nil
	}
	return &Item{
		Description: desc,
	}
}
