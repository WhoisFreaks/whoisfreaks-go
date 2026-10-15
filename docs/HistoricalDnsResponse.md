# HistoricalDnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**HistoricalDnsRecords** | Pointer to [**[]DnsResponse**](DnsResponse.md) |  | [optional] 

## Methods

### NewHistoricalDnsResponse

`func NewHistoricalDnsResponse() *HistoricalDnsResponse`

NewHistoricalDnsResponse instantiates a new HistoricalDnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalDnsResponseWithDefaults

`func NewHistoricalDnsResponseWithDefaults() *HistoricalDnsResponse`

NewHistoricalDnsResponseWithDefaults instantiates a new HistoricalDnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPages

`func (o *HistoricalDnsResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *HistoricalDnsResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *HistoricalDnsResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *HistoricalDnsResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetCurrentPage

`func (o *HistoricalDnsResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *HistoricalDnsResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *HistoricalDnsResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *HistoricalDnsResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetTotalRecords

`func (o *HistoricalDnsResponse) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *HistoricalDnsResponse) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *HistoricalDnsResponse) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *HistoricalDnsResponse) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetHistoricalDnsRecords

`func (o *HistoricalDnsResponse) GetHistoricalDnsRecords() []DnsResponse`

GetHistoricalDnsRecords returns the HistoricalDnsRecords field if non-nil, zero value otherwise.

### GetHistoricalDnsRecordsOk

`func (o *HistoricalDnsResponse) GetHistoricalDnsRecordsOk() (*[]DnsResponse, bool)`

GetHistoricalDnsRecordsOk returns a tuple with the HistoricalDnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalDnsRecords

`func (o *HistoricalDnsResponse) SetHistoricalDnsRecords(v []DnsResponse)`

SetHistoricalDnsRecords sets HistoricalDnsRecords field to given value.

### HasHistoricalDnsRecords

`func (o *HistoricalDnsResponse) HasHistoricalDnsRecords() bool`

HasHistoricalDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


