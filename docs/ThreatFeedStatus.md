# ThreatFeedStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phishing** | Pointer to [**DateRangeStatus**](DateRangeStatus.md) |  | [optional] 
**Malware** | Pointer to [**DateRangeStatus**](DateRangeStatus.md) |  | [optional] 
**Spam** | Pointer to [**DateRangeStatus**](DateRangeStatus.md) |  | [optional] 

## Methods

### NewThreatFeedStatus

`func NewThreatFeedStatus() *ThreatFeedStatus`

NewThreatFeedStatus instantiates a new ThreatFeedStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatFeedStatusWithDefaults

`func NewThreatFeedStatusWithDefaults() *ThreatFeedStatus`

NewThreatFeedStatusWithDefaults instantiates a new ThreatFeedStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhishing

`func (o *ThreatFeedStatus) GetPhishing() DateRangeStatus`

GetPhishing returns the Phishing field if non-nil, zero value otherwise.

### GetPhishingOk

`func (o *ThreatFeedStatus) GetPhishingOk() (*DateRangeStatus, bool)`

GetPhishingOk returns a tuple with the Phishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhishing

`func (o *ThreatFeedStatus) SetPhishing(v DateRangeStatus)`

SetPhishing sets Phishing field to given value.

### HasPhishing

`func (o *ThreatFeedStatus) HasPhishing() bool`

HasPhishing returns a boolean if a field has been set.

### GetMalware

`func (o *ThreatFeedStatus) GetMalware() DateRangeStatus`

GetMalware returns the Malware field if non-nil, zero value otherwise.

### GetMalwareOk

`func (o *ThreatFeedStatus) GetMalwareOk() (*DateRangeStatus, bool)`

GetMalwareOk returns a tuple with the Malware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalware

`func (o *ThreatFeedStatus) SetMalware(v DateRangeStatus)`

SetMalware sets Malware field to given value.

### HasMalware

`func (o *ThreatFeedStatus) HasMalware() bool`

HasMalware returns a boolean if a field has been set.

### GetSpam

`func (o *ThreatFeedStatus) GetSpam() DateRangeStatus`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *ThreatFeedStatus) GetSpamOk() (*DateRangeStatus, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *ThreatFeedStatus) SetSpam(v DateRangeStatus)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *ThreatFeedStatus) HasSpam() bool`

HasSpam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


