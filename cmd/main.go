package main

import (
	"fmt"
	"os"

	certs "github.com/adriantirusli/ssl-checker"
)

var (
	port	string
	addr 	string
)

func main() {
	arg := os.Args[1]
	port = "443"

	var (
		c   *certs.Cert
		err error
	)

	addr = arg + ":" + port 

	if addr != "" {
		c, err = certs.ParseRemoteCertificate(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Usage:")
		return
	}

	fmt.Println(c.Jsonify())
}
