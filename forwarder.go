package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func listenAndServe() {
	if (*listenAddr == "" || *connectAddr == "") ||
		(!*noListenTLS && (*certFilePath == "" || *keyFilePath == "")) {
		flag.Usage()
		os.Exit(2)
	}
	var listen func(network, address string) (net.Listener, error)
	if *noListenTLS {
		listen = net.Listen
	} else {
		var cert tls.Certificate
		cert, err := tls.LoadX509KeyPair(*certFilePath, *keyFilePath)
		if err != nil {
			log.Fatalf("Certificate error: %v\n", err)
		}
		config := &tls.Config{Certificates: []tls.Certificate{cert}}
		listen = func(network, address string) (net.Listener, error) {
			return tls.Listen(network, address, config)
		}
	}
	var dial func(network, address string) (net.Conn, error)
	if *noConnectTLS {
		dial = net.Dial
	} else {
		config := &tls.Config{ServerName: *tlsServerName, InsecureSkipVerify: *insecureTLS}
		dial = func(network, address string) (net.Conn, error) {
			return tls.Dial(network, address, config)
		}
	}
	listener, err := listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("Listen error: %v\n", err)
	}
	log.Printf("Listening at: %v\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v\n", err)
			continue
		}
		go handleConn(conn, dial)
	}
}

func handleConn(src net.Conn, dial func(network, address string) (net.Conn, error)) {
	defer src.Close()
	dst, err := dial("tcp", *connectAddr)
	if err != nil {
		log.Printf("Dial error: %v\n", err)
		return
	}
	defer dst.Close()
	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(dst, src)
		src.Close()
		dst.Close()
		done <- struct{}{}
	}()
	go func() {
		_, _ = io.Copy(src, dst)
		src.Close()
		dst.Close()
		done <- struct{}{}
	}()
	<-done
	<-done
}
