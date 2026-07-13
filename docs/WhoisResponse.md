# WhoisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**QueryTime** | Pointer to **string** |  | [optional] 
**WhoisServer** | Pointer to **string** |  | [optional] 
**DomainRegistered** | Pointer to **string** |  | [optional] 
**SecureDns** | Pointer to **bool** |  | [optional] 
**DomainHandle** | Pointer to **string** |  | [optional] 
**CreateDate** | Pointer to **string** |  | [optional] 
**UpdateDate** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **string** |  | [optional] 
**DomainRegistrar** | Pointer to [**RegistrarInformation**](RegistrarInformation.md) |  | [optional] 
**ResellerContact** | Pointer to [**ResellerContact**](ResellerContact.md) |  | [optional] 
**RegistrantContact** | Pointer to [**PersonalInformation**](PersonalInformation.md) |  | [optional] 
**AdministrativeContact** | Pointer to [**PersonalInformation**](PersonalInformation.md) |  | [optional] 
**TechnicalContact** | Pointer to [**PersonalInformation**](PersonalInformation.md) |  | [optional] 
**BillingContact** | Pointer to [**PersonalInformation**](PersonalInformation.md) |  | [optional] 
**EligibilityInfo** | Pointer to [**EligibilityInfo**](EligibilityInfo.md) |  | [optional] 
**AbuseContact** | Pointer to [**RegistrarInformation**](RegistrarInformation.md) |  | [optional] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**DomainStatus** | Pointer to **[]string** |  | [optional] 
**WhoisRawDomain** | Pointer to **string** |  | [optional] 
**RegistryData** | Pointer to [**RegistryData**](RegistryData.md) |  | [optional] 

## Methods

### NewWhoisResponse

`func NewWhoisResponse() *WhoisResponse`

NewWhoisResponse instantiates a new WhoisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoisResponseWithDefaults

`func NewWhoisResponseWithDefaults() *WhoisResponse`

NewWhoisResponseWithDefaults instantiates a new WhoisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WhoisResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhoisResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhoisResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhoisResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDomainName

`func (o *WhoisResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *WhoisResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *WhoisResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *WhoisResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetQueryTime

`func (o *WhoisResponse) GetQueryTime() string`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *WhoisResponse) GetQueryTimeOk() (*string, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *WhoisResponse) SetQueryTime(v string)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *WhoisResponse) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetWhoisServer

`func (o *WhoisResponse) GetWhoisServer() string`

GetWhoisServer returns the WhoisServer field if non-nil, zero value otherwise.

### GetWhoisServerOk

`func (o *WhoisResponse) GetWhoisServerOk() (*string, bool)`

GetWhoisServerOk returns a tuple with the WhoisServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisServer

`func (o *WhoisResponse) SetWhoisServer(v string)`

SetWhoisServer sets WhoisServer field to given value.

### HasWhoisServer

`func (o *WhoisResponse) HasWhoisServer() bool`

HasWhoisServer returns a boolean if a field has been set.

### GetDomainRegistered

`func (o *WhoisResponse) GetDomainRegistered() string`

GetDomainRegistered returns the DomainRegistered field if non-nil, zero value otherwise.

### GetDomainRegisteredOk

`func (o *WhoisResponse) GetDomainRegisteredOk() (*string, bool)`

GetDomainRegisteredOk returns a tuple with the DomainRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegistered

`func (o *WhoisResponse) SetDomainRegistered(v string)`

SetDomainRegistered sets DomainRegistered field to given value.

### HasDomainRegistered

`func (o *WhoisResponse) HasDomainRegistered() bool`

HasDomainRegistered returns a boolean if a field has been set.

### GetSecureDns

`func (o *WhoisResponse) GetSecureDns() bool`

GetSecureDns returns the SecureDns field if non-nil, zero value otherwise.

### GetSecureDnsOk

`func (o *WhoisResponse) GetSecureDnsOk() (*bool, bool)`

GetSecureDnsOk returns a tuple with the SecureDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureDns

`func (o *WhoisResponse) SetSecureDns(v bool)`

SetSecureDns sets SecureDns field to given value.

### HasSecureDns

`func (o *WhoisResponse) HasSecureDns() bool`

HasSecureDns returns a boolean if a field has been set.

### GetDomainHandle

`func (o *WhoisResponse) GetDomainHandle() string`

GetDomainHandle returns the DomainHandle field if non-nil, zero value otherwise.

### GetDomainHandleOk

`func (o *WhoisResponse) GetDomainHandleOk() (*string, bool)`

GetDomainHandleOk returns a tuple with the DomainHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainHandle

`func (o *WhoisResponse) SetDomainHandle(v string)`

SetDomainHandle sets DomainHandle field to given value.

### HasDomainHandle

`func (o *WhoisResponse) HasDomainHandle() bool`

HasDomainHandle returns a boolean if a field has been set.

### GetCreateDate

`func (o *WhoisResponse) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *WhoisResponse) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *WhoisResponse) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *WhoisResponse) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUpdateDate

