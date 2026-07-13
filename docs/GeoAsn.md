# GeoAsn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**AsnName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**DateAllocated** | Pointer to **string** |  | [optional] 
**AllocationStatus** | Pointer to **string** |  | [optional] 
**NumOfIpv4Routes** | Pointer to **string** |  | [optional] 
**NumOfIpv6Routes** | Pointer to **string** |  | [optional] 
**Rir** | Pointer to **string** |  | [optional] 

## Methods

### NewGeoAsn

`func NewGeoAsn() *GeoAsn`

NewGeoAsn instantiates a new GeoAsn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoAsnWithDefaults

`func NewGeoAsnWithDefaults() *GeoAsn`

NewGeoAsnWithDefaults instantiates a new GeoAsn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *GeoAsn) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *GeoAsn) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *GeoAsn) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *GeoAsn) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetOrganization

`func (o *GeoAsn) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GeoAsn) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GeoAsn) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GeoAsn) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCountry

`func (o *GeoAsn) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GeoAsn) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GeoAsn) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GeoAsn) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAsnName

`func (o *GeoAsn) GetAsnName() string`

GetAsnName returns the AsnName field if non-nil, zero value otherwise.

### GetAsnNameOk

`func (o *GeoAsn) GetAsnNameOk() (*string, bool)`

GetAsnNameOk returns a tuple with the AsnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnName

`func (o *GeoAsn) SetAsnName(v string)`

SetAsnName sets AsnName field to given value.

### HasAsnName

`func (o *GeoAsn) HasAsnName() bool`

HasAsnName returns a boolean if a field has been set.

### GetType

`func (o *GeoAsn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoAsn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoAsn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GeoAsn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomain

`func (o *GeoAsn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GeoAsn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GeoAsn) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GeoAsn) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDateAllocated

`func (o *GeoAsn) GetDateAllocated() string`

GetDateAllocated returns the DateAllocated field if non-nil, zero value otherwise.

### GetDateAllocatedOk

`func (o *GeoAsn) GetDateAllocatedOk() (*string, bool)`

GetDateAllocatedOk returns a tuple with the DateAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAllocated

`func (o *GeoAsn) SetDateAllocated(v string)`

SetDateAllocated sets DateAllocated field to given value.

### HasDateAllocated

`func (o *GeoAsn) HasDateAllocated() bool`

HasDateAllocated returns a boolean if a field has been set.

### GetAllocationStatus

`func (o *GeoAsn) GetAllocationStatus() string`

GetAllocationStatus returns the AllocationStatus field if non-nil, zero value otherwise.

### GetAllocationStatusOk

`func (o *GeoAsn) GetAllocationStatusOk() (*string, bool)`

GetAllocationStatusOk returns a tuple with the AllocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationStatus

`func (o *GeoAsn) SetAllocationStatus(v string)`

SetAllocationStatus sets AllocationStatus field to given value.

### HasAllocationStatus

`func (o *GeoAsn) HasAllocationStatus() bool`

HasAllocationStatus returns a boolean if a field has been set.

### GetNumOfIpv4Routes

`func (o *GeoAsn) GetNumOfIpv4Routes() string`

GetNumOfIpv4Routes returns the NumOfIpv4Routes field if non-nil, zero value otherwise.

### GetNumOfIpv4RoutesOk

`func (o *GeoAsn) GetNumOfIpv4RoutesOk() (*string, bool)`

GetNumOfIpv4RoutesOk returns a tuple with the NumOfIpv4Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv4Routes

`func (o *GeoAsn) SetNumOfIpv4Routes(v string)`

SetNumOfIpv4Routes sets NumOfIpv4Routes field to given value.

### HasNumOfIpv4Routes

`func (o *GeoAsn) HasNumOfIpv4Routes() bool`

HasNumOfIpv4Routes returns a boolean if a field has been set.

### GetNumOfIpv6Routes

`func (o *GeoAsn) GetNumOfIpv6Routes() string`

GetNumOfIpv6Routes returns the NumOfIpv6Routes field if non-nil, zero value otherwise.

### GetNumOfIpv6RoutesOk

`func (o *GeoAsn) GetNumOfIpv6RoutesOk() (*string, bool)`

GetNumOfIpv6RoutesOk returns a tuple with the NumOfIpv6Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv6Routes

`func (o *GeoAsn) SetNumOfIpv6Routes(v string)`

SetNumOfIpv6Routes sets NumOfIpv6Routes field to given value.

### HasNumOfIpv6Routes

`func (o *GeoAsn) HasNumOfIpv6Routes() bool`

HasNumOfIpv6Routes returns a boolean if a field has been set.

### GetRir

`func (o *GeoAsn) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *GeoAsn) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *GeoAsn) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *GeoAsn) HasRir() bool`

HasRir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


