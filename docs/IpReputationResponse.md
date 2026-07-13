# IpReputationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**IpLocation**](IpLocation.md) |  | [optional] 
**Network** | Pointer to [**IpSecurityNetwork**](IpSecurityNetwork.md) |  | [optional] 
**Asn** | Pointer to [**IpSecurityAsn**](IpSecurityAsn.md) |  | [optional] 
**Security** | Pointer to [**IpSecurity**](IpSecurity.md) |  | [optional] 

## Methods

### NewIpReputationResponse

`func NewIpReputationResponse() *IpReputationResponse`

NewIpReputationResponse instantiates a new IpReputationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpReputationResponseWithDefaults

`func NewIpReputationResponseWithDefaults() *IpReputationResponse`

NewIpReputationResponseWithDefaults instantiates a new IpReputationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpReputationResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpReputationResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpReputationResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IpReputationResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLocation

`func (o *IpReputationResponse) GetLocation() IpLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IpReputationResponse) GetLocationOk() (*IpLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IpReputationResponse) SetLocation(v IpLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IpReputationResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNetwork

`func (o *IpReputationResponse) GetNetwork() IpSecurityNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IpReputationResponse) GetNetworkOk() (*IpSecurityNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IpReputationResponse) SetNetwork(v IpSecurityNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IpReputationResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAsn

`func (o *IpReputationResponse) GetAsn() IpSecurityAsn`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *IpReputationResponse) GetAsnOk() (*IpSecurityAsn, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *IpReputationResponse) SetAsn(v IpSecurityAsn)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *IpReputationResponse) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetSecurity

`func (o *IpReputationResponse) GetSecurity() IpSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *IpReputationResponse) GetSecurityOk() (*IpSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *IpReputationResponse) SetSecurity(v IpSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *IpReputationResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


