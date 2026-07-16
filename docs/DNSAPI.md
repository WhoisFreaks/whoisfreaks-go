# \DNSAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsBulk**](DNSAPI.md#DnsBulk) | **Post** /v2.0/dns/bulk/live | Bulk DNS Lookup
[**DnsHistorical**](DNSAPI.md#DnsHistorical) | **Get** /v2.0/dns/historical | Historical DNS Lookup
[**DnsLive**](DNSAPI.md#DnsLive) | **Get** /v2.0/dns/live | Live DNS Lookup
[**DnsReverse**](DNSAPI.md#DnsReverse) | **Get** /v2.1/dns/reverse | Reverse DNS Lookup



## DnsBulk

> BulkDnsResponse DnsBulk(ctx).Type_(type_).DnsBulkRequest(dnsBulkRequest).Format(format).Execute()

Bulk DNS Lookup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/WhoisFreaks/whoisfreaks-go"
)

func main() {
	type_ := "type__example" // string |  (default to "all")
	dnsBulkRequest := *openapiclient.NewDnsBulkRequest() // DnsBulkRequest | 
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.DnsBulk(context.Background()).Type_(type_).DnsBulkRequest(dnsBulkRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.DnsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsBulk`: BulkDnsResponse
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.DnsBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | [default to &quot;all&quot;]
 **dnsBulkRequest** | [**DnsBulkRequest**](DnsBulkRequest.md) |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**BulkDnsResponse**](BulkDnsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsHistorical

> HistoricalDnsResponse DnsHistorical(ctx).DomainName(domainName).Type_(type_).Page(page).Format(format).Execute()

Historical DNS Lookup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/WhoisFreaks/whoisfreaks-go"
)

func main() {
	domainName := "domainName_example" // string | 
	type_ := "type__example" // string |  (default to "all")
	page := int32(56) // int32 |  (optional) (default to 1)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.DnsHistorical(context.Background()).DomainName(domainName).Type_(type_).Page(page).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.DnsHistorical``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsHistorical`: HistoricalDnsResponse
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.DnsHistorical`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsHistoricalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** |  | 
 **type_** | **string** |  | [default to &quot;all&quot;]
 **page** | **int32** |  | [default to 1]
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**HistoricalDnsResponse**](HistoricalDnsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsLive

> DnsResponse DnsLive(ctx).Type_(type_).DomainName(domainName).IpAddress(ipAddress).Format(format).Execute()

Live DNS Lookup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/WhoisFreaks/whoisfreaks-go"
)

func main() {
	type_ := "type__example" // string | all or comma-separated: A,MX,NS,TXT,SOA,SPF,AAAA,CNAME (default to "all")
	domainName := "domainName_example" // string |  (optional)
	ipAddress := "ipAddress_example" // string | Use for PTR lookups (optional)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.DnsLive(context.Background()).Type_(type_).DomainName(domainName).IpAddress(ipAddress).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.DnsLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsLive`: DnsResponse
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.DnsLive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | all or comma-separated: A,MX,NS,TXT,SOA,SPF,AAAA,CNAME | [default to &quot;all&quot;]
 **domainName** | **string** |  | 
 **ipAddress** | **string** | Use for PTR lookups | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**DnsResponse**](DnsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsReverse

> ReverseDnsResponse DnsReverse(ctx).Value(value).Type_(type_).Exact(exact).Page(page).Format(format).Execute()

Reverse DNS Lookup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/WhoisFreaks/whoisfreaks-go"
)

func main() {
	value := "value_example" // string | IP, CIDR, or record value
	type_ := "type__example" // string | 
	exact := true // bool |  (optional) (default to true)
	page := int32(56) // int32 |  (optional) (default to 1)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.DnsReverse(context.Background()).Value(value).Type_(type_).Exact(exact).Page(page).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.DnsReverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsReverse`: ReverseDnsResponse
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.DnsReverse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsReverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** | IP, CIDR, or record value | 
 **type_** | **string** |  | 
 **exact** | **bool** |  | [default to true]
 **page** | **int32** |  | [default to 1]
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**ReverseDnsResponse**](ReverseDnsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

