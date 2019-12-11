# GETTING STARTED WITH PROTO, SERVER, CLIENT

## GETTING STARTED

* Move into project directory
  * `cd 01-code-generation/`

## PROTO

* Open `message.proto` file
  * Details: <https://developers.google.com/protocol-buffers/docs/proto3>
* Generate `message.pb.go` using below command
  * `protoc messagepb/message.proto --go_out=plugins=grpc:.`
* Open generated `message.pb.go` inside `messagepb` package
  * Server specific code:
    * `MessageServiceServer`
    * `RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer)`
  * Client specific code:
    * `MessageServiceClient`
    * `NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient`

## SERVER

* Open `server/server.go`
  * Create **listener** with `Listen(network, address string) (Listener, error)`
  * Create **server** with `grpc.NewServer()`
  * Register **server** with grpc using `RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer)`
  * Start **server** and accept incoming connections using `(s *Server) Serve(lis net.Listener) error`

## CLIENT

* Open `client/client.go`
  * Create **client connection** using `Dial(target string, opts ...DialOption) (*ClientConn, error)`
  * Create **client** using `NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient`
