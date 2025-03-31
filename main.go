package main

import (
	"fmt"
	"log"
	"github.com/Denyme24/go-dns-server/env"
	"github.com/gofiber/fiber/v2"
	"github.com/miekg/dns"
)

// Our DNS records database (in-memory for this example)
var dnsRecords = map[string]string{
	"example.com.":  "192.0.2.1",
	"test.com.":     "203.0.113.42",
	"namanraj.tech": "192.0.2.123",
}

func main() {
	// Start DNS server in a goroutine
	go startDNSServer()

	// Start Fiber web server
	startWebServer()
}

func startDNSServer() {
	// Set up DNS server
	server := &dns.Server{
		Addr: env.GetStringEnv("ADDRESS", "0.0.0.0:9090"), // Bind to all interfaces
		Net:  env.GetStringEnv("PROTOCOLE", "udp"),
	}

	// Handle DNS requests
	dns.HandleFunc(".", handleDNSRequest)

	// Start server
	log.Printf("Starting DNS server on %s...", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start DNS server: %v", err)
	} else {
		log.Printf("DNS server is running")
	}
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true // Mark this server as authoritative
	found := false

	for _, q := range r.Question {
		log.Printf("Query: %s (%s)", q.Name, dns.TypeToString[q.Qtype])
		if q.Qtype == dns.TypeA {
			if ip, exists := dnsRecords[q.Name]; exists {
				rr, _ := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				m.Answer = append(m.Answer, rr)
				found = true
			}
		}
	}

	if !found {
		m.Rcode = dns.RcodeNameError // NXDOMAIN
	}

	w.WriteMsg(m)
}

func startWebServer() {
	app := fiber.New()

	// Web interface to manage DNS records
	app.Get("/records", func(c *fiber.Ctx) error {
		return c.JSON(dnsRecords)
	})

	app.Post("/records", func(c *fiber.Ctx) error {
		type Record struct {
			Domain string `json:"domain"`
			IP     string `json:"ip"`
		}

		var record Record
		if err := c.BodyParser(&record); err != nil {
			return c.Status(400).SendString("Bad request")
		}

		// Ensure domain ends with dot (DNS standard)
		if record.Domain[len(record.Domain)-1] != '.' {
			record.Domain += "."
		}

		dnsRecords[record.Domain] = record.IP
		return c.SendString("Record added successfully")
	})

	log.Printf("Starting web server on :3000...")
	log.Fatal(app.Listen(":3000"))
}
