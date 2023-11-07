# protobuf生成命令

```bash
protoc --go_out=./ *.proto
# gRpc生成
protoc --go_out=plugins=grpc:./ goods.proto
```