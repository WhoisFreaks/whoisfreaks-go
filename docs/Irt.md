# Irt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** |  | [optional] 
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
**Signature** | Pointer to **[]string** |  | [optional] 
**Encryption** | Pointer to **[]string** |  | [optional] 
**Auth** | Pointer to **[]string** |  | [optional] 
**Notify** | Pointer to **[]string** |  | [optional] 
**IrtNfy** | Pointer to **[]string** |  | [optional] 
**MntBy** | Pointer to **[]string** |  | [optional] 
**MntRef** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIrt

`func NewIrt() *Irt`

NewIrt instantiates a new Irt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIrtWithDefaults

`func NewIrtWithDefaults() *Irt`

NewIrtWithDefaults instantiates a new Irt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *Irt) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Irt) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Irt) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *Irt) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetAddress

`func (o *Irt) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Irt) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Irt) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Irt) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStreet

`func (o *Irt) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *Irt) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *Irt) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *Irt) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *Irt) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Irt) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Irt) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Irt) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDistrict

`func (o *Irt) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *Irt) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *Irt) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *Irt) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetState

`func (o *Irt) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Irt) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Irt) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Irt) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *Irt) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Irt) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Irt) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *Irt) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCountry

`func (o *Irt) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Irt) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Irt) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Irt) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *Irt) GetEmail() []string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Irt) GetEmailOk() (*[]string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Irt) SetEmail(v []string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Irt) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAbuseMailbox

`func (o *Irt) GetAbuseMailbox() []string`

GetAbuseMailbox returns the AbuseMailbox field if non-nil, zero value otherwise.

### GetAbuseMailboxOk

`func (o *Irt) GetAbuseMailboxOk() (*[]string, bool)`

GetAbuseMailboxOk returns a tuple with the AbuseMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseMailbox

`func (o *Irt) SetAbuseMailbox(v []string)`

SetAbuseMailbox sets AbuseMailbox field to given value.

### HasAbuseMailbox

`func (o *Irt) HasAbuseMailbox() bool`

HasAbuseMailbox returns a boolean if a field has been set.

### GetPhone

`func (o *Irt) GetPhone() []string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Irt) GetPhoneOk() (*[]string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Irt) SetPhone(v []string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Irt) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFaxNo

`func (o *Irt) GetFaxNo() []string`

GetFaxNo returns the FaxNo field if non-nil, zero value otherwise.

### GetFaxNoOk

`func (o *Irt) GetFaxNoOk() (*[]string, bool)`

GetFaxNoOk returns a tuple with the FaxNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNo

`func (o *Irt) SetFaxNo(v []string)`

SetFaxNo sets FaxNo field to given value.

### HasFaxNo

`func (o *Irt) HasFaxNo() bool`

HasFaxNo returns a boolean if a field has been set.

### GetOrganizations

`func (o *Irt) GetOrganizations() []string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *Irt) GetOrganizationsOk() (*[]string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *Irt) SetOrganizations(v []string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *Irt) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetAdminContacts

`func (o *Irt) GetAdminContacts() []string`

GetAdminContacts returns the AdminContacts field if non-nil, zero value otherwise.

### GetAdminContactsOk

`func (o *Irt) GetAdminContactsOk() (*[]string, bool)`

GetAdminContactsOk returns a tuple with the AdminContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminContacts

`func (o *Irt) SetAdminContacts(v []string)`

SetAdminContacts sets AdminContacts field to given value.

### HasAdminContacts

`func (o *Irt) HasAdminContacts() bool`

HasAdminContacts returns a boolean if a field has been set.

### GetTechContacts

`func (o *Irt) GetTechContacts() []string`

GetTechContacts returns the TechContacts field if non-nil, zero value otherwise.

### GetTechContactsOk

`func (o *Irt) GetTechContactsOk() (*[]string, bool)`

GetTechContactsOk returns a tuple with the TechContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechContacts

`func (o *Irt) SetTechContacts(v []string)`

SetTechContacts sets TechContacts field to given value.

### HasTechContacts

`func (o *Irt) HasTechContacts() bool`

HasTechContacts returns a boolean if a field has been set.

### GetRemarks

`func (o *Irt) GetRemarks() []string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *Irt) GetRemarksOk() (*[]string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *Irt) SetRemarks(v []string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *Irt) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSignature

`func (o *Irt) GetSignature() []string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Irt) GetSignatureOk() (*[]string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Irt) SetSignature(v []string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Irt) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetEncryption

`func (o *Irt) GetEncryption() []string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *Irt) GetEncryptionOk() (*[]string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *Irt) SetEncryption(v []string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *Irt) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetAuth

`func (o *Irt) GetAuth() []string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Irt) GetAuthOk() (*[]string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Irt) SetAuth(v []string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *Irt) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetNotify

`func (o *Irt) GetNotify() []string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *Irt) GetNotifyOk() (*[]string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *Irt) SetNotify(v []string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *Irt) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetIrtNfy

`func (o *Irt) GetIrtNfy() []string`

GetIrtNfy returns the IrtNfy field if non-nil, zero value otherwise.

### GetIrtNfyOk

`func (o *Irt) GetIrtNfyOk() (*[]string, bool)`

GetIrtNfyOk returns a tuple with the IrtNfy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrtNfy

`func (o *Irt) SetIrtNfy(v []string)`

SetIrtNfy sets IrtNfy field to given value.

### HasIrtNfy

`func (o *Irt) HasIrtNfy() bool`

HasIrtNfy returns a boolean if a field has been set.

### GetMntBy

`func (o *Irt) GetMntBy() []string`

GetMntBy returns the MntBy field if non-nil, zero value otherwise.

### GetMntByOk

`func (o *Irt) GetMntByOk() (*[]string, bool)`

GetMntByOk returns a tuple with the MntBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntBy

`func (o *Irt) SetMntBy(v []string)`

SetMntBy sets MntBy field to given value.

### HasMntBy

`func (o *Irt) HasMntBy() bool`

HasMntBy returns a boolean if a field has been set.

### GetMntRef

`func (o *Irt) GetMntRef() []string`

GetMntRef returns the MntRef field if non-nil, zero value otherwise.

### GetMntRefOk

`func (o *Irt) GetMntRefOk() (*[]string, bool)`

GetMntRefOk returns a tuple with the MntRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntRef

`func (o *Irt) SetMntRef(v []string)`

SetMntRef sets MntRef field to given value.

### HasMntRef

`func (o *Irt) HasMntRef() bool`

HasMntRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


