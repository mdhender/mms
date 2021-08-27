/*
 * MMS Example API
 *
 * Documentation for MMS API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CommitsResponse struct {
	Messages []string `json:"messages,omitempty"`

	Rejected []Rejection `json:"rejected,omitempty"`

	Commits []map[string]interface{} `json:"commits,omitempty"`
}
