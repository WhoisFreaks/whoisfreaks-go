# InetNum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartIp** | Pointer to **string** |  | [optional] 
**EndIp** | Pointer to **string** |  | [optional] 
**Cidr** | Pointer to **[]string** |  | [optional] 
**NetName** | Pointer to **string** |  | [optional] 
**NetHandle** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **[]string** |  | [optional] 
**Countries** | Pointer to **[]string** |  | [optional] 
**Geofeed** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **float32** |  | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**SponsoringOrganization** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **[]string** |  | [optional] 
**AssignmentSize** | Pointer to **string** |  | [optional] 
**Notify** | Pointer to **[]string** |  | [optional] 
**MntBy** | Pointer to **[]string** |  | [optional] 
**MntLower** | Pointer to **[]string** |  | [optional] 
**MntDomains** | Pointer to **[]string** |  | [optional] 
**MntRoutes** | Pointer to **[]string** |  | [optional] 
**MntIrt** | Pointer to **[]string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Parents** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInetNum

`func NewInetNum() *InetNum`

NewInetNum instantiates a new InetNum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInetNumWithDefaults

`func NewInetNumWithDefaults() *InetNum`

NewInetNumWithDefaults instantiates a new InetNum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartIp

`func (o *InetNum) GetStartIp() string`

GetStartIp returns the StartIp field if non-nil, zero value otherwise.

### GetStartIpOk

`func (o *InetNum) GetStartIpOk() (*string, bool)`

GetStartIpOk returns a tuple with the StartIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIp

`func (o *InetNum) SetStartIp(v string)`

SetStartIp sets StartIp field to given value.

### HasStartIp

`func (o *InetNum) HasStartIp() bool`

HasStartIp returns a boolean if a field has been set.

### GetEndIp

`func (o *InetNum) GetEndIp() string`

GetEndIp returns the EndIp field if non-nil, zero value otherwise.

### GetEndIpOk

`func (o *InetNum) GetEndIpOk() (*string, bool)`

GetEndIpOk returns a tuple with the EndIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIp

`func (o *InetNum) SetEndIp(v string)`

SetEndIp sets EndIp field to given value.

### HasEndIp

`func (o *InetNum) HasEndIp() bool`

HasEndIp returns a boolean if a field has been set.

### GetCidr

`func (o *InetNum) GetCidr() []string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InetNum) GetCidrOk() (*[]string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InetNum) SetCidr(v []string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InetNum) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetNetName

`func (o *InetNum) GetNetName() string`

GetNetName returns the NetName field if non-nil, zero value otherwise.

### GetNetNameOk

`func (o *InetNum) GetNetNameOk() (*string, bool)`

GetNetNameOk returns a tuple with the NetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetName

`func (o *InetNum) SetNetName(v string)`

SetNetName sets NetName field to given value.

### HasNetName

`func (o *InetNum) HasNetName() bool`

HasNetName returns a boolean if a field has been set.

### GetNetHandle

`func (o *InetNum) GetNetHandle() string`

GetNetHandle returns the NetHandle field if non-nil, zero value otherwise.

### GetNetHandleOk

`func (o *InetNum) GetNetHandleOk() (*string, bool)`

GetNetHandleOk returns a tuple with the NetHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetHandle

`func (o *InetNum) SetNetHandle(v string)`

SetNetHandle sets NetHandle field to given value.

### HasNetHandle

`func (o *InetNum) HasNetHandle() bool`

HasNetHandle returns a boolean if a field has been set.

### GetDescription

`func (o *InetNum) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InetNum) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InetNum) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InetNum) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCountries

`func (o *InetNum) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *InetNum) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *InetNum) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *InetNum) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetGeofeed

`func (o *InetNum) GetGeofeed() string`

GetGeofeed returns the Geofeed field if non-nil, zero value otherwise.

### GetGeofeedOk

`func (o *InetNum) GetGeofeedOk() (*string, bool)`

GetGeofeedOk returns a tuple with the Geofeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofeed

`func (o *InetNum) SetGeofeed(v string)`

SetGeofeed sets Geofeed field to given value.

### HasGeofeed

`func (o *InetNum) HasGeofeed() bool`

HasGeofeed returns a boolean if a field has been set.

### GetLatitude

`func (o *InetNum) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *InetNum) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *InetNum) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *InetNum) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *InetNum) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *InetNum) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *InetNum) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *InetNum) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetCity

`func (o *InetNum) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *InetNum) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *InetNum) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *InetNum) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetLanguages

`func (o *InetNum) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *InetNum) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *InetNum) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *InetNum) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetStatus

`func (o *InetNum) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InetNum) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InetNum) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InetNum) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrganization

`func (o *InetNum) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InetNum) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InetNum) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *InetNum) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSponsoringOrganization

`func (o *InetNum) GetSponsoringOrganization() string`

