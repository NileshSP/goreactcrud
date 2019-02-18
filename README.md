# goreactcrud
Sample web app built using react, go & (hashicorp's)in-memory/sqlite db

Checkout the live [demo](https://goreactcrud.herokuapp.com)

Project is published to heroku using [Dockerfile](https://github.com/NileshSP/goreactcrud/blob/master/Dockerfile)

Clone the repository `https://github.com/NileshSP/goreactcrud.git`

Open terminal one and build as:
 `go build ./server/*.go`

In terminal one, start go server as:
 `go run ./server/*.go`

Open terminal two, start react client app as:
 `cd ./client && npm start`