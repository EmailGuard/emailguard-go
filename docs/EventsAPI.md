# \EventsAPI

All URIs are relative to *http://api.emailguard.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TeamNotificationsEventsGet**](EventsAPI.md#ApiV1TeamNotificationsEventsGet) | **Get** /api/v1/team/notifications/events | List notification events



## ApiV1TeamNotificationsEventsGet

> GetNotificationEventsResponse ApiV1TeamNotificationsEventsGet(ctx).Execute()

List notification events



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
	resp, r, err := apiClient.EventsAPI.ApiV1TeamNotificationsEventsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ApiV1TeamNotificationsEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamNotificationsEventsGet`: GetNotificationEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ApiV1TeamNotificationsEventsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamNotificationsEventsGetRequest struct via the builder pattern


### Return type

[**GetNotificationEventsResponse**](GetNotificationEventsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

