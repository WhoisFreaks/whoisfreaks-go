# AsnInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**DateAllocated** | Pointer to **string** |  | [optional] 
**AsnName** | Pointer to **string** |  | [optional] 
**AllocationStatus** | Pointer to **string** |  | [optional] 
**NumOfIpv4Routes** | Pointer to **string** |  | [optional] 
**NumOfIpv6Routes** | Pointer to **string** |  | [optional] 
**Rir** | Pointer to **string** |  | [optional] 
**Routes** | Pointer to **[]string** |  | [optional] 
**Downstreams** | Pointer to [**[]AsnPeer**](AsnPeer.md) |  | [optional] 
**Upstreams** | Pointer to [**[]AsnPeer**](AsnPeer.md) |  | [optional] 
**Peers** | Pointer to [**[]AsnPeer**](AsnPeer.md) |  | [optional] 
**WhoisResponse** | Pointer to **string** |  | [optional] 

## Methods

### NewAsnInfo

`func NewAsnInfo() *AsnInfo`

NewAsnInfo instantiates a new AsnInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnInfoWithDefaults

`func NewAsnInfoWithDefaults() *AsnInfo`

NewAsnInfoWithDefaults instantiates a new AsnInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *AsnInfo) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *AsnInfo) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *AsnInfo) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *AsnInfo) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetOrganization

`func (o *AsnInfo) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AsnInfo) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AsnInfo) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AsnInfo) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCountry

`func (o *AsnInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AsnInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AsnInfo) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AsnInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetType

`func (o *AsnInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AsnInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AsnInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AsnInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomain

`func (o *AsnInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AsnInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AsnInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AsnInfo) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDateAllocated

`func (o *AsnInfo) GetDateAllocated() string`

GetDateAllocated returns the DateAllocated field if non-nil, zero value otherwise.

### GetDateAllocatedOk

`func (o *AsnInfo) GetDateAllocatedOk() (*string, bool)`

GetDateAllocatedOk returns a tuple with the DateAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAllocated

`func (o *AsnInfo) SetDateAllocated(v string)`

SetDateAllocated sets DateAllocated field to given value.

### HasDateAllocated

`func (o *AsnInfo) HasDateAllocated() bool`

HasDateAllocated returns a boolean if a field has been set.

### GetAsnName

`func (o *AsnInfo) GetAsnName() string`

GetAsnName returns the AsnName field if non-nil, zero value otherwise.

### GetAsnNameOk

`func (o *AsnInfo) GetAsnNameOk() (*string, bool)`

GetAsnNameOk returns a tuple with the AsnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnName

`func (o *AsnInfo) SetAsnName(v string)`

SetAsnName sets AsnName field to given value.

### HasAsnName

`func (o *AsnInfo) HasAsnName() bool`

HasAsnName returns a boolean if a field has been set.

### GetAllocationStatus

`func (o *AsnInfo) GetAllocationStatus() string`

GetAllocationStatus returns the AllocationStatus field if non-nil, zero value otherwise.

### GetAllocationStatusOk

`func (o *AsnInfo) GetAllocationStatusOk() (*string, bool)`

GetAllocationStatusOk returns a tuple with the AllocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationStatus

`func (o *AsnInfo) SetAllocationStatus(v string)`

SetAllocationStatus sets AllocationStatus field to given value.

### HasAllocationStatus

`func (o *AsnInfo) HasAllocationStatus() bool`

HasAllocationStatus returns a boolean if a field has been set.

### GetNumOfIpv4Routes

`func (o *AsnInfo) GetNumOfIpv4Routes() string`

GetNumOfIpv4Routes returns the NumOfIpv4Routes field if non-nil, zero value otherwise.

### GetNumOfIpv4RoutesOk

`func (o *AsnInfo) GetNumOfIpv4RoutesOk() (*string, bool)`

GetNumOfIpv4RoutesOk returns a tuple with the NumOfIpv4Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv4Routes

`func (o *AsnInfo) SetNumOfIpv4Routes(v string)`

SetNumOfIpv4Routes sets NumOfIpv4Routes field to given value.

### HasNumOfIpv4Routes

`func (o *AsnInfo) HasNumOfIpv4Routes() bool`

HasNumOfIpv4Routes returns a boolean if a field has been set.

### GetNumOfIpv6Routes

`func (o *AsnInfo) GetNumOfIpv6Routes() string`

GetNumOfIpv6Routes returns the NumOfIpv6Routes field if non-nil, zero value otherwise.

### GetNumOfIpv6RoutesOk

`func (o *AsnInfo) GetNumOfIpv6RoutesOk() (*string, bool)`

GetNumOfIpv6RoutesOk returns a tuple with the NumOfIpv6Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv6Routes

`func (o *AsnInfo) SetNumOfIpv6Routes(v string)`

SetNumOfIpv6Routes sets NumOfIpv6Routes field to given value.

### HasNumOfIpv6Routes

`func (o *AsnInfo) HasNumOfIpv6Routes() bool`

HasNumOfIpv6Routes returns a boolean if a field has been set.

### GetRir

`func (o *AsnInfo) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *AsnInfo) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *AsnInfo) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *AsnInfo) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRoutes

`func (o *AsnInfo) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *AsnInfo) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *AsnInfo) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *AsnInfo) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetDownstreams

`func (o *AsnInfo) GetDownstreams() []AsnPeer`

GetDownstreams returns the Downstreams field if non-nil, zero value otherwise.

### GetDownstreamsOk

`func (o *AsnInfo) GetDownstreamsOk() (*[]AsnPeer, bool)`

GetDownstreamsOk returns a tuple with the Downstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreams

`func (o *AsnInfo) SetDownstreams(v []AsnPeer)`

SetDownstreams sets Downstreams field to given value.

### HasDownstreams

`func (o *AsnInfo) HasDownstreams() bool`

HasDownstreams returns a boolean if a field has been set.

### GetUpstreams

`func (o *AsnInfo) GetUpstreams() []AsnPeer`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *AsnInfo) GetUpstreamsOk() (*[]AsnPeer, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *AsnInfo) SetUpstreams(v []AsnPeer)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *AsnInfo) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.

### GetPeers

`func (o *AsnInfo) GetPeers() []AsnPeer`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *AsnInfo) GetPeersOk() (*[]AsnPeer, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *AsnInfo) SetPeers(v []AsnPeer)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *AsnInfo) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetWhoisResponse

`func (o *AsnInfo) GetWhoisResponse() string`

GetWhoisResponse returns the WhoisResponse field if non-nil, zero value otherwise.

### GetWhoisResponseOk

`func (o *AsnInfo) GetWhoisResponseOk() (*string, bool)`

GetWhoisResponseOk returns a tuple with the WhoisResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisResponse

`func (o *AsnInfo) SetWhoisResponse(v string)`

SetWhoisResponse sets WhoisResponse field to given value.

### HasWhoisResponse

`func (o *AsnInfo) HasWhoisResponse() bool`

HasWhoisResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


