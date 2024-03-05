package main

import (
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"
)

func main() {
	conn,err := tls.Dial("tcp", "google.com:443", nil)
	must(err, "erver doesn't support SSL certificate err: ")
	// check SSL and site hostname match
	er := conn.VerifyHostname("google.com")
	must(er, "Hostname doesn't match with certificate: ")
	// fmt.Println(conn)

	

	viewPeerCertificates(*conn.ConnectionState().PeerCertificates[0])


	// fmt.Println(conn.ConnectionState().VerifiedChains[0][0].DNSNames)
}

func must(err error, s string) {
	if err != nil {
		panic(s + err.Error())
	}
}

func viewPeerCertificates(c x509.Certificate) {

	nbexpiry := c.NotBefore
	expiry := c.NotAfter
	issuer := c.Issuer
	subject := c.Subject
	extensions := c.Extensions[0]
	// ips := c.IPAddresses
	td := time.Now()
	certFingerprint := sha1.Sum(c.Raw)

	fmt.Println("")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("Check Expiry")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("")

	if td.Before(nbexpiry) {
		fmt.Println("Cert used before commencement\nCertificate Not Before: " + nbexpiry.String())
	} else {
		fmt.Println("Cert Valid. Not before: " + expiry.String())
	}

	if td.After(expiry) {
		fmt.Println("Cert Expired\nCertificate Expired: " + expiry.String())
	} else {
		fmt.Println("Cert Valid. Will expire: " + expiry.String())
	}

	fmt.Println("")
	fmt.Println("--------------------------------------------------------")



	fmt.Printf("\nIssuer: %s\nNot Before: \t%v\nExpiry: \t%v\nSubject: %s\nExtentions: %v\nFingerprint: %v\n", issuer, nbexpiry, expiry, subject, extensions, certFingerprint)

}

