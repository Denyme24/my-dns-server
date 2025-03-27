# Go DNS Server

A simple DNS server written in Go with a web management interface using Fiber. Perfect for learning DNS fundamentals and local network management.

## Features

- 🚀 Basic DNS (A record) resolution
- 🌐 Web-based management API (add/remove records)
- 🔧 In-memory DNS record storage
- 🔄 UDP-based DNS server
- 🔒 Custom port support

## Prerequisites

- [Go 1.16+](https://golang.org/dl/)
- Admin/root access (for port binding)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-dns-server.git
cd go-dns-server
```
2. Install dependencies:
```bash
go get github.com/gofiber/fiber/v2
go get github.com/miekg/dns
```
## Configuration
1. Edit initial DNS records in main.go:
```
var dnsRecords = map[string]string{
    "example.com.": "192.0.2.1",
    "test.com.":    "203.0.113.42",
}
```
2.(Optional) Change ports:
```
// DNS Server port (default: 9090)
server.Addr = ":9090"

// Web Server port (default: 3000)
app.Listen(":3000")
```
## Usage
1. Start the server :
```
go run main.go
```
2. Add DNS records via API: 
``` 
curl -X POST -H "Content-Type: application/json" \
  -d '{"domain":"mysite.com.","ip":"your_ip"}' \
  http://localhost:3000/records
```
3. Query DNS records:
```
dig @127.0.0.1 -p 9090 example.com
# or for Windows:
nslookup -port=9090 example.com 127.0.0.1
```
4.View all records:
```
curl http://localhost:3000/records
```

## Troubleshooting
### Connection refused?
- Check firewall rules (allow UDP/TCP on 9090)
- Run as admin/root (required for ports < 1024)
- Verify server is running:
```
netstat -anu | grep 9090  # Linux
Get-NetUDPEndpoint -LocalPort 9090  # PowerShell
```
### Web interface not working?
- Ensure Fiber server is running on port 3000
- Check for port conflicts

## Things to keep in mind:
- In-memory storage (records lost on restart)
- Only handles A records
  
