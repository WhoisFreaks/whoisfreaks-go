# IpSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreatScore** | Pointer to **int32** |  | [optional] 
**IsTor** | Pointer to **bool** |  | [optional] 
**IsProxy** | Pointer to **bool** |  | [optional] 
**ProxyProviderNames** | Pointer to **[]string** |  | [optional] 
**ProxyConfidenceScore** | Pointer to **int32** |  | [optional] 
**ProxyLastSeen** | Pointer to **string** |  | [optional] 
**IsResidentialProxy** | Pointer to **bool** |  | [optional] 
**IsVpn** | Pointer to **bool** |  | [optional] 
**VpnProviderNames** | Pointer to **[]string** |  | [optional] 
**VpnConfidenceScore** | Pointer to **int32** |  | [optional] 
**VpnLastSeen** | Pointer to **string** |  | [optional] 
**IsRelay** | Pointer to **bool** |  | [optional] 
**RelayProviderName** | Pointer to **string** |  | [optional] 
**IsAnonymous** | Pointer to **bool** |  | [optional] 
**IsKnownAttacker** | Pointer to **bool** |  | [optional] 
**IsBot** | Pointer to **bool** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**IsCloudProvider** | Pointer to **bool** |  | [optional] 
**CloudProviderName** | Pointer to **string** |  | [optional] 

## Methods

### NewIpSecurity

`func NewIpSecurity() *IpSecurity`

NewIpSecurity instantiates a new IpSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSecurityWithDefaults

`func NewIpSecurityWithDefaults() *IpSecurity`

NewIpSecurityWithDefaults instantiates a new IpSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreatScore

`func (o *IpSecurity) GetThreatScore() int32`

GetThreatScore returns the ThreatScore field if non-nil, zero value otherwise.

### GetThreatScoreOk

`func (o *IpSecurity) GetThreatScoreOk() (*int32, bool)`

GetThreatScoreOk returns a tuple with the ThreatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatScore

`func (o *IpSecurity) SetThreatScore(v int32)`

SetThreatScore sets ThreatScore field to given value.

### HasThreatScore

`func (o *IpSecurity) HasThreatScore() bool`

HasThreatScore returns a boolean if a field has been set.

### GetIsTor

`func (o *IpSecurity) GetIsTor() bool`

GetIsTor returns the IsTor field if non-nil, zero value otherwise.

### GetIsTorOk

`func (o *IpSecurity) GetIsTorOk() (*bool, bool)`

GetIsTorOk returns a tuple with the IsTor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTor

`func (o *IpSecurity) SetIsTor(v bool)`

SetIsTor sets IsTor field to given value.

### HasIsTor

`func (o *IpSecurity) HasIsTor() bool`

HasIsTor returns a boolean if a field has been set.

### GetIsProxy

