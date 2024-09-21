# Development setup
To generate some protobuf related code, you need the `protoc` tool, the
`protoc-gen-go` tool and the `protoc-gen-go-vtproto` tool; the latter
two can be installed like this:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.6.0
```

