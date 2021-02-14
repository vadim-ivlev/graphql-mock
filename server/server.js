// Running GraphQL API server
// npm install express express-graphql graphql graphql-list-fields
// https://graphql.org/graphql-js/object-types/

var fs = require('fs')
var express = require('express')
var { graphqlHTTP } = require('express-graphql')
var { buildSchema } = require('graphql')

var getFieldNames = require('graphql-list-fields')
var cors = require('cors')


var fakeDatabase = {'k0':'Hello World!'}



// This class implements the RandomDie GraphQL type
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
    getDie: ({ numSides }, ctx, info) => {
        // get selected fields
        console.log(getFieldNames(info))
        return new RandomDie(numSides)
    },
    getMessage: ({key}) => {
        return fakeDatabase[key]
    },
    setMessage: ({ key, message }) => {
        fakeDatabase[key] = message
        return message
    }
}



// Construct a schema, using GraphQL schema language
var schema = buildSchema(fs.readFileSync('server.graphql', 'utf8'))


var app = express()
app.use(cors())
app.use( '/graphql', graphqlHTTP({
        schema: schema,
        rootValue: root,
        graphiql: true
    })
)
app.listen(4000)
console.log('Running a GraphQL API server at http://localhost:4000/graphql')
