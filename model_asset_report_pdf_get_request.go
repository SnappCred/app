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

// AssetReportPDFGetRequest AssetReportPDFGetRequest defines the request schema for `/asset_report/pdf/get`
type AssetReportPDFGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A token that can be provided to endpoints such as `/asset_report/get` or `/asset_report/pdf/get` to fetch or update an Asset Report.
	AssetReportToken string `json:"asset_report_token"`
	Options *AssetReportPDFGetRequestOptions `json:"options,omitempty"`
}

// NewAssetReportPDFGetRequest instantiates a new AssetReportPDFGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportPDFGetRequest(assetReportToken string) *AssetReportPDFGetRequest {
	this := AssetReportPDFGetRequest{}
	this.AssetReportToken = assetReportToken
	return &this
}

// NewAssetReportPDFGetRequestWithDefaults instantiates a new AssetReportPDFGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportPDFGetRequestWithDefaults() *AssetReportPDFGetRequest {
	this := AssetReportPDFGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportPDFGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportPDFGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportPDFGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportPDFGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportPDFGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportPDFGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportPDFGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportPDFGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAssetReportToken returns the AssetReportToken field value
func (o *AssetReportPDFGetRequest) GetAssetReportToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportPDFGetRequest) GetAssetReportTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportToken, true
}

// SetAssetReportToken sets field value
func (o *AssetReportPDFGetRequest) SetAssetReportToken(v string) {
	o.AssetReportToken = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AssetReportPDFGetRequest) GetOptions() AssetReportPDFGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret AssetReportPDFGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportPDFGetRequest) GetOptionsOk() (*AssetReportPDFGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AssetReportPDFGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AssetReportPDFGetRequestOptions and assigns it to the Options field.
func (o *AssetReportPDFGetRequest) SetOptions(v AssetReportPDFGetRequestOptions) {
	o.Options = &v
}

func (o AssetReportPDFGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportPDFGetRequest struct {
	value *AssetReportPDFGetRequest
	isSet bool
}

func (v NullableAssetReportPDFGetRequest) Get() *AssetReportPDFGetRequest {
	return v.value
}

func (v *NullableAssetReportPDFGetRequest) Set(val *AssetReportPDFGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportPDFGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportPDFGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportPDFGetRequest(val *AssetReportPDFGetRequest) *NullableAssetReportPDFGetRequest {
	return &NullableAssetReportPDFGetRequest{value: val, isSet: true}
}

func (v NullableAssetReportPDFGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportPDFGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

