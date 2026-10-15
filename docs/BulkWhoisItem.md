# BulkWhoisItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
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

### NewBulkWhoisItem

`func NewBulkWhoisItem() *BulkWhoisItem`

NewBulkWhoisItem instantiates a new BulkWhoisItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWhoisItemWithDefaults

`func NewBulkWhoisItemWithDefaults() *BulkWhoisItem`

NewBulkWhoisItemWithDefaults instantiates a new BulkWhoisItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BulkWhoisItem) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWhoisItem) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWhoisItem) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkWhoisItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDomainName

`func (o *BulkWhoisItem) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *BulkWhoisItem) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *BulkWhoisItem) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *BulkWhoisItem) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetError

`func (o *BulkWhoisItem) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkWhoisItem) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkWhoisItem) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkWhoisItem) HasError() bool`

HasError returns a boolean if a field has been set.

### GetQueryTime

`func (o *BulkWhoisItem) GetQueryTime() string`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *BulkWhoisItem) GetQueryTimeOk() (*string, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *BulkWhoisItem) SetQueryTime(v string)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *BulkWhoisItem) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetWhoisServer

`func (o *BulkWhoisItem) GetWhoisServer() string`

GetWhoisServer returns the WhoisServer field if non-nil, zero value otherwise.

### GetWhoisServerOk

`func (o *BulkWhoisItem) GetWhoisServerOk() (*string, bool)`

GetWhoisServerOk returns a tuple with the WhoisServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisServer

`func (o *BulkWhoisItem) SetWhoisServer(v string)`

SetWhoisServer sets WhoisServer field to given value.

### HasWhoisServer

`func (o *BulkWhoisItem) HasWhoisServer() bool`

HasWhoisServer returns a boolean if a field has been set.

### GetDomainRegistered

`func (o *BulkWhoisItem) GetDomainRegistered() string`

GetDomainRegistered returns the DomainRegistered field if non-nil, zero value otherwise.

### GetDomainRegisteredOk

`func (o *BulkWhoisItem) GetDomainRegisteredOk() (*string, bool)`

GetDomainRegisteredOk returns a tuple with the DomainRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegistered

`func (o *BulkWhoisItem) SetDomainRegistered(v string)`

SetDomainRegistered sets DomainRegistered field to given value.

### HasDomainRegistered

`func (o *BulkWhoisItem) HasDomainRegistered() bool`

HasDomainRegistered returns a boolean if a field has been set.

### GetSecureDns

`func (o *BulkWhoisItem) GetSecureDns() bool`

GetSecureDns returns the SecureDns field if non-nil, zero value otherwise.

### GetSecureDnsOk

`func (o *BulkWhoisItem) GetSecureDnsOk() (*bool, bool)`

GetSecureDnsOk returns a tuple with the SecureDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureDns

`func (o *BulkWhoisItem) SetSecureDns(v bool)`

SetSecureDns sets SecureDns field to given value.

### HasSecureDns

`func (o *BulkWhoisItem) HasSecureDns() bool`

HasSecureDns returns a boolean if a field has been set.

### GetDomainHandle

`func (o *BulkWhoisItem) GetDomainHandle() string`

GetDomainHandle returns the DomainHandle field if non-nil, zero value otherwise.

### GetDomainHandleOk

`func (o *BulkWhoisItem) GetDomainHandleOk() (*string, bool)`

GetDomainHandleOk returns a tuple with the DomainHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainHandle

`func (o *BulkWhoisItem) SetDomainHandle(v string)`

SetDomainHandle sets DomainHandle field to given value.

### HasDomainHandle

`func (o *BulkWhoisItem) HasDomainHandle() bool`

HasDomainHandle returns a boolean if a field has been set.

### GetCreateDate

`func (o *BulkWhoisItem) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *BulkWhoisItem) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *BulkWhoisItem) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *BulkWhoisItem) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUpdateDate

