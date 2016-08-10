# gRPC-ing
Battling the first gRPC gym for the first badge!

## Up and running

Install dependencies as described in [here](http://www.grpc.io/docs/).

Generating the gRPC code for client and server use:
```
$ protoc -I=./protobuf --go_out=plugins=grpc:./protobuf ./protobuf/helloworld.proto
```

Run the server:
```
$ go run server/main.go &
```

Run the client, optionally passing in a value:
```
$ go run client/main.go 'Hello, Calvin! 👋'
```
