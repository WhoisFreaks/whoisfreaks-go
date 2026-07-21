# \AccountAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountUsage**](AccountAPI.md#AccountUsage) | **Get** /v1.0/whoisapi/usage | Account Usage
[**DatabaseFileStatus**](AccountAPI.md#DatabaseFileStatus) | **Get** /v3.4/status | Database File Status (Public)
[**RotateApiKey**](AccountAPI.md#RotateApiKey) | **Get** /v1.0/api-key/rotate | Rotate API Key



## AccountUsage

> AccountUsageResponse AccountUsage(ctx).Execute()

Account Usage



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountUsage`: AccountUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountUsageRequest struct via the builder pattern


### Return type

[**AccountUsageResponse**](AccountUsageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabaseFileStatus

> DatabaseFileStatus DatabaseFileStatus(ctx).Execute()

Database File Status (Public)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.DatabaseFileStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DatabaseFileStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatabaseFileStatus`: DatabaseFileStatus
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DatabaseFileStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseFileStatusRequest struct via the builder pattern


### Return type

[**DatabaseFileStatus**](DatabaseFileStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateApiKey

> string RotateApiKey(ctx).Execute()

Rotate API Key



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.RotateApiKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RotateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateApiKey`: string
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RotateApiKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRotateApiKeyRequest struct via the builder pattern


### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

