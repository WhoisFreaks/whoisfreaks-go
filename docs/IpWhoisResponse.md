# IpWhoisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**AsNumber** | Pointer to **string** |  | [optional] 
**QueryTime** | Pointer to **string** |  | [optional] 
**WhoisServer** | Pointer to **string** |  | [optional] 
**WhoisRawResponse** | Pointer to **string** |  | [optional] 
**RWhoisRawResponse** | Pointer to **string** |  | [optional] 
**InetNums** | Pointer to [**[]InetNum**](InetNum.md) |  | [optional] 
**Organization** | Pointer to [**WhoisOrganization**](WhoisOrganization.md) |  | [optional] 
**Irt** | Pointer to [**Irt**](Irt.md) |  | [optional] 
**AdministrativeContacts** | Pointer to [**[]WhoisPerson**](WhoisPerson.md) |  | [optional] 
**TechnicalContacts** | Pointer to [**[]WhoisPerson**](WhoisPerson.md) |  | [optional] 
**AbuseContacts** | Pointer to [**[]WhoisPerson**](WhoisPerson.md) |  | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 

## Methods

### NewIpWhoisResponse

`func NewIpWhoisResponse() *IpWhoisResponse`

NewIpWhoisResponse instantiates a new IpWhoisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpWhoisResponseWithDefaults

`func NewIpWhoisResponseWithDefaults() *IpWhoisResponse`

NewIpWhoisResponseWithDefaults instantiates a new IpWhoisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IpWhoisResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpWhoisResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpWhoisResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IpWhoisResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIpAddress

`func (o *IpWhoisResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IpWhoisResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IpWhoisResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *IpWhoisResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetAsNumber

`func (o *IpWhoisResponse) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *IpWhoisResponse) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *IpWhoisResponse) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *IpWhoisResponse) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetQueryTime

`func (o *IpWhoisResponse) GetQueryTime() string`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *IpWhoisResponse) GetQueryTimeOk() (*string, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *IpWhoisResponse) SetQueryTime(v string)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *IpWhoisResponse) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetWhoisServer

`func (o *IpWhoisResponse) GetWhoisServer() string`

GetWhoisServer returns the WhoisServer field if non-nil, zero value otherwise.

### GetWhoisServerOk

`func (o *IpWhoisResponse) GetWhoisServerOk() (*string, bool)`

GetWhoisServerOk returns a tuple with the WhoisServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisServer

`func (o *IpWhoisResponse) SetWhoisServer(v string)`

SetWhoisServer sets WhoisServer field to given value.

### HasWhoisServer

`func (o *IpWhoisResponse) HasWhoisServer() bool`

HasWhoisServer returns a boolean if a field has been set.

### GetWhoisRawResponse

`func (o *IpWhoisResponse) GetWhoisRawResponse() string`

GetWhoisRawResponse returns the WhoisRawResponse field if non-nil, zero value otherwise.

### GetWhoisRawResponseOk

`func (o *IpWhoisResponse) GetWhoisRawResponseOk() (*string, bool)`

GetWhoisRawResponseOk returns a tuple with the WhoisRawResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisRawResponse

`func (o *IpWhoisResponse) SetWhoisRawResponse(v string)`

SetWhoisRawResponse sets WhoisRawResponse field to given value.

### HasWhoisRawResponse

`func (o *IpWhoisResponse) HasWhoisRawResponse() bool`

HasWhoisRawResponse returns a boolean if a field has been set.

### GetRWhoisRawResponse

`func (o *IpWhoisResponse) GetRWhoisRawResponse() string`

GetRWhoisRawResponse returns the RWhoisRawResponse field if non-nil, zero value otherwise.

### GetRWhoisRawResponseOk

`func (o *IpWhoisResponse) GetRWhoisRawResponseOk() (*string, bool)`

GetRWhoisRawResponseOk returns a tuple with the RWhoisRawResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRWhoisRawResponse

`func (o *IpWhoisResponse) SetRWhoisRawResponse(v string)`

SetRWhoisRawResponse sets RWhoisRawResponse field to given value.

