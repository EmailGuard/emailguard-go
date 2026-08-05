# UpdateTeamInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Contact email address for the team workspace. | [optional] 
**Emoji** | Pointer to **string** | Single grapheme emoji shown as the team avatar in the product UI. | [optional] 
**Name** | **string** | Human-readable team name. | 
**Url** | Pointer to **string** | Public website URL for the team. | [optional] 

## Methods

### NewUpdateTeamInput

`func NewUpdateTeamInput(name string, ) *UpdateTeamInput`

NewUpdateTeamInput instantiates a new UpdateTeamInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamInputWithDefaults

`func NewUpdateTeamInputWithDefaults() *UpdateTeamInput`

NewUpdateTeamInputWithDefaults instantiates a new UpdateTeamInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateTeamInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateTeamInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateTeamInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateTeamInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmoji

`func (o *UpdateTeamInput) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *UpdateTeamInput) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *UpdateTeamInput) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *UpdateTeamInput) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetName

`func (o *UpdateTeamInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTeamInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTeamInput) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *UpdateTeamInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateTeamInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateTeamInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateTeamInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


