# WhoisOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **[]string** |  | [optional] 
**Address** | Pointer to **[]string** |  | [optional] 
**Street** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**District** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**ZipCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **[]string** |  | [optional] 
**Latitude** | Pointer to **float32** |  | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**Email** | Pointer to **[]string** |  | [optional] 
**AbuseMailbox** | Pointer to **[]string** |  | [optional] 
**Phone** | Pointer to **[]string** |  | [optional] 
**FaxNo** | Pointer to **[]string** |  | [optional] 
**Organizations** | Pointer to **[]string** |  | [optional] 
**AdminContacts** | Pointer to **[]string** |  | [optional] 
**TechContacts** | Pointer to **[]string** |  | [optional] 
**AbuseContacts** | Pointer to **[]string** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**Remarks** | Pointer to **[]string** |  | [optional] 
**Notify** | Pointer to **[]string** |  | [optional] 
**RefNfy** | Pointer to **[]string** |  | [optional] 
**MntRef** | Pointer to **[]string** |  | [optional] 
**MntBy** | Pointer to **[]string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewWhoisOrganization

`func NewWhoisOrganization() *WhoisOrganization`

NewWhoisOrganization instantiates a new WhoisOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoisOrganizationWithDefaults

`func NewWhoisOrganizationWithDefaults() *WhoisOrganization`

NewWhoisOrganizationWithDefaults instantiates a new WhoisOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *WhoisOrganization) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *WhoisOrganization) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *WhoisOrganization) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *WhoisOrganization) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetName

`func (o *WhoisOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhoisOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhoisOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhoisOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WhoisOrganization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhoisOrganization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhoisOrganization) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhoisOrganization) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *WhoisOrganization) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhoisOrganization) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhoisOrganization) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhoisOrganization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddress

`func (o *WhoisOrganization) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WhoisOrganization) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WhoisOrganization) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WhoisOrganization) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStreet

`func (o *WhoisOrganization) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *WhoisOrganization) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *WhoisOrganization) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *WhoisOrganization) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *WhoisOrganization) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *WhoisOrganization) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *WhoisOrganization) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *WhoisOrganization) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDistrict

`func (o *WhoisOrganization) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *WhoisOrganization) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *WhoisOrganization) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *WhoisOrganization) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetState

`func (o *WhoisOrganization) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WhoisOrganization) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WhoisOrganization) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WhoisOrganization) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *WhoisOrganization) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *WhoisOrganization) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *WhoisOrganization) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *WhoisOrganization) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCountry

`func (o *WhoisOrganization) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WhoisOrganization) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WhoisOrganization) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *WhoisOrganization) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLatitude

`func (o *WhoisOrganization) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *WhoisOrganization) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *WhoisOrganization) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *WhoisOrganization) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *WhoisOrganization) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *WhoisOrganization) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *WhoisOrganization) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *WhoisOrganization) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetEmail

`func (o *WhoisOrganization) GetEmail() []string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WhoisOrganization) GetEmailOk() (*[]string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WhoisOrganization) SetEmail(v []string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WhoisOrganization) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAbuseMailbox

`func (o *WhoisOrganization) GetAbuseMailbox() []string`

GetAbuseMailbox returns the AbuseMailbox field if non-nil, zero value otherwise.

### GetAbuseMailboxOk

`func (o *WhoisOrganization) GetAbuseMailboxOk() (*[]string, bool)`

GetAbuseMailboxOk returns a tuple with the AbuseMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseMailbox

`func (o *WhoisOrganization) SetAbuseMailbox(v []string)`

SetAbuseMailbox sets AbuseMailbox field to given value.

### HasAbuseMailbox

`func (o *WhoisOrganization) HasAbuseMailbox() bool`

HasAbuseMailbox returns a boolean if a field has been set.

### GetPhone

`func (o *WhoisOrganization) GetPhone() []string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *WhoisOrganization) GetPhoneOk() (*[]string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *WhoisOrganization) SetPhone(v []string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *WhoisOrganization) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFaxNo

`func (o *WhoisOrganization) GetFaxNo() []string`

GetFaxNo returns the FaxNo field if non-nil, zero value otherwise.

### GetFaxNoOk

`func (o *WhoisOrganization) GetFaxNoOk() (*[]string, bool)`

GetFaxNoOk returns a tuple with the FaxNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNo

`func (o *WhoisOrganization) SetFaxNo(v []string)`

SetFaxNo sets FaxNo field to given value.

### HasFaxNo

`func (o *WhoisOrganization) HasFaxNo() bool`

HasFaxNo returns a boolean if a field has been set.

### GetOrganizations

`func (o *WhoisOrganization) GetOrganizations() []string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *WhoisOrganization) GetOrganizationsOk() (*[]string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *WhoisOrganization) SetOrganizations(v []string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *WhoisOrganization) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetAdminContacts

