// Netcat is a read-only TCP client
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	address := fmt.Sprint("localhost:", 8000)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
