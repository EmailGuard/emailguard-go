# \EmailAPI

All URIs are relative to *http://api.emailguard.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1EmailsDetectGet**](EmailAPI.md#ApiV1EmailsDetectGet) | **Get** /api/v1/emails/detect | Detect email characteristics



## ApiV1EmailsDetectGet

> EmailDetectResponse ApiV1EmailsDetectGet(ctx).Email(email).Execute()

Detect email characteristics



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
	email := "email_example" // string | Email address to analyze

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.ApiV1EmailsDetectGet(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.ApiV1EmailsDetectGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1EmailsDetectGet`: EmailDetectResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.ApiV1EmailsDetectGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EmailsDetectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email address to analyze | 

### Return type

[**EmailDetectResponse**](EmailDetectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

