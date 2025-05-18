package cmd

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var tlsCmd = &cobra.Command{
	Use:   "tls",
	Short: "Check the TLS records for a URL address",
	Long: `Check the TLS records for a URL address
	
	EXAMPLE:
	
	go-webscraper tls gork.com
	
	Could return:

	GORK.COM
	TLS Handshake Complete: true
	TLS Version: 772
	Cipher Suite: TLS_AES_128_GCM_SHA256
	Server Name: *.gork.com
	Issuer: E5
	Valid From: 2025-03-19 20:27:27 +0000 UTC
	Valid To: 2025-06-17 20:27:26 +0000 UTC
	DNS Names: [*.gork.com gork.com] 
	`,

	Run: tlsCheck,
}

func init() {
	rootCmd.AddCommand(tlsCmd)

	tlsCmd.Flags().StringVarP(&port, "port-lookup", "p", "443", "Change the portthat is scanned with tls (default is 443)")
	traceCmd.Flags().IntVarP(&timeout, "timeout", "t", 1000, "Timeout of the tls-scan in milliseconds")
}

// these variables are package scoped
var port string

// timeout is declared in the trace.go command

func tlsCheck(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println("Please provide a URL to find the TLS records.")
		return
	}

	for _, domain := range args {
		fmt.Println()
		fmt.Println(strings.ToUpper(domain))
		checkTLS(domain, timeout)
	}
}

func checkTLS(domain string, timeout int) {

	dialer := &net.Dialer{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", domain+":"+port, &tls.Config{})
	if err != nil {
		fmt.Println("Could not establish a connection with", domain, err)
		return
	}
	defer conn.Close()

	state := conn.ConnectionState()

	fmt.Println("TLS Handshake Complete:", state.HandshakeComplete)
	fmt.Println("TLS Version:", state.Version)
	fmt.Println("Cipher Suite:", tls.CipherSuiteName(state.CipherSuite))

	// Inspect certificate
	cert := state.PeerCertificates[0]
	fmt.Println("Server Name:", cert.Subject.CommonName)
	fmt.Println("Issuer:", cert.Issuer.CommonName)
	fmt.Println("Valid From:", cert.NotBefore)
	fmt.Println("Valid To:", cert.NotAfter)
	fmt.Println("DNS Names:", cert.DNSNames)
	fmt.Println(cert.AuthorityKeyId, cert.BasicConstraintsValid, cert.CRLDistributionPoints, cert.EmailAddresses)

}
