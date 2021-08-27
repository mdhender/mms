/*
 * MMS Example API
 *
 * Documentation for MMS API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RefsResponse struct {
	Messages []string `json:"messages,omitempty"`

	Rejected []Rejection `json:"rejected,omitempty"`

	Refs []map[string]interface{} `json:"refs,omitempty"`
}
