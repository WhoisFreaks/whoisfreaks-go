# DatabaseUpdates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainWhois** | Pointer to [**UpdateFrequencies**](UpdateFrequencies.md) |  | [optional] 
**Dns** | Pointer to [**UpdateFrequencies**](UpdateFrequencies.md) |  | [optional] 
**Subdomains** | Pointer to [**UpdateFrequencies**](UpdateFrequencies.md) |  | [optional] 
**IpWhois** | Pointer to [**UpdateFrequencies**](UpdateFrequencies.md) |  | [optional] 

## Methods

### NewDatabaseUpdates

`func NewDatabaseUpdates() *DatabaseUpdates`

NewDatabaseUpdates instantiates a new DatabaseUpdates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseUpdatesWithDefaults

`func NewDatabaseUpdatesWithDefaults() *DatabaseUpdates`

NewDatabaseUpdatesWithDefaults instantiates a new DatabaseUpdates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainWhois

`func (o *DatabaseUpdates) GetDomainWhois() UpdateFrequencies`

GetDomainWhois returns the DomainWhois field if non-nil, zero value otherwise.

### GetDomainWhoisOk

`func (o *DatabaseUpdates) GetDomainWhoisOk() (*UpdateFrequencies, bool)`

GetDomainWhoisOk returns a tuple with the DomainWhois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainWhois

`func (o *DatabaseUpdates) SetDomainWhois(v UpdateFrequencies)`

SetDomainWhois sets DomainWhois field to given value.

### HasDomainWhois

`func (o *DatabaseUpdates) HasDomainWhois() bool`

HasDomainWhois returns a boolean if a field has been set.

### GetDns

`func (o *DatabaseUpdates) GetDns() UpdateFrequencies`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *DatabaseUpdates) GetDnsOk() (*UpdateFrequencies, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *DatabaseUpdates) SetDns(v UpdateFrequencies)`

SetDns sets Dns field to given value.

### HasDns

`func (o *DatabaseUpdates) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetSubdomains

`func (o *DatabaseUpdates) GetSubdomains() UpdateFrequencies`

GetSubdomains returns the Subdomains field if non-nil, zero value otherwise.

### GetSubdomainsOk

`func (o *DatabaseUpdates) GetSubdomainsOk() (*UpdateFrequencies, bool)`

GetSubdomainsOk returns a tuple with the Subdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomains

`func (o *DatabaseUpdates) SetSubdomains(v UpdateFrequencies)`

SetSubdomains sets Subdomains field to given value.

### HasSubdomains

`func (o *DatabaseUpdates) HasSubdomains() bool`

HasSubdomains returns a boolean if a field has been set.

### GetIpWhois

`func (o *DatabaseUpdates) GetIpWhois() UpdateFrequencies`

GetIpWhois returns the IpWhois field if non-nil, zero value otherwise.

### GetIpWhoisOk

`func (o *DatabaseUpdates) GetIpWhoisOk() (*UpdateFrequencies, bool)`

GetIpWhoisOk returns a tuple with the IpWhois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhois

`func (o *DatabaseUpdates) SetIpWhois(v UpdateFrequencies)`

SetIpWhois sets IpWhois field to given value.

### HasIpWhois

`func (o *DatabaseUpdates) HasIpWhois() bool`

HasIpWhois returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


