# GeolocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**IpLocation**](IpLocation.md) |  | [optional] 
**CountryMetadata** | Pointer to [**CountryMetadata**](CountryMetadata.md) |  | [optional] 
**Network** | Pointer to [**GeoNetwork**](GeoNetwork.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 

## Methods

### NewGeolocationResponse

`func NewGeolocationResponse() *GeolocationResponse`

NewGeolocationResponse instantiates a new GeolocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeolocationResponseWithDefaults

`func NewGeolocationResponseWithDefaults() *GeolocationResponse`

NewGeolocationResponseWithDefaults instantiates a new GeolocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GeolocationResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GeolocationResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GeolocationResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GeolocationResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLocation

`func (o *GeolocationResponse) GetLocation() IpLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GeolocationResponse) GetLocationOk() (*IpLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GeolocationResponse) SetLocation(v IpLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GeolocationResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCountryMetadata

`func (o *GeolocationResponse) GetCountryMetadata() CountryMetadata`

GetCountryMetadata returns the CountryMetadata field if non-nil, zero value otherwise.

### GetCountryMetadataOk

`func (o *GeolocationResponse) GetCountryMetadataOk() (*CountryMetadata, bool)`

GetCountryMetadataOk returns a tuple with the CountryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryMetadata

`func (o *GeolocationResponse) SetCountryMetadata(v CountryMetadata)`

SetCountryMetadata sets CountryMetadata field to given value.

### HasCountryMetadata

`func (o *GeolocationResponse) HasCountryMetadata() bool`

HasCountryMetadata returns a boolean if a field has been set.

### GetNetwork

`func (o *GeolocationResponse) GetNetwork() GeoNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GeolocationResponse) GetNetworkOk() (*GeoNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GeolocationResponse) SetNetwork(v GeoNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GeolocationResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCurrency

`func (o *GeolocationResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GeolocationResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GeolocationResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GeolocationResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


