# EmailDetectData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionSource** | Pointer to **string** |  | [optional] 
**Disposable** | Pointer to **bool** |  | [optional] 
**DisposableProvider** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Normalized** | Pointer to **string** |  | [optional] 
**PublicDomain** | Pointer to **bool** |  | [optional] 
**RelayDomain** | Pointer to **bool** |  | [optional] 
**RelayProvider** | Pointer to **string** |  | [optional] 
**RoleAddress** | Pointer to **bool** |  | [optional] 
**Subaddressing** | Pointer to **bool** |  | [optional] 
**SuggestedDomain** | Pointer to **string** | Corrected domain candidate when a domain or TLD typo rule matches. | [optional] 
**SuggestedEmail** | Pointer to **string** | Corrected email candidate when a domain or TLD typo rule matches. | [optional] 
**SyntaxValidation** | Pointer to **bool** |  | [optional] 

## Methods

### NewEmailDetectData

`func NewEmailDetectData() *EmailDetectData`

NewEmailDetectData instantiates a new EmailDetectData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDetectDataWithDefaults

`func NewEmailDetectDataWithDefaults() *EmailDetectData`

NewEmailDetectDataWithDefaults instantiates a new EmailDetectData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectionSource

`func (o *EmailDetectData) GetDetectionSource() string`

GetDetectionSource returns the DetectionSource field if non-nil, zero value otherwise.

### GetDetectionSourceOk

`func (o *EmailDetectData) GetDetectionSourceOk() (*string, bool)`

GetDetectionSourceOk returns a tuple with the DetectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionSource

`func (o *EmailDetectData) SetDetectionSource(v string)`

SetDetectionSource sets DetectionSource field to given value.

### HasDetectionSource

`func (o *EmailDetectData) HasDetectionSource() bool`

HasDetectionSource returns a boolean if a field has been set.

### GetDisposable

`func (o *EmailDetectData) GetDisposable() bool`

GetDisposable returns the Disposable field if non-nil, zero value otherwise.

### GetDisposableOk

`func (o *EmailDetectData) GetDisposableOk() (*bool, bool)`

GetDisposableOk returns a tuple with the Disposable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposable

`func (o *EmailDetectData) SetDisposable(v bool)`

SetDisposable sets Disposable field to given value.

### HasDisposable

`func (o *EmailDetectData) HasDisposable() bool`

HasDisposable returns a boolean if a field has been set.

### GetDisposableProvider

`func (o *EmailDetectData) GetDisposableProvider() string`

GetDisposableProvider returns the DisposableProvider field if non-nil, zero value otherwise.

### GetDisposableProviderOk

`func (o *EmailDetectData) GetDisposableProviderOk() (*string, bool)`

GetDisposableProviderOk returns a tuple with the DisposableProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposableProvider

`func (o *EmailDetectData) SetDisposableProvider(v string)`

SetDisposableProvider sets DisposableProvider field to given value.

### HasDisposableProvider

`func (o *EmailDetectData) HasDisposableProvider() bool`

HasDisposableProvider returns a boolean if a field has been set.

### GetDomain