`func (o *WhoisOrganization) GetAdminContacts() []string`

GetAdminContacts returns the AdminContacts field if non-nil, zero value otherwise.

### GetAdminContactsOk

`func (o *WhoisOrganization) GetAdminContactsOk() (*[]string, bool)`

GetAdminContactsOk returns a tuple with the AdminContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminContacts

`func (o *WhoisOrganization) SetAdminContacts(v []string)`

SetAdminContacts sets AdminContacts field to given value.

### HasAdminContacts

`func (o *WhoisOrganization) HasAdminContacts() bool`

HasAdminContacts returns a boolean if a field has been set.

### GetTechContacts

`func (o *WhoisOrganization) GetTechContacts() []string`

GetTechContacts returns the TechContacts field if non-nil, zero value otherwise.

### GetTechContactsOk

`func (o *WhoisOrganization) GetTechContactsOk() (*[]string, bool)`

GetTechContactsOk returns a tuple with the TechContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechContacts

`func (o *WhoisOrganization) SetTechContacts(v []string)`

SetTechContacts sets TechContacts field to given value.

### HasTechContacts

`func (o *WhoisOrganization) HasTechContacts() bool`

HasTechContacts returns a boolean if a field has been set.

### GetAbuseContacts

`func (o *WhoisOrganization) GetAbuseContacts() []string`

GetAbuseContacts returns the AbuseContacts field if non-nil, zero value otherwise.

### GetAbuseContactsOk

`func (o *WhoisOrganization) GetAbuseContactsOk() (*[]string, bool)`

GetAbuseContactsOk returns a tuple with the AbuseContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseContacts

`func (o *WhoisOrganization) SetAbuseContacts(v []string)`

SetAbuseContacts sets AbuseContacts field to given value.

### HasAbuseContacts

`func (o *WhoisOrganization) HasAbuseContacts() bool`

HasAbuseContacts returns a boolean if a field has been set.

### GetLanguages

`func (o *WhoisOrganization) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *WhoisOrganization) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *WhoisOrganization) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *WhoisOrganization) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetRemarks

`func (o *WhoisOrganization) GetRemarks() []string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *WhoisOrganization) GetRemarksOk() (*[]string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *WhoisOrganization) SetRemarks(v []string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *WhoisOrganization) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetNotify

`func (o *WhoisOrganization) GetNotify() []string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *WhoisOrganization) GetNotifyOk() (*[]string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *WhoisOrganization) SetNotify(v []string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *WhoisOrganization) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetRefNfy

`func (o *WhoisOrganization) GetRefNfy() []string`

GetRefNfy returns the RefNfy field if non-nil, zero value otherwise.

### GetRefNfyOk

`func (o *WhoisOrganization) GetRefNfyOk() (*[]string, bool)`

GetRefNfyOk returns a tuple with the RefNfy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNfy

`func (o *WhoisOrganization) SetRefNfy(v []string)`

SetRefNfy sets RefNfy field to given value.

### HasRefNfy

`func (o *WhoisOrganization) HasRefNfy() bool`

HasRefNfy returns a boolean if a field has been set.

### GetMntRef

`func (o *WhoisOrganization) GetMntRef() []string`

GetMntRef returns the MntRef field if non-nil, zero value otherwise.

### GetMntRefOk

`func (o *WhoisOrganization) GetMntRefOk() (*[]string, bool)`

GetMntRefOk returns a tuple with the MntRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntRef

`func (o *WhoisOrganization) SetMntRef(v []string)`

SetMntRef sets MntRef field to given value.

### HasMntRef

`func (o *WhoisOrganization) HasMntRef() bool`

HasMntRef returns a boolean if a field has been set.

### GetMntBy

`func (o *WhoisOrganization) GetMntBy() []string`

GetMntBy returns the MntBy field if non-nil, zero value otherwise.

### GetMntByOk

`func (o *WhoisOrganization) GetMntByOk() (*[]string, bool)`

GetMntByOk returns a tuple with the MntBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntBy

`func (o *WhoisOrganization) SetMntBy(v []string)`

SetMntBy sets MntBy field to given value.

### HasMntBy

`func (o *WhoisOrganization) HasMntBy() bool`

HasMntBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *WhoisOrganization) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *WhoisOrganization) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *WhoisOrganization) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *WhoisOrganization) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *WhoisOrganization) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *WhoisOrganization) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *WhoisOrganization) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *WhoisOrganization) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetSource

`func (o *WhoisOrganization) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WhoisOrganization) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WhoisOrganization) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *WhoisOrganization) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


