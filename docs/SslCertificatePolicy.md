# SslCertificatePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyQualifier** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSslCertificatePolicy

`func NewSslCertificatePolicy() *SslCertificatePolicy`

NewSslCertificatePolicy instantiates a new SslCertificatePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslCertificatePolicyWithDefaults

`func NewSslCertificatePolicyWithDefaults() *SslCertificatePolicy`

NewSslCertificatePolicyWithDefaults instantiates a new SslCertificatePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *SslCertificatePolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SslCertificatePolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SslCertificatePolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *SslCertificatePolicy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyQualifier

`func (o *SslCertificatePolicy) GetPolicyQualifier() map[string]interface{}`

GetPolicyQualifier returns the PolicyQualifier field if non-nil, zero value otherwise.

### GetPolicyQualifierOk

`func (o *SslCertificatePolicy) GetPolicyQualifierOk() (*map[string]interface{}, bool)`

GetPolicyQualifierOk returns a tuple with the PolicyQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyQualifier

`func (o *SslCertificatePolicy) SetPolicyQualifier(v map[string]interface{})`

SetPolicyQualifier sets PolicyQualifier field to given value.

### HasPolicyQualifier

`func (o *SslCertificatePolicy) HasPolicyQualifier() bool`

HasPolicyQualifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


