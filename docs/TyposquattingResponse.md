# TyposquattingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**HasNextPage** | Pointer to **bool** |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to [**[]TyposquattingDomain**](TyposquattingDomain.md) |  | [optional] 

## Methods

### NewTyposquattingResponse

`func NewTyposquattingResponse() *TyposquattingResponse`

NewTyposquattingResponse instantiates a new TyposquattingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTyposquattingResponseWithDefaults

`func NewTyposquattingResponseWithDefaults() *TyposquattingResponse`

NewTyposquattingResponseWithDefaults instantiates a new TyposquattingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TyposquattingResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TyposquattingResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TyposquattingResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TyposquattingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalRecords

`func (o *TyposquattingResponse) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *TyposquattingResponse) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *TyposquattingResponse) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *TyposquattingResponse) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetCurrentPage

`func (o *TyposquattingResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *TyposquattingResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *TyposquattingResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *TyposquattingResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *TyposquattingResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *TyposquattingResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *TyposquattingResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *TyposquattingResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetHasNextPage

`func (o *TyposquattingResponse) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *TyposquattingResponse) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *TyposquattingResponse) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.

### HasHasNextPage

`func (o *TyposquattingResponse) HasHasNextPage() bool`

HasHasNextPage returns a boolean if a field has been set.

### GetNextPageToken

`func (o *TyposquattingResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *TyposquattingResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *TyposquattingResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *TyposquattingResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetDomains

`func (o *TyposquattingResponse) GetDomains() []TyposquattingDomain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *TyposquattingResponse) GetDomainsOk() (*[]TyposquattingDomain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *TyposquattingResponse) SetDomains(v []TyposquattingDomain)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *TyposquattingResponse) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


