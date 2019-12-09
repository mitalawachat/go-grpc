# UNARY RPC

## GETTING STARTED

* **Unary RPC** calls are the basic **request**/**response** model
* **Client** will send **one** message to the **Server** and will receive **one** response from **Server**
* Unary RPC are very well suited when:
  * Data is small
  * Start with **Unary API** and use **Streaming API** if performance is an issue
* For each RPC call the **protocol buffer** file will contain:
  * **Request** message
  * **Response** message
  * **Service** definition

## PROTO

* Generate `welcome.pb.go` using below command
  * `protoc welcomepb/welcome.proto --go_out=plugins=grpc:.`
* Open generated `welcome.pb.go` inside `welcomepb` package
  * Contains `WelcomeServiceServer` - `RegisterWelcomeServiceServer(s *grpc.Server, srv WelcomeServiceServer)`
  * Contains `WelcomeServiceClient` - `NewWelcomeServiceClient(cc *grpc.ClientConn) WelcomeServiceClient`
  * Contains `Welcome(ctx context.Context, req *WelcomeRequest) (*WelcomeResponse, error)`

## SERVER

* Open `server/server.go`
  * Do steps mentioned in `01-code-generation`
  * Implement method `Welcome(ctx context.Context, request *welcomepb.WelcomeRequest) (*welcomepb.WelcomeResponse, error)`

## CLIENT

* Open `client/client.go`
  * Do steps mentioned in `01-code-generation`
  * Create request object of `welcomepb.WelcomeRequest`
  * Request welcome service using `client.Welcome`