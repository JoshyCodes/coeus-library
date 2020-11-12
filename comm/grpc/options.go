package grpc

import (
	"log"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func SetGrpcOptions(tls, certDir string) []grpc.ServerOption {

//create options object to pass TLS or non secure settings into connection
opts := []grpc.ServerOption{}

//if TLS value true, load SSL certs
if tls == "true" {

	//set filename for SSL cert files
	certFile := fmt.Sprintf("%s/server.crt", certDir)
	keyFile := fmt.Sprintf("%s/server.pem", certDir)

	// set SSL creds from files set in previous filename settings
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	//if error with SSL init, return fatal message with error information 
	if sslErr != nil {
		// log error to console for SSL cert init failure
		log.Fatalf("Failed loading certificates: %v", sslErr)
	}
	
	// re-set options object to pass TLS information initiated in this 'if' statement
	opts = append(opts, grpc.Creds(creds))

	// comment to console that TLS SSSL certs loaded and are ready
	fmt.Println("TLS security is loaded")

}

return opts

}