# \GeolocationAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkGeolocation**](GeolocationAPI.md#BulkGeolocation) | **Post** /v1.0/geolocation | Bulk IP Geolocation
[**Geolocation**](GeolocationAPI.md#Geolocation) | **Get** /v1.0/geolocation | IP Geolocation Lookup



## BulkGeolocation

> []GeolocationResponse BulkGeolocation(ctx).ApiKey(apiKey).BulkGeolocationRequest(bulkGeolocationRequest).Execute()

Bulk IP Geolocation



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
	bulkGeolocationRequest := *openapiclient.NewBulkGeolocationRequest([]string{"Ips_example"}) // BulkGeolocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeolocationAPI.BulkGeolocation(context.Background()).ApiKey(apiKey).BulkGeolocationRequest(bulkGeolocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeolocationAPI.BulkGeolocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkGeolocation`: []GeolocationResponse
	fmt.Fprintf(os.Stdout, "Response from `GeolocationAPI.BulkGeolocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGeolocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **bulkGeolocationRequest** | [**BulkGeolocationRequest**](BulkGeolocationRequest.md) |  | 

### Return type

[**[]GeolocationResponse**](GeolocationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Geolocation

> GeolocationResponse Geolocation(ctx).ApiKey(apiKey).Ip(ip).Execute()

IP Geolocation Lookup



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
	ip := "ip_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeolocationAPI.Geolocation(context.Background()).ApiKey(apiKey).Ip(ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeolocationAPI.Geolocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Geolocation`: GeolocationResponse
	fmt.Fprintf(os.Stdout, "Response from `GeolocationAPI.Geolocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeolocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **ip** | **string** |  | 

### Return type

[**GeolocationResponse**](GeolocationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

