/*
 * MMS Example API
 *
 * Documentation for MMS API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PermissionUpdateRequest struct {
	Action      string       `json:"action"`
	Permissions []Permission `json:"permissions"`
}