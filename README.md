# pesan-grpc-stubs

## generate android stub
```
    ./gradlew :android-stub:build
```

## generate GO stub
```
    cd protos/src/main/proto

    protoc --go_out=../../../../go/ --go_opt=paths=source_relative \
    --go-grpc_out=../../../../go/ --go-grpc_opt=paths=source_relative \
    pesan.proto
```
