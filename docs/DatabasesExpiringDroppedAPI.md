# \DatabasesExpiringDroppedAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbDropped**](DatabasesExpiringDroppedAPI.md#DbDropped) | **Get** /v3.1/download/domainer/dropped | Dropped Domains
[**DbDroppedBacklinks**](DatabasesExpiringDroppedAPI.md#DbDroppedBacklinks) | **Get** /v3.3/download/domainer/dropped/backlinks | Dropped With Backlinks
[**DbDroppedJson**](DatabasesExpiringDroppedAPI.md#DbDroppedJson) | **Get** /v3.1/domains/dropped | Dropped Domains (JSON)
[**DbExpired**](DatabasesExpiringDroppedAPI.md#DbExpired) | **Get** /v3.1/download/domainer/expired | Expiring Domains
[**DbExpiredCleaned**](DatabasesExpiringDroppedAPI.md#DbExpiredCleaned) | **Get** /v3.1/download/domainer/expired/cleaned | Expiring Cleaned WHOIS



## DbDropped

> *os.File DbDropped(ctx).ApiKey(apiKey).Whois(whois).Date(date).Execute()

Dropped Domains



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
	whois := true // bool | 
	date := time.Now() // string | yyyy-MM-dd; omit for latest (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesExpiringDroppedAPI.DbDropped(context.Background()).ApiKey(apiKey).Whois(whois).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesExpiringDroppedAPI.DbDropped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbDropped`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesExpiringDroppedAPI.DbDropped`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbDroppedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **whois** | **bool** |  | 
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


## DbDroppedBacklinks

> *os.File DbDroppedBacklinks(ctx).ApiKey(apiKey).Whois(whois).Date(date).Execute()

Dropped With Backlinks



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
	whois := true // bool |  (optional)
	date := time.Now() // string | yyyy-MM-dd; omit for latest (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesExpiringDroppedAPI.DbDroppedBacklinks(context.Background()).ApiKey(apiKey).Whois(whois).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesExpiringDroppedAPI.DbDroppedBacklinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbDroppedBacklinks`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesExpiringDroppedAPI.DbDroppedBacklinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbDroppedBacklinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **whois** | **bool** |  | 
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


## DbDroppedJson

> []string DbDroppedJson(ctx).ApiKey(apiKey).Date(date).Tlds(tlds).Execute()

Dropped Domains (JSON)



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
	tlds := "tlds_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesExpiringDroppedAPI.DbDroppedJson(context.Background()).ApiKey(apiKey).Date(date).Tlds(tlds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesExpiringDroppedAPI.DbDroppedJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbDroppedJson`: []string
	fmt.Fprintf(os.Stdout, "Response from `DatabasesExpiringDroppedAPI.DbDroppedJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbDroppedJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 
 **tlds** | **string** |  | 

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbExpired

> *os.File DbExpired(ctx).ApiKey(apiKey).Whois(whois).Date(date).Execute()

Expiring Domains



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
	whois := true // bool | 
	date := time.Now() // string | yyyy-MM-dd; omit for latest (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesExpiringDroppedAPI.DbExpired(context.Background()).ApiKey(apiKey).Whois(whois).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesExpiringDroppedAPI.DbExpired``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbExpired`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesExpiringDroppedAPI.DbExpired`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbExpiredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **whois** | **bool** |  | 
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


## DbExpiredCleaned

> *os.File DbExpiredCleaned(ctx).ApiKey(apiKey).Date(date).Execute()

Expiring Cleaned WHOIS



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
	resp, r, err := apiClient.DatabasesExpiringDroppedAPI.DbExpiredCleaned(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesExpiringDroppedAPI.DbExpiredCleaned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbExpiredCleaned`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesExpiringDroppedAPI.DbExpiredCleaned`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbExpiredCleanedRequest struct via the builder pattern


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

