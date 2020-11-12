

protoc:
	## build .pb.go file for Collections protobufs
	protoc ./coeus-library/types/protobuf/collections/collections.proto --go_out=plugins=grpc:.