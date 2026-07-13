# ReputationIntelligence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IocType** | Pointer to **string** |  | [optional] 
**IocValue** | Pointer to **string** |  | [optional] 
**RelatedIocs** | Pointer to [**[]RelatedIoc**](RelatedIoc.md) |  | [optional] 
**FeedTags** | Pointer to **[]string** |  | [optional] 
**StixPattern** | Pointer to **string** |  | [optional] 
**RecommendedAction** | Pointer to **string** |  | [optional] 
**FirstSeen** | Pointer to **string** |  | [optional] 
**LastSeen** | Pointer to **string** |  | [optional] 

## Methods

### NewReputationIntelligence

`func NewReputationIntelligence() *ReputationIntelligence`

NewReputationIntelligence instantiates a new ReputationIntelligence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReputationIntelligenceWithDefaults

`func NewReputationIntelligenceWithDefaults() *ReputationIntelligence`

NewReputationIntelligenceWithDefaults instantiates a new ReputationIntelligence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIocType

`func (o *ReputationIntelligence) GetIocType() string`

GetIocType returns the IocType field if non-nil, zero value otherwise.

### GetIocTypeOk

`func (o *ReputationIntelligence) GetIocTypeOk() (*string, bool)`

GetIocTypeOk returns a tuple with the IocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIocType

`func (o *ReputationIntelligence) SetIocType(v string)`

SetIocType sets IocType field to given value.

### HasIocType

`func (o *ReputationIntelligence) HasIocType() bool`

HasIocType returns a boolean if a field has been set.

### GetIocValue

`func (o *ReputationIntelligence) GetIocValue() string`

GetIocValue returns the IocValue field if non-nil, zero value otherwise.

### GetIocValueOk

`func (o *ReputationIntelligence) GetIocValueOk() (*string, bool)`

GetIocValueOk returns a tuple with the IocValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIocValue

`func (o *ReputationIntelligence) SetIocValue(v string)`

SetIocValue sets IocValue field to given value.

### HasIocValue

`func (o *ReputationIntelligence) HasIocValue() bool`

HasIocValue returns a boolean if a field has been set.

### GetRelatedIocs

`func (o *ReputationIntelligence) GetRelatedIocs() []RelatedIoc`

GetRelatedIocs returns the RelatedIocs field if non-nil, zero value otherwise.

### GetRelatedIocsOk

`func (o *ReputationIntelligence) GetRelatedIocsOk() (*[]RelatedIoc, bool)`

GetRelatedIocsOk returns a tuple with the RelatedIocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIocs

`func (o *ReputationIntelligence) SetRelatedIocs(v []RelatedIoc)`

SetRelatedIocs sets RelatedIocs field to given value.

### HasRelatedIocs

`func (o *ReputationIntelligence) HasRelatedIocs() bool`

HasRelatedIocs returns a boolean if a field has been set.

### GetFeedTags

`func (o *ReputationIntelligence) GetFeedTags() []string`

GetFeedTags returns the FeedTags field if non-nil, zero value otherwise.

### GetFeedTagsOk

`func (o *ReputationIntelligence) GetFeedTagsOk() (*[]string, bool)`

GetFeedTagsOk returns a tuple with the FeedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedTags

`func (o *ReputationIntelligence) SetFeedTags(v []string)`

SetFeedTags sets FeedTags field to given value.

### HasFeedTags

`func (o *ReputationIntelligence) HasFeedTags() bool`

HasFeedTags returns a boolean if a field has been set.

### GetStixPattern

`func (o *ReputationIntelligence) GetStixPattern() string`

GetStixPattern returns the StixPattern field if non-nil, zero value otherwise.

### GetStixPatternOk

`func (o *ReputationIntelligence) GetStixPatternOk() (*string, bool)`

GetStixPatternOk returns a tuple with the StixPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStixPattern

`func (o *ReputationIntelligence) SetStixPattern(v string)`

SetStixPattern sets StixPattern field to given value.

### HasStixPattern

`func (o *ReputationIntelligence) HasStixPattern() bool`

HasStixPattern returns a boolean if a field has been set.

### GetRecommendedAction

`func (o *ReputationIntelligence) GetRecommendedAction() string`

GetRecommendedAction returns the RecommendedAction field if non-nil, zero value otherwise.

### GetRecommendedActionOk

`func (o *ReputationIntelligence) GetRecommendedActionOk() (*string, bool)`

GetRecommendedActionOk returns a tuple with the RecommendedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedAction

`func (o *ReputationIntelligence) SetRecommendedAction(v string)`

SetRecommendedAction sets RecommendedAction field to given value.

### HasRecommendedAction

`func (o *ReputationIntelligence) HasRecommendedAction() bool`

HasRecommendedAction returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ReputationIntelligence) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ReputationIntelligence) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ReputationIntelligence) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ReputationIntelligence) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *ReputationIntelligence) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ReputationIntelligence) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ReputationIntelligence) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ReputationIntelligence) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


