package main

import (
	"crypto/tls"
)

func main() {
	conn,err := tls.Dial("tcp", "example.com:80", nil)
	if err != nil {
		panic("Server doesnt support SSL cert. Err: " + err.Error())
	}

}