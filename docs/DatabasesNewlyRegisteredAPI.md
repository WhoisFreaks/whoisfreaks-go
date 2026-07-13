# \DatabasesNewlyRegisteredAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbNewlyCctld**](DatabasesNewlyRegisteredAPI.md#DbNewlyCctld) | **Get** /v3.1/download/domainer/cctld | Newly Registered ccTLD (CSV)
[**DbNewlyCctldCleaned**](DatabasesNewlyRegisteredAPI.md#DbNewlyCctldCleaned) | **Get** /v3.1/download/domainer/cctld/cleaned | Newly Registered ccTLD Cleaned WHOIS (CSV)
[**DbNewlyCctldJson**](DatabasesNewlyRegisteredAPI.md#DbNewlyCctldJson) | **Get** /v3.1/domains/newly/cctld | Newly Registered ccTLD (JSON)
[**DbNewlyDns**](DatabasesNewlyRegisteredAPI.md#DbNewlyDns) | **Get** /v3.1/download/domainer/newly/dns | Newly Registered With DNS
[**DbNewlyGtld**](DatabasesNewlyRegisteredAPI.md#DbNewlyGtld) | **Get** /v3.1/download/domainer/gtld | Newly Registered gTLD (CSV)
[**DbNewlyGtldCleaned**](DatabasesNewlyRegisteredAPI.md#DbNewlyGtldCleaned) | **Get** /v3.1/download/domainer/gtld/cleaned | Newly Registered gTLD Cleaned WHOIS (CSV)
[**DbNewlyGtldJson**](DatabasesNewlyRegisteredAPI.md#DbNewlyGtldJson) | **Get** /v3.1/domains/newly/gtld | Newly Registered gTLD (JSON)



## DbNewlyCctld

> string DbNewlyCctld(ctx).ApiKey(apiKey).Whois(whois).Date(date).Tlds(tlds).Execute()

Newly Registered ccTLD (CSV)



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
	tlds := "tlds_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesNewlyRegisteredAPI.DbNewlyCctld(context.Background()).ApiKey(apiKey).Whois(whois).Date(date).Tlds(tlds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesNewlyRegisteredAPI.DbNewlyCctld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbNewlyCctld`: string
	fmt.Fprintf(os.Stdout, "Response from `DatabasesNewlyRegisteredAPI.DbNewlyCctld`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbNewlyCctldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **whois** | **bool** |  | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 
 **tlds** | **string** |  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbNewlyCctldCleaned

> string DbNewlyCctldCleaned(ctx).ApiKey(apiKey).Date(date).Execute()

Newly Registered ccTLD Cleaned WHOIS (CSV)



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
	resp, r, err := apiClient.DatabasesNewlyRegisteredAPI.DbNewlyCctldCleaned(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesNewlyRegisteredAPI.DbNewlyCctldCleaned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbNewlyCctldCleaned`: string
	fmt.Fprintf(os.Stdout, "Response from `DatabasesNewlyRegisteredAPI.DbNewlyCctldCleaned`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbNewlyCctldCleanedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbNewlyCctldJson

> []string DbNewlyCctldJson(ctx).ApiKey(apiKey).Date(date).Tlds(tlds).Execute()

Newly Registered ccTLD (JSON)



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
	resp, r, err := apiClient.DatabasesNewlyRegisteredAPI.DbNewlyCctldJson(context.Background()).ApiKey(apiKey).Date(date).Tlds(tlds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesNewlyRegisteredAPI.DbNewlyCctldJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbNewlyCctldJson`: []string
	fmt.Fprintf(os.Stdout, "Response from `DatabasesNewlyRegisteredAPI.DbNewlyCctldJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbNewlyCctldJsonRequest struct via the builder pattern


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


## DbNewlyDns

> []map[string]interface{} DbNewlyDns(ctx).ApiKey(apiKey).Date(date).Execute()

Newly Registered With DNS



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
	resp, r, err := apiClient.DatabasesNewlyRegisteredAPI.DbNewlyDns(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesNewlyRegisteredAPI.DbNewlyDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbNewlyDns`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesNewlyRegisteredAPI.DbNewlyDns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbNewlyDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 

### Return type

**[]map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbNewlyGtld

> string DbNewlyGtld(ctx).ApiKey(apiKey).Whois(whois).Date(date).Tlds(tlds).Execute()

Newly Registered gTLD (CSV)



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
	tlds := "tlds_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesNewlyRegisteredAPI.DbNewlyGtld(context.Background()).ApiKey(apiKey).Whois(whois).Date(date).Tlds(tlds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesNewlyRegisteredAPI.DbNewlyGtld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbNewlyGtld`: string
	fmt.Fprintf(os.Stdout, "Response from `DatabasesNewlyRegisteredAPI.DbNewlyGtld`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbNewlyGtldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **whois** | **bool** |  | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 
 **tlds** | **string** |  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbNewlyGtldCleaned

> string DbNewlyGtldCleaned(ctx).ApiKey(apiKey).Date(date).Execute()

Newly Registered gTLD Cleaned WHOIS (CSV)



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
	resp, r, err := apiClient.DatabasesNewlyRegisteredAPI.DbNewlyGtldCleaned(context.Background()).ApiKey(apiKey).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesNewlyRegisteredAPI.DbNewlyGtldCleaned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbNewlyGtldCleaned`: string
	fmt.Fprintf(os.Stdout, "Response from `DatabasesNewlyRegisteredAPI.DbNewlyGtldCleaned`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbNewlyGtldCleanedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **date** | **string** | yyyy-MM-dd; omit for latest | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbNewlyGtldJson

> []string DbNewlyGtldJson(ctx).ApiKey(apiKey).Date(date).Tlds(tlds).Execute()

Newly Registered gTLD (JSON)



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
	resp, r, err := apiClient.DatabasesNewlyRegisteredAPI.DbNewlyGtldJson(context.Background()).ApiKey(apiKey).Date(date).Tlds(tlds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesNewlyRegisteredAPI.DbNewlyGtldJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbNewlyGtldJson`: []string
	fmt.Fprintf(os.Stdout, "Response from `DatabasesNewlyRegisteredAPI.DbNewlyGtldJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbNewlyGtldJsonRequest struct via the builder pattern


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

