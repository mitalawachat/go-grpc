# gRPC (GoLang)

## API

* API (Application Programming Interface) is a contract that says:
  * Client sends **request**
  * Server sends **response**

## RPC

* RPC - Remote Procedure Call
* In **client** code it looks like we are calling function directly on the **server**

## gRPC

* gRPC is free and open-source framework developed by **Google**
* gRPC is part of the **Cloud Native Computation Foundation (CNCF)** - like Docker & Kubernetes
* gRPC allows us to define **request** and **response** for **RPC** and handled all the rest for us
* gRPC is built on top of **HTTP/2**
* gRPC is fast, efficient, low latency
* gRPC supports **streaming**
* gRPC is language independent
* Easy to build micro-services that interact with each other since code can be generated for almost any language

## STEPS

* Define messages and services using **Protocol Buffers**
* gRPC code will be generated, and we will have to provide an implementation for it
* **.proto** file works for over 12 programming languages (server and client)

## WHY PROTOCOL BUFFERS

* Protocol Buffers are language agnostic
* Code can be generated for almost any language
* Data is **binary**, therefore very efficient to send/receive on network and serialize/de-serialize on a CPU
* Easy to write message definition
* Definition of API is independent from the implementation
* Easy API evolution using rules without breaking existing clients

## HTTP/1.1

* HTTP/1.1 was released in 1997
* HTTP/1.1 opens a **new TCP connection** to a server at each request
* HTTP/1.1 **does not compress headers** (headers are plaintext)
* HTTP/1.1 only works with request/response mechanism (**no server push**)

## HTTP/2

* HTTP/2 is newer standard for internet communications that addesses common pitfalls of HTTP/1.1
* HTTP/2 was released in 2015
* HTTP/2 is battle tested (before that tested by Google under name **SPDY**)
* HTTP/2 is **binary**
* HTTP/2 supports **multiplexing**
* HTTP/2 supports **server push**
* HTTP/2 supports **header compression**
* HTTP/2 is **secure** (**SSL** is not required but recommended by default)
* HTTP1.1 vs HTTP2 Performance Comparision - <https://imagekit.io/demo/http2-vs-http1>
* gRPC leverages HTTP/2 as a backbone for communications

## RPC TYPES

* Unary RPC
  * `rpc ExampleMethod(ExampleRequest) returns (ExampleResponse) {};`
* Server Streaming RPC
  * `rpc ExampleMethod(ExampleRequest) returns (stream ExampleResponse) {};`
* Client Streaming RPC
  * `rpc ExampleMethod(stream ExampleRequest) returns (ExampleResponse) {};`
* Bi-Directional Streaming RPC
  * `rpc ExampleMethod(stream ExampleRequest) returns (stream ExampleResponse) {};`

## gRPC SCALABILITY

* **gRPC servers** are asynchronous by default
  * Do not block threads on request
  * gRPC server can serve millions if requests in parallel
* **gRPC clients** can be asynchronous or synchronous (blocking)
  * The client decides which model works best for the performance needs
  * gRPC client can perform client side load balancing
* Google has *~10 billion* gRPC requests served per second

## gRPC SECURITY

* Security is first class citizen in gRPC as it strongly advocates for you to use **SSL** in your API
* Additionally, using **interceptors**, we can also provide authentication

## gRPC vs REST

| gRPC | REST |
| ---- | ---- |
| Protocol Buffers - smaller, faster | JSON - text based, slower, bigger |
| HTTP/2 (lower latency) - from 2015 | HTTP1.1 (higher latency) - from 1997 |
| Bi-directional | Client -> Server requests only |
| Streaming support | Request/Response support only |
| API Oriented - `what`(no constraints - free design) | CRUD Oriented (Create-Retrieve-Update-Delete / POST-GET-PUT-DELETE) |
| Code Generation through Protocol Buffers in any language - 1st class citizen | Code generation through OpenAPI/Swagger (add-on) - 2nd class citizen |
| RPC based - gRPC does the plumbing for us | HTTP verbs based - we have to write the plumbing or use a 3rd party library |

## ENVIRONMENT SETUP

* Install Golang: <https://golang.org/doc/install>
  * Make sure that the go binaries are in your `PATH`
  * `go version`
* Install VSCode: <https://code.visualstudio.com/>
* Install `protoc`
* `go get -u google.golang.org/grpc`
* `go get -u github.com/golang/protobuf/protoc-gen-go`

## EXAMPLES

* [Code Generation using `.proto` file](01-code-generation/README.md)
* [Unary RPC](02-unary-rpc/README.md)
* [Server Streaming RPC](03-server-streaming-rpc/README.md)
* [Client Streaming RPC](04-client-streaming-rpc/README.md)
* [Bi-Drectional RPC](05-bi-directional-streaming-rpc/README.md)

## GRPC ERROR CODES

* **Guide:** <https://grpc.io/docs/guides/error/>
* If application needs to return extra information on top of an error code, it was use the metadata context.
* Examples: <http://avi.im/grpc-errors/>

## DEADLINES

* Deadlines allow gRPC clients to specify how long they are willing to wait for an RPC to complete before the RPC is terminated with the error `DEADLINE_EXCEEDED`
* gRPC documentation recommends you set a deadline for all client RPC calls
* `Server` should check if the deadline has exceeded and cancel the work it is doing
* Deadlines are propogated across if gRPC calls are chained
  * A -> B -> C (deadline for A is passed to B and then passed to C)
* **Blog:** <https://grpc.io/blog/deadlines/>
