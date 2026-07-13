# RegistryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** |  | [optional] 
**QueryTime** | Pointer to **string** |  | [optional] 
**WhoisServer** | Pointer to **string** |  | [optional] 
**DomainRegistered** | Pointer to **string** |  | [optional] 
**CreateDate** | Pointer to **string** |  | [optional] 
**UpdateDate** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **string** |  | [optional] 
**DomainRegistrar** | Pointer to [**RegistrarInformation**](RegistrarInformation.md) |  | [optional] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**DomainStatus** | Pointer to **[]string** |  | [optional] 
**WhoisRawDomain** | Pointer to **string** |  | [optional] 

## Methods

### NewRegistryData

`func NewRegistryData() *RegistryData`

NewRegistryData instantiates a new RegistryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryDataWithDefaults

`func NewRegistryDataWithDefaults() *RegistryData`

NewRegistryDataWithDefaults instantiates a new RegistryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *RegistryData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *RegistryData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *RegistryData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *RegistryData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetQueryTime

`func (o *RegistryData) GetQueryTime() string`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *RegistryData) GetQueryTimeOk() (*string, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *RegistryData) SetQueryTime(v string)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *RegistryData) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetWhoisServer

`func (o *RegistryData) GetWhoisServer() string`

GetWhoisServer returns the WhoisServer field if non-nil, zero value otherwise.

### GetWhoisServerOk

`func (o *RegistryData) GetWhoisServerOk() (*string, bool)`

GetWhoisServerOk returns a tuple with the WhoisServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisServer

`func (o *RegistryData) SetWhoisServer(v string)`

SetWhoisServer sets WhoisServer field to given value.

### HasWhoisServer

`func (o *RegistryData) HasWhoisServer() bool`

HasWhoisServer returns a boolean if a field has been set.

### GetDomainRegistered

`func (o *RegistryData) GetDomainRegistered() string`

GetDomainRegistered returns the DomainRegistered field if non-nil, zero value otherwise.

### GetDomainRegisteredOk

`func (o *RegistryData) GetDomainRegisteredOk() (*string, bool)`

GetDomainRegisteredOk returns a tuple with the DomainRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegistered

`func (o *RegistryData) SetDomainRegistered(v string)`

SetDomainRegistered sets DomainRegistered field to given value.

### HasDomainRegistered

`func (o *RegistryData) HasDomainRegistered() bool`

HasDomainRegistered returns a boolean if a field has been set.

### GetCreateDate

`func (o *RegistryData) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *RegistryData) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *RegistryData) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *RegistryData) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUpdateDate

`func (o *RegistryData) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *RegistryData) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *RegistryData) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *RegistryData) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *RegistryData) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *RegistryData) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *RegistryData) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *RegistryData) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDomainRegistrar

`func (o *RegistryData) GetDomainRegistrar() RegistrarInformation`

GetDomainRegistrar returns the DomainRegistrar field if non-nil, zero value otherwise.

### GetDomainRegistrarOk

`func (o *RegistryData) GetDomainRegistrarOk() (*RegistrarInformation, bool)`

GetDomainRegistrarOk returns a tuple with the DomainRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegistrar

`func (o *RegistryData) SetDomainRegistrar(v RegistrarInformation)`

SetDomainRegistrar sets DomainRegistrar field to given value.

### HasDomainRegistrar

`func (o *RegistryData) HasDomainRegistrar() bool`

HasDomainRegistrar returns a boolean if a field has been set.

### GetNameServers

`func (o *RegistryData) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *RegistryData) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *RegistryData) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *RegistryData) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetDomainStatus

`func (o *RegistryData) GetDomainStatus() []string`

GetDomainStatus returns the DomainStatus field if non-nil, zero value otherwise.

### GetDomainStatusOk

`func (o *RegistryData) GetDomainStatusOk() (*[]string, bool)`

GetDomainStatusOk returns a tuple with the DomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStatus

`func (o *RegistryData) SetDomainStatus(v []string)`

SetDomainStatus sets DomainStatus field to given value.

### HasDomainStatus

`func (o *RegistryData) HasDomainStatus() bool`

HasDomainStatus returns a boolean if a field has been set.

### GetWhoisRawDomain

`func (o *RegistryData) GetWhoisRawDomain() string`

GetWhoisRawDomain returns the WhoisRawDomain field if non-nil, zero value otherwise.

### GetWhoisRawDomainOk

`func (o *RegistryData) GetWhoisRawDomainOk() (*string, bool)`

GetWhoisRawDomainOk returns a tuple with the WhoisRawDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisRawDomain

`func (o *RegistryData) SetWhoisRawDomain(v string)`

SetWhoisRawDomain sets WhoisRawDomain field to given value.

### HasWhoisRawDomain

`func (o *RegistryData) HasWhoisRawDomain() bool`

HasWhoisRawDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


