# ReputationIndicators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsNewlyRegistered** | Pointer to **bool** |  | [optional] 
**UsesFreeExtension** | Pointer to **bool** |  | [optional] 
**UsesFreeSsl** | Pointer to **bool** |  | [optional] 
**HasPrivacyWhois** | Pointer to **bool** |  | [optional] 
**SslAgeDays** | Pointer to **int32** |  | [optional] 
**HasDmarc** | Pointer to **bool** |  | [optional] 
**HasSpf** | Pointer to **bool** |  | [optional] 
**RedirectsExternally** | Pointer to **bool** |  | [optional] 
**JavascriptObfuscated** | Pointer to **bool** |  | [optional] 
**DomainAgeDays** | Pointer to **int32** |  | [optional] 
**Registrar** | Pointer to **string** |  | [optional] 

## Methods

### NewReputationIndicators

`func NewReputationIndicators() *ReputationIndicators`

NewReputationIndicators instantiates a new ReputationIndicators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReputationIndicatorsWithDefaults

`func NewReputationIndicatorsWithDefaults() *ReputationIndicators`

NewReputationIndicatorsWithDefaults instantiates a new ReputationIndicators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsNewlyRegistered

`func (o *ReputationIndicators) GetIsNewlyRegistered() bool`

GetIsNewlyRegistered returns the IsNewlyRegistered field if non-nil, zero value otherwise.

### GetIsNewlyRegisteredOk

`func (o *ReputationIndicators) GetIsNewlyRegisteredOk() (*bool, bool)`

GetIsNewlyRegisteredOk returns a tuple with the IsNewlyRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewlyRegistered

`func (o *ReputationIndicators) SetIsNewlyRegistered(v bool)`

SetIsNewlyRegistered sets IsNewlyRegistered field to given value.

### HasIsNewlyRegistered

`func (o *ReputationIndicators) HasIsNewlyRegistered() bool`

HasIsNewlyRegistered returns a boolean if a field has been set.

### GetUsesFreeExtension

`func (o *ReputationIndicators) GetUsesFreeExtension() bool`

GetUsesFreeExtension returns the UsesFreeExtension field if non-nil, zero value otherwise.

### GetUsesFreeExtensionOk

`func (o *ReputationIndicators) GetUsesFreeExtensionOk() (*bool, bool)`

GetUsesFreeExtensionOk returns a tuple with the UsesFreeExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesFreeExtension

`func (o *ReputationIndicators) SetUsesFreeExtension(v bool)`

SetUsesFreeExtension sets UsesFreeExtension field to given value.

### HasUsesFreeExtension

`func (o *ReputationIndicators) HasUsesFreeExtension() bool`

HasUsesFreeExtension returns a boolean if a field has been set.

### GetUsesFreeSsl

`func (o *ReputationIndicators) GetUsesFreeSsl() bool`

GetUsesFreeSsl returns the UsesFreeSsl field if non-nil, zero value otherwise.

### GetUsesFreeSslOk

`func (o *ReputationIndicators) GetUsesFreeSslOk() (*bool, bool)`

GetUsesFreeSslOk returns a tuple with the UsesFreeSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesFreeSsl

`func (o *ReputationIndicators) SetUsesFreeSsl(v bool)`

SetUsesFreeSsl sets UsesFreeSsl field to given value.

### HasUsesFreeSsl

`func (o *ReputationIndicators) HasUsesFreeSsl() bool`

HasUsesFreeSsl returns a boolean if a field has been set.

### GetHasPrivacyWhois

`func (o *ReputationIndicators) GetHasPrivacyWhois() bool`

GetHasPrivacyWhois returns the HasPrivacyWhois field if non-nil, zero value otherwise.

### GetHasPrivacyWhoisOk

`func (o *ReputationIndicators) GetHasPrivacyWhoisOk() (*bool, bool)`

GetHasPrivacyWhoisOk returns a tuple with the HasPrivacyWhois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivacyWhois

`func (o *ReputationIndicators) SetHasPrivacyWhois(v bool)`

SetHasPrivacyWhois sets HasPrivacyWhois field to given value.

### HasHasPrivacyWhois

`func (o *ReputationIndicators) HasHasPrivacyWhois() bool`

HasHasPrivacyWhois returns a boolean if a field has been set.

### GetSslAgeDays

