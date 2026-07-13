# \DomainAvailabilityAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDomainAvailabilityV2**](DomainAvailabilityAPI.md#BulkDomainAvailabilityV2) | **Post** /v2.0/domain/availability | Bulk Domain Availability Check
[**DomainAvailabilityV2**](DomainAvailabilityAPI.md#DomainAvailabilityV2) | **Get** /v2.0/domain/availability | Domain Availability Check with Suggestions



## BulkDomainAvailabilityV2

> BulkDomainAvailabilityResponse BulkDomainAvailabilityV2(ctx).ApiKey(apiKey).BulkDomainAvailabilityRequest(bulkDomainAvailabilityRequest).Domain(domain).Format(format).Execute()

Bulk Domain Availability Check



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
	bulkDomainAvailabilityRequest := *openapiclient.NewBulkDomainAvailabilityRequest() // BulkDomainAvailabilityRequest | 
	domain := "google.com" // string | Required for TLD-mode bulk check (base domain). (optional)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAvailabilityAPI.BulkDomainAvailabilityV2(context.Background()).ApiKey(apiKey).BulkDomainAvailabilityRequest(bulkDomainAvailabilityRequest).Domain(domain).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAvailabilityAPI.BulkDomainAvailabilityV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDomainAvailabilityV2`: BulkDomainAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainAvailabilityAPI.BulkDomainAvailabilityV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDomainAvailabilityV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **bulkDomainAvailabilityRequest** | [**BulkDomainAvailabilityRequest**](BulkDomainAvailabilityRequest.md) |  | 
 **domain** | **string** | Required for TLD-mode bulk check (base domain). | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**BulkDomainAvailabilityResponse**](BulkDomainAvailabilityResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainAvailabilityV2

> DomainAvailabilityResponse DomainAvailabilityV2(ctx).ApiKey(apiKey).Domain(domain).Sug(sug).Count(count).Format(format).Execute()

Domain Availability Check with Suggestions



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
	domain := "google.com" // string | The domain name to check
	sug := true // bool | Whether to return TLD suggestions alongside the queried domain. (optional) (default to false)
	count := int32(56) // int32 | Number of TLD suggestions to return when sug=true. Maximum is 100. (optional) (default to 5)
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAvailabilityAPI.DomainAvailabilityV2(context.Background()).ApiKey(apiKey).Domain(domain).Sug(sug).Count(count).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAvailabilityAPI.DomainAvailabilityV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainAvailabilityV2`: DomainAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainAvailabilityAPI.DomainAvailabilityV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainAvailabilityV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **domain** | **string** | The domain name to check | 
 **sug** | **bool** | Whether to return TLD suggestions alongside the queried domain. | [default to false]
 **count** | **int32** | Number of TLD suggestions to return when sug&#x3D;true. Maximum is 100. | [default to 5]
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**DomainAvailabilityResponse**](DomainAvailabilityResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

