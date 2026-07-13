# WhoisHistoricalItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Num** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**QueryTime** | Pointer to **string** |  | [optional] 
**CreateDate** | Pointer to **string** |  | [optional] 
**UpdateDate** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **string** |  | [optional] 
**DomainRegistrar** | Pointer to [**RegistrarInformation**](RegistrarInformation.md) |  | [optional] 
**ResellerContact** | Pointer to [**ResellerContact**](ResellerContact.md) |  | [optional] 
**RegistrantContact** | Pointer to [**PersonalInformation**](PersonalInformation.md) |  | [optional] 
**AdministrativeContact** | Pointer to [**PersonalInformation**](PersonalInformation.md) |  | [optional] 
**TechnicalContact** | Pointer to [**PersonalInformation**](PersonalInformation.md) |  | [optional] 
**BillingContact** | Pointer to [**PersonalInformation**](PersonalInformation.md) |  | [optional] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**DomainStatus** | Pointer to **[]string** |  | [optional] 
**WhoisRawDomain** | Pointer to **string** |  | [optional] 

## Methods

### NewWhoisHistoricalItem

`func NewWhoisHistoricalItem() *WhoisHistoricalItem`

NewWhoisHistoricalItem instantiates a new WhoisHistoricalItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoisHistoricalItemWithDefaults

`func NewWhoisHistoricalItemWithDefaults() *WhoisHistoricalItem`

NewWhoisHistoricalItemWithDefaults instantiates a new WhoisHistoricalItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNum

`func (o *WhoisHistoricalItem) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *WhoisHistoricalItem) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *WhoisHistoricalItem) SetNum(v int32)`

SetNum sets Num field to given value.

### HasNum

`func (o *WhoisHistoricalItem) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetStatus

`func (o *WhoisHistoricalItem) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhoisHistoricalItem) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhoisHistoricalItem) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhoisHistoricalItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDomainName

`func (o *WhoisHistoricalItem) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *WhoisHistoricalItem) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *WhoisHistoricalItem) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *WhoisHistoricalItem) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetQueryTime

`func (o *WhoisHistoricalItem) GetQueryTime() string`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *WhoisHistoricalItem) GetQueryTimeOk() (*string, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *WhoisHistoricalItem) SetQueryTime(v string)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *WhoisHistoricalItem) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetCreateDate

`func (o *WhoisHistoricalItem) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *WhoisHistoricalItem) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *WhoisHistoricalItem) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *WhoisHistoricalItem) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUpdateDate

