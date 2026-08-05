# \ChannelsAPI

All URIs are relative to *http://api.emailguard.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TeamNotificationsChannelsChannelIdDelete**](ChannelsAPI.md#ApiV1TeamNotificationsChannelsChannelIdDelete) | **Delete** /api/v1/team/notifications/channels/{channelId} | Delete notification channel
[**ApiV1TeamNotificationsChannelsChannelIdPut**](ChannelsAPI.md#ApiV1TeamNotificationsChannelsChannelIdPut) | **Put** /api/v1/team/notifications/channels/{channelId} | Update notification channel
[**ApiV1TeamNotificationsChannelsChannelIdTestPost**](ChannelsAPI.md#ApiV1TeamNotificationsChannelsChannelIdTestPost) | **Post** /api/v1/team/notifications/channels/{channelId}/test | Test notification channel
[**ApiV1TeamNotificationsChannelsGet**](ChannelsAPI.md#ApiV1TeamNotificationsChannelsGet) | **Get** /api/v1/team/notifications/channels | List notification channels
[**ApiV1TeamNotificationsChannelsPost**](ChannelsAPI.md#ApiV1TeamNotificationsChannelsPost) | **Post** /api/v1/team/notifications/channels | Create notification channel



## ApiV1TeamNotificationsChannelsChannelIdDelete

> MessageResponse ApiV1TeamNotificationsChannelsChannelIdDelete(ctx, channelId).Execute()

Delete notification channel



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
	channelId := "channelId_example" // string | UUID of the notification channel to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdDelete(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsChannelsChannelIdDelete`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | UUID of the notification channel to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsChannelsChannelIdDeleteRequest struct via the builder pattern


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


## ApiV1TeamNotificationsChannelsChannelIdPut

> GetNotificationChannelResponse ApiV1TeamNotificationsChannelsChannelIdPut(ctx, channelId).UpdateNotificationChannelInput(updateNotificationChannelInput).Execute()

Update notification channel



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
	channelId := "channelId_example" // string | UUID of the notification channel to update
	updateNotificationChannelInput := *openapiclient.NewUpdateNotificationChannelInput() // UpdateNotificationChannelInput | Updated channel name and/or configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdPut(context.Background(), channelId).UpdateNotificationChannelInput(updateNotificationChannelInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsChannelsChannelIdPut`: GetNotificationChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | UUID of the notification channel to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsChannelsChannelIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNotificationChannelInput** | [**UpdateNotificationChannelInput**](UpdateNotificationChannelInput.md) | Updated channel name and/or configuration | 

### Return type

[**GetNotificationChannelResponse**](GetNotificationChannelResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TeamNotificationsChannelsChannelIdTestPost

> MessageResponse ApiV1TeamNotificationsChannelsChannelIdTestPost(ctx, channelId).Execute()

Test notification channel



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
	channelId := "channelId_example" // string | UUID of the webhook channel to test

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdTestPost(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsChannelsChannelIdTestPost`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.ApiV1TeamNotificationsChannelsChannelIdTestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | UUID of the webhook channel to test | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsChannelsChannelIdTestPostRequest struct via the builder pattern


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


## ApiV1TeamNotificationsChannelsGet

> GetNotificationChannelsResponse ApiV1TeamNotificationsChannelsGet(ctx).Execute()

List notification channels



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
	resp, r, err := apiClient.ChannelsAPI.ApiV1TeamNotificationsChannelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.ApiV1TeamNotificationsChannelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsChannelsGet`: GetNotificationChannelsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.ApiV1TeamNotificationsChannelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsChannelsGetRequest struct via the builder pattern


### Return type

[**GetNotificationChannelsResponse**](GetNotificationChannelsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TeamNotificationsChannelsPost

> GetNotificationChannelResponse ApiV1TeamNotificationsChannelsPost(ctx).CreateNotificationChannelInput(createNotificationChannelInput).Execute()

Create notification channel



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
	createNotificationChannelInput := *openapiclient.NewCreateNotificationChannelInput("Name_example", "Type_example") // CreateNotificationChannelInput | Channel name, type, and type-specific configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.ApiV1TeamNotificationsChannelsPost(context.Background()).CreateNotificationChannelInput(createNotificationChannelInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.ApiV1TeamNotificationsChannelsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsChannelsPost`: GetNotificationChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.ApiV1TeamNotificationsChannelsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsChannelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationChannelInput** | [**CreateNotificationChannelInput**](CreateNotificationChannelInput.md) | Channel name, type, and type-specific configuration | 

### Return type

[**GetNotificationChannelResponse**](GetNotificationChannelResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

