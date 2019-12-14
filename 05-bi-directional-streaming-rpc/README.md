# BI-DIRECTIONAL STREAMING RPC

## GETTING STARTED

* **Bi-Directional Streaming RPC** enables client to send **many** messages to the server, and will receive **many** response from the server; server can send response at any point (no need to wait till all messages are received from client)
* Bi-Directional streaming are suited for:
  * Client and Server needs to send lot of data asynchronously (big data)
  * **Chat** protocol
  * Long running connections
* For each RPC call the **protocol buffer** file will contain:
  * **Request** message with keyword **stream**
  * **Response** message with keyword **stream**
  * **Service** definition

## PROTO

* Generate `welcome.pb.go` using below command
  * `protoc welcomepb/welcome.proto --go_out=plugins=grpc:.`
* Open generated `welcome.pb.go` inside `welcomepb` package
  * Server specific code:
    * `WelcomeServiceServer`
    * `WelcomeServiceServer` - `WelcomeEveryone(WelcomeService_WelcomeEveryoneServer) error`
    * `RegisterWelcomeServiceServer(s *grpc.Server, srv WelcomeServiceServer)`
  * Client specific code:
    * `WelcomeServiceClient`
    * `WelcomeServiceClient` - `WelcomeEveryone(ctx context.Context, opts ...grpc.CallOption) (WelcomeService_WelcomeEveryoneClient, error)`
    * `NewWelcomeServiceClient(cc *grpc.ClientConn) WelcomeServiceClient`

## SERVER

* Open `server/server.go`
  * Do steps mentioned in `01-code-generation`
  * Implement method `WelcomeEveryone(WelcomeService_WelcomeEveryoneServer) error`
  * Receive requests using `Recv() (*WelcomeEveryoneRequest, error)`
  * Iterate over stream until you receive error
    * If error **is** `io.EOF` then successfully received all requests from client
    * If error **is not** `io.EOF` then something went wrong
  * Send response using `Send(*WelcomeEveryoneResponse) error`

## CLIENT

* Open `client/client.go`
  * Do steps mentioned in `01-code-generation`
  * Create request object of `welcomepb.WelcomeEveryoneRequest`
  * Send requests using `Send(*WelcomeEveryoneRequest) error`
  * Receive responses using `Recv() (*WelcomeEveryoneResponse, error)`
  * Iterate over stream until you receive error
    * If error **is** `io.EOF` then successfully received all requests from server
    * If error **is not** `io.EOF` then something went wrong
