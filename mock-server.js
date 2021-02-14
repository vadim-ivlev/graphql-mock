// Running Mock GraphQL API server
// https://graphql.org/blog/mocking-with-graphql/

var fs = require('fs')
var express = require('express')
var bodyParser = require('body-parser')
var { mockServer } = require('graphql-tools');

// Construct a schema, using GraphQL schema language
var mockServer = mockServer(fs.readFileSync('server.graphql', 'utf8'));

var app = express()

app.use(function (req, res, next) {
    res.header('Access-Control-Allow-Origin', '*')
    res.header('Access-Control-Allow-Headers', '*')
    res.header('Content-Type', 'application/json; charset=utf-8')
    next()
});

app.use(bodyParser.json());

app.all('/graphql', (req, res) =>{
        mockServer.query(req.body.query || req.query.query).then( result => res.send(result) )        
    }
)
app.listen(4000)
console.log('Running a Mock GraphQL API server at http://localhost:4000/graphql')
