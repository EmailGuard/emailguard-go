# \BillingAPI

All URIs are relative to *http://api.emailguard.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TeamBillingUsageGet**](BillingAPI.md#ApiV1TeamBillingUsageGet) | **Get** /api/v1/team/billing/usage | Get billing usage



## ApiV1TeamBillingUsageGet

> BillingUsageResponse ApiV1TeamBillingUsageGet(ctx).Range_(range_).Start(start).End(end).Execute()

Get billing usage



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
	range_ := "range__example" // string | Period selector: current, previous, or custom (optional)
	start := "start_example" // string | Custom range start (RFC3339 or YYYY-MM-DD); required when range=custom (optional)
	end := "end_example" // string | Custom range end (RFC3339 or YYYY-MM-DD); required when range=custom (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.ApiV1TeamBillingUsageGet(context.Background()).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ApiV1TeamBillingUsageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TeamBillingUsageGet`: BillingUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ApiV1TeamBillingUsageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TeamBillingUsageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Period selector: current, previous, or custom | 
 **start** | **string** | Custom range start (RFC3339 or YYYY-MM-DD); required when range&#x3D;custom | 
 **end** | **string** | Custom range end (RFC3339 or YYYY-MM-DD); required when range&#x3D;custom | 

### Return type

[**BillingUsageResponse**](BillingUsageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

