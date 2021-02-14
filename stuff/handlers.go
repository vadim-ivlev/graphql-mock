package stuff

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// GIN handlers
func GophersGraphQLHandler(schema *graphql.Schema) gin.HandlerFunc {
	r := &relay.Handler{Schema: schema}
	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler(title, endpoint string) gin.HandlerFunc {
	h := playground.Handler(title, endpoint)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
