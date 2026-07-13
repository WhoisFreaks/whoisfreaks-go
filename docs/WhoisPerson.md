# WhoisPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **[]string** |  | [optional] 
**Street** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**District** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**ZipCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **[]string** |  | [optional] 
**AbuseMailbox** | Pointer to **[]string** |  | [optional] 
**Phone** | Pointer to **[]string** |  | [optional] 
**FaxNo** | Pointer to **[]string** |  | [optional] 
**Organizations** | Pointer to **[]string** |  | [optional] 
**AdminContacts** | Pointer to **[]string** |  | [optional] 
**TechContacts** | Pointer to **[]string** |  | [optional] 
**Remarks** | Pointer to **[]string** |  | [optional] 
**Notify** | Pointer to **[]string** |  | [optional] 
**MntBy** | Pointer to **[]string** |  | [optional] 
**MntRef** | Pointer to **[]string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewWhoisPerson

`func NewWhoisPerson() *WhoisPerson`

NewWhoisPerson instantiates a new WhoisPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoisPersonWithDefaults

`func NewWhoisPersonWithDefaults() *WhoisPerson`

NewWhoisPersonWithDefaults instantiates a new WhoisPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *WhoisPerson) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *WhoisPerson) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *WhoisPerson) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *WhoisPerson) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetName

`func (o *WhoisPerson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhoisPerson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhoisPerson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhoisPerson) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *WhoisPerson) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WhoisPerson) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WhoisPerson) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WhoisPerson) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStreet

`func (o *WhoisPerson) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *WhoisPerson) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *WhoisPerson) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *WhoisPerson) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *WhoisPerson) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *WhoisPerson) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *WhoisPerson) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *WhoisPerson) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDistrict

`func (o *WhoisPerson) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *WhoisPerson) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *WhoisPerson) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *WhoisPerson) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetState

`func (o *WhoisPerson) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WhoisPerson) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WhoisPerson) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WhoisPerson) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *WhoisPerson) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *WhoisPerson) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *WhoisPerson) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *WhoisPerson) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCountry

`func (o *WhoisPerson) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WhoisPerson) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WhoisPerson) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *WhoisPerson) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *WhoisPerson) GetEmail() []string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WhoisPerson) GetEmailOk() (*[]string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WhoisPerson) SetEmail(v []string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WhoisPerson) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAbuseMailbox

`func (o *WhoisPerson) GetAbuseMailbox() []string`

GetAbuseMailbox returns the AbuseMailbox field if non-nil, zero value otherwise.

### GetAbuseMailboxOk

`func (o *WhoisPerson) GetAbuseMailboxOk() (*[]string, bool)`

GetAbuseMailboxOk returns a tuple with the AbuseMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseMailbox

`func (o *WhoisPerson) SetAbuseMailbox(v []string)`

SetAbuseMailbox sets AbuseMailbox field to given value.

### HasAbuseMailbox

`func (o *WhoisPerson) HasAbuseMailbox() bool`

HasAbuseMailbox returns a boolean if a field has been set.

### GetPhone

`func (o *WhoisPerson) GetPhone() []string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *WhoisPerson) GetPhoneOk() (*[]string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *WhoisPerson) SetPhone(v []string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *WhoisPerson) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFaxNo

`func (o *WhoisPerson) GetFaxNo() []string`

GetFaxNo returns the FaxNo field if non-nil, zero value otherwise.

### GetFaxNoOk

`func (o *WhoisPerson) GetFaxNoOk() (*[]string, bool)`

GetFaxNoOk returns a tuple with the FaxNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNo

`func (o *WhoisPerson) SetFaxNo(v []string)`

SetFaxNo sets FaxNo field to given value.

### HasFaxNo

`func (o *WhoisPerson) HasFaxNo() bool`

HasFaxNo returns a boolean if a field has been set.

### GetOrganizations

`func (o *WhoisPerson) GetOrganizations() []string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *WhoisPerson) GetOrganizationsOk() (*[]string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *WhoisPerson) SetOrganizations(v []string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *WhoisPerson) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetAdminContacts

`func (o *WhoisPerson) GetAdminContacts() []string`

GetAdminContacts returns the AdminContacts field if non-nil, zero value otherwise.

### GetAdminContactsOk

`func (o *WhoisPerson) GetAdminContactsOk() (*[]string, bool)`

GetAdminContactsOk returns a tuple with the AdminContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminContacts

`func (o *WhoisPerson) SetAdminContacts(v []string)`

SetAdminContacts sets AdminContacts field to given value.

### HasAdminContacts

`func (o *WhoisPerson) HasAdminContacts() bool`

HasAdminContacts returns a boolean if a field has been set.

### GetTechContacts

`func (o *WhoisPerson) GetTechContacts() []string`

GetTechContacts returns the TechContacts field if non-nil, zero value otherwise.

### GetTechContactsOk

`func (o *WhoisPerson) GetTechContactsOk() (*[]string, bool)`

GetTechContactsOk returns a tuple with the TechContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechContacts

`func (o *WhoisPerson) SetTechContacts(v []string)`

SetTechContacts sets TechContacts field to given value.

### HasTechContacts

`func (o *WhoisPerson) HasTechContacts() bool`

HasTechContacts returns a boolean if a field has been set.

### GetRemarks

`func (o *WhoisPerson) GetRemarks() []string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *WhoisPerson) GetRemarksOk() (*[]string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *WhoisPerson) SetRemarks(v []string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *WhoisPerson) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetNotify

`func (o *WhoisPerson) GetNotify() []string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *WhoisPerson) GetNotifyOk() (*[]string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *WhoisPerson) SetNotify(v []string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *WhoisPerson) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetMntBy

`func (o *WhoisPerson) GetMntBy() []string`

GetMntBy returns the MntBy field if non-nil, zero value otherwise.

### GetMntByOk

`func (o *WhoisPerson) GetMntByOk() (*[]string, bool)`

GetMntByOk returns a tuple with the MntBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntBy

`func (o *WhoisPerson) SetMntBy(v []string)`

SetMntBy sets MntBy field to given value.

### HasMntBy

`func (o *WhoisPerson) HasMntBy() bool`

HasMntBy returns a boolean if a field has been set.

### GetMntRef

`func (o *WhoisPerson) GetMntRef() []string`

GetMntRef returns the MntRef field if non-nil, zero value otherwise.

### GetMntRefOk

`func (o *WhoisPerson) GetMntRefOk() (*[]string, bool)`

GetMntRefOk returns a tuple with the MntRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntRef

`func (o *WhoisPerson) SetMntRef(v []string)`

SetMntRef sets MntRef field to given value.

### HasMntRef

`func (o *WhoisPerson) HasMntRef() bool`

HasMntRef returns a boolean if a field has been set.

### GetDateCreated

`func (o *WhoisPerson) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *WhoisPerson) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *WhoisPerson) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *WhoisPerson) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *WhoisPerson) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *WhoisPerson) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *WhoisPerson) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *WhoisPerson) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetSource

`func (o *WhoisPerson) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WhoisPerson) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WhoisPerson) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *WhoisPerson) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


