/*
 * MMS Example API
 *
 * Documentation for MMS API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Rejection struct {
	Object  *interface{} `json:"object,omitempty"`
	Code    int32        `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
}