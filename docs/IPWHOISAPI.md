# \IPWHOISAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IpWhois**](IPWHOISAPI.md#IpWhois) | **Get** /v1.0/ip-whois | IP WHOIS Lookup



## IpWhois

> IpWhoisResponse IpWhois(ctx).Ip(ip).Format(format).Execute()

IP WHOIS Lookup



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
	ip := "ip_example" // string | 
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPWHOISAPI.IpWhois(context.Background()).Ip(ip).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPWHOISAPI.IpWhois``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IpWhois`: IpWhoisResponse
	fmt.Fprintf(os.Stdout, "Response from `IPWHOISAPI.IpWhois`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpWhoisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**IpWhoisResponse**](IpWhoisResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

