# \MembersAPI

All URIs are relative to *http://api.emailguard.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TeamMembersGet**](MembersAPI.md#ApiV1TeamMembersGet) | **Get** /api/v1/team/members | List team members
[**ApiV1TeamMembersUserIdDelete**](MembersAPI.md#ApiV1TeamMembersUserIdDelete) | **Delete** /api/v1/team/members/{userId} | Remove a team member
[**ApiV1TeamMembersUserIdPatch**](MembersAPI.md#ApiV1TeamMembersUserIdPatch) | **Patch** /api/v1/team/members/{userId} | Update a team member role



## ApiV1TeamMembersGet

> ListPublicTeamMembersResponse ApiV1TeamMembersGet(ctx).Execute()

List team members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EmailGuard/emailguard-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembersAPI.ApiV1TeamMembersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ApiV1TeamMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamMembersGet`: ListPublicTeamMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ApiV1TeamMembersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamMembersGetRequest struct via the builder pattern


### Return type

[**ListPublicTeamMembersResponse**](ListPublicTeamMembersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TeamMembersUserIdDelete

> MessageResponse ApiV1TeamMembersUserIdDelete(ctx, userId).Execute()

Remove a team member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EmailGuard/emailguard-go"
)

func main() {
	userId := "userId_example" // string | User account UUID of the member to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembersAPI.ApiV1TeamMembersUserIdDelete(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ApiV1TeamMembersUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamMembersUserIdDelete`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ApiV1TeamMembersUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User account UUID of the member to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamMembersUserIdDeleteRequest struct via the builder pattern


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


## ApiV1TeamMembersUserIdPatch

> MessageResponse ApiV1TeamMembersUserIdPatch(ctx, userId).UpdateTeamMemberInput(updateTeamMemberInput).Execute()

Update a team member role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EmailGuard/emailguard-go"
)

func main() {
	userId := "userId_example" // string | User account UUID of the member to update
	updateTeamMemberInput := *openapiclient.NewUpdateTeamMemberInput("Id_example") // UpdateTeamMemberInput | New role and/or notes for the member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembersAPI.ApiV1TeamMembersUserIdPatch(context.Background(), userId).UpdateTeamMemberInput(updateTeamMemberInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.ApiV1TeamMembersUserIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamMembersUserIdPatch`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MembersAPI.ApiV1TeamMembersUserIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User account UUID of the member to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamMembersUserIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTeamMemberInput** | [**UpdateTeamMemberInput**](UpdateTeamMemberInput.md) | New role and/or notes for the member | 

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

