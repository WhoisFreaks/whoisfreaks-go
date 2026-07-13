# \DatabasesIPGeolocationAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbIpCity**](DatabasesIPGeolocationAPI.md#DbIpCity) | **Get** /v3.3/download/snapshot/ip/city | IP to City Snapshot
[**DbIpCityStatus**](DatabasesIPGeolocationAPI.md#DbIpCityStatus) | **Get** /v3.3/status/snapshot/ip/city | IP to City Snapshot Status
[**DbIpCountry**](DatabasesIPGeolocationAPI.md#DbIpCountry) | **Get** /v3.3/download/snapshot/ip/country | IP to Country Snapshot
[**DbIpCountryStatus**](DatabasesIPGeolocationAPI.md#DbIpCountryStatus) | **Get** /v3.3/status/snapshot/ip/country | IP to Country Snapshot Status



## DbIpCity

> *os.File DbIpCity(ctx).ApiKey(apiKey).Date(date).Execute()

IP to City Snapshot



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
	apiKey := "apiKey_example" // string | Your WHOISFreaks API key
	date := time.Now() // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesIPGeolocationAPI.DbIpCity(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesIPGeolocationAPI.DbIpCity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbIpCity`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesIPGeolocationAPI.DbIpCity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbIpCityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbIpCityStatus

> SnapshotStatus DbIpCityStatus(ctx).ApiKey(apiKey).Execute()

IP to City Snapshot Status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesIPGeolocationAPI.DbIpCityStatus(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesIPGeolocationAPI.DbIpCityStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbIpCityStatus`: SnapshotStatus
	fmt.Fprintf(os.Stdout, "Response from `DatabasesIPGeolocationAPI.DbIpCityStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbIpCityStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 

### Return type

[**SnapshotStatus**](SnapshotStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbIpCountry

> *os.File DbIpCountry(ctx).ApiKey(apiKey).Date(date).Execute()

IP to Country Snapshot



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
	apiKey := "apiKey_example" // string | Your WHOISFreaks API key
	date := time.Now() // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesIPGeolocationAPI.DbIpCountry(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesIPGeolocationAPI.DbIpCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbIpCountry`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesIPGeolocationAPI.DbIpCountry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbIpCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbIpCountryStatus

> SnapshotStatus DbIpCountryStatus(ctx).ApiKey(apiKey).Execute()

IP to Country Snapshot Status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesIPGeolocationAPI.DbIpCountryStatus(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesIPGeolocationAPI.DbIpCountryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbIpCountryStatus`: SnapshotStatus
	fmt.Fprintf(os.Stdout, "Response from `DatabasesIPGeolocationAPI.DbIpCountryStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbIpCountryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 

### Return type

[**SnapshotStatus**](SnapshotStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

