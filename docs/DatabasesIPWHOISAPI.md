# \DatabasesIPWHOISAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbIpWhois**](DatabasesIPWHOISAPI.md#DbIpWhois) | **Get** /v3.3/download/snapshot/ip/whois | IP WHOIS Snapshot
[**DbIpWhoisStatus**](DatabasesIPWHOISAPI.md#DbIpWhoisStatus) | **Get** /v3.3/status/snapshot/ip/whois | IP WHOIS Snapshot Status



## DbIpWhois

> *os.File DbIpWhois(ctx).Date(date).Execute()

IP WHOIS Snapshot



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
	resp, r, err := apiClient.DatabasesIPWHOISAPI.DbIpWhois(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesIPWHOISAPI.DbIpWhois``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbIpWhois`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesIPWHOISAPI.DbIpWhois`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbIpWhoisRequest struct via the builder pattern


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


## DbIpWhoisStatus

> SnapshotStatus DbIpWhoisStatus(ctx).Execute()

IP WHOIS Snapshot Status



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
	resp, r, err := apiClient.DatabasesIPWHOISAPI.DbIpWhoisStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesIPWHOISAPI.DbIpWhoisStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbIpWhoisStatus`: SnapshotStatus
	fmt.Fprintf(os.Stdout, "Response from `DatabasesIPWHOISAPI.DbIpWhoisStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDbIpWhoisStatusRequest struct via the builder pattern


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