### HasRWhoisRawResponse

`func (o *IpWhoisResponse) HasRWhoisRawResponse() bool`

HasRWhoisRawResponse returns a boolean if a field has been set.

### GetInetNums

`func (o *IpWhoisResponse) GetInetNums() []InetNum`

GetInetNums returns the InetNums field if non-nil, zero value otherwise.

### GetInetNumsOk

`func (o *IpWhoisResponse) GetInetNumsOk() (*[]InetNum, bool)`

GetInetNumsOk returns a tuple with the InetNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInetNums

`func (o *IpWhoisResponse) SetInetNums(v []InetNum)`

SetInetNums sets InetNums field to given value.

### HasInetNums

`func (o *IpWhoisResponse) HasInetNums() bool`

HasInetNums returns a boolean if a field has been set.

### GetOrganization

`func (o *IpWhoisResponse) GetOrganization() WhoisOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IpWhoisResponse) GetOrganizationOk() (*WhoisOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IpWhoisResponse) SetOrganization(v WhoisOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IpWhoisResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetIrt

`func (o *IpWhoisResponse) GetIrt() Irt`

GetIrt returns the Irt field if non-nil, zero value otherwise.

### GetIrtOk

`func (o *IpWhoisResponse) GetIrtOk() (*Irt, bool)`

GetIrtOk returns a tuple with the Irt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrt

`func (o *IpWhoisResponse) SetIrt(v Irt)`

SetIrt sets Irt field to given value.

### HasIrt

`func (o *IpWhoisResponse) HasIrt() bool`

HasIrt returns a boolean if a field has been set.

### GetAdministrativeContacts

`func (o *IpWhoisResponse) GetAdministrativeContacts() []WhoisPerson`

GetAdministrativeContacts returns the AdministrativeContacts field if non-nil, zero value otherwise.

### GetAdministrativeContactsOk

`func (o *IpWhoisResponse) GetAdministrativeContactsOk() (*[]WhoisPerson, bool)`

GetAdministrativeContactsOk returns a tuple with the AdministrativeContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeContacts

`func (o *IpWhoisResponse) SetAdministrativeContacts(v []WhoisPerson)`

SetAdministrativeContacts sets AdministrativeContacts field to given value.

### HasAdministrativeContacts

`func (o *IpWhoisResponse) HasAdministrativeContacts() bool`

HasAdministrativeContacts returns a boolean if a field has been set.

### GetTechnicalContacts

`func (o *IpWhoisResponse) GetTechnicalContacts() []WhoisPerson`

GetTechnicalContacts returns the TechnicalContacts field if non-nil, zero value otherwise.

### GetTechnicalContactsOk

`func (o *IpWhoisResponse) GetTechnicalContactsOk() (*[]WhoisPerson, bool)`

GetTechnicalContactsOk returns a tuple with the TechnicalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContacts

`func (o *IpWhoisResponse) SetTechnicalContacts(v []WhoisPerson)`

SetTechnicalContacts sets TechnicalContacts field to given value.

### HasTechnicalContacts

`func (o *IpWhoisResponse) HasTechnicalContacts() bool`

HasTechnicalContacts returns a boolean if a field has been set.

### GetAbuseContacts

`func (o *IpWhoisResponse) GetAbuseContacts() []WhoisPerson`

GetAbuseContacts returns the AbuseContacts field if non-nil, zero value otherwise.

### GetAbuseContactsOk

`func (o *IpWhoisResponse) GetAbuseContactsOk() (*[]WhoisPerson, bool)`

GetAbuseContactsOk returns a tuple with the AbuseContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseContacts

`func (o *IpWhoisResponse) SetAbuseContacts(v []WhoisPerson)`

SetAbuseContacts sets AbuseContacts field to given value.

### HasAbuseContacts

`func (o *IpWhoisResponse) HasAbuseContacts() bool`

HasAbuseContacts returns a boolean if a field has been set.

### GetRoutes

`func (o *IpWhoisResponse) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *IpWhoisResponse) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *IpWhoisResponse) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *IpWhoisResponse) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


