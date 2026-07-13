# DgaScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** |  | [optional] 
**IsDga** | Pointer to **bool** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**DgaFeatures**](DgaFeatures.md) |  | [optional] 
**Interpretation** | Pointer to **string** |  | [optional] 

## Methods

### NewDgaScore

`func NewDgaScore() *DgaScore`

NewDgaScore instantiates a new DgaScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDgaScoreWithDefaults

`func NewDgaScoreWithDefaults() *DgaScore`

NewDgaScoreWithDefaults instantiates a new DgaScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *DgaScore) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DgaScore) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DgaScore) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *DgaScore) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetIsDga

`func (o *DgaScore) GetIsDga() bool`

GetIsDga returns the IsDga field if non-nil, zero value otherwise.

### GetIsDgaOk

`func (o *DgaScore) GetIsDgaOk() (*bool, bool)`

GetIsDgaOk returns a tuple with the IsDga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDga

`func (o *DgaScore) SetIsDga(v bool)`

SetIsDga sets IsDga field to given value.

### HasIsDga

`func (o *DgaScore) HasIsDga() bool`

HasIsDga returns a boolean if a field has been set.

### GetModel

`func (o *DgaScore) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DgaScore) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DgaScore) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DgaScore) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetFeatures

`func (o *DgaScore) GetFeatures() DgaFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DgaScore) GetFeaturesOk() (*DgaFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DgaScore) SetFeatures(v DgaFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DgaScore) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetInterpretation

`func (o *DgaScore) GetInterpretation() string`

GetInterpretation returns the Interpretation field if non-nil, zero value otherwise.

### GetInterpretationOk

`func (o *DgaScore) GetInterpretationOk() (*string, bool)`

GetInterpretationOk returns a tuple with the Interpretation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpretation

`func (o *DgaScore) SetInterpretation(v string)`

SetInterpretation sets Interpretation field to given value.

### HasInterpretation

`func (o *DgaScore) HasInterpretation() bool`

HasInterpretation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


