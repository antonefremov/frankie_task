## Overview

This is a simple RESTful service that listens to a default port and accept the POST method on /isgood path.

It's based on the [Gin](https://github.com/gin-gonic/gin) web server.

### Quick start

1. Clone this project into your Go path directory
2. Install the dependency below
```sh
$ go get -u github.com/gin-gonic/gin
```
3. Open the project folder in your command line and run
```go
go run server.go handlers.go model.go
```
As a result your server should start up on default port :8080

4. Now the server is ready to receive your POST request on ```http://localhost:8080/isgood```. Open another terminal window and send the following curl request containing valid JSON payload
```sh
$ curl -v -X POST   http://localhost:8080/isgood   -H 'content-type: application/json'   -d '{ "checkType": "DEVICE", "activityType": "SIGNUP", "checkSessionKey": "1234", "activityData": [{ "kvpKey": "key1", "kvpValue": "value1", "kvpType": "general.string" }, { "kvpKey": "key2", "kvpValue": "value2", "kvpType": "general.integer" }] }'
```
You should receive a response with status ```200``` and the same valid JSON payload.

### Testing

In order to execute the unit test for the service, please run the following command <ins>in the project folder</ins>
```sh
$ go test -v
```

To get code coverage percentage by the unit tests in project, please run
```sh
$ go test -cover
```

To get a html based code coverage report, please run the command below and open the generated ```cover.html``` file
```sh
$ go test -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html
```

Code coverage of the /isgood path handler is ```100%```.
