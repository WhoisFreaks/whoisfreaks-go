# Subdomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subdomain** | Pointer to **string** |  | [optional] 
**FirstSeen** | Pointer to **string** |  | [optional] 
**LastSeen** | Pointer to **string** |  | [optional] 
**InactiveFrom** | Pointer to **string** |  | [optional] 
**DnsRecords** | Pointer to [**DnsResponse**](DnsResponse.md) |  | [optional] 

## Methods

### NewSubdomain

`func NewSubdomain() *Subdomain`

NewSubdomain instantiates a new Subdomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubdomainWithDefaults

`func NewSubdomainWithDefaults() *Subdomain`

NewSubdomainWithDefaults instantiates a new Subdomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubdomain

`func (o *Subdomain) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Subdomain) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Subdomain) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Subdomain) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetFirstSeen

`func (o *Subdomain) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *Subdomain) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *Subdomain) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *Subdomain) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *Subdomain) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Subdomain) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Subdomain) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Subdomain) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetInactiveFrom

`func (o *Subdomain) GetInactiveFrom() string`

GetInactiveFrom returns the InactiveFrom field if non-nil, zero value otherwise.

### GetInactiveFromOk

`func (o *Subdomain) GetInactiveFromOk() (*string, bool)`

GetInactiveFromOk returns a tuple with the InactiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFrom

`func (o *Subdomain) SetInactiveFrom(v string)`

SetInactiveFrom sets InactiveFrom field to given value.

### HasInactiveFrom

`func (o *Subdomain) HasInactiveFrom() bool`

HasInactiveFrom returns a boolean if a field has been set.

### GetDnsRecords

`func (o *Subdomain) GetDnsRecords() DnsResponse`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *Subdomain) GetDnsRecordsOk() (*DnsResponse, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *Subdomain) SetDnsRecords(v DnsResponse)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *Subdomain) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


