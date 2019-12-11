# CLIENT STREAMING RPC

## GETTING STARTED

* **Client Streaming RPC** enables client to send **many** messages to the server, and will receive **one** response from the server; server can send response at any point (no need to wait till all messages are received from client)
* Client streaming are suited for:
  * Client needs to send lot of data (big data)
  * When the Server processing is expensive and should happen as the client sends data
  * When the Client needs to **PUSH** data to the server without really expecting a response
* For each RPC call the **protocol buffer** file will contain:
  * **Request** message with keyword **stream**
  * **Response** message
  * **Service** definition

## PROTO

* Generate `welcome.pb.go` using below command
  * `protoc welcomepb/welcome.proto --go_out=plugins=grpc:.`
* Open generated `welcome.pb.go` inside `welcomepb` package
  * Server specific code:
    * `WelcomeServiceServer`
    * `WelcomeServiceServer` - `LongWelcome(WelcomeService_LongWelcomeServer) error`
    * `RegisterWelcomeServiceServer(s *grpc.Server, srv WelcomeServiceServer)`
  * Client specific code:
    * `WelcomeServiceClient`
    * `WelcomeServiceClient` - `LongWelcome(ctx context.Context, opts ...grpc.CallOption) (WelcomeService_LongWelcomeClient, error)`
    * `NewWelcomeServiceClient(cc *grpc.ClientConn) WelcomeServiceClient`

## SERVER

* Open `server/server.go`
  * Do steps mentioned in `01-code-generation`
  * Implement method `LongWelcome(WelcomeService_LongWelcomeServer) error`
  * Iterate over stream until you receive error
    * If error **is** `io.EOF` then successfully received all requests from client
    * If error **is not** `io.EOF` then something went wrong
  * Send response using `SendAndClose(*LongWelcomeResponse) error`

## CLIENT

* Open `client/client.go`
  * Do steps mentioned in `01-code-generation`
  * Create request object of `welcomepb.LongWelcomeRequest`
  * Request welcome service using `Send(*LongWelcomeRequest) error`
  * Receive server response using `CloseAndRecv() (*LongWelcomeResponse, error)`
