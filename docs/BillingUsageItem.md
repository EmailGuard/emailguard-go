# BillingUsageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]BillingUsageEvent**](BillingUsageEvent.md) |  | [optional] 
**FeatureKey** | **string** |  | 
**Limit** | Pointer to **int32** |  | [optional] 
**PeriodEnd** | **string** |  | 
**PeriodStart** | **string** |  | 
**Usage** | **int32** |  | 

## Methods

### NewBillingUsageItem

`func NewBillingUsageItem(featureKey string, periodEnd string, periodStart string, usage int32, ) *BillingUsageItem`

NewBillingUsageItem instantiates a new BillingUsageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingUsageItemWithDefaults

`func NewBillingUsageItemWithDefaults() *BillingUsageItem`

NewBillingUsageItemWithDefaults instantiates a new BillingUsageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *BillingUsageItem) GetEvents() []BillingUsageEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BillingUsageItem) GetEventsOk() (*[]BillingUsageEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BillingUsageItem) SetEvents(v []BillingUsageEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *BillingUsageItem) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFeatureKey

`func (o *BillingUsageItem) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *BillingUsageItem) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *BillingUsageItem) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetLimit

`func (o *BillingUsageItem) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BillingUsageItem) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BillingUsageItem) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BillingUsageItem) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *BillingUsageItem) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BillingUsageItem) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BillingUsageItem) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPeriodStart

`func (o *BillingUsageItem) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BillingUsageItem) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BillingUsageItem) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetUsage

`func (o *BillingUsageItem) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BillingUsageItem) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BillingUsageItem) SetUsage(v int32)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


