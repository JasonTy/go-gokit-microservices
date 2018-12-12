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

## Question ##

This protobuf use ```syntax = "proto3"```. But created ```.pb.go``` file used ```const _ = proto.ProtoPackageIsVersion3```

Report error ```undefined```

But no solution was found in ```github.com/golang/protobuf```