`func (o *WhoisHistoricalItem) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *WhoisHistoricalItem) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *WhoisHistoricalItem) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *WhoisHistoricalItem) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *WhoisHistoricalItem) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *WhoisHistoricalItem) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *WhoisHistoricalItem) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *WhoisHistoricalItem) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDomainRegistrar

`func (o *WhoisHistoricalItem) GetDomainRegistrar() RegistrarInformation`

GetDomainRegistrar returns the DomainRegistrar field if non-nil, zero value otherwise.

### GetDomainRegistrarOk

`func (o *WhoisHistoricalItem) GetDomainRegistrarOk() (*RegistrarInformation, bool)`

GetDomainRegistrarOk returns a tuple with the DomainRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegistrar

`func (o *WhoisHistoricalItem) SetDomainRegistrar(v RegistrarInformation)`

SetDomainRegistrar sets DomainRegistrar field to given value.

### HasDomainRegistrar

`func (o *WhoisHistoricalItem) HasDomainRegistrar() bool`

HasDomainRegistrar returns a boolean if a field has been set.

### GetResellerContact

`func (o *WhoisHistoricalItem) GetResellerContact() ResellerContact`

GetResellerContact returns the ResellerContact field if non-nil, zero value otherwise.

### GetResellerContactOk

`func (o *WhoisHistoricalItem) GetResellerContactOk() (*ResellerContact, bool)`

GetResellerContactOk returns a tuple with the ResellerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerContact

`func (o *WhoisHistoricalItem) SetResellerContact(v ResellerContact)`

SetResellerContact sets ResellerContact field to given value.

### HasResellerContact

`func (o *WhoisHistoricalItem) HasResellerContact() bool`

HasResellerContact returns a boolean if a field has been set.

### GetRegistrantContact

`func (o *WhoisHistoricalItem) GetRegistrantContact() PersonalInformation`

GetRegistrantContact returns the RegistrantContact field if non-nil, zero value otherwise.

### GetRegistrantContactOk

`func (o *WhoisHistoricalItem) GetRegistrantContactOk() (*PersonalInformation, bool)`

GetRegistrantContactOk returns a tuple with the RegistrantContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantContact

`func (o *WhoisHistoricalItem) SetRegistrantContact(v PersonalInformation)`

SetRegistrantContact sets RegistrantContact field to given value.

### HasRegistrantContact

`func (o *WhoisHistoricalItem) HasRegistrantContact() bool`

HasRegistrantContact returns a boolean if a field has been set.

### GetAdministrativeContact

`func (o *WhoisHistoricalItem) GetAdministrativeContact() PersonalInformation`

GetAdministrativeContact returns the AdministrativeContact field if non-nil, zero value otherwise.

### GetAdministrativeContactOk

`func (o *WhoisHistoricalItem) GetAdministrativeContactOk() (*PersonalInformation, bool)`

GetAdministrativeContactOk returns a tuple with the AdministrativeContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeContact

`func (o *WhoisHistoricalItem) SetAdministrativeContact(v PersonalInformation)`

SetAdministrativeContact sets AdministrativeContact field to given value.

### HasAdministrativeContact

`func (o *WhoisHistoricalItem) HasAdministrativeContact() bool`

HasAdministrativeContact returns a boolean if a field has been set.

### GetTechnicalContact

`func (o *WhoisHistoricalItem) GetTechnicalContact() PersonalInformation`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *WhoisHistoricalItem) GetTechnicalContactOk() (*PersonalInformation, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *WhoisHistoricalItem) SetTechnicalContact(v PersonalInformation)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *WhoisHistoricalItem) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.

### GetBillingContact

`func (o *WhoisHistoricalItem) GetBillingContact() PersonalInformation`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *WhoisHistoricalItem) GetBillingContactOk() (*PersonalInformation, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *WhoisHistoricalItem) SetBillingContact(v PersonalInformation)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *WhoisHistoricalItem) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetNameServers

`func (o *WhoisHistoricalItem) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *WhoisHistoricalItem) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *WhoisHistoricalItem) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *WhoisHistoricalItem) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetDomainStatus

`func (o *WhoisHistoricalItem) GetDomainStatus() []string`

GetDomainStatus returns the DomainStatus field if non-nil, zero value otherwise.

### GetDomainStatusOk

`func (o *WhoisHistoricalItem) GetDomainStatusOk() (*[]string, bool)`

GetDomainStatusOk returns a tuple with the DomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStatus

`func (o *WhoisHistoricalItem) SetDomainStatus(v []string)`

SetDomainStatus sets DomainStatus field to given value.

### HasDomainStatus

`func (o *WhoisHistoricalItem) HasDomainStatus() bool`

HasDomainStatus returns a boolean if a field has been set.

### GetWhoisRawDomain

`func (o *WhoisHistoricalItem) GetWhoisRawDomain() string`

GetWhoisRawDomain returns the WhoisRawDomain field if non-nil, zero value otherwise.

### GetWhoisRawDomainOk

`func (o *WhoisHistoricalItem) GetWhoisRawDomainOk() (*string, bool)`

GetWhoisRawDomainOk returns a tuple with the WhoisRawDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisRawDomain

`func (o *WhoisHistoricalItem) SetWhoisRawDomain(v string)`

SetWhoisRawDomain sets WhoisRawDomain field to given value.

### HasWhoisRawDomain

`func (o *WhoisHistoricalItem) HasWhoisRawDomain() bool`

HasWhoisRawDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


