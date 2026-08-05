# PublicTeamMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | When the member joined the team (RFC3339). | 
**Email** | **string** | Member login email. | 
**Emoji** | Pointer to **string** | User profile emoji. | [optional] 
**FirstName** | **string** | Member given name. | 
**Id** | **string** | Team membership record ID. | 
**LastName** | **string** | Member family name. | 
**Notes** | Pointer to **string** | Internal notes stored on the membership. | [optional] 
**Role** | **string** | Team role (for example Owner, Admin, Member). | 
**UserId** | **string** | User account ID for this member. | 

## Methods

### NewPublicTeamMemberResponse

`func NewPublicTeamMemberResponse(createdAt string, email string, firstName string, id string, lastName string, role string, userId string, ) *PublicTeamMemberResponse`

NewPublicTeamMemberResponse instantiates a new PublicTeamMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTeamMemberResponseWithDefaults

`func NewPublicTeamMemberResponseWithDefaults() *PublicTeamMemberResponse`

NewPublicTeamMemberResponseWithDefaults instantiates a new PublicTeamMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicTeamMemberResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicTeamMemberResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicTeamMemberResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *PublicTeamMemberResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PublicTeamMemberResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PublicTeamMemberResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmoji

`func (o *PublicTeamMemberResponse) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *PublicTeamMemberResponse) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *PublicTeamMemberResponse) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *PublicTeamMemberResponse) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetFirstName

`func (o *PublicTeamMemberResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PublicTeamMemberResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PublicTeamMemberResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetId

`func (o *PublicTeamMemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicTeamMemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicTeamMemberResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *PublicTeamMemberResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PublicTeamMemberResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PublicTeamMemberResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetNotes

`func (o *PublicTeamMemberResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PublicTeamMemberResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PublicTeamMemberResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PublicTeamMemberResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRole

`func (o *PublicTeamMemberResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PublicTeamMemberResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PublicTeamMemberResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetUserId

`func (o *PublicTeamMemberResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicTeamMemberResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicTeamMemberResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


