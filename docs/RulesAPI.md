# \RulesAPI

All URIs are relative to *http://api.emailguard.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TeamNotificationsRulesGet**](RulesAPI.md#ApiV1TeamNotificationsRulesGet) | **Get** /api/v1/team/notifications/rules | List notification rules
[**ApiV1TeamNotificationsRulesPost**](RulesAPI.md#ApiV1TeamNotificationsRulesPost) | **Post** /api/v1/team/notifications/rules | Create notification rule
[**ApiV1TeamNotificationsRulesRuleIdDelete**](RulesAPI.md#ApiV1TeamNotificationsRulesRuleIdDelete) | **Delete** /api/v1/team/notifications/rules/{ruleId} | Delete notification rule
[**ApiV1TeamNotificationsRulesRuleIdPut**](RulesAPI.md#ApiV1TeamNotificationsRulesRuleIdPut) | **Put** /api/v1/team/notifications/rules/{ruleId} | Update notification rule



## ApiV1TeamNotificationsRulesGet

> GetNotificationRulesResponse ApiV1TeamNotificationsRulesGet(ctx).Execute()

List notification rules



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
	resp, r, err := apiClient.RulesAPI.ApiV1TeamNotificationsRulesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiV1TeamNotificationsRulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsRulesGet`: GetNotificationRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiV1TeamNotificationsRulesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsRulesGetRequest struct via the builder pattern


### Return type

[**GetNotificationRulesResponse**](GetNotificationRulesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TeamNotificationsRulesPost

> IdDataResponse ApiV1TeamNotificationsRulesPost(ctx).CreateNotificationRuleInput(createNotificationRuleInput).Execute()

Create notification rule



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
	createNotificationRuleInput := *openapiclient.NewCreateNotificationRuleInput([]openapiclient.NotificationRuleChannelInput{*openapiclient.NewNotificationRuleChannelInput()}, []string{"EventTypes_example"}, "Name_example") // CreateNotificationRuleInput | Rule name, event types, and channel targets

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.ApiV1TeamNotificationsRulesPost(context.Background()).CreateNotificationRuleInput(createNotificationRuleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiV1TeamNotificationsRulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsRulesPost`: IdDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiV1TeamNotificationsRulesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsRulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationRuleInput** | [**CreateNotificationRuleInput**](CreateNotificationRuleInput.md) | Rule name, event types, and channel targets | 

### Return type

[**IdDataResponse**](IdDataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TeamNotificationsRulesRuleIdDelete

> MessageResponse ApiV1TeamNotificationsRulesRuleIdDelete(ctx, ruleId).Execute()

Delete notification rule



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
	ruleId := "ruleId_example" // string | UUID of the notification rule to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.ApiV1TeamNotificationsRulesRuleIdDelete(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiV1TeamNotificationsRulesRuleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsRulesRuleIdDelete`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiV1TeamNotificationsRulesRuleIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | UUID of the notification rule to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsRulesRuleIdDeleteRequest struct via the builder pattern


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


## ApiV1TeamNotificationsRulesRuleIdPut

> MessageResponse ApiV1TeamNotificationsRulesRuleIdPut(ctx, ruleId).UpdateNotificationRuleInput(updateNotificationRuleInput).Execute()

Update notification rule



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
	ruleId := "ruleId_example" // string | UUID of the notification rule to update
	updateNotificationRuleInput := *openapiclient.NewUpdateNotificationRuleInput() // UpdateNotificationRuleInput | Fields to update on the rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.ApiV1TeamNotificationsRulesRuleIdPut(context.Background(), ruleId).UpdateNotificationRuleInput(updateNotificationRuleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiV1TeamNotificationsRulesRuleIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsRulesRuleIdPut`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiV1TeamNotificationsRulesRuleIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | UUID of the notification rule to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsRulesRuleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNotificationRuleInput** | [**UpdateNotificationRuleInput**](UpdateNotificationRuleInput.md) | Fields to update on the rule | 

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

