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

// Description of an updated application.
type UpdateApp struct {
	// The URI for the application's OAuth redirect
	RedirectUri string `json:"redirect_uri"`

	// The e-mail address of the avatar to use for the application
	AvatarEmail string `json:"avatar_email,omitempty"`

	// The name of the application
	Name string `json:"name"`

	// The URI for the application's homepage
	ApplicationUri string `json:"application_uri"`

	// The human-readable description for the application
	Description string `json:"description,omitempty"`
}
