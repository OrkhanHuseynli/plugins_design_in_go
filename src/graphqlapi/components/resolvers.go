package components

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

type Resolver struct {

}

func (r *Resolver) DefaultResolver(p graphql.ResolveParams) (interface{}, error) {
	log.Println("Calling default resolver")
	author, ok := p.Args["author"].(string)
	fmt.Println(author)
	if ok {
		log.Println("The call was ok")
		res := Payment{"Mars", "Bars"}
		return &res, nil
	}
	return nil, nil
}
