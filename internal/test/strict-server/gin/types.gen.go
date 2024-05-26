// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/stefanobaghino/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package api

// Example defines model for example.
type Example struct {
	Value *string `json:"value,omitempty"`
}

// Reusableresponse defines model for reusableresponse.
type Reusableresponse = Example

// MultipleRequestAndResponseTypesTextBody defines parameters for MultipleRequestAndResponseTypes.
type MultipleRequestAndResponseTypesTextBody = string

// TextExampleTextBody defines parameters for TextExample.
type TextExampleTextBody = string

// HeadersExampleParams defines parameters for HeadersExample.
type HeadersExampleParams struct {
	Header1 string `json:"header1"`
	Header2 *int   `json:"header2,omitempty"`
}

// JSONExampleJSONRequestBody defines body for JSONExample for application/json ContentType.
type JSONExampleJSONRequestBody = Example

// MultipartExampleMultipartRequestBody defines body for MultipartExample for multipart/form-data ContentType.
type MultipartExampleMultipartRequestBody = Example

// MultipartRelatedExampleMultipartRequestBody defines body for MultipartRelatedExample for multipart/related ContentType.
type MultipartRelatedExampleMultipartRequestBody = Example

// MultipleRequestAndResponseTypesJSONRequestBody defines body for MultipleRequestAndResponseTypes for application/json ContentType.
type MultipleRequestAndResponseTypesJSONRequestBody = Example

// MultipleRequestAndResponseTypesFormdataRequestBody defines body for MultipleRequestAndResponseTypes for application/x-www-form-urlencoded ContentType.
type MultipleRequestAndResponseTypesFormdataRequestBody = Example

// MultipleRequestAndResponseTypesMultipartRequestBody defines body for MultipleRequestAndResponseTypes for multipart/form-data ContentType.
type MultipleRequestAndResponseTypesMultipartRequestBody = Example

// MultipleRequestAndResponseTypesTextRequestBody defines body for MultipleRequestAndResponseTypes for text/plain ContentType.
type MultipleRequestAndResponseTypesTextRequestBody = MultipleRequestAndResponseTypesTextBody

// ReusableResponsesJSONRequestBody defines body for ReusableResponses for application/json ContentType.
type ReusableResponsesJSONRequestBody = Example

// TextExampleTextRequestBody defines body for TextExample for text/plain ContentType.
type TextExampleTextRequestBody = TextExampleTextBody

// URLEncodedExampleFormdataRequestBody defines body for URLEncodedExample for application/x-www-form-urlencoded ContentType.
type URLEncodedExampleFormdataRequestBody = Example

// HeadersExampleJSONRequestBody defines body for HeadersExample for application/json ContentType.
type HeadersExampleJSONRequestBody = Example

// UnionExampleJSONRequestBody defines body for UnionExample for application/json ContentType.
type UnionExampleJSONRequestBody = Example
