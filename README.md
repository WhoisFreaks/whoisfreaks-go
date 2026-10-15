# WhoisFreaks Go SDK: WHOIS API Client for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/WhoisFreaks/whoisfreaks-go.svg)](https://pkg.go.dev/github.com/WhoisFreaks/whoisfreaks-go)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![API status](https://img.shields.io/badge/API%20status-live-brightgreen.svg)](https://status.whois.whoisfreaks.com/)

The official Go client for the [WhoisFreaks WHOIS API](https://whoisfreaks.com/products/whois-api). Look up live WHOIS records for any domain across 1,529+ TLDs, pull WHOIS history back to 1986, search registrants with reverse WHOIS and run bulk WHOIS lookups from Go code.

The same package covers the [DNS lookup API](https://whoisfreaks.com/products/dns-checker-api), [domain availability API](https://whoisfreaks.com/products/domain-availability-api), [typosquatting API](https://whoisfreaks.com/products/domain-typosquats-api), [SSL certificate API](https://whoisfreaks.com/products/ssl-certificate-api), IP geolocation and IP reputation.

A live WHOIS lookup returns structured JSON like this (trimmed):

```json
{
  "status": true,
  "domain_name": "whoisfreaks.com",
  "domain_registered": "yes",
  "create_date": "2019-03-19",
  "update_date": "2026-02-17",
  "expiry_date": "2027-03-19",
  "domain_registrar": {
    "iana_id": "1068",
    "registrar_name": "NAMECHEAP INC",
    "normalized_name": "namecheap, inc.",
    "website_url": "http://www.namecheap.com"
  },
  "registrant_contact": {
    "city": "Lewes",
    "state": "Delaware",
    "zip_code": "19958",
    "country_name": "United States",
    "country_code": "US"
  },
  "name_servers": ["albert.ns.cloudflare.com", "sue.ns.cloudflare.com"],
  "domain_status": ["clienttransferprohibited"]
}
```

## Contents

- [Installation](#installation)
- [Get an API key](#get-an-api-key)
- [WHOIS lookup](#whois-lookup)
- [Normalized registrar, address and date fields](#normalized-registrar-address-and-date-fields)
- [Bulk WHOIS lookup](#bulk-whois-lookup)
- [WHOIS history](#whois-history)
- [Reverse WHOIS](#reverse-whois)
- [DNS lookup](#dns-lookup)
- [Domain availability check](#domain-availability-check)
- [Typosquatting detection](#typosquatting-detection)
- [Error handling](#error-handling)
- [All endpoints](#all-endpoints)
- [Requirements](#requirements)
- [Support](#support)

## Installation

```sh
go get github.com/WhoisFreaks/whoisfreaks-go@latest
```

## Get an API key

Create a free account at [billing.whoisfreaks.com/signup](https://billing.whoisfreaks.com/signup). New accounts get 500 free credits with no credit card, usable across the live WHOIS, WHOIS history and reverse WHOIS APIs.

Keep the key out of source control. The examples read it from an environment variable:

```sh
export WHOISFREAKS_API_KEY="your_api_key"
```

Every example below assumes this setup:

```go
import (
	"context"
	"errors"
	"fmt"
	"os"

	whoisfreaks "github.com/WhoisFreaks/whoisfreaks-go"
)

ctx := context.WithValue(context.Background(), whoisfreaks.ContextAPIKeys,
	map[string]whoisfreaks.APIKey{"ApiKeyAuth": {Key: os.Getenv("WHOISFREAKS_API_KEY")}})
client := whoisfreaks.NewAPIClient(whoisfreaks.NewConfiguration())
```

## WHOIS lookup

Each request queries WHOIS servers and RDAP live, so you get the current record rather than a cached copy. One lookup costs 1 credit.

```go
record, _, err := client.WHOISAPI.WhoisLive(ctx).DomainName("whoisfreaks.com").Execute()
if err != nil {
	panic(err)
}

fmt.Println(record.GetCreateDate(), record.GetExpiryDate()) // 2019-03-19 2027-03-19
registrar := record.GetDomainRegistrar()
fmt.Println(registrar.GetRegistrarName())                // NAMECHEAP INC
fmt.Println(record.GetNameServers())
```

## Normalized registrar, address and date fields

Registries and registrars write the same facts in different formats. The API returns normalized fields so records from different TLDs line up without your own parsing code.

| What you get | Field | Example |
|---|---|---|
| One registrar name, however the registry spells it | `domain_registrar.normalized_name` | `namecheap, inc.` for both "NAMECHEAP INC" and "NameCheap, Inc" |
| A parsed and geocoded contact address | `registrant_contact.street`, `city`, `state`, `zip_code`, `country_name`, `country_code` | `Lewes`, `Delaware`, `19958`, `United States`, `US` |
| One date format | `create_date`, `update_date`, `expiry_date` | `2019-03-19` |

```go
registrant := record.GetRegistrantContact()

fmt.Println(registrar.GetNormalizedName())                 // namecheap, inc.
fmt.Println(registrant.GetCity(), registrant.GetState())   // Lewes Delaware
fmt.Println(registrant.GetCountryCode())                   // US
```

Use these fields to group a domain portfolio by registrar, filter registrants by country code or sort domains by expiry date across every TLD.

## Bulk WHOIS lookup

Look up to 100 domains in one call.

```go
request := whoisfreaks.NewBulkWhoisRequest([]string{"example.com", "whoisfreaks.com", "github.com"})
result, _, err := client.WHOISAPI.BulkWhois(ctx).BulkWhoisRequest(*request).Execute()
```

For larger jobs, the [Bulk WHOIS API](https://whoisfreaks.com/products/bulk-whois-api) accepts file uploads of up to 3M domains.

## WHOIS history

Get every recorded ownership change for a domain in one request. A new record is stored only when the WHOIS data changed, and history goes back to 1986.

```go
history, _, err := client.WHOISAPI.WhoisHistory(ctx).DomainName("whoisfreaks.com").Execute()
```

## Reverse WHOIS

Find every domain registered with a keyword, such as a company name or registrant email, across 4.5B+ WHOIS records.

```go
matches, _, err := client.WHOISAPI.WhoisReverse(ctx).Keyword("whoisfreaks").Page(1).Execute()
```

## DNS lookup

Pull live DNS records for a domain. Pass `all` or a comma-separated list of record types: `A`, `AAAA`, `MX`, `NS`, `TXT`, `SOA`, `SPF`, `CNAME`.

```go
records, _, err := client.DNSAPI.DnsLive(ctx).Type_("A,MX,TXT").DomainName("whoisfreaks.com").Execute()
```

The same class also covers bulk DNS, historical DNS and reverse DNS lookups.

## Domain availability check

Check whether a domain is registered and, with suggestions turned on, get up to 100 available alternatives on other TLDs.

```go
result, _, err := client.DomainAvailabilityAPI.DomainAvailabilityV2(ctx).
	Domain("whoisfreaks.io").Sug(true).Count(5).Execute()
```

## Typosquatting detection

Find registered domains that look like typo variants of a brand name. Search by a brand `keyword`, or by a wildcard `pattern` such as `whois*` (up to 3 asterisks). Results come 100 domains per page with registration date, expiry date and drop status, across 1,529+ TLDs.

```go
typos, _, err := client.TyposquattingAPI.Typosquatting(ctx).Keyword("whoisfreaks").Execute()
```

To get the next page, pass the `nextPageToken` value from the response as `PageToken`, along with the same keyword or pattern.

## Error handling

Failed requests return an `error`. For HTTP errors it is a `*whoisfreaks.GenericOpenAPIError` carrying the response body, and the returned `*http.Response` has the status code. Network errors have no response.

```go
_, httpResp, err := client.WHOISAPI.WhoisLive(ctx).DomainName("whoisfreaks.com").Execute()
if err != nil {
	var apiErr *whoisfreaks.GenericOpenAPIError
	if errors.As(err, &apiErr) && httpResp != nil {
		fmt.Fprintln(os.Stderr, httpResp.StatusCode, string(apiErr.Body()))
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
}
```

Live API uptime and incident history are on the [status page](https://status.whois.whoisfreaks.com/).

## All endpoints

All URIs are relative to `https://api.whoisfreaks.com`. Each method links to its parameters and response model.

### WHOIS

| Class | Method | HTTP request | Description |
|---|---|---|---|
| WHOISAPI | [BulkWhois](docs/WHOISAPI.md#bulkwhois) | POST /v2.0/bulkwhois/live | Bulk WHOIS Lookup |
| WHOISAPI | [WhoisHistory](docs/WHOISAPI.md#whoishistory) | GET /v2.0/whois/history | Historical WHOIS records for a domain |
| WHOISAPI | [WhoisLive](docs/WHOISAPI.md#whoislive) | GET /v2.0/whois/live | Live WHOIS Lookup |
| WHOISAPI | [WhoisReverse](docs/WHOISAPI.md#whoisreverse) | GET /v2.0/whois/reverse | Reverse WHOIS lookup by keyword |
| ASNWHOISAPI | [AsnWhois](docs/ASNWHOISAPI.md#asnwhois) | GET /v2.0/asn-whois | ASN WHOIS Lookup |
| IPWHOISAPI | [IpWhois](docs/IPWHOISAPI.md#ipwhois) | GET /v1.0/ip-whois | IP WHOIS Lookup |

### DNS, SSL and subdomains

| Class | Method | HTTP request | Description |
|---|---|---|---|
| DNSAPI | [DnsBulk](docs/DNSAPI.md#dnsbulk) | POST /v2.0/dns/bulk/live | Bulk DNS Lookup |
| DNSAPI | [DnsHistorical](docs/DNSAPI.md#dnshistorical) | GET /v2.0/dns/historical | Historical DNS Lookup |
| DNSAPI | [DnsLive](docs/DNSAPI.md#dnslive) | GET /v2.0/dns/live | Live DNS Lookup |
| DNSAPI | [DnsReverse](docs/DNSAPI.md#dnsreverse) | GET /v2.1/dns/reverse | Reverse DNS Lookup |
| SSLAPI | [SslLookup](docs/SSLAPI.md#ssllookup) | GET /v1.0/ssl/live | SSL Certificate Lookup |
| SubdomainsAPI | [Subdomains](docs/SubdomainsAPI.md#subdomains) | GET /v1.0/subdomains | Subdomains Lookup |

### Domain availability, typosquatting and reputation

| Class | Method | HTTP request | Description |
|---|---|---|---|
| DomainAvailabilityAPI | [BulkDomainAvailabilityV2](docs/DomainAvailabilityAPI.md#bulkdomainavailabilityv2) | POST /v2.0/domain/availability | Bulk Domain Availability Check |
| DomainAvailabilityAPI | [DomainAvailabilityV2](docs/DomainAvailabilityAPI.md#domainavailabilityv2) | GET /v2.0/domain/availability | Domain Availability Check with Suggestions |
| TyposquattingAPI | [Typosquatting](docs/TyposquattingAPI.md#typosquatting) | GET /v3.0/domain/typos | Typosquatting Lookup |
| DomainReputationAPI | [DomainReputation](docs/DomainReputationAPI.md#domainreputation) | GET /v1/domain/security | Domain Reputation Lookup |

### IP geolocation and reputation

| Class | Method | HTTP request | Description |
|---|---|---|---|
| GeolocationAPI | [BulkGeolocation](docs/GeolocationAPI.md#bulkgeolocation) | POST /v1.0/geolocation | Bulk IP Geolocation |
| GeolocationAPI | [Geolocation](docs/GeolocationAPI.md#geolocation) | GET /v1.0/geolocation | IP Geolocation Lookup |
| IPReputationAPI | [BulkIpReputation](docs/IPReputationAPI.md#bulkipreputation) | POST /v1.0/security | Bulk IP Reputation |
| IPReputationAPI | [IpReputation](docs/IPReputationAPI.md#ipreputation) | GET /v1.0/security | IP Reputation Lookup |

### Account

| Class | Method | HTTP request | Description |
|---|---|---|---|
| AccountAPI | [AccountUsage](docs/AccountAPI.md#accountusage) | GET /v1.0/whoisapi/usage | Account Usage |
| AccountAPI | [DatabaseFileStatus](docs/AccountAPI.md#databasefilestatus) | GET /v3.4/status | Database File Status (Public) |
| AccountAPI | [RotateApiKey](docs/AccountAPI.md#rotateapikey) | GET /v1.0/api-key/rotate | Rotate API Key |

### Database downloads

WHOIS, DNS, subdomain, IP, newly registered, expiring and dropped domain databases, plus the phishing, malware and spam threat feeds, are available through the `Databases*` classes. See the [full method list](docs/) for every download and status endpoint.

## Requirements

Go 1.18 or later.

## Support

- WHOIS API documentation: [whoisfreaks.com/documentation/whois-api](https://whoisfreaks.com/documentation/whois-api)
- All API documentation: [whoisfreaks.com/documentation](https://whoisfreaks.com/documentation)
- Product page: [WhoisFreaks WHOIS API](https://whoisfreaks.com/products/whois-api)
- API status: [status.whois.whoisfreaks.com](https://status.whois.whoisfreaks.com/)
- Email: support@whoisfreaks.com
- Quick start: [WHOIS API examples](https://github.com/WhoisFreaks/whois-api)
- Other languages: [Python](https://github.com/WhoisFreaks/whoisfreaks-python), [JavaScript](https://github.com/WhoisFreaks/whoisfreaks-javascript), [TypeScript](https://github.com/WhoisFreaks/whoisfreaks-typescript), [Java](https://github.com/WhoisFreaks/whoisfreaks-java), [Go](https://github.com/WhoisFreaks/whoisfreaks-go), [C#](https://github.com/WhoisFreaks/whoisfreaks-csharp), [Ruby](https://github.com/WhoisFreaks/whoisfreaks-ruby), [PHP](https://github.com/WhoisFreaks/whoisfreaks-php), [Swift](https://github.com/WhoisFreaks/whoisfreaks-swift), [Kotlin](https://github.com/WhoisFreaks/whoisfreaks-kotlin)

## License

MIT. See [LICENSE](LICENSE).
