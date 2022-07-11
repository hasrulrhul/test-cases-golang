
# GOKA PROTOBUF

### For Test Cases

### Install Apache Kafka 
```
docker-compose up --build -d
```

### Go plugins for the protocol compiler
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Go package
```
$ go get google.golang.org/grpc
$ go get google.golang.org/protobuf
```

### Generate protobuf and run server
```
make gen
```

### Run Server
```
make run-server

http://localhost:2000/
```

### Run Client
```
make run-client

http://localhost:1234/
```
