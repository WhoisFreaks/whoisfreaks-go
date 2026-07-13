# SubdomainsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**CurrentPage** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**QueryTime** | Pointer to **string** |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**Subdomains** | Pointer to [**[]Subdomain**](Subdomain.md) |  | [optional] 

## Methods

### NewSubdomainsResponse

`func NewSubdomainsResponse() *SubdomainsResponse`

NewSubdomainsResponse instantiates a new SubdomainsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubdomainsResponseWithDefaults

`func NewSubdomainsResponseWithDefaults() *SubdomainsResponse`

NewSubdomainsResponseWithDefaults instantiates a new SubdomainsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SubdomainsResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SubdomainsResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SubdomainsResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SubdomainsResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetStatus

`func (o *SubdomainsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubdomainsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubdomainsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubdomainsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrentPage

`func (o *SubdomainsResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *SubdomainsResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *SubdomainsResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *SubdomainsResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *SubdomainsResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *SubdomainsResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *SubdomainsResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *SubdomainsResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetQueryTime

`func (o *SubdomainsResponse) GetQueryTime() string`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *SubdomainsResponse) GetQueryTimeOk() (*string, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *SubdomainsResponse) SetQueryTime(v string)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *SubdomainsResponse) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetTotalRecords

`func (o *SubdomainsResponse) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *SubdomainsResponse) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *SubdomainsResponse) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *SubdomainsResponse) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetSubdomains

`func (o *SubdomainsResponse) GetSubdomains() []Subdomain`

GetSubdomains returns the Subdomains field if non-nil, zero value otherwise.

### GetSubdomainsOk

`func (o *SubdomainsResponse) GetSubdomainsOk() (*[]Subdomain, bool)`

GetSubdomainsOk returns a tuple with the Subdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomains

`func (o *SubdomainsResponse) SetSubdomains(v []Subdomain)`

SetSubdomains sets Subdomains field to given value.

### HasSubdomains

`func (o *SubdomainsResponse) HasSubdomains() bool`

HasSubdomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


