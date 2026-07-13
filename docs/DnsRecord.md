# DnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Strings** | Pointer to **[]string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Admin** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **int32** |  | [optional] 
**Refresh** | Pointer to **int32** |  | [optional] 
**Retry** | Pointer to **int32** |  | [optional] 
**Expire** | Pointer to **int32** |  | [optional] 
**Minimum** | Pointer to **int32** |  | [optional] 

## Methods

### NewDnsRecord

`func NewDnsRecord() *DnsRecord`

NewDnsRecord instantiates a new DnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordWithDefaults

`func NewDnsRecordWithDefaults() *DnsRecord`

NewDnsRecordWithDefaults instantiates a new DnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DnsRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DnsRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DnsRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *DnsRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAddress

`func (o *DnsRecord) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DnsRecord) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DnsRecord) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DnsRecord) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPriority

`func (o *DnsRecord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DnsRecord) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DnsRecord) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DnsRecord) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetExchange

`func (o *DnsRecord) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *DnsRecord) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *DnsRecord) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *DnsRecord) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetTarget

`func (o *DnsRecord) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DnsRecord) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DnsRecord) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DnsRecord) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStrings

`func (o *DnsRecord) GetStrings() []string`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *DnsRecord) GetStringsOk() (*[]string, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *DnsRecord) SetStrings(v []string)`

SetStrings sets Strings field to given value.

### HasStrings

`func (o *DnsRecord) HasStrings() bool`

HasStrings returns a boolean if a field has been set.

### GetHost

`func (o *DnsRecord) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DnsRecord) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DnsRecord) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DnsRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAdmin

`func (o *DnsRecord) GetAdmin() string`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *DnsRecord) GetAdminOk() (*string, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *DnsRecord) SetAdmin(v string)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *DnsRecord) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetSerial

`func (o *DnsRecord) GetSerial() int32`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DnsRecord) GetSerialOk() (*int32, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DnsRecord) SetSerial(v int32)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DnsRecord) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetRefresh

`func (o *DnsRecord) GetRefresh() int32`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *DnsRecord) GetRefreshOk() (*int32, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *DnsRecord) SetRefresh(v int32)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *DnsRecord) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRetry

`func (o *DnsRecord) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *DnsRecord) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *DnsRecord) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *DnsRecord) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetExpire

`func (o *DnsRecord) GetExpire() int32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *DnsRecord) GetExpireOk() (*int32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *DnsRecord) SetExpire(v int32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *DnsRecord) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetMinimum

`func (o *DnsRecord) GetMinimum() int32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *DnsRecord) GetMinimumOk() (*int32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *DnsRecord) SetMinimum(v int32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *DnsRecord) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


