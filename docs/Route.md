# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Route** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **[]string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Pingable** | Pointer to **[]string** |  | [optional] 
**PingHdl** | Pointer to **[]string** |  | [optional] 
**Holes** | Pointer to **[]string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to **[]string** |  | [optional] 
**MemberOf** | Pointer to **[]string** |  | [optional] 
**Inject** | Pointer to **[]string** |  | [optional] 
**AggrMtd** | Pointer to **string** |  | [optional] 
**AggrBndry** | Pointer to **string** |  | [optional] 
**ExportComps** | Pointer to **string** |  | [optional] 
**Components** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **[]string** |  | [optional] 
**Notify** | Pointer to **[]string** |  | [optional] 
**MntLower** | Pointer to **[]string** |  | [optional] 
**MntRoutes** | Pointer to **[]string** |  | [optional] 
**MntBy** | Pointer to **[]string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoute

`func (o *Route) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *Route) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *Route) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *Route) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetDescription

`func (o *Route) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Route) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Route) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Route) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrigin

`func (o *Route) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Route) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Route) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Route) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPingable

`func (o *Route) GetPingable() []string`

GetPingable returns the Pingable field if non-nil, zero value otherwise.

### GetPingableOk

`func (o *Route) GetPingableOk() (*[]string, bool)`

GetPingableOk returns a tuple with the Pingable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingable

`func (o *Route) SetPingable(v []string)`

SetPingable sets Pingable field to given value.

### HasPingable

`func (o *Route) HasPingable() bool`

HasPingable returns a boolean if a field has been set.

### GetPingHdl

`func (o *Route) GetPingHdl() []string`

GetPingHdl returns the PingHdl field if non-nil, zero value otherwise.

### GetPingHdlOk

`func (o *Route) GetPingHdlOk() (*[]string, bool)`

GetPingHdlOk returns a tuple with the PingHdl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingHdl

`func (o *Route) SetPingHdl(v []string)`

SetPingHdl sets PingHdl field to given value.

### HasPingHdl

`func (o *Route) HasPingHdl() bool`

HasPingHdl returns a boolean if a field has been set.

### GetHoles

`func (o *Route) GetHoles() []string`

GetHoles returns the Holes field if non-nil, zero value otherwise.

### GetHolesOk

`func (o *Route) GetHolesOk() (*[]string, bool)`

GetHolesOk returns a tuple with the Holes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoles

`func (o *Route) SetHoles(v []string)`

SetHoles sets Holes field to given value.

### HasHoles

`func (o *Route) HasHoles() bool`

HasHoles returns a boolean if a field has been set.

### GetCountry

`func (o *Route) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Route) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Route) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Route) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetOrganizations

`func (o *Route) GetOrganizations() []string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *Route) GetOrganizationsOk() (*[]string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *Route) SetOrganizations(v []string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *Route) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetMemberOf

`func (o *Route) GetMemberOf() []string`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *Route) GetMemberOfOk() (*[]string, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *Route) SetMemberOf(v []string)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *Route) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetInject

`func (o *Route) GetInject() []string`

GetInject returns the Inject field if non-nil, zero value otherwise.

### GetInjectOk

`func (o *Route) GetInjectOk() (*[]string, bool)`

GetInjectOk returns a tuple with the Inject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInject

`func (o *Route) SetInject(v []string)`

SetInject sets Inject field to given value.

### HasInject

`func (o *Route) HasInject() bool`

HasInject returns a boolean if a field has been set.

### GetAggrMtd

`func (o *Route) GetAggrMtd() string`

GetAggrMtd returns the AggrMtd field if non-nil, zero value otherwise.

### GetAggrMtdOk

`func (o *Route) GetAggrMtdOk() (*string, bool)`

GetAggrMtdOk returns a tuple with the AggrMtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggrMtd

`func (o *Route) SetAggrMtd(v string)`

SetAggrMtd sets AggrMtd field to given value.

### HasAggrMtd

`func (o *Route) HasAggrMtd() bool`

HasAggrMtd returns a boolean if a field has been set.

### GetAggrBndry

`func (o *Route) GetAggrBndry() string`

GetAggrBndry returns the AggrBndry field if non-nil, zero value otherwise.

### GetAggrBndryOk

`func (o *Route) GetAggrBndryOk() (*string, bool)`

GetAggrBndryOk returns a tuple with the AggrBndry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggrBndry

`func (o *Route) SetAggrBndry(v string)`

SetAggrBndry sets AggrBndry field to given value.

### HasAggrBndry

`func (o *Route) HasAggrBndry() bool`

HasAggrBndry returns a boolean if a field has been set.

### GetExportComps

`func (o *Route) GetExportComps() string`

GetExportComps returns the ExportComps field if non-nil, zero value otherwise.

### GetExportCompsOk

`func (o *Route) GetExportCompsOk() (*string, bool)`

GetExportCompsOk returns a tuple with the ExportComps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportComps

`func (o *Route) SetExportComps(v string)`

SetExportComps sets ExportComps field to given value.

### HasExportComps

`func (o *Route) HasExportComps() bool`

HasExportComps returns a boolean if a field has been set.

### GetComponents

`func (o *Route) GetComponents() string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Route) GetComponentsOk() (*string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Route) SetComponents(v string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *Route) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetRemarks

`func (o *Route) GetRemarks() []string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *Route) GetRemarksOk() (*[]string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *Route) SetRemarks(v []string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *Route) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetNotify

`func (o *Route) GetNotify() []string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *Route) GetNotifyOk() (*[]string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *Route) SetNotify(v []string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *Route) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetMntLower

`func (o *Route) GetMntLower() []string`

GetMntLower returns the MntLower field if non-nil, zero value otherwise.

### GetMntLowerOk

`func (o *Route) GetMntLowerOk() (*[]string, bool)`

GetMntLowerOk returns a tuple with the MntLower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntLower

`func (o *Route) SetMntLower(v []string)`

SetMntLower sets MntLower field to given value.

### HasMntLower

`func (o *Route) HasMntLower() bool`

HasMntLower returns a boolean if a field has been set.

### GetMntRoutes

`func (o *Route) GetMntRoutes() []string`

GetMntRoutes returns the MntRoutes field if non-nil, zero value otherwise.

### GetMntRoutesOk

`func (o *Route) GetMntRoutesOk() (*[]string, bool)`

GetMntRoutesOk returns a tuple with the MntRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntRoutes

`func (o *Route) SetMntRoutes(v []string)`

SetMntRoutes sets MntRoutes field to given value.

### HasMntRoutes

`func (o *Route) HasMntRoutes() bool`

HasMntRoutes returns a boolean if a field has been set.

### GetMntBy

`func (o *Route) GetMntBy() []string`

GetMntBy returns the MntBy field if non-nil, zero value otherwise.

### GetMntByOk

`func (o *Route) GetMntByOk() (*[]string, bool)`

GetMntByOk returns a tuple with the MntBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMntBy

`func (o *Route) SetMntBy(v []string)`

SetMntBy sets MntBy field to given value.

### HasMntBy

`func (o *Route) HasMntBy() bool`

HasMntBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *Route) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Route) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Route) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Route) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *Route) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *Route) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *Route) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *Route) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetSource

`func (o *Route) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Route) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Route) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Route) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


