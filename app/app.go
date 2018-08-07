package app

import (
	"fmt"
	"net/http"

	"github.com/dalloriam/graphkit"
	"github.com/dalloriam/websynth/app/audio"
	"github.com/dalloriam/websynth/app/resolver"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type App struct {
	cfg Config
}

func New(cfg Config) *App {
	return &App{cfg}
}

func (a *App) Run() error {
	schema, err := graphkit.LoadSchema(a.cfg.SchemaDirectory)

	if err != nil {
		return err
	}

	audioSys, err := audio.NewSystem(a.cfg.Audio)
	if err != nil {
		return err
	}

	schemaRes, err := graphql.ParseSchema(schema.Raw, resolver.New(audioSys))
	if err != nil {
		return err
	}
	http.Handle(a.cfg.GQLRoute, &relay.Handler{Schema: schemaRes})

	fmt.Printf("Running on http://%s%s\n", a.cfg.Host, a.cfg.GQLRoute)
	return http.ListenAndServe(a.cfg.Host, nil)
}
