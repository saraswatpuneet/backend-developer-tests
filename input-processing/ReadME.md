# Input-Processing: a gRPC based bi-directional streaming via STDIN/STDOUT

## General Layout
```
1. streamserver: a gRPC server that receives a stream of messages from STDIN and sends a stream of messages to STDOUT, while also exposing RPC server over http via grpc gateway based reverse proxy
2. streamclient: a gRPC client that sends a stream of messages to the streamserver, client reads from stdin via bufio, streams to running grpc server and prints any responses received from the server to stdout
Note: currently the rpc port :9192 and http port: 8091 are hardcoded in the streamserver and streamclient that can be exposed via env vars or secrets
```

## Usage and Testing
```
1. streamserver: to build and run the server run
    1. cd streamserver/
    2. make run: will build server binary and run it
    3. go run main.go : alternative to make run
2. streamclient: to build and run the client run
    1. cd streamclient/
    2. make run: will build client binary and run it
    3. go run main.go : alternative to make run
    4. streamclient will connect to server via grpc connection and takes input from command line and streams to server, server will return lines which has word "error"
3. Calling via http: (checkout postman file in /postman folder)
    POST: http://localhost:8091/v1/lines
    Body: {"message": "hello world we have an error \n another error"}
    Expected output: two lines of response from server since error is present in the input separated by newline
```
## Multiple Streams/ Multiple Clients
```
1. Currently client can be initialized multiple times, grpc transport layer should allow
    upto 100 concurrent streams. Another layer of scalable and large scale application can be achieved by creating and managing a pool of servers exposed over multiple ports with smart load balancing and failover mechanism.
2. With respect to RestAPI  the stream terminates when connection is closed prompting a    
    request/response behavior.
```

## Datastreaming
```
1. A bi-directional grpc stream enables streaming large data saying reading a file and writing contents to a stream.
2 grpc stream generally limits data size to 4MB so for such dev only applications, we can increasing the limit to desired level but ideally we should implement a chunking mechanism or define high level interface that can be extended to create custom chunks.