package components

import "github.com/graphql-go/graphql"

type Payment struct {
	Author 	string `json:"author"`
	Sum 	string `json:"sum"`
}

var PaymentGrObj = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PAYMENT_TABLE",
		Fields: graphql.Fields{
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"sum": &graphql.Field{
				Type: graphql.String,
			},
		},
	})