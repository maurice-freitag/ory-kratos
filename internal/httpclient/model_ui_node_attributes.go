/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: v0.8.0-alpha.4.pre.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// UiNodeAttributes - struct for UiNodeAttributes
type UiNodeAttributes struct {
	UiNodeAnchorAttributes *UiNodeAnchorAttributes
	UiNodeImageAttributes  *UiNodeImageAttributes
	UiNodeInputAttributes  *UiNodeInputAttributes
	UiNodeScriptAttributes *UiNodeScriptAttributes
	UiNodeTextAttributes   *UiNodeTextAttributes
}

// UiNodeAnchorAttributesAsUiNodeAttributes is a convenience function that returns UiNodeAnchorAttributes wrapped in UiNodeAttributes
func UiNodeAnchorAttributesAsUiNodeAttributes(v *UiNodeAnchorAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeAnchorAttributes: v,
	}
}

// UiNodeImageAttributesAsUiNodeAttributes is a convenience function that returns UiNodeImageAttributes wrapped in UiNodeAttributes
func UiNodeImageAttributesAsUiNodeAttributes(v *UiNodeImageAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeImageAttributes: v,
	}
}

// UiNodeInputAttributesAsUiNodeAttributes is a convenience function that returns UiNodeInputAttributes wrapped in UiNodeAttributes
func UiNodeInputAttributesAsUiNodeAttributes(v *UiNodeInputAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeInputAttributes: v,
	}
}

// UiNodeScriptAttributesAsUiNodeAttributes is a convenience function that returns UiNodeScriptAttributes wrapped in UiNodeAttributes
func UiNodeScriptAttributesAsUiNodeAttributes(v *UiNodeScriptAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeScriptAttributes: v,
	}
}

// UiNodeTextAttributesAsUiNodeAttributes is a convenience function that returns UiNodeTextAttributes wrapped in UiNodeAttributes
func UiNodeTextAttributesAsUiNodeAttributes(v *UiNodeTextAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeTextAttributes: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UiNodeAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UiNodeAnchorAttributes
	err = newStrictDecoder(data).Decode(&dst.UiNodeAnchorAttributes)
	if err == nil {
		jsonUiNodeAnchorAttributes, _ := json.Marshal(dst.UiNodeAnchorAttributes)
		if string(jsonUiNodeAnchorAttributes) == "{}" { // empty struct
			dst.UiNodeAnchorAttributes = nil
		} else {
			match++
		}
	} else {
		dst.UiNodeAnchorAttributes = nil
	}

	// try to unmarshal data into UiNodeImageAttributes
	err = newStrictDecoder(data).Decode(&dst.UiNodeImageAttributes)
	if err == nil {
		jsonUiNodeImageAttributes, _ := json.Marshal(dst.UiNodeImageAttributes)
		if string(jsonUiNodeImageAttributes) == "{}" { // empty struct
			dst.UiNodeImageAttributes = nil
		} else {
			match++
		}
	} else {
		dst.UiNodeImageAttributes = nil
	}

	// try to unmarshal data into UiNodeInputAttributes
	err = newStrictDecoder(data).Decode(&dst.UiNodeInputAttributes)
	if err == nil {
		jsonUiNodeInputAttributes, _ := json.Marshal(dst.UiNodeInputAttributes)
		if string(jsonUiNodeInputAttributes) == "{}" { // empty struct
			dst.UiNodeInputAttributes = nil
		} else {
			match++
		}
	} else {
		dst.UiNodeInputAttributes = nil
	}

	// try to unmarshal data into UiNodeScriptAttributes
	err = newStrictDecoder(data).Decode(&dst.UiNodeScriptAttributes)
	if err == nil {
		jsonUiNodeScriptAttributes, _ := json.Marshal(dst.UiNodeScriptAttributes)
		if string(jsonUiNodeScriptAttributes) == "{}" { // empty struct
			dst.UiNodeScriptAttributes = nil
		} else {
			match++
		}
	} else {
		dst.UiNodeScriptAttributes = nil
	}

	// try to unmarshal data into UiNodeTextAttributes
	err = newStrictDecoder(data).Decode(&dst.UiNodeTextAttributes)
	if err == nil {
		jsonUiNodeTextAttributes, _ := json.Marshal(dst.UiNodeTextAttributes)
		if string(jsonUiNodeTextAttributes) == "{}" { // empty struct
			dst.UiNodeTextAttributes = nil
		} else {
			match++
		}
	} else {
		dst.UiNodeTextAttributes = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UiNodeAnchorAttributes = nil
		dst.UiNodeImageAttributes = nil
		dst.UiNodeInputAttributes = nil
		dst.UiNodeScriptAttributes = nil
		dst.UiNodeTextAttributes = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(UiNodeAttributes)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(UiNodeAttributes)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UiNodeAttributes) MarshalJSON() ([]byte, error) {
	if src.UiNodeAnchorAttributes != nil {
		return json.Marshal(&src.UiNodeAnchorAttributes)
	}

	if src.UiNodeImageAttributes != nil {
		return json.Marshal(&src.UiNodeImageAttributes)
	}

	if src.UiNodeInputAttributes != nil {
		return json.Marshal(&src.UiNodeInputAttributes)
	}

	if src.UiNodeScriptAttributes != nil {
		return json.Marshal(&src.UiNodeScriptAttributes)
	}

	if src.UiNodeTextAttributes != nil {
		return json.Marshal(&src.UiNodeTextAttributes)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UiNodeAttributes) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UiNodeAnchorAttributes != nil {
		return obj.UiNodeAnchorAttributes
	}

	if obj.UiNodeImageAttributes != nil {
		return obj.UiNodeImageAttributes
	}

	if obj.UiNodeInputAttributes != nil {
		return obj.UiNodeInputAttributes
	}

	if obj.UiNodeScriptAttributes != nil {
		return obj.UiNodeScriptAttributes
	}

	if obj.UiNodeTextAttributes != nil {
		return obj.UiNodeTextAttributes
	}

	// all schemas are nil
	return nil
}

type NullableUiNodeAttributes struct {
	value *UiNodeAttributes
	isSet bool
}

func (v NullableUiNodeAttributes) Get() *UiNodeAttributes {
	return v.value
}

func (v *NullableUiNodeAttributes) Set(val *UiNodeAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeAttributes(val *UiNodeAttributes) *NullableUiNodeAttributes {
	return &NullableUiNodeAttributes{value: val, isSet: true}
}

func (v NullableUiNodeAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
