# \SubdomainsAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Subdomains**](SubdomainsAPI.md#Subdomains) | **Get** /v1.0/subdomains | Subdomains Lookup



## Subdomains

> SubdomainsResponse Subdomains(ctx).Domain(domain).After(after).Before(before).Status(status).Page(page).Format(format).Execute()

Subdomains Lookup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/WhoisFreaks/whoisfreaks-go"
)

func main() {
	domain := "domain_example" // string | 
	after := time.Now() // string |  (optional)
	before := time.Now() // string |  (optional)
	status := "status_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainsAPI.Subdomains(context.Background()).Domain(domain).After(after).Before(before).Status(status).Page(page).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainsAPI.Subdomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Subdomains`: SubdomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainsAPI.Subdomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** |  | 
 **after** | **string** |  | 
 **before** | **string** |  | 
 **status** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**SubdomainsResponse**](SubdomainsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