`func (o *IpSecurity) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *IpSecurity) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *IpSecurity) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *IpSecurity) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### GetProxyProviderNames

`func (o *IpSecurity) GetProxyProviderNames() []string`

GetProxyProviderNames returns the ProxyProviderNames field if non-nil, zero value otherwise.

### GetProxyProviderNamesOk

`func (o *IpSecurity) GetProxyProviderNamesOk() (*[]string, bool)`

GetProxyProviderNamesOk returns a tuple with the ProxyProviderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProviderNames

`func (o *IpSecurity) SetProxyProviderNames(v []string)`

SetProxyProviderNames sets ProxyProviderNames field to given value.

### HasProxyProviderNames

`func (o *IpSecurity) HasProxyProviderNames() bool`

HasProxyProviderNames returns a boolean if a field has been set.

### GetProxyConfidenceScore

`func (o *IpSecurity) GetProxyConfidenceScore() int32`

GetProxyConfidenceScore returns the ProxyConfidenceScore field if non-nil, zero value otherwise.

### GetProxyConfidenceScoreOk

`func (o *IpSecurity) GetProxyConfidenceScoreOk() (*int32, bool)`

GetProxyConfidenceScoreOk returns a tuple with the ProxyConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfidenceScore

`func (o *IpSecurity) SetProxyConfidenceScore(v int32)`

SetProxyConfidenceScore sets ProxyConfidenceScore field to given value.

### HasProxyConfidenceScore

`func (o *IpSecurity) HasProxyConfidenceScore() bool`

HasProxyConfidenceScore returns a boolean if a field has been set.

### GetProxyLastSeen

`func (o *IpSecurity) GetProxyLastSeen() string`

GetProxyLastSeen returns the ProxyLastSeen field if non-nil, zero value otherwise.

### GetProxyLastSeenOk

`func (o *IpSecurity) GetProxyLastSeenOk() (*string, bool)`

GetProxyLastSeenOk returns a tuple with the ProxyLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyLastSeen

`func (o *IpSecurity) SetProxyLastSeen(v string)`

SetProxyLastSeen sets ProxyLastSeen field to given value.

### HasProxyLastSeen

`func (o *IpSecurity) HasProxyLastSeen() bool`

HasProxyLastSeen returns a boolean if a field has been set.

### GetIsResidentialProxy

`func (o *IpSecurity) GetIsResidentialProxy() bool`

GetIsResidentialProxy returns the IsResidentialProxy field if non-nil, zero value otherwise.

### GetIsResidentialProxyOk

`func (o *IpSecurity) GetIsResidentialProxyOk() (*bool, bool)`

GetIsResidentialProxyOk returns a tuple with the IsResidentialProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResidentialProxy

`func (o *IpSecurity) SetIsResidentialProxy(v bool)`

SetIsResidentialProxy sets IsResidentialProxy field to given value.

### HasIsResidentialProxy

`func (o *IpSecurity) HasIsResidentialProxy() bool`

HasIsResidentialProxy returns a boolean if a field has been set.

### GetIsVpn

`func (o *IpSecurity) GetIsVpn() bool`

GetIsVpn returns the IsVpn field if non-nil, zero value otherwise.

### GetIsVpnOk

`func (o *IpSecurity) GetIsVpnOk() (*bool, bool)`

GetIsVpnOk returns a tuple with the IsVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpn

`func (o *IpSecurity) SetIsVpn(v bool)`

SetIsVpn sets IsVpn field to given value.

### HasIsVpn

`func (o *IpSecurity) HasIsVpn() bool`

HasIsVpn returns a boolean if a field has been set.

### GetVpnProviderNames

`func (o *IpSecurity) GetVpnProviderNames() []string`

GetVpnProviderNames returns the VpnProviderNames field if non-nil, zero value otherwise.

### GetVpnProviderNamesOk

`func (o *IpSecurity) GetVpnProviderNamesOk() (*[]string, bool)`

GetVpnProviderNamesOk returns a tuple with the VpnProviderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProviderNames

`func (o *IpSecurity) SetVpnProviderNames(v []string)`

SetVpnProviderNames sets VpnProviderNames field to given value.

### HasVpnProviderNames

`func (o *IpSecurity) HasVpnProviderNames() bool`

HasVpnProviderNames returns a boolean if a field has been set.

### GetVpnConfidenceScore

`func (o *IpSecurity) GetVpnConfidenceScore() int32`

GetVpnConfidenceScore returns the VpnConfidenceScore field if non-nil, zero value otherwise.

### GetVpnConfidenceScoreOk

`func (o *IpSecurity) GetVpnConfidenceScoreOk() (*int32, bool)`

GetVpnConfidenceScoreOk returns a tuple with the VpnConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConfidenceScore

`func (o *IpSecurity) SetVpnConfidenceScore(v int32)`

SetVpnConfidenceScore sets VpnConfidenceScore field to given value.

### HasVpnConfidenceScore

`func (o *IpSecurity) HasVpnConfidenceScore() bool`

HasVpnConfidenceScore returns a boolean if a field has been set.

### GetVpnLastSeen

`func (o *IpSecurity) GetVpnLastSeen() string`

GetVpnLastSeen returns the VpnLastSeen field if non-nil, zero value otherwise.

### GetVpnLastSeenOk

`func (o *IpSecurity) GetVpnLastSeenOk() (*string, bool)`

GetVpnLastSeenOk returns a tuple with the VpnLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnLastSeen

`func (o *IpSecurity) SetVpnLastSeen(v string)`

SetVpnLastSeen sets VpnLastSeen field to given value.

### HasVpnLastSeen

`func (o *IpSecurity) HasVpnLastSeen() bool`

HasVpnLastSeen returns a boolean if a field has been set.

### GetIsRelay

`func (o *IpSecurity) GetIsRelay() bool`

GetIsRelay returns the IsRelay field if non-nil, zero value otherwise.

### GetIsRelayOk

`func (o *IpSecurity) GetIsRelayOk() (*bool, bool)`

GetIsRelayOk returns a tuple with the IsRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRelay

`func (o *IpSecurity) SetIsRelay(v bool)`

SetIsRelay sets IsRelay field to given value.

### HasIsRelay

`func (o *IpSecurity) HasIsRelay() bool`

HasIsRelay returns a boolean if a field has been set.

### GetRelayProviderName

`func (o *IpSecurity) GetRelayProviderName() string`

GetRelayProviderName returns the RelayProviderName field if non-nil, zero value otherwise.

### GetRelayProviderNameOk

`func (o *IpSecurity) GetRelayProviderNameOk() (*string, bool)`

GetRelayProviderNameOk returns a tuple with the RelayProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayProviderName

`func (o *IpSecurity) SetRelayProviderName(v string)`

SetRelayProviderName sets RelayProviderName field to given value.

### HasRelayProviderName

`func (o *IpSecurity) HasRelayProviderName() bool`

HasRelayProviderName returns a boolean if a field has been set.

### GetIsAnonymous

`func (o *IpSecurity) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *IpSecurity) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *IpSecurity) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.

