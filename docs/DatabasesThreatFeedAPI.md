# \DatabasesThreatFeedAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadThreatFeedMalware**](DatabasesThreatFeedAPI.md#DownloadThreatFeedMalware) | **Get** /v3.4/download/threat-feed/malware | Download the daily malware threat feed (CSV)
[**DownloadThreatFeedMalwareSample**](DatabasesThreatFeedAPI.md#DownloadThreatFeedMalwareSample) | **Get** /v3.4/download/threat-feed/malware/sample | Download a sample of the malware threat feed (CSV)
[**DownloadThreatFeedPhishing**](DatabasesThreatFeedAPI.md#DownloadThreatFeedPhishing) | **Get** /v3.4/download/threat-feed/phishing | Download the daily phishing threat feed (CSV)
[**DownloadThreatFeedPhishingSample**](DatabasesThreatFeedAPI.md#DownloadThreatFeedPhishingSample) | **Get** /v3.4/download/threat-feed/phishing/sample | Download a sample of the phishing threat feed (CSV)
[**DownloadThreatFeedSpam**](DatabasesThreatFeedAPI.md#DownloadThreatFeedSpam) | **Get** /v3.4/download/threat-feed/spam | Download the daily spam threat feed (CSV)
[**DownloadThreatFeedSpamSample**](DatabasesThreatFeedAPI.md#DownloadThreatFeedSpamSample) | **Get** /v3.4/download/threat-feed/spam/sample | Download a sample of the spam threat feed (CSV)



## DownloadThreatFeedMalware

> *os.File DownloadThreatFeedMalware(ctx).Date(date).Execute()

Download the daily malware threat feed (CSV)

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
	date := "date_example" // string | Feed date (yyyy-MM-dd); defaults to latest available (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesThreatFeedAPI.DownloadThreatFeedMalware(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesThreatFeedAPI.DownloadThreatFeedMalware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadThreatFeedMalware`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesThreatFeedAPI.DownloadThreatFeedMalware`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadThreatFeedMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | Feed date (yyyy-MM-dd); defaults to latest available | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadThreatFeedMalwareSample

> *os.File DownloadThreatFeedMalwareSample(ctx).Execute()

Download a sample of the malware threat feed (CSV)

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
	resp, r, err := apiClient.DatabasesThreatFeedAPI.DownloadThreatFeedMalwareSample(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesThreatFeedAPI.DownloadThreatFeedMalwareSample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadThreatFeedMalwareSample`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesThreatFeedAPI.DownloadThreatFeedMalwareSample`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadThreatFeedMalwareSampleRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadThreatFeedPhishing

> *os.File DownloadThreatFeedPhishing(ctx).Date(date).Execute()

Download the daily phishing threat feed (CSV)

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
	date := "date_example" // string | Feed date (yyyy-MM-dd); defaults to latest available (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesThreatFeedAPI.DownloadThreatFeedPhishing(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesThreatFeedAPI.DownloadThreatFeedPhishing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadThreatFeedPhishing`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesThreatFeedAPI.DownloadThreatFeedPhishing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadThreatFeedPhishingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | Feed date (yyyy-MM-dd); defaults to latest available | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadThreatFeedPhishingSample

> *os.File DownloadThreatFeedPhishingSample(ctx).Execute()

Download a sample of the phishing threat feed (CSV)

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
	resp, r, err := apiClient.DatabasesThreatFeedAPI.DownloadThreatFeedPhishingSample(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesThreatFeedAPI.DownloadThreatFeedPhishingSample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadThreatFeedPhishingSample`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesThreatFeedAPI.DownloadThreatFeedPhishingSample`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadThreatFeedPhishingSampleRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadThreatFeedSpam

> *os.File DownloadThreatFeedSpam(ctx).Date(date).Execute()

Download the daily spam threat feed (CSV)

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
	date := "date_example" // string | Feed date (yyyy-MM-dd); defaults to latest available (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesThreatFeedAPI.DownloadThreatFeedSpam(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesThreatFeedAPI.DownloadThreatFeedSpam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadThreatFeedSpam`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesThreatFeedAPI.DownloadThreatFeedSpam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadThreatFeedSpamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | Feed date (yyyy-MM-dd); defaults to latest available | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadThreatFeedSpamSample

> *os.File DownloadThreatFeedSpamSample(ctx).Execute()

Download a sample of the spam threat feed (CSV)

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
	resp, r, err := apiClient.DatabasesThreatFeedAPI.DownloadThreatFeedSpamSample(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesThreatFeedAPI.DownloadThreatFeedSpamSample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadThreatFeedSpamSample`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DatabasesThreatFeedAPI.DownloadThreatFeedSpamSample`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadThreatFeedSpamSampleRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

