// Running GraphQL API server
var fs = require('fs')
var express = require('express')
var { graphqlHTTP } = require('express-graphql')
var { buildSchema } = require('graphql')
var cors = require('cors')

var fakeDatabase = { k0: 'Hello World!' }

// The root provides the top-level API endpoints
var root = {
    getMessage: ({ key }) => {
        return fakeDatabase[key]
    },
    setMessage: ({ key, message }) => {        
        return fakeDatabase[key] = message
    }
}

// Construct a schema, using GraphQL schema language
var schema = buildSchema(fs.readFileSync('server1.graphql', 'utf8'))

var app = express()
app.use(cors())
app.use( '/graphql', graphqlHTTP({ 
        schema: schema, 
        rootValue: root, 
        graphiql: true, 
    })
)
console.log('Running a GraphQL API server at http://localhost:4000/graphql')
app.listen(4000)
