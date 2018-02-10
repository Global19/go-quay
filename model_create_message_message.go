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

// A single message
type CreateMessageMessage struct {
	// The actual message
	Content string `json:"content"`

	// The media type of the message
	MediaType string `json:"media_type"`

	// The severity of the message
	Severity string `json:"severity"`
}
