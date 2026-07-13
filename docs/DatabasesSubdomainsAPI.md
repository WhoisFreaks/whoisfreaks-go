# \DatabasesSubdomainsAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbSubdomainsDaily**](DatabasesSubdomainsAPI.md#DbSubdomainsDaily) | **Get** /v3.2/download/dbupdate/daily/subdomains | Subdomains Daily
[**DbSubdomainsMonthly**](DatabasesSubdomainsAPI.md#DbSubdomainsMonthly) | **Get** /v3.2/download/dbupdate/monthly/subdomains | Subdomains Monthly
[**DbSubdomainsWeekly**](DatabasesSubdomainsAPI.md#DbSubdomainsWeekly) | **Get** /v3.2/download/dbupdate/weekly/subdomains | Subdomains Weekly



## DbSubdomainsDaily

> *os.File DbSubdomainsDaily(ctx).ApiKey(apiKey).Date(date).Execute()

Subdomains Daily



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
	date := time.Now() // string | yyyy-MM-dd; omit for latest (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesSubdomainsAPI.DbSubdomainsDaily(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesSubdomainsAPI.DbSubdomainsDaily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbSubdomainsDaily`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesSubdomainsAPI.DbSubdomainsDaily`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbSubdomainsDailyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 

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


## DbSubdomainsMonthly

> *os.File DbSubdomainsMonthly(ctx).ApiKey(apiKey).Date(date).Execute()

Subdomains Monthly



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
	date := time.Now() // string | yyyy-MM-dd; omit for latest (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesSubdomainsAPI.DbSubdomainsMonthly(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesSubdomainsAPI.DbSubdomainsMonthly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbSubdomainsMonthly`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesSubdomainsAPI.DbSubdomainsMonthly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbSubdomainsMonthlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 

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


## DbSubdomainsWeekly

> *os.File DbSubdomainsWeekly(ctx).ApiKey(apiKey).Date(date).Execute()

Subdomains Weekly



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
	date := time.Now() // string | yyyy-MM-dd; omit for latest (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesSubdomainsAPI.DbSubdomainsWeekly(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesSubdomainsAPI.DbSubdomainsWeekly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbSubdomainsWeekly`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesSubdomainsAPI.DbSubdomainsWeekly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbSubdomainsWeeklyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 

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

