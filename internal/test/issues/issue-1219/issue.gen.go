// Package issue1219 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/stefanobaghino/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package issue1219

import (
	"encoding/json"
	"fmt"
)

// DefaultAddtional1 defines model for DefaultAddtional1.
type DefaultAddtional1 struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
}

// DefaultAddtional2 defines model for DefaultAddtional2.
type DefaultAddtional2 struct {
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeDefaultDefault defines model for MergeDefaultDefault.
type MergeDefaultDefault struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeDefaultWithAny defines model for MergeDefaultWithAny.
type MergeDefaultWithAny struct {
	Field1               *int                   `json:"field1,omitempty"`
	Field2               *string                `json:"field2,omitempty"`
	FieldA               *int                   `json:"fieldA,omitempty"`
	FieldB               *string                `json:"fieldB,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MergeDefaultWithString defines model for MergeDefaultWithString.
type MergeDefaultWithString struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// MergeDefaultWithout defines model for MergeDefaultWithout.
type MergeDefaultWithout struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithAnyDefault defines model for MergeWithAnyDefault.
type MergeWithAnyDefault struct {
	Field1               *int                   `json:"field1,omitempty"`
	Field2               *string                `json:"field2,omitempty"`
	FieldA               *int                   `json:"fieldA,omitempty"`
	FieldB               *string                `json:"fieldB,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MergeWithAnyWithAny defines model for MergeWithAnyWithAny.
type MergeWithAnyWithAny struct {
	Field1               *int                   `json:"field1,omitempty"`
	Field2               *string                `json:"field2,omitempty"`
	FieldA               *int                   `json:"fieldA,omitempty"`
	FieldB               *string                `json:"fieldB,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MergeWithAnyWithString defines model for MergeWithAnyWithString.
type MergeWithAnyWithString struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// MergeWithAnyWithout defines model for MergeWithAnyWithout.
type MergeWithAnyWithout struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithStringDefault defines model for MergeWithStringDefault.
type MergeWithStringDefault struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// MergeWithStringWithAny defines model for MergeWithStringWithAny.
type MergeWithStringWithAny struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// MergeWithStringWithout defines model for MergeWithStringWithout.
type MergeWithStringWithout struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithoutDefault defines model for MergeWithoutDefault.
type MergeWithoutDefault struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithoutWithAny defines model for MergeWithoutWithAny.
type MergeWithoutWithAny struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithoutWithString defines model for MergeWithoutWithString.
type MergeWithoutWithString struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithoutWithout defines model for MergeWithoutWithout.
type MergeWithoutWithout struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// WithAnyAddtional1 defines model for WithAnyAddtional1.
type WithAnyAddtional1 struct {
	Field1               *int                   `json:"field1,omitempty"`
	Field2               *string                `json:"field2,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// WithAnyAddtional2 defines model for WithAnyAddtional2.
type WithAnyAddtional2 struct {
	FieldA               *int                   `json:"fieldA,omitempty"`
	FieldB               *string                `json:"fieldB,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// WithStringAddtional1 defines model for WithStringAddtional1.
type WithStringAddtional1 struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// WithStringAddtional2 defines model for WithStringAddtional2.
type WithStringAddtional2 struct {
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// WithoutAddtional1 defines model for WithoutAddtional1.
type WithoutAddtional1 struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
}

// WithoutAddtional2 defines model for WithoutAddtional2.
type WithoutAddtional2 struct {
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// Getter for additional properties for MergeDefaultWithAny. Returns the specified
// element and whether it was found
func (a MergeDefaultWithAny) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeDefaultWithAny
func (a *MergeDefaultWithAny) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeDefaultWithAny to handle AdditionalProperties
func (a *MergeDefaultWithAny) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeDefaultWithAny to handle AdditionalProperties
func (a MergeDefaultWithAny) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeDefaultWithString. Returns the specified
// element and whether it was found
func (a MergeDefaultWithString) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeDefaultWithString
func (a *MergeDefaultWithString) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeDefaultWithString to handle AdditionalProperties
func (a *MergeDefaultWithString) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeDefaultWithString to handle AdditionalProperties
func (a MergeDefaultWithString) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithAnyDefault. Returns the specified
// element and whether it was found
func (a MergeWithAnyDefault) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithAnyDefault
func (a *MergeWithAnyDefault) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithAnyDefault to handle AdditionalProperties
func (a *MergeWithAnyDefault) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithAnyDefault to handle AdditionalProperties
func (a MergeWithAnyDefault) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithAnyWithAny. Returns the specified
// element and whether it was found
func (a MergeWithAnyWithAny) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithAnyWithAny
func (a *MergeWithAnyWithAny) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithAnyWithAny to handle AdditionalProperties
func (a *MergeWithAnyWithAny) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithAnyWithAny to handle AdditionalProperties
func (a MergeWithAnyWithAny) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithAnyWithString. Returns the specified
// element and whether it was found
func (a MergeWithAnyWithString) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithAnyWithString
func (a *MergeWithAnyWithString) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithAnyWithString to handle AdditionalProperties
func (a *MergeWithAnyWithString) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithAnyWithString to handle AdditionalProperties
func (a MergeWithAnyWithString) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithStringDefault. Returns the specified
// element and whether it was found
func (a MergeWithStringDefault) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithStringDefault
func (a *MergeWithStringDefault) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithStringDefault to handle AdditionalProperties
func (a *MergeWithStringDefault) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithStringDefault to handle AdditionalProperties
func (a MergeWithStringDefault) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithStringWithAny. Returns the specified
// element and whether it was found
func (a MergeWithStringWithAny) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithStringWithAny
func (a *MergeWithStringWithAny) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithStringWithAny to handle AdditionalProperties
func (a *MergeWithStringWithAny) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithStringWithAny to handle AdditionalProperties
func (a MergeWithStringWithAny) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for WithAnyAddtional1. Returns the specified
// element and whether it was found
func (a WithAnyAddtional1) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for WithAnyAddtional1
func (a *WithAnyAddtional1) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for WithAnyAddtional1 to handle AdditionalProperties
func (a *WithAnyAddtional1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for WithAnyAddtional1 to handle AdditionalProperties
func (a WithAnyAddtional1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for WithAnyAddtional2. Returns the specified
// element and whether it was found
func (a WithAnyAddtional2) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for WithAnyAddtional2
func (a *WithAnyAddtional2) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for WithAnyAddtional2 to handle AdditionalProperties
func (a *WithAnyAddtional2) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for WithAnyAddtional2 to handle AdditionalProperties
func (a WithAnyAddtional2) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for WithStringAddtional1. Returns the specified
// element and whether it was found
func (a WithStringAddtional1) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for WithStringAddtional1
func (a *WithStringAddtional1) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for WithStringAddtional1 to handle AdditionalProperties
func (a *WithStringAddtional1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for WithStringAddtional1 to handle AdditionalProperties
func (a WithStringAddtional1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for WithStringAddtional2. Returns the specified
// element and whether it was found
func (a WithStringAddtional2) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for WithStringAddtional2
func (a *WithStringAddtional2) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for WithStringAddtional2 to handle AdditionalProperties
func (a *WithStringAddtional2) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for WithStringAddtional2 to handle AdditionalProperties
func (a WithStringAddtional2) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
