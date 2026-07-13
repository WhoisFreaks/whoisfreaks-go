# \TyposquattingAPI

All URIs are relative to *https://api.whoisfreaks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Typosquatting**](TyposquattingAPI.md#Typosquatting) | **Get** /v3.0/domain/typos | Typosquatting Lookup



## Typosquatting

> TyposquattingResponse Typosquatting(ctx).ApiKey(apiKey).Keyword(keyword).Pattern(pattern).PageToken(pageToken).Execute()

Typosquatting Lookup



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
	keyword := "keyword_example" // string |  (optional)
	pattern := "pattern_example" // string |  (optional)
	pageToken := "pageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TyposquattingAPI.Typosquatting(context.Background()).ApiKey(apiKey).Keyword(keyword).Pattern(pattern).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TyposquattingAPI.Typosquatting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Typosquatting`: TyposquattingResponse
	fmt.Fprintf(os.Stdout, "Response from `TyposquattingAPI.Typosquatting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTyposquattingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | Your WHOISFreaks API key | 
 **keyword** | **string** |  | 
 **pattern** | **string** |  | 
 **pageToken** | **string** |  | 

### Return type

[**TyposquattingResponse**](TyposquattingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

