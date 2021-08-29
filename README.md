# go-gRPC

## Generating proto file
``cd`` into ``hello/hellopb/`` and type ``protoc --go_out=plugins=grpc:. hellopb/hello.proto``.
The result gives you a `hello.pb.go` file. If you look into the file, you will find the magic of gRPC is that all these informations are represented in go specific languages, and if you use other languages, the same protobuf will be represented in other specific language, and this is also why protobuf can cross-communicate between languages.

### THe basic point here is:
``In gRPC a client application can directly call methods on a server application on a different machine as if it was a local object, making it easier for you to create distributed applications and services.``