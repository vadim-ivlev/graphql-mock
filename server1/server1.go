package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var fakeDatabase = map[string]string{}

type root struct{}

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
var schemaBytes, _ = ioutil.ReadFile("server1.graphql")
var schema = graphql.MustParseSchema(string(schemaBytes), &root{}, graphql.UseStringDescriptions())

// GIN handler
func ginGraphQLHandler() gin.HandlerFunc {
	r := &relay.Handler{Schema: schema}
	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	fmt.Println("Running a GraphQL API server at http://localhost:4000/graphql")
	r := gin.New()
	r.POST("/graphql", ginGraphQLHandler())
	r.Run(":4000")
}
