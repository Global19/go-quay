/*
 * Quay Frontend
 *
 * This API allows you to perform many of the operations required to work with Quay repositories, users, and organizations. You can find out more at <a href=\"https://quay.io\">Quay</a>.
 *
 * API version: v1
 * Contact: support@quay.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quay

// Description of an error
type ApiErrorDescription struct {
	// A reference to the error type resource
	Type_ string `json:"type"`

	// A more detailed description of the error that may include help for fixing the issue.
	Description string `json:"description"`

	// The title of the error. Can be used to uniquely identify the kind of error.
	Title string `json:"title"`
}
