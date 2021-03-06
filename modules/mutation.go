package modules

import (
	"golang-graphql-talk/modules/post/resolvers"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Mutation",
	Description: "Mutation do graphQL",
	Fields: graphql.Fields{
		"insertPost": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"message": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve:     resolvers.InsertPostResolver,
			Description: "Adiciona um novo post",
		},
	},
})
