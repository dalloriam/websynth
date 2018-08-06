package app

import (
	"fmt"
	"net/http"

	"github.com/dalloriam/graphkit"
	"github.com/dalloriam/websynth/app/resolver"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

const (
	gqlRoute        = "/synth"
	host            = "0.0.0.0:8080"
	schemaDirectory = "./schema"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	schema, err := graphkit.LoadSchema(schemaDirectory)

	if err != nil {
		return err
	}

	synth, err := getSynth()
	if err != nil {
		return err
	}

	schemaRes, err := graphql.ParseSchema(schema.Raw, resolver.New(synth))
	if err != nil {
		return err
	}
	http.Handle(gqlRoute, &relay.Handler{Schema: schemaRes})

	fmt.Printf("Running on %s\n", host)
	return http.ListenAndServe(host, nil)
}
