# DnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**QueryTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**DomainRegistered** | Pointer to **bool** |  | [optional] 
**DnsTypes** | Pointer to **map[string]interface{}** |  | [optional] 
**DnsRecords** | Pointer to [**[]DnsRecord**](DnsRecord.md) |  | [optional] 

## Methods

### NewDnsResponse

`func NewDnsResponse() *DnsResponse`

NewDnsResponse instantiates a new DnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsResponseWithDefaults

`func NewDnsResponseWithDefaults() *DnsResponse`

NewDnsResponseWithDefaults instantiates a new DnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *DnsResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DnsResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DnsResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DnsResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetIpAddress

`func (o *DnsResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DnsResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DnsResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DnsResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetQueryTime

`func (o *DnsResponse) GetQueryTime() string`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *DnsResponse) GetQueryTimeOk() (*string, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *DnsResponse) SetQueryTime(v string)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *DnsResponse) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetStatus

`func (o *DnsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DnsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DnsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DnsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDomainRegistered

`func (o *DnsResponse) GetDomainRegistered() bool`

GetDomainRegistered returns the DomainRegistered field if non-nil, zero value otherwise.

### GetDomainRegisteredOk

`func (o *DnsResponse) GetDomainRegisteredOk() (*bool, bool)`

GetDomainRegisteredOk returns a tuple with the DomainRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegistered

`func (o *DnsResponse) SetDomainRegistered(v bool)`

SetDomainRegistered sets DomainRegistered field to given value.

### HasDomainRegistered

`func (o *DnsResponse) HasDomainRegistered() bool`

HasDomainRegistered returns a boolean if a field has been set.

### GetDnsTypes

`func (o *DnsResponse) GetDnsTypes() map[string]interface{}`

GetDnsTypes returns the DnsTypes field if non-nil, zero value otherwise.

### GetDnsTypesOk

`func (o *DnsResponse) GetDnsTypesOk() (*map[string]interface{}, bool)`

GetDnsTypesOk returns a tuple with the DnsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTypes

`func (o *DnsResponse) SetDnsTypes(v map[string]interface{})`

SetDnsTypes sets DnsTypes field to given value.

### HasDnsTypes

`func (o *DnsResponse) HasDnsTypes() bool`

HasDnsTypes returns a boolean if a field has been set.

### GetDnsRecords

`func (o *DnsResponse) GetDnsRecords() []DnsRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *DnsResponse) GetDnsRecordsOk() (*[]DnsRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *DnsResponse) SetDnsRecords(v []DnsRecord)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *DnsResponse) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


