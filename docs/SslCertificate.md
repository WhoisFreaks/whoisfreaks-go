# SslCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainOrder** | Pointer to **string** |  | [optional] 
**AuthenticationType** | Pointer to **string** |  | [optional] 
**ValidityStartDate** | Pointer to **string** |  | [optional] 
**ValidityEndDate** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SignatureAlgorithm** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to [**SslUnitInfo**](SslUnitInfo.md) |  | [optional] 
**Issuer** | Pointer to [**SslUnitInfo**](SslUnitInfo.md) |  | [optional] 
**PublicKey** | Pointer to [**SslPublicKeyInfo**](SslPublicKeyInfo.md) |  | [optional] 
**Extensions** | Pointer to [**SslExtensionsInfo**](SslExtensionsInfo.md) |  | [optional] 
**PemRaw** | Pointer to **string** |  | [optional] 

## Methods

### NewSslCertificate

`func NewSslCertificate() *SslCertificate`

NewSslCertificate instantiates a new SslCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslCertificateWithDefaults

`func NewSslCertificateWithDefaults() *SslCertificate`

NewSslCertificateWithDefaults instantiates a new SslCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainOrder

`func (o *SslCertificate) GetChainOrder() string`

GetChainOrder returns the ChainOrder field if non-nil, zero value otherwise.

### GetChainOrderOk

`func (o *SslCertificate) GetChainOrderOk() (*string, bool)`

GetChainOrderOk returns a tuple with the ChainOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainOrder

`func (o *SslCertificate) SetChainOrder(v string)`

SetChainOrder sets ChainOrder field to given value.

### HasChainOrder

`func (o *SslCertificate) HasChainOrder() bool`

HasChainOrder returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *SslCertificate) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *SslCertificate) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *SslCertificate) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *SslCertificate) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetValidityStartDate

`func (o *SslCertificate) GetValidityStartDate() string`

GetValidityStartDate returns the ValidityStartDate field if non-nil, zero value otherwise.

### GetValidityStartDateOk

`func (o *SslCertificate) GetValidityStartDateOk() (*string, bool)`

GetValidityStartDateOk returns a tuple with the ValidityStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityStartDate

`func (o *SslCertificate) SetValidityStartDate(v string)`

SetValidityStartDate sets ValidityStartDate field to given value.

### HasValidityStartDate

`func (o *SslCertificate) HasValidityStartDate() bool`

HasValidityStartDate returns a boolean if a field has been set.

### GetValidityEndDate

`func (o *SslCertificate) GetValidityEndDate() string`

GetValidityEndDate returns the ValidityEndDate field if non-nil, zero value otherwise.

### GetValidityEndDateOk

`func (o *SslCertificate) GetValidityEndDateOk() (*string, bool)`

GetValidityEndDateOk returns a tuple with the ValidityEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityEndDate

`func (o *SslCertificate) SetValidityEndDate(v string)`

SetValidityEndDate sets ValidityEndDate field to given value.

### HasValidityEndDate

`func (o *SslCertificate) HasValidityEndDate() bool`

HasValidityEndDate returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SslCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SslCertificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SslCertificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SslCertificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *SslCertificate) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SslCertificate) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SslCertificate) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SslCertificate) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSubject

`func (o *SslCertificate) GetSubject() SslUnitInfo`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SslCertificate) GetSubjectOk() (*SslUnitInfo, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SslCertificate) SetSubject(v SslUnitInfo)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SslCertificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetIssuer

`func (o *SslCertificate) GetIssuer() SslUnitInfo`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SslCertificate) GetIssuerOk() (*SslUnitInfo, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SslCertificate) SetIssuer(v SslUnitInfo)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SslCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetPublicKey

`func (o *SslCertificate) GetPublicKey() SslPublicKeyInfo`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SslCertificate) GetPublicKeyOk() (*SslPublicKeyInfo, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SslCertificate) SetPublicKey(v SslPublicKeyInfo)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SslCertificate) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetExtensions

`func (o *SslCertificate) GetExtensions() SslExtensionsInfo`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SslCertificate) GetExtensionsOk() (*SslExtensionsInfo, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SslCertificate) SetExtensions(v SslExtensionsInfo)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *SslCertificate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetPemRaw

`func (o *SslCertificate) GetPemRaw() string`

GetPemRaw returns the PemRaw field if non-nil, zero value otherwise.

### GetPemRawOk

`func (o *SslCertificate) GetPemRawOk() (*string, bool)`

GetPemRawOk returns a tuple with the PemRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemRaw

`func (o *SslCertificate) SetPemRaw(v string)`

SetPemRaw sets PemRaw field to given value.

### HasPemRaw

`func (o *SslCertificate) HasPemRaw() bool`

HasPemRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


