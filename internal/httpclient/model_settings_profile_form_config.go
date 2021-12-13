/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: v0.8.1-alpha.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SettingsProfileFormConfig struct for SettingsProfileFormConfig
type SettingsProfileFormConfig struct {
	// Action should be used as the form action URL `<form action=\"{{ .Action }}\" method=\"post\">`.
	Action   string   `json:"action"`
	Messages []UiText `json:"messages,omitempty"`
	// Method is the form method (e.g. POST)
	Method string   `json:"method"`
	Nodes  []UiNode `json:"nodes"`
}

// NewSettingsProfileFormConfig instantiates a new SettingsProfileFormConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsProfileFormConfig(action string, method string, nodes []UiNode) *SettingsProfileFormConfig {
	this := SettingsProfileFormConfig{}
	this.Action = action
	this.Method = method
	this.Nodes = nodes
	return &this
}

// NewSettingsProfileFormConfigWithDefaults instantiates a new SettingsProfileFormConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsProfileFormConfigWithDefaults() *SettingsProfileFormConfig {
	this := SettingsProfileFormConfig{}
	return &this
}

// GetAction returns the Action field value
func (o *SettingsProfileFormConfig) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *SettingsProfileFormConfig) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *SettingsProfileFormConfig) SetAction(v string) {
	o.Action = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *SettingsProfileFormConfig) GetMessages() []UiText {
	if o == nil || o.Messages == nil {
		var ret []UiText
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsProfileFormConfig) GetMessagesOk() ([]UiText, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *SettingsProfileFormConfig) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []UiText and assigns it to the Messages field.
func (o *SettingsProfileFormConfig) SetMessages(v []UiText) {
	o.Messages = v
}

// GetMethod returns the Method field value
func (o *SettingsProfileFormConfig) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SettingsProfileFormConfig) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SettingsProfileFormConfig) SetMethod(v string) {
	o.Method = v
}

// GetNodes returns the Nodes field value
func (o *SettingsProfileFormConfig) GetNodes() []UiNode {
	if o == nil {
		var ret []UiNode
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *SettingsProfileFormConfig) GetNodesOk() ([]UiNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *SettingsProfileFormConfig) SetNodes(v []UiNode) {
	o.Nodes = v
}

func (o SettingsProfileFormConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["nodes"] = o.Nodes
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsProfileFormConfig struct {
	value *SettingsProfileFormConfig
	isSet bool
}

func (v NullableSettingsProfileFormConfig) Get() *SettingsProfileFormConfig {
	return v.value
}

func (v *NullableSettingsProfileFormConfig) Set(val *SettingsProfileFormConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsProfileFormConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsProfileFormConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsProfileFormConfig(val *SettingsProfileFormConfig) *NullableSettingsProfileFormConfig {
	return &NullableSettingsProfileFormConfig{value: val, isSet: true}
}

func (v NullableSettingsProfileFormConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsProfileFormConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
