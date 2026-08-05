# InviteTeamMemberInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address that will receive the invite. | 
**Role** | **string** | Team role to assign when the invite is accepted (for example Owner, Admin, Member). See [Team roles and permissions](/docs/knowledge-base/team-roles-and-permissions). | 

## Methods

### NewInviteTeamMemberInput

`func NewInviteTeamMemberInput(email string, role string, ) *InviteTeamMemberInput`

NewInviteTeamMemberInput instantiates a new InviteTeamMemberInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteTeamMemberInputWithDefaults

`func NewInviteTeamMemberInputWithDefaults() *InviteTeamMemberInput`

NewInviteTeamMemberInputWithDefaults instantiates a new InviteTeamMemberInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteTeamMemberInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteTeamMemberInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteTeamMemberInput) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InviteTeamMemberInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteTeamMemberInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteTeamMemberInput) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


