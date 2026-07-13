# \DatabasesWHOISAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbWhoisDaily**](DatabasesWHOISAPI.md#DbWhoisDaily) | **Get** /v3.3/download/dbupdate/daily/domains/whois | WHOIS Database Daily
[**DbWhoisMonthly**](DatabasesWHOISAPI.md#DbWhoisMonthly) | **Get** /v3.3/download/dbupdate/monthly/domains/whois | WHOIS Database Monthly
[**DbWhoisWeekly**](DatabasesWHOISAPI.md#DbWhoisWeekly) | **Get** /v3.3/download/dbupdate/weekly/domains/whois | WHOIS Database Weekly



## DbWhoisDaily

> *os.File DbWhoisDaily(ctx).ApiKey(apiKey).Date(date).Execute()

WHOIS Database Daily



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
	resp, r, err := apiClient.DatabasesWHOISAPI.DbWhoisDaily(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesWHOISAPI.DbWhoisDaily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbWhoisDaily`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesWHOISAPI.DbWhoisDaily`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbWhoisDailyRequest struct via the builder pattern


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


## DbWhoisMonthly

> *os.File DbWhoisMonthly(ctx).ApiKey(apiKey).Date(date).Execute()

WHOIS Database Monthly



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
	resp, r, err := apiClient.DatabasesWHOISAPI.DbWhoisMonthly(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesWHOISAPI.DbWhoisMonthly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbWhoisMonthly`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesWHOISAPI.DbWhoisMonthly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbWhoisMonthlyRequest struct via the builder pattern


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


## DbWhoisWeekly

> *os.File DbWhoisWeekly(ctx).ApiKey(apiKey).Date(date).Execute()

WHOIS Database Weekly



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
	resp, r, err := apiClient.DatabasesWHOISAPI.DbWhoisWeekly(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesWHOISAPI.DbWhoisWeekly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbWhoisWeekly`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesWHOISAPI.DbWhoisWeekly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbWhoisWeeklyRequest struct via the builder pattern


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

