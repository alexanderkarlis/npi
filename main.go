package main

import (
	"fmt"

	server "github.com/alexanderkarlis/npi/npiserver"
)

func main() {
	fmt.Println("starting server...")
	server.Serve()
}
