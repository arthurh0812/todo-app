package main

import (
	"cloud.google.com/go/firestore"
	goContext "context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"github.com/arthurh0812/framework"
	"github.com/arthurh0812/framework/context"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/graphql-go/graphql"
)

var store *firestore.Client

var userSchema graphql.Schema

func init() {
	ctx := goContext.Background()
	sak := option.WithCredentialsFile("./service-account-key.json")

	firebaseApp, err := firebase.NewApp(ctx, nil, sak)
	if err != nil {
		log.Fatalf("failed to start the firebase application: %v", err)
	}

	store, err = firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("failed to connect to firestore: %v", err)
	}

	userSchema, err = UserSchema()
	if err != nil {
		log.Fatalf("failed to create user schema: %v", err)
	}
}

func executeRequest(r GraphQLRequest, schema graphql.Schema, ctx goContext.Context) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Context: ctx,
		Schema: schema,
		RequestString: r.Query,
		VariableValues: r.Variables,
	})
	return result
}

func main() {
	app := framework.New()
	defer store.Close()

	miller := &User{
		FirstName: "Mike",
		LastName: "Miller",
		BirthDate: time.Now().Add(-20 * 365 * 24 * time.Hour),
		Email: "mike.miller@web.de",
		ID: "hello",
	}

	_, err := store.Collection("users").Doc("0").Set(goContext.Background(), miller)
	if err != nil {
		log.Fatalf("failed to write data to firestore: %v", err)
	}

	app.Register(http.MethodGet, "/hello", func(ctx *context.Context) {
		d := context.NewDirector(ctx)

		_, _ = d.WriteString("Hello World!\n")

		query := store.Collection("users").Where("firstName", "==", "Mike")
		users := query.Documents(goContext.Background())
		for next, err := users.Next(); next != nil && err == nil; {
			for k, v := range next.Data() {
				d.Writef("%v: %v\n", k, v)
			}
			break
		}
	})

	app.Register(http.MethodGet, "/items", func(ctx *context.Context) {
		var r GraphQLRequest
		err := json.NewDecoder(ctx.Request().Body).Decode(&r)
		defer ctx.Request().Body.Close()
		if err != nil {
			return
		}

		result := executeRequest(r, itemSchema, ctx.Request().Context())
		err = json.NewEncoder(ctx).Encode(result)
		if err != nil {
			log.Fatal(err)
		}
	})

	app.Register(http.MethodGet, "/users", func(ctx *context.Context) {
		d := context.NewDirector(ctx)
		var r GraphQLRequest

		err := json.NewDecoder(ctx.Body()).Decode(&r)
		if err != nil {
			d.Errorf(http.StatusInternalServerError, "error: reading request body: %v", err)
		}

		res := executeRequest(r, userSchema, goContext.Background())
		err = json.NewEncoder(ctx).Encode(res)
		if err != nil {
			log.Fatalf("error: encoding JSON result: %v", err)
		}
	})

	err = app.Build()
	if err != nil {
		log.Fatal(err)
	}

	err = importItemsFromJSON("./items.json", &items)
	if err != nil {
		log.Fatal(err)
	}

	err = app.ListenAndServe("localhost:10000")
	if err != nil {
		log.Fatal(err)
	}
}

func importItemsFromJSON(filename string, result interface{}) error {
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