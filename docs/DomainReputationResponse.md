# DomainReputationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to [**DomainReputationInput**](DomainReputationInput.md) |  | [optional] 
**AssessedAt** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ProcessingTimeMs** | Pointer to **int32** |  | [optional] 
**RiskCategory** | Pointer to [**RiskCategory**](RiskCategory.md) |  | [optional] 
**DgaScore** | Pointer to [**DgaScore**](DgaScore.md) |  | [optional] 
**TrustSignals** | Pointer to [**TrustSignals**](TrustSignals.md) |  | [optional] 
**Intelligence** | Pointer to [**ReputationIntelligence**](ReputationIntelligence.md) |  | [optional] 
**EvidenceSummary** | Pointer to [**EvidenceSummary**](EvidenceSummary.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDomainReputationResponse

`func NewDomainReputationResponse() *DomainReputationResponse`

NewDomainReputationResponse instantiates a new DomainReputationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainReputationResponseWithDefaults

`func NewDomainReputationResponseWithDefaults() *DomainReputationResponse`

NewDomainReputationResponseWithDefaults instantiates a new DomainReputationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *DomainReputationResponse) GetInput() DomainReputationInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DomainReputationResponse) GetInputOk() (*DomainReputationInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DomainReputationResponse) SetInput(v DomainReputationInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *DomainReputationResponse) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetAssessedAt

`func (o *DomainReputationResponse) GetAssessedAt() string`

GetAssessedAt returns the AssessedAt field if non-nil, zero value otherwise.

### GetAssessedAtOk

`func (o *DomainReputationResponse) GetAssessedAtOk() (*string, bool)`

GetAssessedAtOk returns a tuple with the AssessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessedAt

`func (o *DomainReputationResponse) SetAssessedAt(v string)`

SetAssessedAt sets AssessedAt field to given value.

### HasAssessedAt

`func (o *DomainReputationResponse) HasAssessedAt() bool`

HasAssessedAt returns a boolean if a field has been set.

### GetVersion

`func (o *DomainReputationResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DomainReputationResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DomainReputationResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DomainReputationResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *DomainReputationResponse) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *DomainReputationResponse) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *DomainReputationResponse) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *DomainReputationResponse) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.

### GetRiskCategory

`func (o *DomainReputationResponse) GetRiskCategory() RiskCategory`

GetRiskCategory returns the RiskCategory field if non-nil, zero value otherwise.

### GetRiskCategoryOk

`func (o *DomainReputationResponse) GetRiskCategoryOk() (*RiskCategory, bool)`

GetRiskCategoryOk returns a tuple with the RiskCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskCategory

`func (o *DomainReputationResponse) SetRiskCategory(v RiskCategory)`

SetRiskCategory sets RiskCategory field to given value.

### HasRiskCategory

`func (o *DomainReputationResponse) HasRiskCategory() bool`

HasRiskCategory returns a boolean if a field has been set.

### GetDgaScore

`func (o *DomainReputationResponse) GetDgaScore() DgaScore`

GetDgaScore returns the DgaScore field if non-nil, zero value otherwise.

### GetDgaScoreOk

`func (o *DomainReputationResponse) GetDgaScoreOk() (*DgaScore, bool)`

GetDgaScoreOk returns a tuple with the DgaScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDgaScore

`func (o *DomainReputationResponse) SetDgaScore(v DgaScore)`

SetDgaScore sets DgaScore field to given value.

### HasDgaScore

`func (o *DomainReputationResponse) HasDgaScore() bool`

HasDgaScore returns a boolean if a field has been set.

### GetTrustSignals

`func (o *DomainReputationResponse) GetTrustSignals() TrustSignals`

GetTrustSignals returns the TrustSignals field if non-nil, zero value otherwise.

### GetTrustSignalsOk

`func (o *DomainReputationResponse) GetTrustSignalsOk() (*TrustSignals, bool)`

GetTrustSignalsOk returns a tuple with the TrustSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustSignals

`func (o *DomainReputationResponse) SetTrustSignals(v TrustSignals)`

SetTrustSignals sets TrustSignals field to given value.

### HasTrustSignals

`func (o *DomainReputationResponse) HasTrustSignals() bool`

HasTrustSignals returns a boolean if a field has been set.

### GetIntelligence

`func (o *DomainReputationResponse) GetIntelligence() ReputationIntelligence`

GetIntelligence returns the Intelligence field if non-nil, zero value otherwise.

### GetIntelligenceOk

`func (o *DomainReputationResponse) GetIntelligenceOk() (*ReputationIntelligence, bool)`

GetIntelligenceOk returns a tuple with the Intelligence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelligence

`func (o *DomainReputationResponse) SetIntelligence(v ReputationIntelligence)`

SetIntelligence sets Intelligence field to given value.

### HasIntelligence

`func (o *DomainReputationResponse) HasIntelligence() bool`

HasIntelligence returns a boolean if a field has been set.

### GetEvidenceSummary

`func (o *DomainReputationResponse) GetEvidenceSummary() EvidenceSummary`

GetEvidenceSummary returns the EvidenceSummary field if non-nil, zero value otherwise.

### GetEvidenceSummaryOk

`func (o *DomainReputationResponse) GetEvidenceSummaryOk() (*EvidenceSummary, bool)`

GetEvidenceSummaryOk returns a tuple with the EvidenceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSummary

`func (o *DomainReputationResponse) SetEvidenceSummary(v EvidenceSummary)`

SetEvidenceSummary sets EvidenceSummary field to given value.

### HasEvidenceSummary

`func (o *DomainReputationResponse) HasEvidenceSummary() bool`

HasEvidenceSummary returns a boolean if a field has been set.

### GetErrors

`func (o *DomainReputationResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DomainReputationResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DomainReputationResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DomainReputationResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


