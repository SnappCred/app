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

// AccountIdentityMatchScoreAllOf struct for AccountIdentityMatchScoreAllOf
type AccountIdentityMatchScoreAllOf struct {
	LegalName NullableNameMatchScore `json:"legal_name,omitempty"`
	PhoneNumber NullablePhoneNumberMatchScore `json:"phone_number,omitempty"`
	EmailAddress NullableEmailAddressMatchScore `json:"email_address,omitempty"`
	Address NullableAddressMatchScore `json:"address,omitempty"`
}

// NewAccountIdentityMatchScoreAllOf instantiates a new AccountIdentityMatchScoreAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountIdentityMatchScoreAllOf() *AccountIdentityMatchScoreAllOf {
	this := AccountIdentityMatchScoreAllOf{}
	return &this
}

// NewAccountIdentityMatchScoreAllOfWithDefaults instantiates a new AccountIdentityMatchScoreAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountIdentityMatchScoreAllOfWithDefaults() *AccountIdentityMatchScoreAllOf {
	this := AccountIdentityMatchScoreAllOf{}
	return &this
}

// GetLegalName returns the LegalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityMatchScoreAllOf) GetLegalName() NameMatchScore {
	if o == nil || o.LegalName.Get() == nil {
		var ret NameMatchScore
		return ret
	}
	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScoreAllOf) GetLegalNameOk() (*NameMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// HasLegalName returns a boolean if a field has been set.
func (o *AccountIdentityMatchScoreAllOf) HasLegalName() bool {
	if o != nil && o.LegalName.IsSet() {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given NullableNameMatchScore and assigns it to the LegalName field.
func (o *AccountIdentityMatchScoreAllOf) SetLegalName(v NameMatchScore) {
	o.LegalName.Set(&v)
}
// SetLegalNameNil sets the value for LegalName to be an explicit nil
func (o *AccountIdentityMatchScoreAllOf) SetLegalNameNil() {
	o.LegalName.Set(nil)
}

// UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
func (o *AccountIdentityMatchScoreAllOf) UnsetLegalName() {
	o.LegalName.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityMatchScoreAllOf) GetPhoneNumber() PhoneNumberMatchScore {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret PhoneNumberMatchScore
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScoreAllOf) GetPhoneNumberOk() (*PhoneNumberMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AccountIdentityMatchScoreAllOf) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullablePhoneNumberMatchScore and assigns it to the PhoneNumber field.
func (o *AccountIdentityMatchScoreAllOf) SetPhoneNumber(v PhoneNumberMatchScore) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *AccountIdentityMatchScoreAllOf) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *AccountIdentityMatchScoreAllOf) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityMatchScoreAllOf) GetEmailAddress() EmailAddressMatchScore {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret EmailAddressMatchScore
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScoreAllOf) GetEmailAddressOk() (*EmailAddressMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *AccountIdentityMatchScoreAllOf) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableEmailAddressMatchScore and assigns it to the EmailAddress field.
func (o *AccountIdentityMatchScoreAllOf) SetEmailAddress(v EmailAddressMatchScore) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *AccountIdentityMatchScoreAllOf) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *AccountIdentityMatchScoreAllOf) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityMatchScoreAllOf) GetAddress() AddressMatchScore {
	if o == nil || o.Address.Get() == nil {
		var ret AddressMatchScore
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityMatchScoreAllOf) GetAddressOk() (*AddressMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *AccountIdentityMatchScoreAllOf) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableAddressMatchScore and assigns it to the Address field.
func (o *AccountIdentityMatchScoreAllOf) SetAddress(v AddressMatchScore) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *AccountIdentityMatchScoreAllOf) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *AccountIdentityMatchScoreAllOf) UnsetAddress() {
	o.Address.Unset()
}

func (o AccountIdentityMatchScoreAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LegalName.IsSet() {
		toSerialize["legal_name"] = o.LegalName.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAccountIdentityMatchScoreAllOf struct {
	value *AccountIdentityMatchScoreAllOf
	isSet bool
}

func (v NullableAccountIdentityMatchScoreAllOf) Get() *AccountIdentityMatchScoreAllOf {
	return v.value
}

func (v *NullableAccountIdentityMatchScoreAllOf) Set(val *AccountIdentityMatchScoreAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountIdentityMatchScoreAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountIdentityMatchScoreAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountIdentityMatchScoreAllOf(val *AccountIdentityMatchScoreAllOf) *NullableAccountIdentityMatchScoreAllOf {
	return &NullableAccountIdentityMatchScoreAllOf{value: val, isSet: true}
}

func (v NullableAccountIdentityMatchScoreAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountIdentityMatchScoreAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

