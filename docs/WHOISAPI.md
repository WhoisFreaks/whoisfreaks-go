# \WHOISAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkWhois**](WHOISAPI.md#BulkWhois) | **Post** /v2.0/bulkwhois/live | Bulk WHOIS Lookup
[**WhoisHistoricalOrReverse**](WHOISAPI.md#WhoisHistoricalOrReverse) | **Get** /v1.0/whois | WHOIS Historical or Reverse Lookup
[**WhoisLive**](WHOISAPI.md#WhoisLive) | **Get** /v2.0/whois/live | Live WHOIS Lookup



## BulkWhois

> BulkWhoisResponse BulkWhois(ctx).ApiKey(apiKey).BulkWhoisRequest(bulkWhoisRequest).Format(format).Execute()

Bulk WHOIS Lookup



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
	apiKey := "apiKey_example" // string | Your WHOISFreaks API key
	bulkWhoisRequest := *openapiclient.NewBulkWhoisRequest([]string{"DomainNames_example"}) // BulkWhoisRequest | 
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WHOISAPI.BulkWhois(context.Background()).ApiKey(apiKey).BulkWhoisRequest(bulkWhoisRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WHOISAPI.BulkWhois``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkWhois`: BulkWhoisResponse
	fmt.Fprintf(os.Stdout, "Response from `WHOISAPI.BulkWhois`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkWhoisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **bulkWhoisRequest** | [**BulkWhoisRequest**](BulkWhoisRequest.md) |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**BulkWhoisResponse**](BulkWhoisResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WhoisHistoricalOrReverse

> WhoisHistoricalResponse WhoisHistoricalOrReverse(ctx).ApiKey(apiKey).Whois(whois).DomainName(domainName).Keyword(keyword).Email(email).Owner(owner).Company(company).Mode(mode).Exact(exact).Page(page).Format(format).Execute()

WHOIS Historical or Reverse Lookup



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
	apiKey := "apiKey_example" // string | Your WHOISFreaks API key
	whois := "whois_example" // string | 
	domainName := "domainName_example" // string | Required for historical lookup (optional)
	keyword := "keyword_example" // string | For reverse — domain keyword search (optional)
	email := "email_example" // string | For reverse — registrant email search (optional)
	owner := "owner_example" // string | For reverse — registrant name search (optional)
	company := "company_example" // string | For reverse — company name search (optional)
	mode := "mode_example" // string |  (optional) (default to "default")
	exact := true // bool |  (optional) (default to true)
	page := int32(56) // int32 |  (optional) (default to 1)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WHOISAPI.WhoisHistoricalOrReverse(context.Background()).ApiKey(apiKey).Whois(whois).DomainName(domainName).Keyword(keyword).Email(email).Owner(owner).Company(company).Mode(mode).Exact(exact).Page(page).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WHOISAPI.WhoisHistoricalOrReverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WhoisHistoricalOrReverse`: WhoisHistoricalResponse
	fmt.Fprintf(os.Stdout, "Response from `WHOISAPI.WhoisHistoricalOrReverse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWhoisHistoricalOrReverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **whois** | **string** |  | 
 **domainName** | **string** | Required for historical lookup | 
 **keyword** | **string** | For reverse — domain keyword search | 
 **email** | **string** | For reverse — registrant email search | 
 **owner** | **string** | For reverse — registrant name search | 
 **company** | **string** | For reverse — company name search | 
 **mode** | **string** |  | [default to &quot;default&quot;]
 **exact** | **bool** |  | [default to true]
 **page** | **int32** |  | [default to 1]
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**WhoisHistoricalResponse**](WhoisHistoricalResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WhoisLive

> WhoisResponse WhoisLive(ctx).ApiKey(apiKey).DomainName(domainName).Format(format).Execute()

Live WHOIS Lookup



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
	apiKey := "apiKey_example" // string | Your WHOISFreaks API key
	domainName := "whoisfreaks.com" // string | 
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WHOISAPI.WhoisLive(context.Background()).ApiKey(apiKey).DomainName(domainName).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WHOISAPI.WhoisLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WhoisLive`: WhoisResponse
	fmt.Fprintf(os.Stdout, "Response from `WHOISAPI.WhoisLive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWhoisLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **domainName** | **string** |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**WhoisResponse**](WhoisResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

