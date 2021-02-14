package main

import (
	"fmt"
	"graphql-mock/stuff"
	"io/ioutil"
	"math/rand"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
)

var fakeDatabase = map[string]string{}

// represents a die
type randomDie struct{ numSides int32 }

func (r *randomDie) NumSides() int32 {
	return r.numSides
}

func (r *randomDie) RollOnce() int32 {
	return 1 + rand.Int31n(r.numSides)
}

func (r *randomDie) Roll(args struct{ NumRolls int32 }) []int32 {
	output := make([]int32, 0, args.NumRolls)
	for i := int32(0); i < args.NumRolls; i++ {
		output = append(output, r.RollOnce())
	}
	return output
}

type root struct{}

func (r *root) GetDie(args struct{ NumSides int32 }) *randomDie {
	return &randomDie{numSides: args.NumSides}
}

func (r *root) GetMessage(args struct{ Key string }) *string {
	s := fakeDatabase[args.Key]
	return &s
}

func (r *root) SetMessage(args struct{ Key, Message string }) *string {
	fakeDatabase[args.Key] = args.Message
	s := fakeDatabase[args.Key]
	return &s
}

// Construct a schema, using GraphQL schema language
var schemaBytes, _ = ioutil.ReadFile("server.graphql")
var schema = graphql.MustParseSchema(string(schemaBytes), &root{}, graphql.UseStringDescriptions())

func main() {
	fmt.Println("Running a GraphQL API server at http://localhost:4000/graphql")
	r := gin.New()
	r.POST("/graphql", stuff.GophersGraphQLHandler(schema))
	r.GET("/", stuff.PlaygroundHandler("GraphQL playground", "/graphql"))
	r.Run(":4000")
}
