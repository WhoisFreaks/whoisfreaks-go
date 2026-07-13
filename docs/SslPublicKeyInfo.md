# SslPublicKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySize** | Pointer to **string** |  | [optional] 
**KeyAlgorithm** | Pointer to **string** |  | [optional] 
**PemRaw** | Pointer to **string** |  | [optional] 

## Methods

### NewSslPublicKeyInfo

`func NewSslPublicKeyInfo() *SslPublicKeyInfo`

NewSslPublicKeyInfo instantiates a new SslPublicKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslPublicKeyInfoWithDefaults

`func NewSslPublicKeyInfoWithDefaults() *SslPublicKeyInfo`

NewSslPublicKeyInfoWithDefaults instantiates a new SslPublicKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySize

`func (o *SslPublicKeyInfo) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *SslPublicKeyInfo) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *SslPublicKeyInfo) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *SslPublicKeyInfo) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *SslPublicKeyInfo) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *SslPublicKeyInfo) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *SslPublicKeyInfo) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *SslPublicKeyInfo) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetPemRaw

`func (o *SslPublicKeyInfo) GetPemRaw() string`

GetPemRaw returns the PemRaw field if non-nil, zero value otherwise.

### GetPemRawOk

`func (o *SslPublicKeyInfo) GetPemRawOk() (*string, bool)`

GetPemRawOk returns a tuple with the PemRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemRaw

`func (o *SslPublicKeyInfo) SetPemRaw(v string)`

SetPemRaw sets PemRaw field to given value.

### HasPemRaw

`func (o *SslPublicKeyInfo) HasPemRaw() bool`

HasPemRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


