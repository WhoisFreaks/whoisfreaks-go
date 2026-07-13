# BulkDomainAvailabilityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainNames** | Pointer to **[]string** | Mode 1 — list of full domain names. Cannot be used with tld. | [optional] 
**Tld** | Pointer to **[]string** | Mode 2 — TLDs to check against the domain query param. Cannot be used with domainNames. | [optional] 

## Methods

### NewBulkDomainAvailabilityRequest

`func NewBulkDomainAvailabilityRequest() *BulkDomainAvailabilityRequest`

NewBulkDomainAvailabilityRequest instantiates a new BulkDomainAvailabilityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDomainAvailabilityRequestWithDefaults

`func NewBulkDomainAvailabilityRequestWithDefaults() *BulkDomainAvailabilityRequest`

NewBulkDomainAvailabilityRequestWithDefaults instantiates a new BulkDomainAvailabilityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainNames

`func (o *BulkDomainAvailabilityRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *BulkDomainAvailabilityRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *BulkDomainAvailabilityRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *BulkDomainAvailabilityRequest) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetTld

`func (o *BulkDomainAvailabilityRequest) GetTld() []string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *BulkDomainAvailabilityRequest) GetTldOk() (*[]string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *BulkDomainAvailabilityRequest) SetTld(v []string)`

SetTld sets Tld field to given value.

### HasTld

`func (o *BulkDomainAvailabilityRequest) HasTld() bool`

HasTld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


