package main

import (
	"net"
	"fmt"
	"os"
	"time"
	"go-networking/server/names"
)

/** Day timeserver*/
func main(){
	if len(os.Args) != 2 {
		fmt.Printf("Usage: <protocol> [tcp|udp]")
		os.Exit(1)
	}

	switch os.Args[1] {
		case "tcp":
			StartTCPServer()
		case "udp":
			StartUDPServer()
	}
}

func StartTCPServer(){
	service := names.TCP_SERVER_PORT
	/*tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)*/

	listener,err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil{
			continue
		}

		daytime := time.Now().String()
		go handleClient(conn, daytime)
	}
}

/**This method will start the UDP serverit will use the ListenPacket() in net package to
   get the PacketConnection  */
func StartUDPServer(){
	service := names.UDP_SERVER_PORT

	conn, err := net.ListenPacket("udp", service)
	checkError(err)
	for{
		handleUDPClient(conn)
	}
}
/** as UDP is a connectionless protocol and does not require session like TCP the source and destination are already
	mentioned in the request. */
func handleUDPClient(conn net.PacketConn){
	var buf [512]byte
	_,addr, err := conn.ReadFrom(buf[0:])
	if err != nil{
		return
	}
	daytime := time.Now().String()
	conn.WriteTo([]byte(daytime), addr)
}


func handleClient(conn net.Conn, message string){
	fmt.Printf("Processing message %s\n", message)
	defer conn.Close()
	conn.Write([]byte(message))
}

func checkError(err error) {
	if err != nil{
		fmt.Fprint(os.Stderr, "Fatal Error : %s", err)
		os.Exit(1)
	}
}
