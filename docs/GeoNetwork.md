# GeoNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to [**GeoAsn**](GeoAsn.md) |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**GeoCompany**](GeoCompany.md) |  | [optional] 

## Methods

### NewGeoNetwork

`func NewGeoNetwork() *GeoNetwork`

NewGeoNetwork instantiates a new GeoNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoNetworkWithDefaults

`func NewGeoNetworkWithDefaults() *GeoNetwork`

NewGeoNetworkWithDefaults instantiates a new GeoNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *GeoNetwork) GetAsn() GeoAsn`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *GeoNetwork) GetAsnOk() (*GeoAsn, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *GeoNetwork) SetAsn(v GeoAsn)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *GeoNetwork) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetConnectionType

`func (o *GeoNetwork) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GeoNetwork) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GeoNetwork) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GeoNetwork) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetCompany

`func (o *GeoNetwork) GetCompany() GeoCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GeoNetwork) GetCompanyOk() (*GeoCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GeoNetwork) SetCompany(v GeoCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GeoNetwork) HasCompany() bool`

HasCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


