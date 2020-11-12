

protoc:
	## build .pb.go file for Collections protobufs
	protoc ./types/protobuf/collections/collections.proto --go_out=plugins=grpc:.