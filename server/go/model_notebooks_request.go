/*
 * MMS Example API
 *
 * Documentation for MMS API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NotebooksRequest struct {
	Source string `json:"source,omitempty"`

	Comment string `json:"comment,omitempty"`

	Notebooks []map[string]interface{} `json:"notebooks"`
}