GetSponsoringOrganization returns the SponsoringOrganization field if non-nil, zero value otherwise.

### GetSponsoringOrganizationOk

`func (o *InetNum) GetSponsoringOrganizationOk() (*string, bool)`

GetSponsoringOrganizationOk returns a tuple with the SponsoringOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsoringOrganization

`func (o *InetNum) SetSponsoringOrganization(v string)`

SetSponsoringOrganization sets SponsoringOrganization field to given value.

### HasSponsoringOrganization

`func (o *InetNum) HasSponsoringOrganization() bool`

HasSponsoringOrganization returns a boolean if a field has been set.

### GetRemarks

`func (o *InetNum) GetRemarks() []string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *InetNum) GetRemarksOk() (*[]string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *InetNum) SetRemarks(v []string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *InetNum) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetAssignmentSize

`func (o *InetNum) GetAssignmentSize() string`

GetAssignmentSize returns the AssignmentSize field if non-nil, zero value otherwise.

### GetAssignmentSizeOk

`func (o *InetNum) GetAssignmentSizeOk() (*string, bool)`

GetAssignmentSizeOk returns a tuple with the AssignmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSize

`func (o *InetNum) SetAssignmentSize(v string)`

SetAssignmentSize sets AssignmentSize field to given value.

### HasAssignmentSize

`func (o *InetNum) HasAssignmentSize() bool`

HasAssignmentSize returns a boolean if a field has been set.

### GetNotify

`func (o *InetNum) GetNotify() []string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *InetNum) GetNotifyOk() (*[]string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *InetNum) SetNotify(v []string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *InetNum) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetMntBy

`func (o *InetNum) GetMntBy() []string`

GetMntBy returns the MntBy field if non-nil, zero value otherwise.

### GetMntByOk

`func (o *InetNum) GetMntByOk() (*[]string, bool)`

GetMntByOk returns a tuple with the MntBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntBy

`func (o *InetNum) SetMntBy(v []string)`

SetMntBy sets MntBy field to given value.

### HasMntBy

`func (o *InetNum) HasMntBy() bool`

HasMntBy returns a boolean if a field has been set.

### GetMntLower

`func (o *InetNum) GetMntLower() []string`

GetMntLower returns the MntLower field if non-nil, zero value otherwise.

### GetMntLowerOk

`func (o *InetNum) GetMntLowerOk() (*[]string, bool)`

GetMntLowerOk returns a tuple with the MntLower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntLower

`func (o *InetNum) SetMntLower(v []string)`

SetMntLower sets MntLower field to given value.

### HasMntLower

`func (o *InetNum) HasMntLower() bool`

HasMntLower returns a boolean if a field has been set.

### GetMntDomains

`func (o *InetNum) GetMntDomains() []string`

GetMntDomains returns the MntDomains field if non-nil, zero value otherwise.

### GetMntDomainsOk

`func (o *InetNum) GetMntDomainsOk() (*[]string, bool)`

GetMntDomainsOk returns a tuple with the MntDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntDomains

`func (o *InetNum) SetMntDomains(v []string)`

SetMntDomains sets MntDomains field to given value.

### HasMntDomains

`func (o *InetNum) HasMntDomains() bool`

HasMntDomains returns a boolean if a field has been set.

### GetMntRoutes

`func (o *InetNum) GetMntRoutes() []string`

GetMntRoutes returns the MntRoutes field if non-nil, zero value otherwise.

### GetMntRoutesOk

`func (o *InetNum) GetMntRoutesOk() (*[]string, bool)`

GetMntRoutesOk returns a tuple with the MntRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntRoutes

`func (o *InetNum) SetMntRoutes(v []string)`

SetMntRoutes sets MntRoutes field to given value.

### HasMntRoutes

`func (o *InetNum) HasMntRoutes() bool`

HasMntRoutes returns a boolean if a field has been set.

### GetMntIrt

`func (o *InetNum) GetMntIrt() []string`

GetMntIrt returns the MntIrt field if non-nil, zero value otherwise.

### GetMntIrtOk

`func (o *InetNum) GetMntIrtOk() (*[]string, bool)`

GetMntIrtOk returns a tuple with the MntIrt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntIrt

`func (o *InetNum) SetMntIrt(v []string)`

SetMntIrt sets MntIrt field to given value.

### HasMntIrt

`func (o *InetNum) HasMntIrt() bool`

HasMntIrt returns a boolean if a field has been set.

### GetDateCreated

`func (o *InetNum) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InetNum) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InetNum) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InetNum) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *InetNum) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *InetNum) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *InetNum) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *InetNum) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetSource

`func (o *InetNum) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InetNum) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InetNum) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InetNum) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetParents

`func (o *InetNum) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *InetNum) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *InetNum) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *InetNum) HasParents() bool`

HasParents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


