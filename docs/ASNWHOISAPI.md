# \ASNWHOISAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsnWhois**](ASNWHOISAPI.md#AsnWhois) | **Get** /v2.0/asn-whois | ASN WHOIS Lookup



## AsnWhois

> AsnWhoisResponse AsnWhois(ctx).Asn(asn).Format(format).Execute()

ASN WHOIS Lookup



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
	asn := "as15169" // string | 
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ASNWHOISAPI.AsnWhois(context.Background()).Asn(asn).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNWHOISAPI.AsnWhois``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsnWhois`: AsnWhoisResponse
	fmt.Fprintf(os.Stdout, "Response from `ASNWHOISAPI.AsnWhois`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsnWhoisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asn** | **string** |  | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**AsnWhoisResponse**](AsnWhoisResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

