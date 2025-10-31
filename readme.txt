1. make sure u have go installed

> go version
go version go1.25.3 darwin/arm64

2. run go run main.go

3. create module with:

> go mod init github.com/softwarebyze/go-service-template
go: creating new go.mod: module github.com/softwarebyze/go-service-template
go: to add module requirements and sums:
        go mod tidy

4. for hot reloading, i used nodemon:

> nodemon --exec go run main.go --signal SIGTERM
(without --signal SIGTERM, it wasnt reloading for me)

5. /hello endpoint works

> curl localhost:8080/hi
hi⏎                                 