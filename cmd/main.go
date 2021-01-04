package main

import (
	"fmt"
	"os"

	certs "github.com/adriantirusli/ssl-checker"
)

func main() {
	addr := os.Args[1]

	var (
		c   *certs.Cert
		err error
	)

	addr = addr + ":" + "443" 

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
