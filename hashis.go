package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func HandleConnection(conn net.Conn) {
	log.Printf("Accepted connection from %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	queryStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading query string")
	}

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strings.TrimSpacequeryStr)))
	hashedStr := fmt.Sprintf("%x", sha1Hash.Sum(nil))

	writer := bufio.NewWriter(conn)
	n, err := writer.WriteString(fmt.Sprintf("%s sha1 %s", hashedStr, queryStr))
	if err != nil {
		log.Fatal("Couldn't write to connection.")
	} else {
		log.Printf("Wrote %d bytes", n)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal("Could not flush bytes to connection")
	}
	log.Println("Closing connection")
	conn.Close()
}

func main() {
	log.Println("Starting HashIs")

	ln, err := net.Listen("tcp", ":43")
	if err != nil {
		log.Fatal("Error listening.")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Can't accept.")
			return
		}

		go HandleConnection(conn)
	}
}
