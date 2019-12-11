# SERVER STREAMING RPC

## GETTING STARTED

* **Server Streaming RPC** enables client to send **one** message to the server, and will receive **many** responses from the server; possibly an infinite number
* Server streaming are suited for:
  * Server needs to send lot of data (big data)
  * When the server needs to **PUSH** data to client without having the client request for more (like feed, chat, etc)
* For each RPC call the **protocol buffer** file will contain:
  * **Request** message
  * **Response** message with keyword **stream**
  * **Service** definition

## PROTO

* Generate `welcome.pb.go` using below command
  * `protoc welcomepb/welcome.proto --go_out=plugins=grpc:.`
* Open generated `welcome.pb.go` inside `welcomepb` package
  * Server specific code:
    * `RegisterWelcomeServiceServer(s *grpc.Server, srv WelcomeServiceServer)`
    * `WelcomeServiceServer` - `WelcomeManyTimes(*WelcomeManyTimesRequest, WelcomeService_WelcomeManyTimesServer) error`
  * Client specific code:
    * `NewWelcomeServiceClient(cc *grpc.ClientConn) WelcomeServiceClient`
    * `WelcomeServiceClient` - `WelcomeManyTimes(ctx context.Context, in *WelcomeManyTimesRequest, opts ...grpc.CallOption) (WelcomeService_WelcomeManyTimesClient, error)`

## SERVER

* Open `server/server.go`
  * Do steps mentioned in `01-code-generation`
  * Implement method `WelcomeManyTimes(*WelcomeManyTimesRequest, WelcomeService_WelcomeManyTimesServer) error`

## CLIENT

* Open `client/client.go`
  * Do steps mentioned in `01-code-generation`
  * Create request object of `welcomepb.WelcomeManyTimesRequest`
  * Request welcome service using `client.WelcomeManyTimes`
  * Iterate over response utill you receive error
    * If error **is** `io.EOF` then successfully received all response from server
    * If error **is not** `io.EOF` then something went wrong
