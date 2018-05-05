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

// System status response
type SystemStatusResponse struct {

	Busy bool `json:"busy,omitempty"`

	Up bool `json:"up,omitempty"`

	Message string `json:"message,omitempty"`

	Detail *interface{} `json:"detail,omitempty"`
}
