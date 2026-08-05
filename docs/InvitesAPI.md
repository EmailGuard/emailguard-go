# \InvitesAPI

All URIs are relative to *http://api.emailguard.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TeamMembersInvitesGet**](InvitesAPI.md#ApiV1TeamMembersInvitesGet) | **Get** /api/v1/team/members/invites | List pending team invites
[**ApiV1TeamMembersInvitesInviteIdDelete**](InvitesAPI.md#ApiV1TeamMembersInvitesInviteIdDelete) | **Delete** /api/v1/team/members/invites/{inviteId} | Cancel a pending invite
[**ApiV1TeamMembersInvitesPost**](InvitesAPI.md#ApiV1TeamMembersInvitesPost) | **Post** /api/v1/team/members/invites | Invite a team member



## ApiV1TeamMembersInvitesGet

> ListPublicTeamInvitesResponse ApiV1TeamMembersInvitesGet(ctx).Execute()

List pending team invites



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EmailGuard/emailguard-go/emailguardsdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.ApiV1TeamMembersInvitesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.ApiV1TeamMembersInvitesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamMembersInvitesGet`: ListPublicTeamInvitesResponse
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.ApiV1TeamMembersInvitesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamMembersInvitesGetRequest struct via the builder pattern


### Return type

[**ListPublicTeamInvitesResponse**](ListPublicTeamInvitesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TeamMembersInvitesInviteIdDelete

> MessageResponse ApiV1TeamMembersInvitesInviteIdDelete(ctx, inviteId).Execute()

Cancel a pending invite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EmailGuard/emailguard-go/emailguardsdk"
)

func main() {
	inviteId := "inviteId_example" // string | UUID of the pending invite to cancel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.ApiV1TeamMembersInvitesInviteIdDelete(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.ApiV1TeamMembersInvitesInviteIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamMembersInvitesInviteIdDelete`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.ApiV1TeamMembersInvitesInviteIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string** | UUID of the pending invite to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamMembersInvitesInviteIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TeamMembersInvitesPost

> MessageResponse ApiV1TeamMembersInvitesPost(ctx).InviteTeamMemberInput(inviteTeamMemberInput).Execute()

Invite a team member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EmailGuard/emailguard-go/emailguardsdk"
)

func main() {
	inviteTeamMemberInput := *openapiclient.NewInviteTeamMemberInput("Email_example", "Role_example") // InviteTeamMemberInput | Invitee email and role to assign on acceptance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.ApiV1TeamMembersInvitesPost(context.Background()).InviteTeamMemberInput(inviteTeamMemberInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.ApiV1TeamMembersInvitesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamMembersInvitesPost`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.ApiV1TeamMembersInvitesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamMembersInvitesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteTeamMemberInput** | [**InviteTeamMemberInput**](InviteTeamMemberInput.md) | Invitee email and role to assign on acceptance | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

