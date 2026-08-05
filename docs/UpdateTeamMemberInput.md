# UpdateTeamMemberInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal team member record ID. When calling the public API, this is set from the path userId. | 
**Notes** | Pointer to **string** | Internal notes about the member. Omit to leave unchanged. | [optional] 
**Role** | Pointer to **string** | New team role (for example Owner, Admin, Member). Omit to leave unchanged. See [Team roles and permissions](/docs/knowledge-base/team-roles-and-permissions). | [optional] 

## Methods

### NewUpdateTeamMemberInput

`func NewUpdateTeamMemberInput(id string, ) *UpdateTeamMemberInput`

NewUpdateTeamMemberInput instantiates a new UpdateTeamMemberInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamMemberInputWithDefaults

`func NewUpdateTeamMemberInputWithDefaults() *UpdateTeamMemberInput`

NewUpdateTeamMemberInputWithDefaults instantiates a new UpdateTeamMemberInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTeamMemberInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTeamMemberInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTeamMemberInput) SetId(v string)`

SetId sets Id field to given value.


### GetNotes

`func (o *UpdateTeamMemberInput) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateTeamMemberInput) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateTeamMemberInput) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateTeamMemberInput) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRole

`func (o *UpdateTeamMemberInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateTeamMemberInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateTeamMemberInput) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateTeamMemberInput) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


