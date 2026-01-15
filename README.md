# Go Stdlib Vulnerability Test

This is a test repository to demonstrate how Dependabot reports Go standard library vulnerabilities.

## Purpose

This project uses **Go 1.17** which contains multiple known vulnerabilities in the Go standard library:

- **GO-2022-0434** (CVE-2022-27536) - crypto/tls certificate verification panic
- **GO-2022-0435** (CVE-2022-28327) - crypto/elliptic P256 panic
- **GO-2022-0520** (CVE-2022-32148) - net/http/httputil X-Forwarded-For leak
- **GO-2022-0521** - Additional stdlib vulnerability
- **GO-2022-0523** - Additional stdlib vulnerability
- **GO-2022-0525** - Additional stdlib vulnerability
- **GO-2022-0532** - Additional stdlib vulnerability

## How Dependabot Reports Stdlib Vulnerabilities

Dependabot scans the `go` version directive in `go.mod` and checks if that Go version has known vulnerabilities in its standard library. When it detects vulnerable stdlib versions, it creates security alerts.

## Testing

To trigger Dependabot alerts:
1. Push this repository to GitHub
2. Enable Dependabot security updates
3. Wait for Dependabot to scan
4. Check the Security tab for alerts

The key is the `go 1.17` directive in `go.mod` - this old version has multiple stdlib vulnerabilities that Dependabot will detect.
