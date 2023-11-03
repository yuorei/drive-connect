package resolver

import (
	"drive-connect-bff/application"
	"drive-connect-bff/client"
)

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app    *application.App
	client *client.Client
}

func NewResolver() *Resolver {
	return &Resolver{
		app:    application.NewApplication(),
		client: client.NewClient(),
	}
}
