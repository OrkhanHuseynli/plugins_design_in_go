package graphqlapi

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"sample_graphql_in_go/src/graphqlapi/components"
)


func NewGraphQLProcessor () *GraphQLProcessor {
	return &GraphQLProcessor{}
}

type GraphQLProcessor struct {

}

func (p GraphQLProcessor) ExecuteQuery(query string){
	// Schema
	fields := graphql.Fields{
		"payment": components.GetPaymentField(),
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	log.Println("Getting to Query")

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphqlapi operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}
}