`func (o *WhoisResponse) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *WhoisResponse) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *WhoisResponse) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *WhoisResponse) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *WhoisResponse) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *WhoisResponse) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *WhoisResponse) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *WhoisResponse) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDomainRegistrar

`func (o *WhoisResponse) GetDomainRegistrar() RegistrarInformation`

GetDomainRegistrar returns the DomainRegistrar field if non-nil, zero value otherwise.

### GetDomainRegistrarOk

`func (o *WhoisResponse) GetDomainRegistrarOk() (*RegistrarInformation, bool)`

GetDomainRegistrarOk returns a tuple with the DomainRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegistrar

`func (o *WhoisResponse) SetDomainRegistrar(v RegistrarInformation)`

SetDomainRegistrar sets DomainRegistrar field to given value.

### HasDomainRegistrar

`func (o *WhoisResponse) HasDomainRegistrar() bool`

HasDomainRegistrar returns a boolean if a field has been set.

### GetResellerContact

`func (o *WhoisResponse) GetResellerContact() ResellerContact`

GetResellerContact returns the ResellerContact field if non-nil, zero value otherwise.

### GetResellerContactOk

`func (o *WhoisResponse) GetResellerContactOk() (*ResellerContact, bool)`

GetResellerContactOk returns a tuple with the ResellerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerContact

`func (o *WhoisResponse) SetResellerContact(v ResellerContact)`

SetResellerContact sets ResellerContact field to given value.

### HasResellerContact

`func (o *WhoisResponse) HasResellerContact() bool`

HasResellerContact returns a boolean if a field has been set.

### GetRegistrantContact

`func (o *WhoisResponse) GetRegistrantContact() PersonalInformation`

GetRegistrantContact returns the RegistrantContact field if non-nil, zero value otherwise.

### GetRegistrantContactOk

`func (o *WhoisResponse) GetRegistrantContactOk() (*PersonalInformation, bool)`

GetRegistrantContactOk returns a tuple with the RegistrantContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantContact

`func (o *WhoisResponse) SetRegistrantContact(v PersonalInformation)`

SetRegistrantContact sets RegistrantContact field to given value.

### HasRegistrantContact

`func (o *WhoisResponse) HasRegistrantContact() bool`

HasRegistrantContact returns a boolean if a field has been set.

### GetAdministrativeContact

`func (o *WhoisResponse) GetAdministrativeContact() PersonalInformation`

GetAdministrativeContact returns the AdministrativeContact field if non-nil, zero value otherwise.

### GetAdministrativeContactOk

`func (o *WhoisResponse) GetAdministrativeContactOk() (*PersonalInformation, bool)`

GetAdministrativeContactOk returns a tuple with the AdministrativeContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeContact

`func (o *WhoisResponse) SetAdministrativeContact(v PersonalInformation)`

SetAdministrativeContact sets AdministrativeContact field to given value.

### HasAdministrativeContact

`func (o *WhoisResponse) HasAdministrativeContact() bool`

HasAdministrativeContact returns a boolean if a field has been set.

### GetTechnicalContact

`func (o *WhoisResponse) GetTechnicalContact() PersonalInformation`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *WhoisResponse) GetTechnicalContactOk() (*PersonalInformation, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *WhoisResponse) SetTechnicalContact(v PersonalInformation)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *WhoisResponse) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.

### GetBillingContact

`func (o *WhoisResponse) GetBillingContact() PersonalInformation`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *WhoisResponse) GetBillingContactOk() (*PersonalInformation, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *WhoisResponse) SetBillingContact(v PersonalInformation)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *WhoisResponse) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetEligibilityInfo

`func (o *WhoisResponse) GetEligibilityInfo() EligibilityInfo`

GetEligibilityInfo returns the EligibilityInfo field if non-nil, zero value otherwise.

### GetEligibilityInfoOk

`func (o *WhoisResponse) GetEligibilityInfoOk() (*EligibilityInfo, bool)`

GetEligibilityInfoOk returns a tuple with the EligibilityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibilityInfo

`func (o *WhoisResponse) SetEligibilityInfo(v EligibilityInfo)`

SetEligibilityInfo sets EligibilityInfo field to given value.

### HasEligibilityInfo

`func (o *WhoisResponse) HasEligibilityInfo() bool`

HasEligibilityInfo returns a boolean if a field has been set.

### GetAbuseContact

`func (o *WhoisResponse) GetAbuseContact() RegistrarInformation`

GetAbuseContact returns the AbuseContact field if non-nil, zero value otherwise.

### GetAbuseContactOk

`func (o *WhoisResponse) GetAbuseContactOk() (*RegistrarInformation, bool)`

GetAbuseContactOk returns a tuple with the AbuseContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseContact

`func (o *WhoisResponse) SetAbuseContact(v RegistrarInformation)`

SetAbuseContact sets AbuseContact field to given value.

### HasAbuseContact

`func (o *WhoisResponse) HasAbuseContact() bool`

HasAbuseContact returns a boolean if a field has been set.

### GetNameServers

`func (o *WhoisResponse) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *WhoisResponse) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *WhoisResponse) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *WhoisResponse) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetDomainStatus

