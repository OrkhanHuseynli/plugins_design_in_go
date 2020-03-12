package components

import (
	"github.com/graphql-go/graphql"
)

func GetPaymentField() *graphql.Field {
	resolver := Resolver{}
	return &graphql.Field{
		Type: PaymentGrObj,
		Args: graphql.FieldConfigArgument{
			"author": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.DefaultResolver,
	}
}
