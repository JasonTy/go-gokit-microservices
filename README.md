## Create protobuf ##

```$xslt
protoc --go_out=plugins=grpc:. helloword.proto
```

Use ```go-kit``` restructure grpc-server

## Components used ##

```
1.github.com/go-kit/kit/transport/grpc

2.github.com/go-kit/kit/endpoint
```
