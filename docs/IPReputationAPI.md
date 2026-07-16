# \IPReputationAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkIpReputation**](IPReputationAPI.md#BulkIpReputation) | **Post** /v1.0/security | Bulk IP Reputation
[**IpReputation**](IPReputationAPI.md#IpReputation) | **Get** /v1.0/security | IP Reputation Lookup



## BulkIpReputation

> []IpReputationResponse BulkIpReputation(ctx).BulkGeolocationRequest(bulkGeolocationRequest).Execute()

Bulk IP Reputation



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
	bulkGeolocationRequest := *openapiclient.NewBulkGeolocationRequest([]string{"Ips_example"}) // BulkGeolocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPReputationAPI.BulkIpReputation(context.Background()).BulkGeolocationRequest(bulkGeolocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPReputationAPI.BulkIpReputation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkIpReputation`: []IpReputationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPReputationAPI.BulkIpReputation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkIpReputationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkGeolocationRequest** | [**BulkGeolocationRequest**](BulkGeolocationRequest.md) |  | 

### Return type

[**[]IpReputationResponse**](IpReputationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpReputation

> IpReputationResponse IpReputation(ctx).Ip(ip).Execute()

IP Reputation Lookup



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPReputationAPI.IpReputation(context.Background()).Ip(ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPReputationAPI.IpReputation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IpReputation`: IpReputationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPReputationAPI.IpReputation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpReputationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 

### Return type

[**IpReputationResponse**](IpReputationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

