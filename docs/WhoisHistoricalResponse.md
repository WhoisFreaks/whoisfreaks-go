# WhoisHistoricalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Whois** | Pointer to **string** |  | [optional] 
**TotalRecords** | Pointer to **string** |  | [optional] 
**WhoisDomainsHistorical** | Pointer to [**[]WhoisHistoricalItem**](WhoisHistoricalItem.md) |  | [optional] 

## Methods

### NewWhoisHistoricalResponse

`func NewWhoisHistoricalResponse() *WhoisHistoricalResponse`

NewWhoisHistoricalResponse instantiates a new WhoisHistoricalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoisHistoricalResponseWithDefaults

`func NewWhoisHistoricalResponseWithDefaults() *WhoisHistoricalResponse`

NewWhoisHistoricalResponseWithDefaults instantiates a new WhoisHistoricalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WhoisHistoricalResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhoisHistoricalResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhoisHistoricalResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhoisHistoricalResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWhois

`func (o *WhoisHistoricalResponse) GetWhois() string`

GetWhois returns the Whois field if non-nil, zero value otherwise.

### GetWhoisOk

`func (o *WhoisHistoricalResponse) GetWhoisOk() (*string, bool)`

GetWhoisOk returns a tuple with the Whois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhois

`func (o *WhoisHistoricalResponse) SetWhois(v string)`

SetWhois sets Whois field to given value.

### HasWhois

`func (o *WhoisHistoricalResponse) HasWhois() bool`

HasWhois returns a boolean if a field has been set.

### GetTotalRecords

`func (o *WhoisHistoricalResponse) GetTotalRecords() string`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *WhoisHistoricalResponse) GetTotalRecordsOk() (*string, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *WhoisHistoricalResponse) SetTotalRecords(v string)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *WhoisHistoricalResponse) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetWhoisDomainsHistorical

`func (o *WhoisHistoricalResponse) GetWhoisDomainsHistorical() []WhoisHistoricalItem`

GetWhoisDomainsHistorical returns the WhoisDomainsHistorical field if non-nil, zero value otherwise.

### GetWhoisDomainsHistoricalOk

`func (o *WhoisHistoricalResponse) GetWhoisDomainsHistoricalOk() (*[]WhoisHistoricalItem, bool)`

GetWhoisDomainsHistoricalOk returns a tuple with the WhoisDomainsHistorical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisDomainsHistorical

`func (o *WhoisHistoricalResponse) SetWhoisDomainsHistorical(v []WhoisHistoricalItem)`

SetWhoisDomainsHistorical sets WhoisDomainsHistorical field to given value.

### HasWhoisDomainsHistorical

`func (o *WhoisHistoricalResponse) HasWhoisDomainsHistorical() bool`

HasWhoisDomainsHistorical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


