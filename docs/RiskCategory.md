# RiskCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verdict** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **float32** |  | [optional] 
**PrimaryThreat** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**ThreatTypes** | Pointer to **[]string** |  | [optional] 
**Sources** | Pointer to [**[]ThreatSource**](ThreatSource.md) |  | [optional] 
**PivotMatches** | Pointer to [**[]PivotMatch**](PivotMatch.md) |  | [optional] 

## Methods

### NewRiskCategory

`func NewRiskCategory() *RiskCategory`

NewRiskCategory instantiates a new RiskCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskCategoryWithDefaults

`func NewRiskCategoryWithDefaults() *RiskCategory`

NewRiskCategoryWithDefaults instantiates a new RiskCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdict

`func (o *RiskCategory) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *RiskCategory) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *RiskCategory) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *RiskCategory) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetConfidence

`func (o *RiskCategory) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RiskCategory) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RiskCategory) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RiskCategory) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetPrimaryThreat

`func (o *RiskCategory) GetPrimaryThreat() string`

GetPrimaryThreat returns the PrimaryThreat field if non-nil, zero value otherwise.

### GetPrimaryThreatOk

`func (o *RiskCategory) GetPrimaryThreatOk() (*string, bool)`

GetPrimaryThreatOk returns a tuple with the PrimaryThreat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryThreat

`func (o *RiskCategory) SetPrimaryThreat(v string)`

SetPrimaryThreat sets PrimaryThreat field to given value.

### HasPrimaryThreat

`func (o *RiskCategory) HasPrimaryThreat() bool`

HasPrimaryThreat returns a boolean if a field has been set.

### GetSeverity

`func (o *RiskCategory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RiskCategory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RiskCategory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RiskCategory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetThreatTypes

`func (o *RiskCategory) GetThreatTypes() []string`

GetThreatTypes returns the ThreatTypes field if non-nil, zero value otherwise.

### GetThreatTypesOk

`func (o *RiskCategory) GetThreatTypesOk() (*[]string, bool)`

GetThreatTypesOk returns a tuple with the ThreatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatTypes

`func (o *RiskCategory) SetThreatTypes(v []string)`

SetThreatTypes sets ThreatTypes field to given value.

### HasThreatTypes

`func (o *RiskCategory) HasThreatTypes() bool`

HasThreatTypes returns a boolean if a field has been set.

### GetSources

`func (o *RiskCategory) GetSources() []ThreatSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RiskCategory) GetSourcesOk() (*[]ThreatSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RiskCategory) SetSources(v []ThreatSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *RiskCategory) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetPivotMatches

`func (o *RiskCategory) GetPivotMatches() []PivotMatch`

GetPivotMatches returns the PivotMatches field if non-nil, zero value otherwise.

### GetPivotMatchesOk

`func (o *RiskCategory) GetPivotMatchesOk() (*[]PivotMatch, bool)`

GetPivotMatchesOk returns a tuple with the PivotMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotMatches

`func (o *RiskCategory) SetPivotMatches(v []PivotMatch)`

SetPivotMatches sets PivotMatches field to given value.

### HasPivotMatches

`func (o *RiskCategory) HasPivotMatches() bool`

HasPivotMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


