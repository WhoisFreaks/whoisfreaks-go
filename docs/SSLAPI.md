# \SSLAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SslLookup**](SSLAPI.md#SslLookup) | **Get** /v1.0/ssl/live | SSL Certificate Lookup



## SslLookup

> SslResponse SslLookup(ctx).DomainName(domainName).Chain(chain).SslRaw(sslRaw).Format(format).Execute()

SSL Certificate Lookup



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
	chain := true // bool |  (optional) (default to false)
	sslRaw := true // bool |  (optional) (default to false)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSLAPI.SslLookup(context.Background()).DomainName(domainName).Chain(chain).SslRaw(sslRaw).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSLAPI.SslLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SslLookup`: SslResponse
	fmt.Fprintf(os.Stdout, "Response from `SSLAPI.SslLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** |  | 
 **chain** | **bool** |  | [default to false]
 **sslRaw** | **bool** |  | [default to false]
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**SslResponse**](SslResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

