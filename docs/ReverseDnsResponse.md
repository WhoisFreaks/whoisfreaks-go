# ReverseDnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRecords** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 
**ReverseDnsRecords** | Pointer to [**[]DnsResponse**](DnsResponse.md) |  | [optional] 

## Methods

### NewReverseDnsResponse

`func NewReverseDnsResponse() *ReverseDnsResponse`

NewReverseDnsResponse instantiates a new ReverseDnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseDnsResponseWithDefaults

`func NewReverseDnsResponseWithDefaults() *ReverseDnsResponse`

NewReverseDnsResponseWithDefaults instantiates a new ReverseDnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRecords

`func (o *ReverseDnsResponse) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *ReverseDnsResponse) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *ReverseDnsResponse) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *ReverseDnsResponse) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetTotalPages

`func (o *ReverseDnsResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ReverseDnsResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ReverseDnsResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ReverseDnsResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ReverseDnsResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ReverseDnsResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ReverseDnsResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ReverseDnsResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetReverseDnsRecords

`func (o *ReverseDnsResponse) GetReverseDnsRecords() []DnsResponse`

GetReverseDnsRecords returns the ReverseDnsRecords field if non-nil, zero value otherwise.

### GetReverseDnsRecordsOk

`func (o *ReverseDnsResponse) GetReverseDnsRecordsOk() (*[]DnsResponse, bool)`

GetReverseDnsRecordsOk returns a tuple with the ReverseDnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDnsRecords

`func (o *ReverseDnsResponse) SetReverseDnsRecords(v []DnsResponse)`

SetReverseDnsRecords sets ReverseDnsRecords field to given value.

### HasReverseDnsRecords

`func (o *ReverseDnsResponse) HasReverseDnsRecords() bool`

HasReverseDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


