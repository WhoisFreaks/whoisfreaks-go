# \DatabasesASNWHOISAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbAsnWhois**](DatabasesASNWHOISAPI.md#DbAsnWhois) | **Get** /v3.3/download/snapshot/asn/whois | ASN WHOIS Snapshot
[**DbAsnWhoisStatus**](DatabasesASNWHOISAPI.md#DbAsnWhoisStatus) | **Get** /v3.3/status/snapshot/asn/whois | ASN WHOIS Snapshot Status



## DbAsnWhois

> *os.File DbAsnWhois(ctx).Date(date).Execute()

ASN WHOIS Snapshot



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
	date := time.Now() // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesASNWHOISAPI.DbAsnWhois(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesASNWHOISAPI.DbAsnWhois``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbAsnWhois`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesASNWHOISAPI.DbAsnWhois`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbAsnWhoisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## DbAsnWhoisStatus

> SnapshotStatus DbAsnWhoisStatus(ctx).Execute()

ASN WHOIS Snapshot Status



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
	resp, r, err := apiClient.DatabasesASNWHOISAPI.DbAsnWhoisStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesASNWHOISAPI.DbAsnWhoisStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbAsnWhoisStatus`: SnapshotStatus
	fmt.Fprintf(os.Stdout, "Response from `DatabasesASNWHOISAPI.DbAsnWhoisStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDbAsnWhoisStatusRequest struct via the builder pattern


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

