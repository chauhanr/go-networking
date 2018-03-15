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
		fmt.Printf("Usage: Choose between (TCP, UDP protocol servers)")
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
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener,err := net.ListenTCP("tcp", tcpAddr)
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

func StartUDPServer(){
	udpAddr, err := net.ResolveUDPAddr("udp4", names.UDP_SERVER_PORT)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for{
		handleUDPClient(conn)
	}
}

func handleUDPClient(conn *net.UDPConn){
	var buf [512]byte

	_,addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil{
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
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