`func (o *BulkWhoisItem) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *BulkWhoisItem) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *BulkWhoisItem) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *BulkWhoisItem) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *BulkWhoisItem) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *BulkWhoisItem) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *BulkWhoisItem) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *BulkWhoisItem) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDomainRegistrar

`func (o *BulkWhoisItem) GetDomainRegistrar() RegistrarInformation`

GetDomainRegistrar returns the DomainRegistrar field if non-nil, zero value otherwise.

### GetDomainRegistrarOk

`func (o *BulkWhoisItem) GetDomainRegistrarOk() (*RegistrarInformation, bool)`

GetDomainRegistrarOk returns a tuple with the DomainRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegistrar

`func (o *BulkWhoisItem) SetDomainRegistrar(v RegistrarInformation)`

SetDomainRegistrar sets DomainRegistrar field to given value.

### HasDomainRegistrar

`func (o *BulkWhoisItem) HasDomainRegistrar() bool`

HasDomainRegistrar returns a boolean if a field has been set.

### GetResellerContact

`func (o *BulkWhoisItem) GetResellerContact() ResellerContact`

GetResellerContact returns the ResellerContact field if non-nil, zero value otherwise.

### GetResellerContactOk

`func (o *BulkWhoisItem) GetResellerContactOk() (*ResellerContact, bool)`

GetResellerContactOk returns a tuple with the ResellerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerContact

`func (o *BulkWhoisItem) SetResellerContact(v ResellerContact)`

SetResellerContact sets ResellerContact field to given value.

### HasResellerContact

`func (o *BulkWhoisItem) HasResellerContact() bool`

HasResellerContact returns a boolean if a field has been set.

### GetRegistrantContact

`func (o *BulkWhoisItem) GetRegistrantContact() PersonalInformation`

GetRegistrantContact returns the RegistrantContact field if non-nil, zero value otherwise.

### GetRegistrantContactOk

`func (o *BulkWhoisItem) GetRegistrantContactOk() (*PersonalInformation, bool)`

GetRegistrantContactOk returns a tuple with the RegistrantContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantContact

`func (o *BulkWhoisItem) SetRegistrantContact(v PersonalInformation)`

SetRegistrantContact sets RegistrantContact field to given value.

### HasRegistrantContact

`func (o *BulkWhoisItem) HasRegistrantContact() bool`

HasRegistrantContact returns a boolean if a field has been set.

### GetAdministrativeContact

`func (o *BulkWhoisItem) GetAdministrativeContact() PersonalInformation`

GetAdministrativeContact returns the AdministrativeContact field if non-nil, zero value otherwise.

### GetAdministrativeContactOk

`func (o *BulkWhoisItem) GetAdministrativeContactOk() (*PersonalInformation, bool)`

GetAdministrativeContactOk returns a tuple with the AdministrativeContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeContact

`func (o *BulkWhoisItem) SetAdministrativeContact(v PersonalInformation)`

SetAdministrativeContact sets AdministrativeContact field to given value.

### HasAdministrativeContact

`func (o *BulkWhoisItem) HasAdministrativeContact() bool`

HasAdministrativeContact returns a boolean if a field has been set.

### GetTechnicalContact

`func (o *BulkWhoisItem) GetTechnicalContact() PersonalInformation`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *BulkWhoisItem) GetTechnicalContactOk() (*PersonalInformation, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *BulkWhoisItem) SetTechnicalContact(v PersonalInformation)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *BulkWhoisItem) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.

### GetBillingContact

`func (o *BulkWhoisItem) GetBillingContact() PersonalInformation`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *BulkWhoisItem) GetBillingContactOk() (*PersonalInformation, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *BulkWhoisItem) SetBillingContact(v PersonalInformation)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *BulkWhoisItem) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetEligibilityInfo

`func (o *BulkWhoisItem) GetEligibilityInfo() EligibilityInfo`

GetEligibilityInfo returns the EligibilityInfo field if non-nil, zero value otherwise.

### GetEligibilityInfoOk

`func (o *BulkWhoisItem) GetEligibilityInfoOk() (*EligibilityInfo, bool)`

GetEligibilityInfoOk returns a tuple with the EligibilityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibilityInfo

`func (o *BulkWhoisItem) SetEligibilityInfo(v EligibilityInfo)`

SetEligibilityInfo sets EligibilityInfo field to given value.

### HasEligibilityInfo

`func (o *BulkWhoisItem) HasEligibilityInfo() bool`

HasEligibilityInfo returns a boolean if a field has been set.

### GetAbuseContact

`func (o *BulkWhoisItem) GetAbuseContact() RegistrarInformation`

GetAbuseContact returns the AbuseContact field if non-nil, zero value otherwise.

### GetAbuseContactOk

`func (o *BulkWhoisItem) GetAbuseContactOk() (*RegistrarInformation, bool)`

GetAbuseContactOk returns a tuple with the AbuseContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseContact

`func (o *BulkWhoisItem) SetAbuseContact(v RegistrarInformation)`

SetAbuseContact sets AbuseContact field to given value.

### HasAbuseContact

`func (o *BulkWhoisItem) HasAbuseContact() bool`

HasAbuseContact returns a boolean if a field has been set.

### GetNameServers

`func (o *BulkWhoisItem) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *BulkWhoisItem) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *BulkWhoisItem) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *BulkWhoisItem) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetDomainStatus

`func (o *BulkWhoisItem) GetDomainStatus() []string`

GetDomainStatus returns the DomainStatus field if non-nil, zero value otherwise.

### GetDomainStatusOk

`func (o *BulkWhoisItem) GetDomainStatusOk() (*[]string, bool)`

GetDomainStatusOk returns a tuple with the DomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStatus

`func (o *BulkWhoisItem) SetDomainStatus(v []string)`

SetDomainStatus sets DomainStatus field to given value.

### HasDomainStatus

`func (o *BulkWhoisItem) HasDomainStatus() bool`

HasDomainStatus returns a boolean if a field has been set.

### GetWhoisRawDomain

`func (o *BulkWhoisItem) GetWhoisRawDomain() string`

GetWhoisRawDomain returns the WhoisRawDomain field if non-nil, zero value otherwise.

### GetWhoisRawDomainOk

`func (o *BulkWhoisItem) GetWhoisRawDomainOk() (*string, bool)`

GetWhoisRawDomainOk returns a tuple with the WhoisRawDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisRawDomain

`func (o *BulkWhoisItem) SetWhoisRawDomain(v string)`

SetWhoisRawDomain sets WhoisRawDomain field to given value.

### HasWhoisRawDomain

`func (o *BulkWhoisItem) HasWhoisRawDomain() bool`

HasWhoisRawDomain returns a boolean if a field has been set.

### GetRegistryData

`func (o *BulkWhoisItem) GetRegistryData() RegistryData`

GetRegistryData returns the RegistryData field if non-nil, zero value otherwise.

### GetRegistryDataOk

`func (o *BulkWhoisItem) GetRegistryDataOk() (*RegistryData, bool)`

GetRegistryDataOk returns a tuple with the RegistryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryData

`func (o *BulkWhoisItem) SetRegistryData(v RegistryData)`

SetRegistryData sets RegistryData field to given value.

### HasRegistryData

`func (o *BulkWhoisItem) HasRegistryData() bool`

HasRegistryData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


