/*
 * MMS Example API
 *
 * Documentation for MMS API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type User struct {
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Admin     bool   `json:"admin,omitempty"`
	Password  string `json:"password,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	FullName  string `json:"fullName,omitempty"`
}
