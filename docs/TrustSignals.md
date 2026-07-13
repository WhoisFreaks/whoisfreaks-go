# TrustSignals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrustScore** | Pointer to **int32** |  | [optional] 
**TrustBand** | Pointer to **string** |  | [optional] 
**Signals** | Pointer to [**ReputationSignals**](ReputationSignals.md) |  | [optional] 
**Indicators** | Pointer to [**ReputationIndicators**](ReputationIndicators.md) |  | [optional] 

## Methods

### NewTrustSignals

`func NewTrustSignals() *TrustSignals`

NewTrustSignals instantiates a new TrustSignals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustSignalsWithDefaults

`func NewTrustSignalsWithDefaults() *TrustSignals`

NewTrustSignalsWithDefaults instantiates a new TrustSignals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrustScore

`func (o *TrustSignals) GetTrustScore() int32`

GetTrustScore returns the TrustScore field if non-nil, zero value otherwise.

### GetTrustScoreOk

`func (o *TrustSignals) GetTrustScoreOk() (*int32, bool)`

GetTrustScoreOk returns a tuple with the TrustScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustScore

`func (o *TrustSignals) SetTrustScore(v int32)`

SetTrustScore sets TrustScore field to given value.

### HasTrustScore

`func (o *TrustSignals) HasTrustScore() bool`

HasTrustScore returns a boolean if a field has been set.

### GetTrustBand

`func (o *TrustSignals) GetTrustBand() string`

GetTrustBand returns the TrustBand field if non-nil, zero value otherwise.

### GetTrustBandOk

`func (o *TrustSignals) GetTrustBandOk() (*string, bool)`

GetTrustBandOk returns a tuple with the TrustBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustBand

`func (o *TrustSignals) SetTrustBand(v string)`

SetTrustBand sets TrustBand field to given value.

### HasTrustBand

`func (o *TrustSignals) HasTrustBand() bool`

HasTrustBand returns a boolean if a field has been set.

### GetSignals

`func (o *TrustSignals) GetSignals() ReputationSignals`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *TrustSignals) GetSignalsOk() (*ReputationSignals, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *TrustSignals) SetSignals(v ReputationSignals)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *TrustSignals) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### GetIndicators

`func (o *TrustSignals) GetIndicators() ReputationIndicators`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *TrustSignals) GetIndicatorsOk() (*ReputationIndicators, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *TrustSignals) SetIndicators(v ReputationIndicators)`

SetIndicators sets Indicators field to given value.

### HasIndicators

`func (o *TrustSignals) HasIndicators() bool`

HasIndicators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


