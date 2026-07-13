# DgaFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainLength** | Pointer to **int32** |  | [optional] 
**VowelConsonantRatio** | Pointer to **float32** |  | [optional] 
**NgramPerplexity** | Pointer to **float32** |  | [optional] 
**ShannonEntropy** | Pointer to **float32** |  | [optional] 
**DigitLetterRatio** | Pointer to **float32** |  | [optional] 
**ConsonantStreakMax** | Pointer to **int32** |  | [optional] 
**TldInKnownDgaSet** | Pointer to **bool** |  | [optional] 

## Methods

### NewDgaFeatures

`func NewDgaFeatures() *DgaFeatures`

NewDgaFeatures instantiates a new DgaFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDgaFeaturesWithDefaults

`func NewDgaFeaturesWithDefaults() *DgaFeatures`

NewDgaFeaturesWithDefaults instantiates a new DgaFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainLength

`func (o *DgaFeatures) GetDomainLength() int32`

GetDomainLength returns the DomainLength field if non-nil, zero value otherwise.

### GetDomainLengthOk

`func (o *DgaFeatures) GetDomainLengthOk() (*int32, bool)`

GetDomainLengthOk returns a tuple with the DomainLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainLength

`func (o *DgaFeatures) SetDomainLength(v int32)`

SetDomainLength sets DomainLength field to given value.

### HasDomainLength

`func (o *DgaFeatures) HasDomainLength() bool`

HasDomainLength returns a boolean if a field has been set.

### GetVowelConsonantRatio

`func (o *DgaFeatures) GetVowelConsonantRatio() float32`

GetVowelConsonantRatio returns the VowelConsonantRatio field if non-nil, zero value otherwise.

### GetVowelConsonantRatioOk

`func (o *DgaFeatures) GetVowelConsonantRatioOk() (*float32, bool)`

GetVowelConsonantRatioOk returns a tuple with the VowelConsonantRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVowelConsonantRatio

`func (o *DgaFeatures) SetVowelConsonantRatio(v float32)`

SetVowelConsonantRatio sets VowelConsonantRatio field to given value.

### HasVowelConsonantRatio

`func (o *DgaFeatures) HasVowelConsonantRatio() bool`

HasVowelConsonantRatio returns a boolean if a field has been set.

### GetNgramPerplexity

`func (o *DgaFeatures) GetNgramPerplexity() float32`

GetNgramPerplexity returns the NgramPerplexity field if non-nil, zero value otherwise.

### GetNgramPerplexityOk

`func (o *DgaFeatures) GetNgramPerplexityOk() (*float32, bool)`

GetNgramPerplexityOk returns a tuple with the NgramPerplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgramPerplexity

`func (o *DgaFeatures) SetNgramPerplexity(v float32)`

SetNgramPerplexity sets NgramPerplexity field to given value.

### HasNgramPerplexity

`func (o *DgaFeatures) HasNgramPerplexity() bool`

HasNgramPerplexity returns a boolean if a field has been set.

### GetShannonEntropy

`func (o *DgaFeatures) GetShannonEntropy() float32`

GetShannonEntropy returns the ShannonEntropy field if non-nil, zero value otherwise.

### GetShannonEntropyOk

`func (o *DgaFeatures) GetShannonEntropyOk() (*float32, bool)`

GetShannonEntropyOk returns a tuple with the ShannonEntropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShannonEntropy

`func (o *DgaFeatures) SetShannonEntropy(v float32)`

SetShannonEntropy sets ShannonEntropy field to given value.

### HasShannonEntropy

`func (o *DgaFeatures) HasShannonEntropy() bool`

HasShannonEntropy returns a boolean if a field has been set.

### GetDigitLetterRatio

`func (o *DgaFeatures) GetDigitLetterRatio() float32`

GetDigitLetterRatio returns the DigitLetterRatio field if non-nil, zero value otherwise.

### GetDigitLetterRatioOk

`func (o *DgaFeatures) GetDigitLetterRatioOk() (*float32, bool)`

GetDigitLetterRatioOk returns a tuple with the DigitLetterRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitLetterRatio

`func (o *DgaFeatures) SetDigitLetterRatio(v float32)`

SetDigitLetterRatio sets DigitLetterRatio field to given value.

### HasDigitLetterRatio

`func (o *DgaFeatures) HasDigitLetterRatio() bool`

HasDigitLetterRatio returns a boolean if a field has been set.

### GetConsonantStreakMax

`func (o *DgaFeatures) GetConsonantStreakMax() int32`

GetConsonantStreakMax returns the ConsonantStreakMax field if non-nil, zero value otherwise.

### GetConsonantStreakMaxOk

`func (o *DgaFeatures) GetConsonantStreakMaxOk() (*int32, bool)`

GetConsonantStreakMaxOk returns a tuple with the ConsonantStreakMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsonantStreakMax

`func (o *DgaFeatures) SetConsonantStreakMax(v int32)`

SetConsonantStreakMax sets ConsonantStreakMax field to given value.

### HasConsonantStreakMax

`func (o *DgaFeatures) HasConsonantStreakMax() bool`

HasConsonantStreakMax returns a boolean if a field has been set.

### GetTldInKnownDgaSet

`func (o *DgaFeatures) GetTldInKnownDgaSet() bool`

GetTldInKnownDgaSet returns the TldInKnownDgaSet field if non-nil, zero value otherwise.

### GetTldInKnownDgaSetOk

`func (o *DgaFeatures) GetTldInKnownDgaSetOk() (*bool, bool)`

GetTldInKnownDgaSetOk returns a tuple with the TldInKnownDgaSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTldInKnownDgaSet

`func (o *DgaFeatures) SetTldInKnownDgaSet(v bool)`

SetTldInKnownDgaSet sets TldInKnownDgaSet field to given value.

### HasTldInKnownDgaSet

`func (o *DgaFeatures) HasTldInKnownDgaSet() bool`

HasTldInKnownDgaSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


