module github.com/linkall-labs/vanus/proto

go 1.18

require (
	github.com/golang/mock v1.6.0
	github.com/linkall-labs/vanus/raft v0.5.7
	google.golang.org/grpc v1.56.3
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)

replace (
	github.com/linkall-labs/vanus/observability => ../observability
	github.com/linkall-labs/vanus/raft => ../raft
)
