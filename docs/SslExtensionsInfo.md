# SslExtensionsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorityKeyIdentifier** | Pointer to **string** |  | [optional] 
**SubjectKeyIdentifier** | Pointer to **string** |  | [optional] 
**KeyUsages** | Pointer to **[]string** |  | [optional] 
**ExtendedKeyUsages** | Pointer to **[]string** |  | [optional] 
**CrlDistributionPoints** | Pointer to **[]string** |  | [optional] 
**AuthorityInfoAccess** | Pointer to [**SslAuthorityInfo**](SslAuthorityInfo.md) |  | [optional] 
**SubjectAlternativeNames** | Pointer to [**SslAlternateNames**](SslAlternateNames.md) |  | [optional] 
**CertificatePolicies** | Pointer to [**[]SslCertificatePolicy**](SslCertificatePolicy.md) |  | [optional] 

## Methods

### NewSslExtensionsInfo

`func NewSslExtensionsInfo() *SslExtensionsInfo`

NewSslExtensionsInfo instantiates a new SslExtensionsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslExtensionsInfoWithDefaults

`func NewSslExtensionsInfoWithDefaults() *SslExtensionsInfo`

NewSslExtensionsInfoWithDefaults instantiates a new SslExtensionsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorityKeyIdentifier

`func (o *SslExtensionsInfo) GetAuthorityKeyIdentifier() string`

GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field if non-nil, zero value otherwise.

### GetAuthorityKeyIdentifierOk

`func (o *SslExtensionsInfo) GetAuthorityKeyIdentifierOk() (*string, bool)`

GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyIdentifier

`func (o *SslExtensionsInfo) SetAuthorityKeyIdentifier(v string)`

SetAuthorityKeyIdentifier sets AuthorityKeyIdentifier field to given value.

### HasAuthorityKeyIdentifier

`func (o *SslExtensionsInfo) HasAuthorityKeyIdentifier() bool`

HasAuthorityKeyIdentifier returns a boolean if a field has been set.

### GetSubjectKeyIdentifier

`func (o *SslExtensionsInfo) GetSubjectKeyIdentifier() string`

GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field if non-nil, zero value otherwise.

### GetSubjectKeyIdentifierOk

`func (o *SslExtensionsInfo) GetSubjectKeyIdentifierOk() (*string, bool)`

GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyIdentifier

`func (o *SslExtensionsInfo) SetSubjectKeyIdentifier(v string)`

SetSubjectKeyIdentifier sets SubjectKeyIdentifier field to given value.

### HasSubjectKeyIdentifier

`func (o *SslExtensionsInfo) HasSubjectKeyIdentifier() bool`

HasSubjectKeyIdentifier returns a boolean if a field has been set.

### GetKeyUsages

`func (o *SslExtensionsInfo) GetKeyUsages() []string`

GetKeyUsages returns the KeyUsages field if non-nil, zero value otherwise.

### GetKeyUsagesOk

`func (o *SslExtensionsInfo) GetKeyUsagesOk() (*[]string, bool)`

GetKeyUsagesOk returns a tuple with the KeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsages

`func (o *SslExtensionsInfo) SetKeyUsages(v []string)`

SetKeyUsages sets KeyUsages field to given value.

### HasKeyUsages

`func (o *SslExtensionsInfo) HasKeyUsages() bool`

HasKeyUsages returns a boolean if a field has been set.

### GetExtendedKeyUsages

`func (o *SslExtensionsInfo) GetExtendedKeyUsages() []string`

GetExtendedKeyUsages returns the ExtendedKeyUsages field if non-nil, zero value otherwise.

### GetExtendedKeyUsagesOk

`func (o *SslExtensionsInfo) GetExtendedKeyUsagesOk() (*[]string, bool)`

GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsages

`func (o *SslExtensionsInfo) SetExtendedKeyUsages(v []string)`

SetExtendedKeyUsages sets ExtendedKeyUsages field to given value.

### HasExtendedKeyUsages

`func (o *SslExtensionsInfo) HasExtendedKeyUsages() bool`

HasExtendedKeyUsages returns a boolean if a field has been set.

### GetCrlDistributionPoints

`func (o *SslExtensionsInfo) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *SslExtensionsInfo) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *SslExtensionsInfo) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.

### HasCrlDistributionPoints

`func (o *SslExtensionsInfo) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.

### GetAuthorityInfoAccess

`func (o *SslExtensionsInfo) GetAuthorityInfoAccess() SslAuthorityInfo`

GetAuthorityInfoAccess returns the AuthorityInfoAccess field if non-nil, zero value otherwise.

### GetAuthorityInfoAccessOk

`func (o *SslExtensionsInfo) GetAuthorityInfoAccessOk() (*SslAuthorityInfo, bool)`

GetAuthorityInfoAccessOk returns a tuple with the AuthorityInfoAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityInfoAccess

`func (o *SslExtensionsInfo) SetAuthorityInfoAccess(v SslAuthorityInfo)`

SetAuthorityInfoAccess sets AuthorityInfoAccess field to given value.

### HasAuthorityInfoAccess

`func (o *SslExtensionsInfo) HasAuthorityInfoAccess() bool`

HasAuthorityInfoAccess returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *SslExtensionsInfo) GetSubjectAlternativeNames() SslAlternateNames`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *SslExtensionsInfo) GetSubjectAlternativeNamesOk() (*SslAlternateNames, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *SslExtensionsInfo) SetSubjectAlternativeNames(v SslAlternateNames)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *SslExtensionsInfo) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetCertificatePolicies

`func (o *SslExtensionsInfo) GetCertificatePolicies() []SslCertificatePolicy`

GetCertificatePolicies returns the CertificatePolicies field if non-nil, zero value otherwise.

### GetCertificatePoliciesOk

`func (o *SslExtensionsInfo) GetCertificatePoliciesOk() (*[]SslCertificatePolicy, bool)`

GetCertificatePoliciesOk returns a tuple with the CertificatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePolicies

`func (o *SslExtensionsInfo) SetCertificatePolicies(v []SslCertificatePolicy)`

SetCertificatePolicies sets CertificatePolicies field to given value.

### HasCertificatePolicies

`func (o *SslExtensionsInfo) HasCertificatePolicies() bool`

HasCertificatePolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


