# ReverseWhoisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResult** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 
**WhoisDomainsHistorical** | Pointer to [**[]WhoisHistoricalItem**](WhoisHistoricalItem.md) |  | [optional] 

## Methods

### NewReverseWhoisResponse

`func NewReverseWhoisResponse() *ReverseWhoisResponse`

NewReverseWhoisResponse instantiates a new ReverseWhoisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseWhoisResponseWithDefaults

`func NewReverseWhoisResponseWithDefaults() *ReverseWhoisResponse`

NewReverseWhoisResponseWithDefaults instantiates a new ReverseWhoisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResult

`func (o *ReverseWhoisResponse) GetTotalResult() int32`

GetTotalResult returns the TotalResult field if non-nil, zero value otherwise.

### GetTotalResultOk

`func (o *ReverseWhoisResponse) GetTotalResultOk() (*int32, bool)`

GetTotalResultOk returns a tuple with the TotalResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResult

`func (o *ReverseWhoisResponse) SetTotalResult(v int32)`

SetTotalResult sets TotalResult field to given value.

### HasTotalResult

`func (o *ReverseWhoisResponse) HasTotalResult() bool`

HasTotalResult returns a boolean if a field has been set.

### GetTotalPages

`func (o *ReverseWhoisResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ReverseWhoisResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ReverseWhoisResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ReverseWhoisResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ReverseWhoisResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ReverseWhoisResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ReverseWhoisResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ReverseWhoisResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetWhoisDomainsHistorical

`func (o *ReverseWhoisResponse) GetWhoisDomainsHistorical() []WhoisHistoricalItem`

GetWhoisDomainsHistorical returns the WhoisDomainsHistorical field if non-nil, zero value otherwise.

### GetWhoisDomainsHistoricalOk

`func (o *ReverseWhoisResponse) GetWhoisDomainsHistoricalOk() (*[]WhoisHistoricalItem, bool)`

GetWhoisDomainsHistoricalOk returns a tuple with the WhoisDomainsHistorical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisDomainsHistorical

`func (o *ReverseWhoisResponse) SetWhoisDomainsHistorical(v []WhoisHistoricalItem)`

SetWhoisDomainsHistorical sets WhoisDomainsHistorical field to given value.

### HasWhoisDomainsHistorical

`func (o *ReverseWhoisResponse) HasWhoisDomainsHistorical() bool`

HasWhoisDomainsHistorical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


