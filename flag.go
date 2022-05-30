package main

import (
	"flag"
)

var (
	listenAddr    = flag.String("l", "", "Listen at the specific address")
	connectAddr   = flag.String("c", "", "Connect to the specific address")
	noListenTLS   = flag.Bool("noltls", false, "Do not listen TLS")
	noConnectTLS  = flag.Bool("noctls", false, "Do not connect TLS")
	certFilePath  = flag.String("cert", "", "Specify the path to the certificate file")
	keyFilePath   = flag.String("key", "", "Specify the path to the private key file")
	insecureTLS   = flag.Bool("insecuretls", false, "Allow insecure TLS")
	tlsServerName = flag.String("servername", "", "Specify the server name in the certificate presented by the server")
)
