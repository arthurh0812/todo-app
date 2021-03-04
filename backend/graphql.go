package main

import (
	"errors"
	"github.com/graphql-go/graphql"
	"log"
)

var ErrItemIDDoesNotExist = errors.New("querying item: error: the ID argument could not be found in the database")

var ErrItemIDArgumentNotSpecified = errors.New("reading query: error: no ID argument was specified")

var ErrInputItemNotSpecified = errors.New("reading mutation: error: no input item was specified")

var ErrInputItemIncomplete = errors.New("reading input item: error: the input item lacks some fields")

// GraphQLRequest with three components: "query", "operation" and "variables"
type GraphQLRequest struct {
	Query string `json:"query"`
	Mutation string `json:"mutation"`

	Variables map[string]interface{} `json:"variables"`
}

var graphqlUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"birthDate": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var itemInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ItemInput",
	Fields: graphql.Fields{
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"authorID": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var itemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Item",
	Fields: graphql.Fields{
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
		"author": &graphql.Field{
			Type: graphqlUser,
		},
	},
})

var itemQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "ItemQuery",
	Fields: graphql.Fields{
		"item": &graphql.Field{
			Type: itemType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
					DefaultValue: -1,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"]
				if !ok {
					return nil, ErrItemIDArgumentNotSpecified
				}
				return resolveItemID(id.(int))
			},
		},
	},
})

var itemMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "ItemMutation",
	Fields: graphql.Fields{
		"createItem": &graphql.Field{
			Type: itemType,
			Args: graphql.FieldConfigArgument{
				"item": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(itemInputType),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				item, ok := p.Args["item"]
				if !ok {
					return nil, ErrInputItemNotSpecified
				}
				log.Printf("%T", item)
				return resolveItem(item)
			},
		},
	},
})

func resolveItemID(id int) (interface{}, error) {
	item := getItemByID(id)
	if item == nil {
		return nil, ErrItemIDDoesNotExist
	}
	return item, nil
}

func resolveItem(item interface{}) (interface{}, error){
	return nil, errors.New("temporary error")
	//return createItem(description., authorID.(int)), nil
}

var itemSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: itemQuery,
	Mutation: itemMutation,
	Types: []graphql.Type{
		itemInputType,
	},
})