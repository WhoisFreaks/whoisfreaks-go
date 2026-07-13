# BulkWhoisItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkWhoisItem

`func NewBulkWhoisItem() *BulkWhoisItem`

NewBulkWhoisItem instantiates a new BulkWhoisItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWhoisItemWithDefaults

`func NewBulkWhoisItemWithDefaults() *BulkWhoisItem`

NewBulkWhoisItemWithDefaults instantiates a new BulkWhoisItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BulkWhoisItem) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWhoisItem) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWhoisItem) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkWhoisItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDomainName

`func (o *BulkWhoisItem) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *BulkWhoisItem) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *BulkWhoisItem) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *BulkWhoisItem) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetError

`func (o *BulkWhoisItem) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkWhoisItem) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkWhoisItem) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkWhoisItem) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