### HasIsAnonymous

`func (o *IpSecurity) HasIsAnonymous() bool`

HasIsAnonymous returns a boolean if a field has been set.

### GetIsKnownAttacker

`func (o *IpSecurity) GetIsKnownAttacker() bool`

GetIsKnownAttacker returns the IsKnownAttacker field if non-nil, zero value otherwise.

### GetIsKnownAttackerOk

`func (o *IpSecurity) GetIsKnownAttackerOk() (*bool, bool)`

GetIsKnownAttackerOk returns a tuple with the IsKnownAttacker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKnownAttacker

`func (o *IpSecurity) SetIsKnownAttacker(v bool)`

SetIsKnownAttacker sets IsKnownAttacker field to given value.

### HasIsKnownAttacker

`func (o *IpSecurity) HasIsKnownAttacker() bool`

HasIsKnownAttacker returns a boolean if a field has been set.

### GetIsBot

`func (o *IpSecurity) GetIsBot() bool`

GetIsBot returns the IsBot field if non-nil, zero value otherwise.

### GetIsBotOk

`func (o *IpSecurity) GetIsBotOk() (*bool, bool)`

GetIsBotOk returns a tuple with the IsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBot

`func (o *IpSecurity) SetIsBot(v bool)`

SetIsBot sets IsBot field to given value.

### HasIsBot

`func (o *IpSecurity) HasIsBot() bool`

HasIsBot returns a boolean if a field has been set.

### GetIsSpam

`func (o *IpSecurity) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *IpSecurity) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *IpSecurity) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *IpSecurity) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetIsCloudProvider

`func (o *IpSecurity) GetIsCloudProvider() bool`

GetIsCloudProvider returns the IsCloudProvider field if non-nil, zero value otherwise.

### GetIsCloudProviderOk

`func (o *IpSecurity) GetIsCloudProviderOk() (*bool, bool)`

GetIsCloudProviderOk returns a tuple with the IsCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudProvider

`func (o *IpSecurity) SetIsCloudProvider(v bool)`

SetIsCloudProvider sets IsCloudProvider field to given value.

### HasIsCloudProvider

`func (o *IpSecurity) HasIsCloudProvider() bool`

HasIsCloudProvider returns a boolean if a field has been set.

### GetCloudProviderName

`func (o *IpSecurity) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *IpSecurity) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *IpSecurity) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.

### HasCloudProviderName

`func (o *IpSecurity) HasCloudProviderName() bool`

HasCloudProviderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


