

protoc:
	## build .pb.go file for Collections protobufs
	protoc ./types/protobuf/collections/collections.proto --go_out=plugins=grpc:.

	## build .pb.go file for Authority protobufs
	protoc ./types/protobuf/authority/authority.proto --go_out=plugins=grpc:.