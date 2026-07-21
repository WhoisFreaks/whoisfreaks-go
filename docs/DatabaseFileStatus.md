# DatabaseFileStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Newly** | Pointer to [**NewlyStatus**](NewlyStatus.md) |  | [optional] 
**Expired** | Pointer to [**DateRangeStatus**](DateRangeStatus.md) |  | [optional] 
**CleanedExpired** | Pointer to [**DateRangeStatus**](DateRangeStatus.md) |  | [optional] 
**Dropped** | Pointer to [**DateRangeStatus**](DateRangeStatus.md) |  | [optional] 
**DroppedWithBacklinks** | Pointer to [**DateRangeStatus**](DateRangeStatus.md) |  | [optional] 
**ThreatFeed** | Pointer to [**ThreatFeedStatus**](ThreatFeedStatus.md) |  | [optional] 
**DatabaseUpdates** | Pointer to [**DatabaseUpdates**](DatabaseUpdates.md) |  | [optional] 

## Methods

### NewDatabaseFileStatus

`func NewDatabaseFileStatus() *DatabaseFileStatus`

NewDatabaseFileStatus instantiates a new DatabaseFileStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseFileStatusWithDefaults

`func NewDatabaseFileStatusWithDefaults() *DatabaseFileStatus`

NewDatabaseFileStatusWithDefaults instantiates a new DatabaseFileStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewly

`func (o *DatabaseFileStatus) GetNewly() NewlyStatus`

GetNewly returns the Newly field if non-nil, zero value otherwise.

### GetNewlyOk

`func (o *DatabaseFileStatus) GetNewlyOk() (*NewlyStatus, bool)`

GetNewlyOk returns a tuple with the Newly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewly

`func (o *DatabaseFileStatus) SetNewly(v NewlyStatus)`

SetNewly sets Newly field to given value.

### HasNewly

`func (o *DatabaseFileStatus) HasNewly() bool`

HasNewly returns a boolean if a field has been set.

### GetExpired

`func (o *DatabaseFileStatus) GetExpired() DateRangeStatus`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *DatabaseFileStatus) GetExpiredOk() (*DateRangeStatus, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *DatabaseFileStatus) SetExpired(v DateRangeStatus)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *DatabaseFileStatus) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetCleanedExpired

`func (o *DatabaseFileStatus) GetCleanedExpired() DateRangeStatus`

GetCleanedExpired returns the CleanedExpired field if non-nil, zero value otherwise.

### GetCleanedExpiredOk

`func (o *DatabaseFileStatus) GetCleanedExpiredOk() (*DateRangeStatus, bool)`

GetCleanedExpiredOk returns a tuple with the CleanedExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanedExpired

`func (o *DatabaseFileStatus) SetCleanedExpired(v DateRangeStatus)`

SetCleanedExpired sets CleanedExpired field to given value.

### HasCleanedExpired

`func (o *DatabaseFileStatus) HasCleanedExpired() bool`

HasCleanedExpired returns a boolean if a field has been set.

### GetDropped

`func (o *DatabaseFileStatus) GetDropped() DateRangeStatus`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *DatabaseFileStatus) GetDroppedOk() (*DateRangeStatus, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *DatabaseFileStatus) SetDropped(v DateRangeStatus)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *DatabaseFileStatus) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetDroppedWithBacklinks

`func (o *DatabaseFileStatus) GetDroppedWithBacklinks() DateRangeStatus`

GetDroppedWithBacklinks returns the DroppedWithBacklinks field if non-nil, zero value otherwise.

### GetDroppedWithBacklinksOk

`func (o *DatabaseFileStatus) GetDroppedWithBacklinksOk() (*DateRangeStatus, bool)`

GetDroppedWithBacklinksOk returns a tuple with the DroppedWithBacklinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedWithBacklinks

`func (o *DatabaseFileStatus) SetDroppedWithBacklinks(v DateRangeStatus)`

SetDroppedWithBacklinks sets DroppedWithBacklinks field to given value.

### HasDroppedWithBacklinks

`func (o *DatabaseFileStatus) HasDroppedWithBacklinks() bool`

HasDroppedWithBacklinks returns a boolean if a field has been set.

### GetThreatFeed

`func (o *DatabaseFileStatus) GetThreatFeed() ThreatFeedStatus`

GetThreatFeed returns the ThreatFeed field if non-nil, zero value otherwise.

### GetThreatFeedOk

`func (o *DatabaseFileStatus) GetThreatFeedOk() (*ThreatFeedStatus, bool)`

GetThreatFeedOk returns a tuple with the ThreatFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatFeed

`func (o *DatabaseFileStatus) SetThreatFeed(v ThreatFeedStatus)`

SetThreatFeed sets ThreatFeed field to given value.

### HasThreatFeed

`func (o *DatabaseFileStatus) HasThreatFeed() bool`

HasThreatFeed returns a boolean if a field has been set.

### GetDatabaseUpdates

`func (o *DatabaseFileStatus) GetDatabaseUpdates() DatabaseUpdates`

GetDatabaseUpdates returns the DatabaseUpdates field if non-nil, zero value otherwise.

### GetDatabaseUpdatesOk

`func (o *DatabaseFileStatus) GetDatabaseUpdatesOk() (*DatabaseUpdates, bool)`

GetDatabaseUpdatesOk returns a tuple with the DatabaseUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUpdates

`func (o *DatabaseFileStatus) SetDatabaseUpdates(v DatabaseUpdates)`

SetDatabaseUpdates sets DatabaseUpdates field to given value.

### HasDatabaseUpdates

`func (o *DatabaseFileStatus) HasDatabaseUpdates() bool`

HasDatabaseUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


