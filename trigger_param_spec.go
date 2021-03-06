/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.5
 * Contact: nurmi@anchore.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TriggerParamSpec struct {

	// Parameter name as it appears in policy document
	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	// An example value for the parameter (encoded as a string if the parameter is an object or list type)
	Example string `json:"example,omitempty"`

	// State of the trigger parameter
	State string `json:"state,omitempty"`

	// The name of another trigger that supercedes this on functionally if this is deprecated
	SupercededBy string `json:"superceded_by,omitempty"`

	// Is this a required parameter or optional
	Required bool `json:"required,omitempty"`

	// If present, a definition for validation of input. Typically a jsonschema object that can be used to validate an input against.
	Validator *interface{} `json:"validator,omitempty"`
}
