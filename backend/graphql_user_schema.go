package main

import (
	"context"
	"github.com/graphql-go/graphql"
)

var fieldArgumentID = &graphql.ArgumentConfig{
	Type:         graphql.String,
	DefaultValue: "",
}

var userInputTypeConfig = graphql.InputObjectConfig{
	Name: "UserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"firstName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
			DefaultValue: "",
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
			DefaultValue: "",
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
			DefaultValue: "",
		},
		"birthDate": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
			DefaultValue: "January 1, 2000 at 0:00:00 AM UTC+1",
		},
	},
}

var userTypeConfig = graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"birthDate": &graphql.Field{
			Type: graphql.String,
		},
	},
}

var userType = graphql.NewObject(userTypeConfig)

var queryUserConfig = graphql.ObjectConfig{
	Name: "MyQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": fieldArgumentID,
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(string)
				return GetUserByID(context.Background(), id), nil
			},
		},
	},
}

var mutationUserConfig = graphql.ObjectConfig{
	Name: "MyMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewInputObject(userInputTypeConfig),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				user, _ := params.Args["user"]
				return CreateUser(context.Background(), user.(map[string]interface{}))
			},
		},
	},
}

var userSchemaConfig = graphql.SchemaConfig{
	Query:    graphql.NewObject(queryUserConfig),
	Mutation: graphql.NewObject(mutationUserConfig),
}

func UserSchema() (graphql.Schema, error) {
	return graphql.NewSchema(userSchemaConfig)
}
