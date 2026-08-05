# BillingUsageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delta** | **int32** |  | 
**EventType** | **string** |  | 
**Metadata** | **string** |  | 
**UsageAfter** | **int32** |  | 
**UsageBefore** | **int32** |  | 

## Methods

### NewBillingUsageEvent

`func NewBillingUsageEvent(delta int32, eventType string, metadata string, usageAfter int32, usageBefore int32, ) *BillingUsageEvent`

NewBillingUsageEvent instantiates a new BillingUsageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingUsageEventWithDefaults

`func NewBillingUsageEventWithDefaults() *BillingUsageEvent`

NewBillingUsageEventWithDefaults instantiates a new BillingUsageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelta

`func (o *BillingUsageEvent) GetDelta() int32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *BillingUsageEvent) GetDeltaOk() (*int32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *BillingUsageEvent) SetDelta(v int32)`

SetDelta sets Delta field to given value.


### GetEventType

`func (o *BillingUsageEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BillingUsageEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BillingUsageEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMetadata

`func (o *BillingUsageEvent) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingUsageEvent) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingUsageEvent) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetUsageAfter

`func (o *BillingUsageEvent) GetUsageAfter() int32`

GetUsageAfter returns the UsageAfter field if non-nil, zero value otherwise.

### GetUsageAfterOk

`func (o *BillingUsageEvent) GetUsageAfterOk() (*int32, bool)`

GetUsageAfterOk returns a tuple with the UsageAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAfter

`func (o *BillingUsageEvent) SetUsageAfter(v int32)`

SetUsageAfter sets UsageAfter field to given value.


### GetUsageBefore

`func (o *BillingUsageEvent) GetUsageBefore() int32`

GetUsageBefore returns the UsageBefore field if non-nil, zero value otherwise.

### GetUsageBeforeOk

`func (o *BillingUsageEvent) GetUsageBeforeOk() (*int32, bool)`

GetUsageBeforeOk returns a tuple with the UsageBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBefore

`func (o *BillingUsageEvent) SetUsageBefore(v int32)`

SetUsageBefore sets UsageBefore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


