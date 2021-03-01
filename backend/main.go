package main

import (
	goContext "context"
	"encoding/json"
	"github.com/arthurh0812/framework"
	"github.com/arthurh0812/framework/context"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
)

type Episode int

const (
	EpisodeJedi Episode = iota
	EpisodeEmpire
	EpisodeNewHope
)

func executeRequest(r GraphQLRequest, schema graphql.Schema, ctx goContext.Context) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Context: ctx,
		Schema: schema,
		RequestString: r.Query,
		VariableValues: r.Variables,
		OperationName: r.Operation,
	})
	if len(result.Errors) != 0 {
		log.Printf("invalid result: %v", result.Errors)
	}
	return result
}

func main() {
	app := framework.New()

	app.Register(http.MethodGet, "/hello", func(ctx *context.Context) {
		d := context.NewDirector(ctx)

		_, _ = d.WriteString("Hello World!")
	})

	app.Register(http.MethodGet, "/", func(ctx *context.Context) {
		var r GraphQLRequest
		err := json.NewDecoder(ctx.Request().Body).Decode(&r)
		defer ctx.Request().Body.Close()
		if err != nil {
			return
		}


		result := executeRequest(r, schema, ctx.Request().Context())
		err = json.NewEncoder(ctx).Encode(result)
		if err != nil {
			log.Fatal(err)
		}
	})

	err := app.Build()
	if err != nil {
		log.Fatal(err)
	}

	err = importJSONFromFile("./data.json", &data)
	if err != nil {
		log.Fatal(err)
	}

	err = app.ListenAndServe("localhost:10000")
	if err != nil {
		log.Fatal(err)
	}
}

func importJSONFromFile(filename string, result interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(result)
	if err != nil {
		return err
	}
	return nil
}