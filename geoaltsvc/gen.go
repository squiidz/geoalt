//go:generate protoc geoaltsvc.proto --swift_out=. --swiftgrpc_out=Client=true,Server=false:.
//go:generate protoc --proto_path=. --go_out=plugins=grpc:. geoaltsvc.proto
//go:generate protoc -I=. geoaltsvc.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:.
package geoaltsvc
