# RegistrarInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IdType** | Pointer to **string** |  | [optional] 
**IanaId** | Pointer to **string** |  | [optional] 
**RegistryId** | Pointer to **string** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RegistrarName** | Pointer to **string** |  | [optional] 
**NormalizedName** | Pointer to **string** |  | [optional] 
**WhoisServer** | Pointer to **string** |  | [optional] 
**RdapServer** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**AuthoritativeRegistryName** | Pointer to **string** |  | [optional] 
**OrganizationNumber** | Pointer to **string** |  | [optional] 
**IsSponsor** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegistrarInformation

`func NewRegistrarInformation() *RegistrarInformation`

NewRegistrarInformation instantiates a new RegistrarInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrarInformationWithDefaults

`func NewRegistrarInformationWithDefaults() *RegistrarInformation`

NewRegistrarInformationWithDefaults instantiates a new RegistrarInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistrarInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistrarInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistrarInformation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistrarInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdType

`func (o *RegistrarInformation) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *RegistrarInformation) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *RegistrarInformation) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *RegistrarInformation) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetIanaId

`func (o *RegistrarInformation) GetIanaId() string`

GetIanaId returns the IanaId field if non-nil, zero value otherwise.

### GetIanaIdOk

`func (o *RegistrarInformation) GetIanaIdOk() (*string, bool)`

GetIanaIdOk returns a tuple with the IanaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIanaId

`func (o *RegistrarInformation) SetIanaId(v string)`

SetIanaId sets IanaId field to given value.

### HasIanaId

`func (o *RegistrarInformation) HasIanaId() bool`

HasIanaId returns a boolean if a field has been set.

### GetRegistryId

`func (o *RegistrarInformation) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *RegistrarInformation) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *RegistrarInformation) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *RegistrarInformation) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetHandle

`func (o *RegistrarInformation) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *RegistrarInformation) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *RegistrarInformation) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *RegistrarInformation) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetStatus

`func (o *RegistrarInformation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistrarInformation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistrarInformation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegistrarInformation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegistrarName

`func (o *RegistrarInformation) GetRegistrarName() string`

GetRegistrarName returns the RegistrarName field if non-nil, zero value otherwise.

### GetRegistrarNameOk

`func (o *RegistrarInformation) GetRegistrarNameOk() (*string, bool)`

GetRegistrarNameOk returns a tuple with the RegistrarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrarName

`func (o *RegistrarInformation) SetRegistrarName(v string)`

SetRegistrarName sets RegistrarName field to given value.

### HasRegistrarName

`func (o *RegistrarInformation) HasRegistrarName() bool`

HasRegistrarName returns a boolean if a field has been set.

### GetNormalizedName

`func (o *RegistrarInformation) GetNormalizedName() string`

GetNormalizedName returns the NormalizedName field if non-nil, zero value otherwise.

### GetNormalizedNameOk

`func (o *RegistrarInformation) GetNormalizedNameOk() (*string, bool)`

GetNormalizedNameOk returns a tuple with the NormalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedName

`func (o *RegistrarInformation) SetNormalizedName(v string)`

SetNormalizedName sets NormalizedName field to given value.

### HasNormalizedName

`func (o *RegistrarInformation) HasNormalizedName() bool`

HasNormalizedName returns a boolean if a field has been set.

### GetWhoisServer

`func (o *RegistrarInformation) GetWhoisServer() string`

GetWhoisServer returns the WhoisServer field if non-nil, zero value otherwise.

### GetWhoisServerOk

`func (o *RegistrarInformation) GetWhoisServerOk() (*string, bool)`

GetWhoisServerOk returns a tuple with the WhoisServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisServer

`func (o *RegistrarInformation) SetWhoisServer(v string)`

SetWhoisServer sets WhoisServer field to given value.

### HasWhoisServer

`func (o *RegistrarInformation) HasWhoisServer() bool`

HasWhoisServer returns a boolean if a field has been set.

### GetRdapServer

`func (o *RegistrarInformation) GetRdapServer() string`

GetRdapServer returns the RdapServer field if non-nil, zero value otherwise.

### GetRdapServerOk

`func (o *RegistrarInformation) GetRdapServerOk() (*string, bool)`

GetRdapServerOk returns a tuple with the RdapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdapServer

`func (o *RegistrarInformation) SetRdapServer(v string)`

SetRdapServer sets RdapServer field to given value.

### HasRdapServer

`func (o *RegistrarInformation) HasRdapServer() bool`

HasRdapServer returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *RegistrarInformation) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *RegistrarInformation) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *RegistrarInformation) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *RegistrarInformation) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetEmailAddress

`func (o *RegistrarInformation) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *RegistrarInformation) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *RegistrarInformation) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *RegistrarInformation) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *RegistrarInformation) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RegistrarInformation) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RegistrarInformation) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RegistrarInformation) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAuthoritativeRegistryName

`func (o *RegistrarInformation) GetAuthoritativeRegistryName() string`

GetAuthoritativeRegistryName returns the AuthoritativeRegistryName field if non-nil, zero value otherwise.

### GetAuthoritativeRegistryNameOk

`func (o *RegistrarInformation) GetAuthoritativeRegistryNameOk() (*string, bool)`

GetAuthoritativeRegistryNameOk returns a tuple with the AuthoritativeRegistryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritativeRegistryName

`func (o *RegistrarInformation) SetAuthoritativeRegistryName(v string)`

SetAuthoritativeRegistryName sets AuthoritativeRegistryName field to given value.

### HasAuthoritativeRegistryName

`func (o *RegistrarInformation) HasAuthoritativeRegistryName() bool`

HasAuthoritativeRegistryName returns a boolean if a field has been set.

### GetOrganizationNumber

`func (o *RegistrarInformation) GetOrganizationNumber() string`

GetOrganizationNumber returns the OrganizationNumber field if non-nil, zero value otherwise.

### GetOrganizationNumberOk

`func (o *RegistrarInformation) GetOrganizationNumberOk() (*string, bool)`

GetOrganizationNumberOk returns a tuple with the OrganizationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationNumber

`func (o *RegistrarInformation) SetOrganizationNumber(v string)`

SetOrganizationNumber sets OrganizationNumber field to given value.

### HasOrganizationNumber

`func (o *RegistrarInformation) HasOrganizationNumber() bool`

HasOrganizationNumber returns a boolean if a field has been set.

### GetIsSponsor

`func (o *RegistrarInformation) GetIsSponsor() bool`

GetIsSponsor returns the IsSponsor field if non-nil, zero value otherwise.

### GetIsSponsorOk

`func (o *RegistrarInformation) GetIsSponsorOk() (*bool, bool)`

GetIsSponsorOk returns a tuple with the IsSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSponsor

`func (o *RegistrarInformation) SetIsSponsor(v bool)`

SetIsSponsor sets IsSponsor field to given value.

### HasIsSponsor

`func (o *RegistrarInformation) HasIsSponsor() bool`

HasIsSponsor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


