# DomainAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**DomainAvailableResponse** | Pointer to [**[]DomainAvailabilityItem**](DomainAvailabilityItem.md) |  | [optional] 

## Methods

### NewDomainAvailabilityResponse

`func NewDomainAvailabilityResponse() *DomainAvailabilityResponse`

NewDomainAvailabilityResponse instantiates a new DomainAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAvailabilityResponseWithDefaults

`func NewDomainAvailabilityResponseWithDefaults() *DomainAvailabilityResponse`

NewDomainAvailabilityResponseWithDefaults instantiates a new DomainAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainAvailabilityResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainAvailabilityResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainAvailabilityResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainAvailabilityResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetAvailability

`func (o *DomainAvailabilityResponse) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *DomainAvailabilityResponse) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *DomainAvailabilityResponse) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *DomainAvailabilityResponse) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetMessage

`func (o *DomainAvailabilityResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DomainAvailabilityResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DomainAvailabilityResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DomainAvailabilityResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDomainAvailableResponse

`func (o *DomainAvailabilityResponse) GetDomainAvailableResponse() []DomainAvailabilityItem`

GetDomainAvailableResponse returns the DomainAvailableResponse field if non-nil, zero value otherwise.

### GetDomainAvailableResponseOk

`func (o *DomainAvailabilityResponse) GetDomainAvailableResponseOk() (*[]DomainAvailabilityItem, bool)`

GetDomainAvailableResponseOk returns a tuple with the DomainAvailableResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAvailableResponse

`func (o *DomainAvailabilityResponse) SetDomainAvailableResponse(v []DomainAvailabilityItem)`

SetDomainAvailableResponse sets DomainAvailableResponse field to given value.

### HasDomainAvailableResponse

`func (o *DomainAvailabilityResponse) HasDomainAvailableResponse() bool`

HasDomainAvailableResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


