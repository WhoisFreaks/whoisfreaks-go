# \DomainReputationAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainReputation**](DomainReputationAPI.md#DomainReputation) | **Get** /v1/domain/security | Domain Reputation Lookup



## DomainReputation

> DomainReputationResponse DomainReputation(ctx).ApiKey(apiKey).DomainName(domainName).Format(format).Execute()

Domain Reputation Lookup



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
	domainName := "amazon.com" // string | The domain name to assess
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainReputationAPI.DomainReputation(context.Background()).ApiKey(apiKey).DomainName(domainName).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainReputationAPI.DomainReputation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainReputation`: DomainReputationResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainReputationAPI.DomainReputation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainReputationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **domainName** | **string** | The domain name to assess | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**DomainReputationResponse**](DomainReputationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

