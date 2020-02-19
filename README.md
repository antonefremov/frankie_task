## Overview

This is a simple RESTful service that listens to a default port and accepts POST requests on the ```/isgood``` path. It accepts JSON payloads described in the [Swagger file](https://github.com/antonefremov/frankie_task/blob/master/swagger.yaml) definition.

The server is based on the [Gin](https://github.com/gin-gonic/gin) framework.

### Quick start

1. Clone this project into your Go path directory
2. Install the dependency below
```sh
$ go get -u github.com/gin-gonic/gin
```
3. Open the project folder in your terminal and run
```go
go run server.go handlers.go model.go
```
As a result a server should start up on default port :8080

4. Now the server is ready to receive your POST requests on ```http://localhost:8080/isgood```. Open another terminal window and send the following curl request containing valid JSON payload
```sh
$ curl -v -X POST   http://localhost:8080/isgood   -H 'content-type: application/json'   -d '{ "checkType": "DEVICE", "activityType": "SIGNUP", "checkSessionKey": "1234", "activityData": [{ "kvpKey": "key1", "kvpValue": "value1", "kvpType": "general.string" }, { "kvpKey": "key2", "kvpValue": "value2", "kvpType": "general.integer" }] }'
```
You should receive a response with status ```200``` and the same valid JSON payload as below
```sh
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 19 Feb 2020 10:37:19 GMT
< Content-Length: 220
< 
{
    "checkType":"DEVICE",
    "activityType":"SIGNUP",
    "checkSessionKey":"1234",
    "activityData": [
        {
            "kvpKey":"key1",
            "kvpValue":"value1",
            "kvpType":"general.string"
        },
        {
            "kvpKey":"key2",
            "kvpValue":"value2",
            "kvpType":"general.integer"
        }
    ]
}
```

### Testing

In order to execute the unit tests for the service, please run the following terminal command <ins>in the project folder</ins>
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

Code coverage of the /isgood path [handler](https://github.com/antonefremov/frankie_task/blob/master/handlers.go#L10) is ```100%```

![Code coverage image](/handler_code_coverage.jpg)