package main

import (
	"fmt"
	"graphql-mock/stuff"
	"io/ioutil"
	"math/rand"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// represents a die
type randomDie struct{ numSides int32 }

// constructor
func newRandomDie(numSides int32) *randomDie {
	rd := randomDie{numSides: numSides}
	return &rd
}

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

func (*root) GetDie(args struct{ NumSides int32 }) *randomDie {
	return newRandomDie(args.NumSides)
}

// Construct a schema, using GraphQL schema language
var schemaBytes, _ = ioutil.ReadFile("server2.graphql")
var schema = graphql.MustParseSchema(string(schemaBytes), &root{}, graphql.UseStringDescriptions())

// GIN handler
func gopherGraphQLHandler() gin.HandlerFunc {
	r := &relay.Handler{Schema: schema}
	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// r := gin.New()
	r := gin.Default()
	r.POST("/graphql", gopherGraphQLHandler())
	r.GET("/", stuff.PlaygroundHandler("GraphQL playground", "/graphql"))

	fmt.Println(`
	Running a GraphQL playground http://localhost:4000/
	GraphQL end point            http://localhost:4000/graphql
	`)

	r.Run(":4000")
}
