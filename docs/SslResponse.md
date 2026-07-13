# SslResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** |  | [optional] 
**QueryTime** | Pointer to **string** |  | [optional] 
**SslCertificates** | Pointer to [**[]SslCertificate**](SslCertificate.md) |  | [optional] 
**SslRaw** | Pointer to **string** |  | [optional] 

## Methods

### NewSslResponse

`func NewSslResponse() *SslResponse`

NewSslResponse instantiates a new SslResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslResponseWithDefaults

`func NewSslResponseWithDefaults() *SslResponse`

NewSslResponseWithDefaults instantiates a new SslResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *SslResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SslResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SslResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *SslResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetQueryTime

`func (o *SslResponse) GetQueryTime() string`

GetQueryTime returns the QueryTime field if non-nil, zero value otherwise.

### GetQueryTimeOk

`func (o *SslResponse) GetQueryTimeOk() (*string, bool)`

GetQueryTimeOk returns a tuple with the QueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTime

`func (o *SslResponse) SetQueryTime(v string)`

SetQueryTime sets QueryTime field to given value.

### HasQueryTime

`func (o *SslResponse) HasQueryTime() bool`

HasQueryTime returns a boolean if a field has been set.

### GetSslCertificates

`func (o *SslResponse) GetSslCertificates() []SslCertificate`

GetSslCertificates returns the SslCertificates field if non-nil, zero value otherwise.

### GetSslCertificatesOk

`func (o *SslResponse) GetSslCertificatesOk() (*[]SslCertificate, bool)`

GetSslCertificatesOk returns a tuple with the SslCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificates

`func (o *SslResponse) SetSslCertificates(v []SslCertificate)`

SetSslCertificates sets SslCertificates field to given value.

### HasSslCertificates

`func (o *SslResponse) HasSslCertificates() bool`

HasSslCertificates returns a boolean if a field has been set.

### GetSslRaw

`func (o *SslResponse) GetSslRaw() string`

GetSslRaw returns the SslRaw field if non-nil, zero value otherwise.

### GetSslRawOk

`func (o *SslResponse) GetSslRawOk() (*string, bool)`

GetSslRawOk returns a tuple with the SslRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRaw

`func (o *SslResponse) SetSslRaw(v string)`

SetSslRaw sets SslRaw field to given value.

### HasSslRaw

`func (o *SslResponse) HasSslRaw() bool`

HasSslRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