`func (o *ReputationIndicators) GetSslAgeDays() int32`

GetSslAgeDays returns the SslAgeDays field if non-nil, zero value otherwise.

### GetSslAgeDaysOk

`func (o *ReputationIndicators) GetSslAgeDaysOk() (*int32, bool)`

GetSslAgeDaysOk returns a tuple with the SslAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAgeDays

`func (o *ReputationIndicators) SetSslAgeDays(v int32)`

SetSslAgeDays sets SslAgeDays field to given value.

### HasSslAgeDays

`func (o *ReputationIndicators) HasSslAgeDays() bool`

HasSslAgeDays returns a boolean if a field has been set.

### GetHasDmarc

`func (o *ReputationIndicators) GetHasDmarc() bool`

GetHasDmarc returns the HasDmarc field if non-nil, zero value otherwise.

### GetHasDmarcOk

`func (o *ReputationIndicators) GetHasDmarcOk() (*bool, bool)`

GetHasDmarcOk returns a tuple with the HasDmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDmarc

`func (o *ReputationIndicators) SetHasDmarc(v bool)`

SetHasDmarc sets HasDmarc field to given value.

### HasHasDmarc

`func (o *ReputationIndicators) HasHasDmarc() bool`

HasHasDmarc returns a boolean if a field has been set.

### GetHasSpf

`func (o *ReputationIndicators) GetHasSpf() bool`

GetHasSpf returns the HasSpf field if non-nil, zero value otherwise.

### GetHasSpfOk

`func (o *ReputationIndicators) GetHasSpfOk() (*bool, bool)`

GetHasSpfOk returns a tuple with the HasSpf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpf

`func (o *ReputationIndicators) SetHasSpf(v bool)`

SetHasSpf sets HasSpf field to given value.

### HasHasSpf

`func (o *ReputationIndicators) HasHasSpf() bool`

HasHasSpf returns a boolean if a field has been set.

### GetRedirectsExternally

`func (o *ReputationIndicators) GetRedirectsExternally() bool`

GetRedirectsExternally returns the RedirectsExternally field if non-nil, zero value otherwise.

### GetRedirectsExternallyOk

`func (o *ReputationIndicators) GetRedirectsExternallyOk() (*bool, bool)`

GetRedirectsExternallyOk returns a tuple with the RedirectsExternally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectsExternally

`func (o *ReputationIndicators) SetRedirectsExternally(v bool)`

SetRedirectsExternally sets RedirectsExternally field to given value.

### HasRedirectsExternally

`func (o *ReputationIndicators) HasRedirectsExternally() bool`

HasRedirectsExternally returns a boolean if a field has been set.

### GetJavascriptObfuscated

`func (o *ReputationIndicators) GetJavascriptObfuscated() bool`

GetJavascriptObfuscated returns the JavascriptObfuscated field if non-nil, zero value otherwise.

### GetJavascriptObfuscatedOk

`func (o *ReputationIndicators) GetJavascriptObfuscatedOk() (*bool, bool)`

GetJavascriptObfuscatedOk returns a tuple with the JavascriptObfuscated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptObfuscated

`func (o *ReputationIndicators) SetJavascriptObfuscated(v bool)`

SetJavascriptObfuscated sets JavascriptObfuscated field to given value.

### HasJavascriptObfuscated

`func (o *ReputationIndicators) HasJavascriptObfuscated() bool`

HasJavascriptObfuscated returns a boolean if a field has been set.

### GetDomainAgeDays

`func (o *ReputationIndicators) GetDomainAgeDays() int32`

GetDomainAgeDays returns the DomainAgeDays field if non-nil, zero value otherwise.

### GetDomainAgeDaysOk

`func (o *ReputationIndicators) GetDomainAgeDaysOk() (*int32, bool)`

GetDomainAgeDaysOk returns a tuple with the DomainAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAgeDays

`func (o *ReputationIndicators) SetDomainAgeDays(v int32)`

SetDomainAgeDays sets DomainAgeDays field to given value.

### HasDomainAgeDays

`func (o *ReputationIndicators) HasDomainAgeDays() bool`

HasDomainAgeDays returns a boolean if a field has been set.

### GetRegistrar

`func (o *ReputationIndicators) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *ReputationIndicators) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *ReputationIndicators) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *ReputationIndicators) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


