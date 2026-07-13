# AccountUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**ApiCredits** | Pointer to [**ApiCredits**](ApiCredits.md) |  | [optional] 
**ApiSubscription** | Pointer to [**ApiSubscription**](ApiSubscription.md) |  | [optional] 

## Methods

### NewAccountUsageResponse

`func NewAccountUsageResponse() *AccountUsageResponse`

NewAccountUsageResponse instantiates a new AccountUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUsageResponseWithDefaults

`func NewAccountUsageResponseWithDefaults() *AccountUsageResponse`

NewAccountUsageResponseWithDefaults instantiates a new AccountUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *AccountUsageResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AccountUsageResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AccountUsageResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AccountUsageResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiCredits

`func (o *AccountUsageResponse) GetApiCredits() ApiCredits`

GetApiCredits returns the ApiCredits field if non-nil, zero value otherwise.

### GetApiCreditsOk

`func (o *AccountUsageResponse) GetApiCreditsOk() (*ApiCredits, bool)`

GetApiCreditsOk returns a tuple with the ApiCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredits

`func (o *AccountUsageResponse) SetApiCredits(v ApiCredits)`

SetApiCredits sets ApiCredits field to given value.

### HasApiCredits

`func (o *AccountUsageResponse) HasApiCredits() bool`

HasApiCredits returns a boolean if a field has been set.

### GetApiSubscription

`func (o *AccountUsageResponse) GetApiSubscription() ApiSubscription`

GetApiSubscription returns the ApiSubscription field if non-nil, zero value otherwise.

### GetApiSubscriptionOk

`func (o *AccountUsageResponse) GetApiSubscriptionOk() (*ApiSubscription, bool)`

GetApiSubscriptionOk returns a tuple with the ApiSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSubscription

`func (o *AccountUsageResponse) SetApiSubscription(v ApiSubscription)`

SetApiSubscription sets ApiSubscription field to given value.

### HasApiSubscription

`func (o *AccountUsageResponse) HasApiSubscription() bool`

HasApiSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