`func (o *EmailDetectData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailDetectData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailDetectData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EmailDetectData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEmail

`func (o *EmailDetectData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailDetectData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailDetectData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailDetectData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNormalized

`func (o *EmailDetectData) GetNormalized() string`

GetNormalized returns the Normalized field if non-nil, zero value otherwise.

### GetNormalizedOk

`func (o *EmailDetectData) GetNormalizedOk() (*string, bool)`

GetNormalizedOk returns a tuple with the Normalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalized

`func (o *EmailDetectData) SetNormalized(v string)`

SetNormalized sets Normalized field to given value.

### HasNormalized

`func (o *EmailDetectData) HasNormalized() bool`

HasNormalized returns a boolean if a field has been set.

### GetPublicDomain

`func (o *EmailDetectData) GetPublicDomain() bool`

GetPublicDomain returns the PublicDomain field if non-nil, zero value otherwise.

### GetPublicDomainOk

`func (o *EmailDetectData) GetPublicDomainOk() (*bool, bool)`

GetPublicDomainOk returns a tuple with the PublicDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDomain

`func (o *EmailDetectData) SetPublicDomain(v bool)`

SetPublicDomain sets PublicDomain field to given value.

### HasPublicDomain

`func (o *EmailDetectData) HasPublicDomain() bool`

HasPublicDomain returns a boolean if a field has been set.

### GetRelayDomain

`func (o *EmailDetectData) GetRelayDomain() bool`

GetRelayDomain returns the RelayDomain field if non-nil, zero value otherwise.

### GetRelayDomainOk

`func (o *EmailDetectData) GetRelayDomainOk() (*bool, bool)`

GetRelayDomainOk returns a tuple with the RelayDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayDomain

`func (o *EmailDetectData) SetRelayDomain(v bool)`

SetRelayDomain sets RelayDomain field to given value.

### HasRelayDomain

`func (o *EmailDetectData) HasRelayDomain() bool`

HasRelayDomain returns a boolean if a field has been set.

### GetRelayProvider

`func (o *EmailDetectData) GetRelayProvider() string`

GetRelayProvider returns the RelayProvider field if non-nil, zero value otherwise.

### GetRelayProviderOk

`func (o *EmailDetectData) GetRelayProviderOk() (*string, bool)`

GetRelayProviderOk returns a tuple with the RelayProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayProvider

`func (o *EmailDetectData) SetRelayProvider(v string)`

SetRelayProvider sets RelayProvider field to given value.

### HasRelayProvider

`func (o *EmailDetectData) HasRelayProvider() bool`

HasRelayProvider returns a boolean if a field has been set.

### GetRoleAddress

`func (o *EmailDetectData) GetRoleAddress() bool`

GetRoleAddress returns the RoleAddress field if non-nil, zero value otherwise.

### GetRoleAddressOk

`func (o *EmailDetectData) GetRoleAddressOk() (*bool, bool)`

GetRoleAddressOk returns a tuple with the RoleAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAddress

`func (o *EmailDetectData) SetRoleAddress(v bool)`

SetRoleAddress sets RoleAddress field to given value.

### HasRoleAddress

`func (o *EmailDetectData) HasRoleAddress() bool`

HasRoleAddress returns a boolean if a field has been set.

### GetSubaddressing

`func (o *EmailDetectData) GetSubaddressing() bool`

GetSubaddressing returns the Subaddressing field if non-nil, zero value otherwise.

### GetSubaddressingOk

`func (o *EmailDetectData) GetSubaddressingOk() (*bool, bool)`

GetSubaddressingOk returns a tuple with the Subaddressing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubaddressing

`func (o *EmailDetectData) SetSubaddressing(v bool)`

SetSubaddressing sets Subaddressing field to given value.

### HasSubaddressing

`func (o *EmailDetectData) HasSubaddressing() bool`

HasSubaddressing returns a boolean if a field has been set.

### GetSuggestedDomain

`func (o *EmailDetectData) GetSuggestedDomain() string`

GetSuggestedDomain returns the SuggestedDomain field if non-nil, zero value otherwise.

### GetSuggestedDomainOk

`func (o *EmailDetectData) GetSuggestedDomainOk() (*string, bool)`

GetSuggestedDomainOk returns a tuple with the SuggestedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedDomain

`func (o *EmailDetectData) SetSuggestedDomain(v string)`

SetSuggestedDomain sets SuggestedDomain field to given value.

### HasSuggestedDomain

`func (o *EmailDetectData) HasSuggestedDomain() bool`

HasSuggestedDomain returns a boolean if a field has been set.

### GetSuggestedEmail

`func (o *EmailDetectData) GetSuggestedEmail() string`

GetSuggestedEmail returns the SuggestedEmail field if non-nil, zero value otherwise.

### GetSuggestedEmailOk

`func (o *EmailDetectData) GetSuggestedEmailOk() (*string, bool)`

GetSuggestedEmailOk returns a tuple with the SuggestedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedEmail

`func (o *EmailDetectData) SetSuggestedEmail(v string)`

SetSuggestedEmail sets SuggestedEmail field to given value.

### HasSuggestedEmail

`func (o *EmailDetectData) HasSuggestedEmail() bool`

HasSuggestedEmail returns a boolean if a field has been set.

### GetSyntaxValidation

`func (o *EmailDetectData) GetSyntaxValidation() bool`

GetSyntaxValidation returns the SyntaxValidation field if non-nil, zero value otherwise.

### GetSyntaxValidationOk

`func (o *EmailDetectData) GetSyntaxValidationOk() (*bool, bool)`

GetSyntaxValidationOk returns a tuple with the SyntaxValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxValidation

`func (o *EmailDetectData) SetSyntaxValidation(v bool)`

SetSyntaxValidation sets SyntaxValidation field to given value.

### HasSyntaxValidation

`func (o *EmailDetectData) HasSyntaxValidation() bool`

HasSyntaxValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


