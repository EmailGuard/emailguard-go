# \TeamsAPI

All URIs are relative to *http://api.emailguard.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TeamGet**](TeamsAPI.md#ApiV1TeamGet) | **Get** /api/v1/team | Get team
[**ApiV1TeamPatch**](TeamsAPI.md#ApiV1TeamPatch) | **Patch** /api/v1/team | Update team



## ApiV1TeamGet

> GetPublicTeamResponse ApiV1TeamGet(ctx).Execute()

Get team



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
	resp, r, err := apiClient.TeamsAPI.ApiV1TeamGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.ApiV1TeamGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamGet`: GetPublicTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.ApiV1TeamGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamGetRequest struct via the builder pattern


### Return type

[**GetPublicTeamResponse**](GetPublicTeamResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TeamPatch

> GetPublicTeamResponse ApiV1TeamPatch(ctx).UpdateTeamInput(updateTeamInput).Execute()

Update team



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
	updateTeamInput := *openapiclient.NewUpdateTeamInput("Name_example") // UpdateTeamInput | Team profile fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.ApiV1TeamPatch(context.Background()).UpdateTeamInput(updateTeamInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.ApiV1TeamPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamPatch`: GetPublicTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.ApiV1TeamPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTeamInput** | [**UpdateTeamInput**](UpdateTeamInput.md) | Team profile fields to update | 

### Return type

[**GetPublicTeamResponse**](GetPublicTeamResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

