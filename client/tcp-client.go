package main

import (
	"net"
	"fmt"
	"os"
	"io/ioutil"
	"go-networking/server/names"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Printf("Usage: Choose between (TCP, UDP protocol servers)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "tcp":
		TCPClient(os.Args[1])
	case "udp":
		UDPClient(os.Args[1])
	}
}

func TCPClient(proto string){
	/*tcpAddr, err := net.ResolveTCPAddr("tcp", names.TCP_SERVER_PORT)
	checkError(err)*/

	conn, err := net.Dial(proto, names.TCP_SERVER_PORT)
	checkError(err)

	_,err = conn.Write([]byte(""))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Printf("%s \n",string(result))

}

func UDPClient(proto string){
	/*udpAddr, err := net.ResolveUDPAddr("udp4", names.UDP_SERVER_PORT)
	checkError(err)*/

	conn, err := net.Dial(proto, names.UDP_SERVER_PORT)
	checkError(err)

	_,err = conn.Write([]byte(""))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	fmt.Printf("%s\n", string(buf[0:n]))
}

func checkError(err error) {
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal Error : %s", err)
		os.Exit(1)
	}
}