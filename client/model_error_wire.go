/*
 * WIRE API
 *
 * Moov WIRE implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ErrorWire ErrorWire
type ErrorWire struct {
	// ErrorCategory * `E` - Data Error * `F` - Insufficient Balance * `H` - Accountability Error * `I` - In Process or Intercepted * `W` - Cutoff Hour Error * `X` - Duplicate IMAD
	ErrorCategory string `json:"errorCategory,omitempty"`
	// ErrorCode
	ErrorCode string `json:"errorCode,omitempty"`
	// ErrorDescription
	ErrorDescription string `json:"errorDescription,omitempty"`
}
