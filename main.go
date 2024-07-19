package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:400")
	if err != nil {
		fmt.Println("Failed to bind to port 400:", err)
		os.Exit(1)
	}
	fmt.Println("Listening on port 400")
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			request := string(buf[:n])
			solicitud := strings.SplitN(request, "\r\n", 2)[0]
			fmt.Println("Solicitud recibida:", solicitud)
			response := "HTTP/1.1 200 OK\r\n" + "\r\n"
			conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Error escribiendo la respuesta:", err)
				break
			}
		}
		if err != nil {
			fmt.Println("Error escribiendo leyendo la solicitud:", err)
			break
		}
	}
	fmt.Println("Conexión con:", conn.RemoteAddr(), "cerrada.\n")
}
