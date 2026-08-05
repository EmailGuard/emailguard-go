# PublicTeamInviteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address the invite was sent to. | 
**ExpiresAt** | **string** | When the invite expires if not accepted (RFC3339). | 
**Id** | **string** | Invite record ID. | 
**Role** | **string** | Role the invitee will receive when they accept. | 

## Methods

### NewPublicTeamInviteResponse

`func NewPublicTeamInviteResponse(email string, expiresAt string, id string, role string, ) *PublicTeamInviteResponse`

NewPublicTeamInviteResponse instantiates a new PublicTeamInviteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTeamInviteResponseWithDefaults

`func NewPublicTeamInviteResponseWithDefaults() *PublicTeamInviteResponse`

NewPublicTeamInviteResponseWithDefaults instantiates a new PublicTeamInviteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PublicTeamInviteResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PublicTeamInviteResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PublicTeamInviteResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpiresAt

`func (o *PublicTeamInviteResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PublicTeamInviteResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PublicTeamInviteResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *PublicTeamInviteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicTeamInviteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicTeamInviteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRole

`func (o *PublicTeamInviteResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PublicTeamInviteResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PublicTeamInviteResponse) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


