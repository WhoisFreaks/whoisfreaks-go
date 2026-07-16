# \WHOISAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkWhois**](WHOISAPI.md#BulkWhois) | **Post** /v2.0/bulkwhois/live | Bulk WHOIS Lookup
[**WhoisHistory**](WHOISAPI.md#WhoisHistory) | **Get** /v2.0/whois/history | Historical WHOIS records for a domain
[**WhoisLive**](WHOISAPI.md#WhoisLive) | **Get** /v2.0/whois/live | Live WHOIS Lookup
[**WhoisReverse**](WHOISAPI.md#WhoisReverse) | **Get** /v2.0/whois/reverse | Reverse WHOIS lookup by keyword



## BulkWhois

> BulkWhoisResponse BulkWhois(ctx).BulkWhoisRequest(bulkWhoisRequest).Format(format).Execute()

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
	bulkWhoisRequest := *openapiclient.NewBulkWhoisRequest([]string{"DomainNames_example"}) // BulkWhoisRequest | 
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WHOISAPI.BulkWhois(context.Background()).BulkWhoisRequest(bulkWhoisRequest).Format(format).Execute()
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


## WhoisHistory

> WhoisHistoricalResponse WhoisHistory(ctx).DomainName(domainName).Page(page).Format(format).Execute()

Historical WHOIS records for a domain

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
	domainName := "domainName_example" // string | Domain to fetch historical WHOIS records for
	page := int32(56) // int32 | Page number (optional)
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WHOISAPI.WhoisHistory(context.Background()).DomainName(domainName).Page(page).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WHOISAPI.WhoisHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WhoisHistory`: WhoisHistoricalResponse
	fmt.Fprintf(os.Stdout, "Response from `WHOISAPI.WhoisHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWhoisHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** | Domain to fetch historical WHOIS records for | 
 **page** | **int32** | Page number | 
 **format** | **string** |  | 

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

> WhoisResponse WhoisLive(ctx).DomainName(domainName).Format(format).Execute()

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
	domainName := "whoisfreaks.com" // string | 
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WHOISAPI.WhoisLive(context.Background()).DomainName(domainName).Format(format).Execute()
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


## WhoisReverse

> ReverseWhoisResponse WhoisReverse(ctx).Keyword(keyword).Page(page).Format(format).Execute()

Reverse WHOIS lookup by keyword

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
	keyword := "keyword_example" // string | Keyword to search across WHOIS records
	page := int32(56) // int32 | Page number (optional)
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WHOISAPI.WhoisReverse(context.Background()).Keyword(keyword).Page(page).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WHOISAPI.WhoisReverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WhoisReverse`: ReverseWhoisResponse
	fmt.Fprintf(os.Stdout, "Response from `WHOISAPI.WhoisReverse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWhoisReverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyword** | **string** | Keyword to search across WHOIS records | 
 **page** | **int32** | Page number | 
 **format** | **string** |  | 

### Return type

[**ReverseWhoisResponse**](ReverseWhoisResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

