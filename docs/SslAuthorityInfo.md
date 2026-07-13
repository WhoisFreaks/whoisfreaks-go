# SslAuthorityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuers** | Pointer to **[]string** |  | [optional] 
**Ocsp** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSslAuthorityInfo

`func NewSslAuthorityInfo() *SslAuthorityInfo`

NewSslAuthorityInfo instantiates a new SslAuthorityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslAuthorityInfoWithDefaults

`func NewSslAuthorityInfoWithDefaults() *SslAuthorityInfo`

NewSslAuthorityInfoWithDefaults instantiates a new SslAuthorityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuers

`func (o *SslAuthorityInfo) GetIssuers() []string`

GetIssuers returns the Issuers field if non-nil, zero value otherwise.

### GetIssuersOk

`func (o *SslAuthorityInfo) GetIssuersOk() (*[]string, bool)`

GetIssuersOk returns a tuple with the Issuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuers

`func (o *SslAuthorityInfo) SetIssuers(v []string)`

SetIssuers sets Issuers field to given value.

### HasIssuers

`func (o *SslAuthorityInfo) HasIssuers() bool`

HasIssuers returns a boolean if a field has been set.

### GetOcsp

`func (o *SslAuthorityInfo) GetOcsp() []string`

GetOcsp returns the Ocsp field if non-nil, zero value otherwise.

### GetOcspOk

`func (o *SslAuthorityInfo) GetOcspOk() (*[]string, bool)`

GetOcspOk returns a tuple with the Ocsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcsp

`func (o *SslAuthorityInfo) SetOcsp(v []string)`

SetOcsp sets Ocsp field to given value.

### HasOcsp

`func (o *SslAuthorityInfo) HasOcsp() bool`

HasOcsp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


