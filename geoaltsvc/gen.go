//go:generate protoc geoaltsvc.proto --swift_out=. --swiftgrpc_out=Client=true,Server=false:.
//go:generate protoc --proto_path=. --go_out=plugins=grpc:. geoaltsvc.proto
package geoaltsvc
