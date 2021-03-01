package main

import (
	"errors"
	"github.com/graphql-go/graphql"
)

var ErrUserIDDoesNotExist = errors.New("querying user: error: the ID argument could not be found in the database")

var ErrDroidModelDoesNotExist = errors.New("querying droid: error: the model argument could not be found in the database")

// GraphQLRequest with three components: "query", "operation" and "variables"
type GraphQLRequest struct {
	Query string `json:"query"`
	Operation string `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

var StringField = func() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
	}
}

var episodeType = graphql.NewEnum(graphql.EnumConfig{
	Name: "Episode",
	Values: graphql.EnumValueConfigMap{
		"JEDI": &graphql.EnumValueConfig{
			Value: EpisodeJedi,
		},
		"EMPIRE": &graphql.EnumValueConfig{
			Value: EpisodeEmpire,
		},
		"NEW_HOPE": &graphql.EnumValueConfig{
			Value: EpisodeNewHope,
		},
	},
})

var modelType = graphql.NewEnum(graphql.EnumConfig{
	Name: "Model",
	Values: graphql.EnumValueConfigMap{
		"C3-PO": &graphql.EnumValueConfig{
			Value: DroidModelC3PO,
		},
		"R2-D2": &graphql.EnumValueConfig{
			Value: DroidModelR2D2,
		},
	},
})

var droidType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Droid",
	Fields: graphql.Fields{
		"name": StringField(),
		"model": &graphql.Field{
			Type: modelType,
		},
	},
})

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": StringField(),
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"episode": &graphql.Field{
			Type: episodeType,
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
					DefaultValue: "1",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(string)
				if !ok {
					return nil, ErrUserIDDoesNotExist
				}
				user := getUserByID(id)
				if user == nil {
					return nil, ErrUserIDDoesNotExist
				}
				return user, nil
			},
		},
		"droid": &graphql.Field{
			Type: droidType,
			Args: graphql.FieldConfigArgument{
				"model": &graphql.ArgumentConfig{
					Type: modelType,
					DefaultValue: "C3-PO",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				model, ok := p.Args["model"]
				if !ok {
					return nil, ErrDroidModelDoesNotExist
				}
				droidModel := DroidModel(model.(string))
				if droid := getDroidByModel(droidModel); droid != nil {
					return droid, nil
				}
				return nil, ErrDroidModelDoesNotExist
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})