`func (o *WhoisResponse) GetDomainStatus() []string`

GetDomainStatus returns the DomainStatus field if non-nil, zero value otherwise.

### GetDomainStatusOk

`func (o *WhoisResponse) GetDomainStatusOk() (*[]string, bool)`

GetDomainStatusOk returns a tuple with the DomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStatus

`func (o *WhoisResponse) SetDomainStatus(v []string)`

SetDomainStatus sets DomainStatus field to given value.

### HasDomainStatus

`func (o *WhoisResponse) HasDomainStatus() bool`

HasDomainStatus returns a boolean if a field has been set.

### GetWhoisRawDomain

`func (o *WhoisResponse) GetWhoisRawDomain() string`

GetWhoisRawDomain returns the WhoisRawDomain field if non-nil, zero value otherwise.

### GetWhoisRawDomainOk

`func (o *WhoisResponse) GetWhoisRawDomainOk() (*string, bool)`

GetWhoisRawDomainOk returns a tuple with the WhoisRawDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisRawDomain

`func (o *WhoisResponse) SetWhoisRawDomain(v string)`

SetWhoisRawDomain sets WhoisRawDomain field to given value.

### HasWhoisRawDomain

`func (o *WhoisResponse) HasWhoisRawDomain() bool`

HasWhoisRawDomain returns a boolean if a field has been set.

### GetRegistryData

`func (o *WhoisResponse) GetRegistryData() RegistryData`

GetRegistryData returns the RegistryData field if non-nil, zero value otherwise.

### GetRegistryDataOk

`func (o *WhoisResponse) GetRegistryDataOk() (*RegistryData, bool)`

GetRegistryDataOk returns a tuple with the RegistryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryData

`func (o *WhoisResponse) SetRegistryData(v RegistryData)`

SetRegistryData sets RegistryData field to given value.

### HasRegistryData

`func (o *WhoisResponse) HasRegistryData() bool`

HasRegistryData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


