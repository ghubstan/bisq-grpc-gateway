# bisq-grpc-gateway

Golang based grpc gateway to support Bisq REST clients (POC)

Routes HTTP 1.1 requests to a remote server's single gRPC service method, which is responsible for detecting if the 
request is from this gateway, and wrapping the response in json.  

There is a single gRPC Service `CallService` with one method `Call(String params)`.  The `CallService` returns a
`Response` protobuf message with a single field `result`, or an error. 

The gateway/proxy takes care of mapping gRPC error status codes to HTTP status codes, but the remote server is responsible
for translating a core api exception into a gRPC `StatusRuntimeException` with the appropriate `grpc.io.Status`.   

There is a single HTTP 1.1 endpoint `http://host:port/v1/call` which accepts a request body containing `params`;  extra 
path elements and query strings are ignored, and will likely cause errors.

## Run HTTP 1.1 Gateway Proxy

Start your gRPC server first, then the proxy:

`go run main.go proxy`

## Example HTTP 1.1 Requests 

### getversion 

`curl -v -X POST http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "getversion"}'`

### setwalletpassword 

`curl -v -X POST http://localhost:8080/v1/call -H "Authorization:xyz"  -H "Content-Type: application/json" -d '{"params": "setwalletpassword newpassword"}'`

## Development

There are sample server and client programs for dev/test purposes.  Both can be started from main.go:
 
`go run main.go server`

`go run main.go client [params]`


## TODO
Explain go runtime setup basics, bisq-grpc-gateway install and run steps.

## TODO 
Explain server/proxy address configuration.

## TODO 
Show response examples.
