# HistoricalDnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **int32** |  | [optional] 
**CurrenPage** | Pointer to **int32** |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**HistoricalDns** | Pointer to [**[]DnsResponse**](DnsResponse.md) |  | [optional] 

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

### GetCurrenPage

`func (o *HistoricalDnsResponse) GetCurrenPage() int32`

GetCurrenPage returns the CurrenPage field if non-nil, zero value otherwise.

### GetCurrenPageOk

`func (o *HistoricalDnsResponse) GetCurrenPageOk() (*int32, bool)`

GetCurrenPageOk returns a tuple with the CurrenPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrenPage

`func (o *HistoricalDnsResponse) SetCurrenPage(v int32)`

SetCurrenPage sets CurrenPage field to given value.

### HasCurrenPage

`func (o *HistoricalDnsResponse) HasCurrenPage() bool`

HasCurrenPage returns a boolean if a field has been set.

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

### GetHistoricalDns

`func (o *HistoricalDnsResponse) GetHistoricalDns() []DnsResponse`

GetHistoricalDns returns the HistoricalDns field if non-nil, zero value otherwise.

### GetHistoricalDnsOk

`func (o *HistoricalDnsResponse) GetHistoricalDnsOk() (*[]DnsResponse, bool)`

GetHistoricalDnsOk returns a tuple with the HistoricalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalDns

`func (o *HistoricalDnsResponse) SetHistoricalDns(v []DnsResponse)`

SetHistoricalDns sets HistoricalDns field to given value.

### HasHistoricalDns

`func (o *HistoricalDnsResponse) HasHistoricalDns() bool`

HasHistoricalDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


