# Go Stdlib Vulnerability Test

This is a test repository to demonstrate how Dependabot reports Go standard library vulnerabilities.

## Primary Target: GO-2021-0226

This project specifically targets **GO-2021-0226** - a Cross-site Scripting (XSS) vulnerability in Go's standard library.

### GO-2021-0226 Details

- **Affected Package:** Go stdlib (standard library)
- **Vulnerable Versions:** Go < 1.14.8 and Go 1.15.0
- **Affected Modules:** `net/http/cgi` and `net/http/fcgi`
- **CVE:** Related to XSS vulnerability when Content-Type headers are not explicitly set
- **Fix:** Fixed in Go 1.14.8 and 1.15.1+

**This repository uses Go 1.15.0** which is vulnerable to GO-2021-0226.

## Additional Vulnerabilities

This project also contains other known vulnerabilities:

- **GO-2022-0434** (CVE-2022-27536) - crypto/tls certificate verification panic
- **GO-2022-0435** (CVE-2022-28327) - crypto/elliptic P256 panic
- **GO-2022-0520** (CVE-2022-32148) - net/http/httputil X-Forwarded-For leak
- **Vulnerable third-party dependencies:**
  - `github.com/gin-gonic/gin v1.7.0`
  - `github.com/gorilla/websocket v1.4.0`
  - `golang.org/x/crypto` (old version)
  - `golang.org/x/net` (old version)

## How Dependabot Reports Stdlib Vulnerabilities

Dependabot scans the `go` version directive in `go.mod` and checks if that Go version has known vulnerabilities in its standard library. When it detects vulnerable stdlib versions, it creates security alerts.

## Files

- **go.mod** - Specifies Go 1.15.0 and vulnerable dependencies
- **go.sum** - Checksums for all dependencies (required for Dependabot scanning)
- **main.go** - Sample code that uses vulnerable packages including `net/http/cgi` and `net/http/fcgi`

## Testing

To trigger Dependabot alerts:
1. Push this repository to GitHub
2. Enable Dependabot security updates
3. Wait for Dependabot to scan
4. Check the Security tab for alerts

The key is the `go 1.15` directive in `go.mod` - this version (1.15.0) is vulnerable to GO-2021-0226.

## References

- [OSV GO-2021-0226](https://osv.dev/vulnerability/GO-2021-0226)
- [Go Vulnerability Database](https://vuln.go.dev/)
