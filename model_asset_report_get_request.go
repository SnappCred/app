/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.391.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AssetReportGetRequest AssetReportGetRequest defines the request schema for `/asset_report/get`
type AssetReportGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A token that can be provided to endpoints such as `/asset_report/get` or `/asset_report/pdf/get` to fetch or update an Asset Report.
	AssetReportToken *string `json:"asset_report_token,omitempty"`
	// The user token associated with the User for which to create an asset report for. The latest asset report associated with the User will be returned
	UserToken *string `json:"user_token,omitempty"`
	// `true` if you would like to retrieve the Asset Report with Insights, `false` otherwise. This field defaults to `false` if omitted.
	IncludeInsights *bool `json:"include_insights,omitempty"`
	// `true` to fetch \"fast\" version of asset report. Defaults to false if omitted. Can only be used if `/asset_report/create` was called with `options.add_ons` set to `[\"fast_assets\"]`.
	FastReport *bool `json:"fast_report,omitempty"`
	Options *AssetReportGetRequestOptions `json:"options,omitempty"`
}

// NewAssetReportGetRequest instantiates a new AssetReportGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportGetRequest() *AssetReportGetRequest {
	this := AssetReportGetRequest{}
	var includeInsights bool = false
	this.IncludeInsights = &includeInsights
	var fastReport bool = false
	this.FastReport = &fastReport
	return &this
}

// NewAssetReportGetRequestWithDefaults instantiates a new AssetReportGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportGetRequestWithDefaults() *AssetReportGetRequest {
	this := AssetReportGetRequest{}
	var includeInsights bool = false
	this.IncludeInsights = &includeInsights
	var fastReport bool = false
	this.FastReport = &fastReport
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAssetReportToken returns the AssetReportToken field value if set, zero value otherwise.
func (o *AssetReportGetRequest) GetAssetReportToken() string {
	if o == nil || o.AssetReportToken == nil {
		var ret string
		return ret
	}
	return *o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportGetRequest) GetAssetReportTokenOk() (*string, bool) {
	if o == nil || o.AssetReportToken == nil {
		return nil, false
	}
	return o.AssetReportToken, true
}

// HasAssetReportToken returns a boolean if a field has been set.
func (o *AssetReportGetRequest) HasAssetReportToken() bool {
	if o != nil && o.AssetReportToken != nil {
		return true
	}

	return false
}

// SetAssetReportToken gets a reference to the given string and assigns it to the AssetReportToken field.
func (o *AssetReportGetRequest) SetAssetReportToken(v string) {
	o.AssetReportToken = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *AssetReportGetRequest) GetUserToken() string {
	if o == nil || o.UserToken == nil {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil || o.UserToken == nil {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *AssetReportGetRequest) HasUserToken() bool {
	if o != nil && o.UserToken != nil {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *AssetReportGetRequest) SetUserToken(v string) {
	o.UserToken = &v
}

// GetIncludeInsights returns the IncludeInsights field value if set, zero value otherwise.
func (o *AssetReportGetRequest) GetIncludeInsights() bool {
	if o == nil || o.IncludeInsights == nil {
		var ret bool
		return ret
	}
	return *o.IncludeInsights
}

// GetIncludeInsightsOk returns a tuple with the IncludeInsights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportGetRequest) GetIncludeInsightsOk() (*bool, bool) {
	if o == nil || o.IncludeInsights == nil {
		return nil, false
	}
	return o.IncludeInsights, true
}

// HasIncludeInsights returns a boolean if a field has been set.
func (o *AssetReportGetRequest) HasIncludeInsights() bool {
	if o != nil && o.IncludeInsights != nil {
		return true
	}

	return false
}

// SetIncludeInsights gets a reference to the given bool and assigns it to the IncludeInsights field.
func (o *AssetReportGetRequest) SetIncludeInsights(v bool) {
	o.IncludeInsights = &v
}

// GetFastReport returns the FastReport field value if set, zero value otherwise.
func (o *AssetReportGetRequest) GetFastReport() bool {
	if o == nil || o.FastReport == nil {
		var ret bool
		return ret
	}
	return *o.FastReport
}

// GetFastReportOk returns a tuple with the FastReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportGetRequest) GetFastReportOk() (*bool, bool) {
	if o == nil || o.FastReport == nil {
		return nil, false
	}
	return o.FastReport, true
}

// HasFastReport returns a boolean if a field has been set.
func (o *AssetReportGetRequest) HasFastReport() bool {
	if o != nil && o.FastReport != nil {
		return true
	}

	return false
}

// SetFastReport gets a reference to the given bool and assigns it to the FastReport field.
func (o *AssetReportGetRequest) SetFastReport(v bool) {
	o.FastReport = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AssetReportGetRequest) GetOptions() AssetReportGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret AssetReportGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportGetRequest) GetOptionsOk() (*AssetReportGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AssetReportGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AssetReportGetRequestOptions and assigns it to the Options field.
func (o *AssetReportGetRequest) SetOptions(v AssetReportGetRequestOptions) {
	o.Options = &v
}

func (o AssetReportGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.AssetReportToken != nil {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if o.UserToken != nil {
		toSerialize["user_token"] = o.UserToken
	}
	if o.IncludeInsights != nil {
		toSerialize["include_insights"] = o.IncludeInsights
	}
	if o.FastReport != nil {
		toSerialize["fast_report"] = o.FastReport
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportGetRequest struct {
	value *AssetReportGetRequest
	isSet bool
}

func (v NullableAssetReportGetRequest) Get() *AssetReportGetRequest {
	return v.value
}

func (v *NullableAssetReportGetRequest) Set(val *AssetReportGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportGetRequest(val *AssetReportGetRequest) *NullableAssetReportGetRequest {
	return &NullableAssetReportGetRequest{value: val, isSet: true}
}

func (v NullableAssetReportGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

