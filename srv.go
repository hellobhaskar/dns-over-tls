package main

import (
	"crypto/tls"
	"github.com/miekg/dns"
	"log"
	"os"
	"strconv"
	"sync"
)

//Refs:

//Hardcoded Cloudflare DNS
const (
	UPSTREAM_SERVER = "1.1.1.1:853"
)

type RequestHandler struct {
}

func (this *RequestHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	log.Println("Recieved DNS Request")
	c := new(dns.Client)
	c.Net = "tcp4-tls"
	c.TLSConfig = &tls.Config{
		MinVersion:               tls.VersionTLS12,
		InsecureSkipVerify:       true,
		PreferServerCipherSuites: true,
		ServerName:               "cloudflare-dns.com",
	}

	m := new(dns.Msg)

	response, _, err := c.Exchange(r, UPSTREAM_SERVER)

	if err != nil {
		log.Println("Failed to handle Request", err)
		dns.HandleFailed(w, m)
		return
	}

	if response == nil || r.Rcode != dns.RcodeSuccess {
		log.Println("Query failed : %v", response)
		log.Println(err)
		return
	}

	w.WriteMsg(response)

}

func TCPServerInterface() {
	log.Println("TCP Server Started")
	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "tcp"}
	srv.Handler = &RequestHandler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed to Start TCP Server", err)
		os.Exit(-1)
	}

}

func UDPServerInterface() {
	log.Println("UDP Server Started")
	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = &RequestHandler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed to Start UDP Server", err)
		os.Exit(-1)
	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	log.Println("Starting Servers")
	go TCPServerInterface()
	go UDPServerInterface()

	wg.Wait()
	log.Println("Server Shutdown")
}
