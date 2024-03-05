package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	conn,err := tls.Dial("tcp", "google.com:443", nil)
	must(err, "erver doesn't support SSL certificate err: ")
	// check SSL and site hostname match
	er := conn.VerifyHostname("go4ogle.com")
	must(er, "Hostname doesn't match with certificate: ")
	fmt.Println(conn)
}

func must(err error, s string) {
	if err != nil {
		panic(s + err.Error())
	}
}