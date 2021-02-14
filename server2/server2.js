// Running GraphQL API server
// npm install express express-graphql graphql cors
// https://graphql.org/graphql-js/object-types/

var fs = require('fs')
var express = require('express')
var { graphqlHTTP } = require('express-graphql')
var { buildSchema } = require('graphql')
var cors = require('cors')


// Represents a die
class RandomDie {
    constructor (numSides) {
        this.numSides = numSides
    }

    rollOnce () {
        return 1 + Math.floor(Math.random() * this.numSides)
    }

    roll ({ numRolls }) {
        var output = []
        for (var i = 0; i < numRolls; i++) {
            output.push(this.rollOnce())
        }
        return output
    }
}

// The root provides the top-level API endpoints
var root = {
    getDie: ({ numSides } ) => {
       return new RandomDie(numSides)
    },
}

// Construct a schema, using GraphQL schema language
var schema = buildSchema(fs.readFileSync('server2.graphql', 'utf8'))

var app = express()
app.use(cors())
app.use( '/graphql', graphqlHTTP({
        schema: schema,
        rootValue: root,
        graphiql: true
    })
)
console.log('Running a GraphQL API server at http://localhost:4000/graphql')
app.listen(4000)
