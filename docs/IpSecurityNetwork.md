# IpSecurityNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** |  | [optional] 
**Route** | Pointer to **string** |  | [optional] 
**IsAnycast** | Pointer to **bool** |  | [optional] 

## Methods

### NewIpSecurityNetwork

`func NewIpSecurityNetwork() *IpSecurityNetwork`

NewIpSecurityNetwork instantiates a new IpSecurityNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSecurityNetworkWithDefaults

`func NewIpSecurityNetworkWithDefaults() *IpSecurityNetwork`

NewIpSecurityNetworkWithDefaults instantiates a new IpSecurityNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *IpSecurityNetwork) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *IpSecurityNetwork) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *IpSecurityNetwork) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *IpSecurityNetwork) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetRoute

`func (o *IpSecurityNetwork) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *IpSecurityNetwork) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *IpSecurityNetwork) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *IpSecurityNetwork) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetIsAnycast

`func (o *IpSecurityNetwork) GetIsAnycast() bool`

GetIsAnycast returns the IsAnycast field if non-nil, zero value otherwise.

### GetIsAnycastOk

`func (o *IpSecurityNetwork) GetIsAnycastOk() (*bool, bool)`

GetIsAnycastOk returns a tuple with the IsAnycast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnycast

`func (o *IpSecurityNetwork) SetIsAnycast(v bool)`

SetIsAnycast sets IsAnycast field to given value.

### HasIsAnycast

`func (o *IpSecurityNetwork) HasIsAnycast() bool`

HasIsAnycast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


